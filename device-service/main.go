package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/IlhamNur/mertani-device/controllers"
	"github.com/IlhamNur/mertani-device/models"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=admin dbname=mertani port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Device{})

	r := gin.Default()
	r.GET("/devices", func(c *gin.Context) { controllers.GetDevices(c, db) })
	r.GET("/devices/:id", func(c *gin.Context) { controllers.GetDevice(c, db) })
	r.POST("/devices", func(c *gin.Context) { controllers.CreateDevice(c, db) })
	r.PUT("/devices/:id", func(c *gin.Context) { controllers.UpdateDevice(c, db) })
	r.DELETE("/devices/:id", func(c *gin.Context) { controllers.DeleteDevice(c, db) })

	r.Run(":8080")
}
