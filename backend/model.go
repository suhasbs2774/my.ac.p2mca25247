package main

type Notification struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	IsRead    bool   `json:"is_read"`
	CreatedAt string `json:"created_at"`
}
