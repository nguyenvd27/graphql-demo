package model

type Event struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}
