package model

type Event struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user"`
}

type NewEvent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}
