package worker

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/IlhamNur/mertani-device/models"
	"github.com/robfig/cron/v3"
)

var (
	BaseBackoffSeconds = 60
	MaxRetriesDefault  = 5
	HttpTimeoutSeconds = 10
)

func StartDeliveryWorker(db *gorm.DB) {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		if err := processPendingDeliveries(db); err != nil {
			fmt.Println("worker error:", err)
		}
	})
	c.Start()
}

func processPendingDeliveries(db *gorm.DB) error {
	now := time.Now()
	var logs []models.DeliveryLog

	if err := db.
		Where("status IN ? AND next_attempt_at <= ? AND retry_count < max_retries", []string{"pending", "failed"}, now).
		Order("next_attempt_at ASC").
		Limit(20).
		Find(&logs).Error; err != nil {
		return err
	}

	for _, l := range logs {
		if err := processSingleDelivery(db, l); err != nil {
			fmt.Printf("delivery id=%d failed processing: %v\n", l.ID, err)
		}
	}
	return nil
}

func processSingleDelivery(db *gorm.DB, log models.DeliveryLog) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var dl models.DeliveryLog
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&dl, log.ID).Error; err != nil {
			return err
		}

		if dl.Status == "sent" {
			return nil
		}

		payloadBytes, err := dl.Payload.MarshalJSON()
		if err != nil {
			return err
		}

		err = sendHTTP(dl.ClientURL, payloadBytes)
		now := time.Now()
		dl.LastAttemptAt = &now

		if err == nil {
			dl.Status = "sent"
			dl.UpdatedAt = time.Now()
			if err := tx.Save(&dl).Error; err != nil {
				return err
			}
			return nil
		}

		errStr := err.Error()
		dl.RetryCount += 1
		dl.LastError = &errStr

		if dl.RetryCount >= dl.MaxRetries {
			dl.Status = "permanent_failed"
			dl.UpdatedAt = time.Now()
			if err := tx.Save(&dl).Error; err != nil {
				return err
			}
			return nil
		}

		dl.Status = "failed"
		backoff := time.Duration(BaseBackoffSeconds*(1<<(dl.RetryCount-1))) * time.Second
		maxBackoff := 24 * time.Hour
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		dl.NextAttemptAt = time.Now().Add(backoff)
		dl.UpdatedAt = time.Now()

		if err := tx.Save(&dl).Error; err != nil {
			return err
		}
		return nil
	})
}

func sendHTTP(url string, body []byte) error {
	if url == "" {
		return errors.New("client_url empty")
	}

	timeout := time.Duration(HttpTimeoutSeconds) * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("http status %d", resp.StatusCode)
}
