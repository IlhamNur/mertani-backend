package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/IlhamNur/mertani-device/controllers"
	"github.com/IlhamNur/mertani-device/models"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: Root .env not found, using system environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	db.AutoMigrate(&models.Device{}, &models.DeliveryLog{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/devices", func(c *gin.Context) { controllers.GetDevices(c, db) })
	r.GET("/devices/:id", func(c *gin.Context) { controllers.GetDevice(c, db) })
	r.POST("/devices", func(c *gin.Context) { controllers.CreateDevice(c, db) })
	r.PUT("/devices/:id", func(c *gin.Context) { controllers.UpdateDevice(c, db) })
	r.DELETE("/devices/:id", func(c *gin.Context) { controllers.DeleteDevice(c, db) })
	r.GET("/delivery-logs", func(c *gin.Context) { controllers.GetDeliveryLogs(c, db) })

	port := os.Getenv("DEVICE_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Device Service running on port %s", port)
	r.Run(":" + port)
}
