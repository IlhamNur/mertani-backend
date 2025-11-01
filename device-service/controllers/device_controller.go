package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/IlhamNur/mertani-device/models"
)

func GetDevices(c *gin.Context, db *gorm.DB) {
	var devices []models.Device
	db.Find(&devices)
	c.JSON(http.StatusOK, devices)
}

func GetDevice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var device models.Device
	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, device)
}

func CreateDevice(c *gin.Context, db *gorm.DB) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()
	db.Create(&device)
	c.JSON(http.StatusCreated, device)
}

func UpdateDevice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var device models.Device
	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device.UpdatedAt = time.Now()
	db.Save(&device)
	c.JSON(http.StatusOK, device)
}

func DeleteDevice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Device{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
