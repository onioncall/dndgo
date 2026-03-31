package models

type Note struct {
	Title   string `json:"title" clover:"title"`
	Content string `json:"content" clover:"content"`
}
