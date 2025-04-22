package repositories

import (
	"database/sql"

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

	return db.Ping() // Проверяем соединение
}

// Получаем текущее подключение
func GetDB() *sql.DB {
	return db
}
