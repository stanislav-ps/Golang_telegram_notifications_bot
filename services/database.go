package services

import (
	"database/sql"
	"obit_bot/utils"

	_ "github.com/mattn/go-sqlite3"
)

// db - объект базы данных
var db *sql.DB

// OpenDB открывает соединение с базой данных
func OpenDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		utils.ErrorLogger.Println(err)
		return err
	}
	return nil
}

// GetConn возвращает соединение с базой данных
func GetConn() (*sql.DB, error) {
	if db == nil {
		if err := OpenDB(); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// CreateNotification сохраняет уведомление в базе данных
func CreateNotification(userID int, dateTime, text string, frequency int) error {
	// Получаем соединение с базой данных
	conn, err := GetConn()
	if err != nil {
		utils.ErrorLogger.Println(err)
		return err
	}

	// Создаем подготовленное выражение для вставки записи в базу данных
	stmt, err := conn.Prepare("INSERT INTO notifications(user_id, date_time, text, frequency) VALUES(?, ?, ?, ?)")
	if err != nil {
		utils.ErrorLogger.Println(err)
		return err
	}

	// Выполняем запрос на вставку записи в базу данных
	_, err = stmt.Exec(userID, dateTime, text, frequency)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return err
	}

	return nil
}
