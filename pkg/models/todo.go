package models

import "github.com/jinzhu/gorm"

// Todo very basic table to test a few features
type Todo struct {
	gorm.Model
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
