# **Документация** 


## **Оглавление**

1. [Общее описание](#Общее%20описание)
    
2. [Технический стек](#технический-стек)
    
3. [Установка и запуск](#установка-и-запуск)
    
4. [API Endpoints](#api-endpoints)
    
5. [Тестирование](#тестирование)

6. [Способы связи](#способы-связи)

---

## **Общее описание**

Тестовое задание.

Автор проекта: <b><u>LOG</u></b>  
---

## **Технический стек**

- **Язык:** Go 1.22.4+
- **Фреймворк:** стандартная библиотека net/http
- **База данных:** PostgreSQL
- **ORM:** gorm.io/gorm
- **Драйвер для БД** gorm.io/driver/postgres v1.6.0	
- **Контейнеризация:** Docker, docker-compose


---

## Установка и запуск

### Предварительные требования
- Go 1.22.4 или выше
- PostgreSQL 12+
---

### Шаги установки

#### 1. **Клонирование репозитория**
   ```bash
   git clone https://github.com/logd4n/job-test.git

   cd job-test
   ```

#### <br>2. **Запуск сервера**</br>
   Так как Dockerfile и docker-compose уже настроены, Вам нужно выполнить только запуск сервера:
   ```bash
   docker-compose up -d --build
   ```
---

## API Endpoints

Для взаимодействия с приложением необходимо отправлять HTTP-запросы на определенные URI.

### **Создать чат:**
Метод POST
```
localhost:8080/chats
```

``
Тело запроса должно содержать название будущего чата
``

### **Отправить сообщение в чат:**
Метод POST
```
localhost:8080/chats/{id}/message

Пример: localhost:8080/chats/1/message
```

``
Тело запроса должно включать в себя всё содержимое сообщения
``

### **Получить чат и последние N сообщений:**
Метод GET
```
localhost:8080/chats/{id}

Пример: localhost:8080/chats/1
```

``
Ответ возвращается в формате JSON: id, []messages
``

### **Удалить чат:**
Метод DELETE
```
localhost:8080/chats/{id}

Пример: localhost:8080/chats/1
```

``
Вместе с чатом удалятся все его сообщения 
``

---

## Тестирование
Для теста используется httptest.
Были созданы небольшой обработчик и файл, тестирующий этот обработчик.

``handler_test.go:``
```go
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
```

``handlers.go``
```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод "+r.Method+" не поддерживается!", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
```
---

## **Способы связи:**
 - **Telegram:** @logd4n
 - **VK:** @logd4n