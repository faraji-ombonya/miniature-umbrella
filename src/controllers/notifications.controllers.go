package controllers

import (
	"errors"

	"github.com/faraji-fuji/miniature-umbrella/src/filters"
	"github.com/faraji-fuji/miniature-umbrella/src/models"
	"github.com/faraji-fuji/miniature-umbrella/src/serializers"
	"github.com/faraji-fuji/miniature-umbrella/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetNotification ...
func GetNotification(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification

	result := models.Sql.First(&notification, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"detail": "record not found"})
		return
	}

	c.JSON(200, notification)
}

// GetNotifications ...
func GetNotifications(c *gin.Context) {
	var filter filters.Notification
	c.ShouldBind(&filter)

	var notifications []models.Notification
	result := models.Sql.Where(&filter).Find(&notifications)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"detail": "record not found",
		})

		return
	}

	c.JSON(200, notifications)
}

// CreateNotification ...
func CreateNotification(c *gin.Context) {
	var req serializers.CreateNotificationRequest

	c.ShouldBind(&req)

	notification := models.Notification{
		Sender:   req.Sender,
		Receiver: req.Receiver,
		Body:     req.Body,
		Channel:  req.Channel,
		Address:  req.Address,
	}

	result := models.Sql.Create(&notification)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
	}

	go func() {
		utils.SendToExchange(notification)
	}()

	c.JSON(200, notification)
}

// UpdateNotification ...
func UpdateNotification(c *gin.Context) {

	var req serializers.CreateNotificationRequest
	c.ShouldBind(&req)

	var notification models.Notification
	id := c.Param("id")

	result := models.Sql.First(&notification, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"detail": "record not found"})
		return
	}

	if req.Address != "" {
		notification.Address = req.Address
	}

	if req.Body != "" {
		notification.Body = req.Body
	}

	if req.Channel != "" {
		notification.Channel = req.Channel
	}

	if req.Receiver != "" {
		notification.Receiver = req.Receiver
	}

	if req.Sender != "" {
		notification.Sender = req.Sender
	}

	models.Sql.Save(&notification)

	c.JSON(200, notification)
}

// DeleteNotification ...
func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification

	result := models.Sql.First(&notification, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"detail": "record not found"})
		return
	}

	models.Sql.Delete(&notification)

	c.JSON(200, gin.H{"detail": "record deleted"})
}
