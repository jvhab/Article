package model

type Article struct {
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	Body        string   `json:"body" db:"body"`
	Counts      int      `json:"counts" db:"counts"`
	TagList     []string `json:"tag_list" db:"tag_list"`
}
