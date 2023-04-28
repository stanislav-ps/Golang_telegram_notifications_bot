package utils

import (
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger // Переменная для хранения объекта логгера информационных сообщений
	ErrorLogger *log.Logger // Переменная для хранения объекта логгера сообщений об ошибках
)

func init() {
	// Создание файла для записи логов
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Создание объектов логгеров для информационных сообщений и сообщений об ошибках
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Запуск функции для очистки логов раз в неделю
	go func() {
		for range time.Tick(7 * 24 * time.Hour) {
			err := os.Truncate("logfile.log", 0)
			if err != nil {
				log.Println("Failed to truncate log file:", err)
			} else {
				log.Println("Log file truncated")
			}
		}
	}()
}
