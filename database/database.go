package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Инициализация подключения к БД
func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	// Проверка соединения
	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("Successfully connected to database!")
	return nil

}

// Получаем текущее подключение
func GetDB() *sql.DB {
	return db
}

// CloseDB закрывает соединение с базой данных
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
