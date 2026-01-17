package main

import (
	"job-test/internal/server"
	"log"
)

func main() {
	//Запускаем сервер
	err := server.StartServer()
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v\n", err.Error())
	}
}
