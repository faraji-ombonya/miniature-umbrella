package routes

import (
	"github.com/faraji-fuji/miniature-umbrella/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	notifications := router.Group("/notifications")
	{
		notifications.GET("/", controllers.GetNotifications)
		notifications.GET("/:id", controllers.GetNotification)
		notifications.POST("/", controllers.CreateNotification)
		notifications.PUT("/:id", controllers.UpdateNotification)
		notifications.DELETE("/:id", controllers.DeleteNotification)
	}

	router.Run(":8080")
}
