package database

import (
	"fmt"
	"job-test/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	dataBase, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	db = dataBase

	err = migrate()
	if err != nil {
		return err
	}

	return nil
}

func migrate() error {
	err := db.AutoMigrate(&models.Chat{}, &models.Message{})
	if err != nil {
		return err
	}

	return nil
}
