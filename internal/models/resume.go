package models

type Resume struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	UserID      int    `json:"user_id"`
}