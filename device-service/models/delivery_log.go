package models

import (
	"time"

	"gorm.io/datatypes"
)

type DeliveryLog struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	DeviceID      *uint          `json:"device_id"`
	Payload       datatypes.JSON `gorm:"type:jsonb" json:"payload"`
	ClientURL     string         `json:"client_url"`
	Status        string         `json:"status"`
	RetryCount    int            `json:"retry_count"`
	MaxRetries    int            `json:"max_retries"`
	NextAttemptAt time.Time      `json:"next_attempt_at"`
	LastAttemptAt *time.Time     `json:"last_attempt_at"`
	LastError     *string        `json:"last_error"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
