package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateNotification(c *gin.Context) {
	var n Notification

	if err := c.BindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(
		"INSERT INTO notifications(user_id, title, message, type) VALUES (?, ?, ?, ?)",
		n.UserID, n.Title, n.Message, n.Type,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	n.ID = int(id)

	c.JSON(201, n)
}

// GET ALL
func GetNotifications(c *gin.Context) {
	userID := c.Query("user_id")

	rows, err := db.Query("SELECT * FROM notifications WHERE user_id=?", userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []Notification

	for rows.Next() {
		var n Notification
		rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Message, &n.Type, &n.IsRead, &n.CreatedAt)
		list = append(list, n)
	}

	c.JSON(200, list)
}

// MARK AS READ
func MarkAsRead(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("UPDATE notifications SET is_read=TRUE WHERE id=?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Marked as read"})
}

// DELETE
func DeleteNotification(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM notifications WHERE id=?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted"})
}

func UnreadCount(c *gin.Context) {
	userID := c.Query("user_id")

	var count int
	db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id=? AND is_read=FALSE", userID).Scan(&count)

	c.JSON(200, gin.H{"unread_count": count})
}
