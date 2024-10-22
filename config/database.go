package config

import (
	"bookstore-api/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL não está definida")
	}

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Configurações adicionais do banco de dados
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{}, &models.Author{}, &models.Publisher{}).Error
}