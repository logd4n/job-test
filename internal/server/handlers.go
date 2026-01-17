package server

import (
	"encoding/json"
	"io"
	"job-test/internal/database"
	"job-test/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Метод %v не поддерживается!\n", r.Method)
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	/*
		if r.Header.Get("Content-Type") != "text/plain" {
			log.Printf("Content-Type должен быть text/plain!\n")
			http.Error(w, "Content-Type должен быть text/plain!", http.StatusBadRequest)
			return
		}
	*/

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка чтения запроса: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	chat := models.Chat{
		Title: string(body),
	}
	err = database.CreateChat(&chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet && r.Method != http.MethodDelete {
		log.Printf("Метод %v не поддерживается!\n", r.Method)
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/") //[  chats X message] X - chat id
	if len(urlParts) < 3 {
		http.Error(w, "Неверный URL!", http.StatusBadRequest)
		return
	}

	chatID, err := strconv.Atoi(urlParts[2])
	if err != nil {
		log.Printf("Не удалось получить ID!\n")
		http.Error(w, "Не удалось получить ID: "+err.Error(), http.StatusBadGateway)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка чтения запроса: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	if r.Method == http.MethodPost {
		message := models.Message{
			Chat_ID: uint(chatID),
			Text:    strings.TrimSpace(string(body)),
		}

		err = database.SendMessage(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}

	if r.Method == http.MethodGet {
		query := r.URL.Query()

		limit := query.Get("limit")
		limitInt := 20

		if limit != "" {
			if limitInt, err = strconv.Atoi(limit); err != nil {
				log.Printf("Не удалось получить limit: %v\n", limitInt)
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}

		if limitInt > 100 {
			limitInt = 100
		}

		chat, err := database.GetChat(uint(chatID), limitInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chat)
	}

	if r.Method == http.MethodDelete {
		err = database.DeleteChat(uint(chatID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
