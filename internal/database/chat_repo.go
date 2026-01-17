package database

import (
	"job-test/internal/models"
	"log"
)

func CreateChat(chat *models.Chat) error {
	result := db.Create(chat)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Чат %s успешно создан!\n", chat.Title)
	return nil
}

func GetChats() error {
	var chats []models.Chat
	result := db.Find(&chats)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Чаты получены:\n%v", chats)
	return nil
}

func GetChat(chat_ID uint, limit int) (*models.ResponseChat, error) {
	var chat models.Chat
	result := db.Preload("Messages").Where("id = ?", chat_ID).Order("created_at desc").Limit(limit).Find(&chat)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("Chat: %d, Messages: [%v]", chat.ID, chat.Messages)
	return &models.ResponseChat{
		ID:       chat_ID,
		Messages: chat.Messages,
	}, nil
}

func DeleteChat(chatID uint) error {
	result := db.Delete(&models.Chat{}, chatID)
	if result.Error != nil {
		log.Printf("Не удалось удалить чат!\n")
		return result.Error
	}
	log.Printf("Чат удален!\n")
	return nil
}
