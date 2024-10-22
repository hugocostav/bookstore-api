package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}