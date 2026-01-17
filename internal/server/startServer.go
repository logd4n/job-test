package server

import (
	"job-test/internal/database"
	"log"
	"net/http"
)

func StartServer() error {
	err := database.Connect()
	if err != nil {
		return err
	}
	log.Printf("Подключение к БД выполнено успешно!\n")

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/chats", CreateChatHandler)
	http.HandleFunc("/chats/", SendMessageHandler)

	log.Printf("Сервер запущен!\n\n")
	err = http.ListenAndServe(":8080", nil)
	return err
}
