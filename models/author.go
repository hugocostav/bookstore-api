package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name        string `json:"name"`
	Biography   string `json:"biography"`
	DateOfBirth string `json:"date_of_birth"`
}