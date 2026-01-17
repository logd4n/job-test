package database

import (
	"job-test/internal/models"
	"log"
)

func SendMessage(message *models.Message) error {
	result := db.Create(message)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Сообщение отправлено! ID чата: %v", message.Chat_ID)
	return nil
}
