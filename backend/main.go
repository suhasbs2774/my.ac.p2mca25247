package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Root test
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
