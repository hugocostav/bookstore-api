package models

import "github.com/jinzhu/gorm"

type Publisher struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
}