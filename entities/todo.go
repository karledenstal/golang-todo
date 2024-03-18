package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
