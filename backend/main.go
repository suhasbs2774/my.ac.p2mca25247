package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()

	r := gin.Default()

	// Root test
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})

	api := r.Group("/api")
	{
		api.POST("/notifications", CreateNotification)
		api.GET("/notifications", GetNotifications)
		api.PUT("/notifications/:id/read", MarkAsRead)
		api.DELETE("/notifications/:id", DeleteNotification)
		api.GET("/notifications/unread-count", UnreadCount)
	}
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "POST working"})
	})

	r.Run(":8080")
}
