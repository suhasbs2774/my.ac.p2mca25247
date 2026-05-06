package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server running"})
	})

	api := r.Group("/api")
	{
		api.POST("/notifications", CreateNotification)
		api.GET("/notifications", GetNotifications)
		api.POST("/notify-all", NotifyAll)
	}

	r.Run(":8080")
}
