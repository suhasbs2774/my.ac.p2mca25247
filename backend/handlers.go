package main

import (
	"github.com/gin-gonic/gin"
)

func CreateNotification(c *gin.Context) {
	var n Notification

	if err := c.BindJSON(&n); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Exec(
		"INSERT INTO notifications(user_id, title, message, type) VALUES (?, ?, ?, ?)",
		n.UserID, n.Title, n.Message, n.Type,
	)

	c.JSON(200, gin.H{"message": "Notification created"})
}

func GetNotifications(c *gin.Context) {
	userID := c.Query("user_id")

	rows, _ := db.Query("SELECT * FROM notifications WHERE user_id=?", userID)
	defer rows.Close()

	var list []Notification

	for rows.Next() {
		var n Notification
		rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Message, &n.Type, &n.IsRead, &n.CreatedAt)
		list = append(list, n)
	}

	c.JSON(200, list)
}

func NotifyAll(c *gin.Context) {
	var req struct {
		Title   string `json:"title"`
		Message string `json:"message"`
		Type    string `json:"type"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rows, _ := db.Query("SELECT email FROM users")
	defer rows.Close()

	for rows.Next() {
		var email string
		rows.Scan(&email)

		SendEmail(email, req.Title, req.Message)

		db.Exec(
			"INSERT INTO notifications(user_id, title, message, type) VALUES (?, ?, ?, ?)",
			email, req.Title, req.Message, req.Type,
		)
	}

	c.JSON(200, gin.H{"message": "Sent to all students"})
}
