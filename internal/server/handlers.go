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

// Маленький обработчик для тестирования
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Обработчик для localhost:8080/chats
func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	// Запрещаем все методы, кроме POST
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

	// Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка чтения запроса: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Создаем структуру чата
	chat := models.Chat{
		Title: string(body),
	}
	// Отправляем запрос
	err = database.CreateChat(&chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
Обработчик для:
1. localhost:8080/chats/{id}/message
2. localhost:8080/chats/{id}&limit=N
3. localhost:8080/chats/{id}
*/
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST, GET, DELETE
	if r.Method != http.MethodPost && r.Method != http.MethodGet && r.Method != http.MethodDelete {
		log.Printf("Метод %v не поддерживается!\n", r.Method)
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	// Сплитуем адрес для получения ID
	urlParts := strings.Split(r.URL.Path, "/") //[  chats X message] X - chat id
	if len(urlParts) < 3 {
		http.Error(w, "Неверный URL!", http.StatusBadRequest)
		return
	}

	// Из сплита получаем ID чата
	chatID, err := strconv.Atoi(urlParts[2])
	if err != nil {
		log.Printf("Не удалось получить ID!\n")
		http.Error(w, "Не удалось получить ID: "+err.Error(), http.StatusBadGateway)
		return
	}

	// Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка чтения запроса: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Проверка метода
	if r.Method == http.MethodPost {
		message := models.Message{
			Chat_ID: uint(chatID),
			Text:    strings.TrimSpace(string(body)),
		}

		// Отправялем запрос
		err = database.SendMessage(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}

	// Проверка метода
	if r.Method == http.MethodGet {
		// Устанавливаем query для получения limit из URI
		query := r.URL.Query()

		// Получаем limit
		limit := query.Get("limit")

		// Устанавливаем дефолтное значение
		limitInt := 20

		// Проверяем limit из URI
		if limit != "" {
			if limitInt, err = strconv.Atoi(limit); err != nil {
				log.Printf("Не удалось получить limit: %v\n", limitInt)
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}

		// Устанавливаем максимальное значение
		if limitInt > 100 {
			limitInt = 100
		}

		// Отправляем запрос
		chat, err := database.GetChat(uint(chatID), limitInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		//Возвращаем ответ в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chat)
	}

	// Проверка метода
	if r.Method == http.MethodDelete {
		// Отправляем запрос
		err = database.DeleteChat(uint(chatID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		// Возращаем статус 204
		w.WriteHeader(http.StatusNoContent)
	}
}
