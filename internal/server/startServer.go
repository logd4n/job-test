package server

import (
	"job-test/internal/database"
	"log"
	"net/http"
)

func StartServer() error {
	// Подключаемся к БД
	err := database.Connect()
	if err != nil {
		return err
	}
	log.Printf("Подключение к БД выполнено успешно!\n")

	// Обработчики
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/chats", CreateChatHandler)
	http.HandleFunc("/chats/", SendMessageHandler)

	// Запускаем сервер localhost:8080
	log.Printf("Сервер запущен!\n\n")
	err = http.ListenAndServe(":8080", nil)
	return err
}
