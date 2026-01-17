package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"job-test/internal/server"
)

func TestHelloHandler(t *testing.T) {
	// Создаем HTTP запрос
	req := httptest.NewRequest("GET", "/hello", nil)

	// Создаем ResponseRecorder (записываем ответ)
	rr := httptest.NewRecorder()

	// Вызываем обработчик напрямую
	server.HelloHandler(rr, req)

	// Проверяем статус-код
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %v, а получили %v", http.StatusOK, rr.Code)
	}
}
