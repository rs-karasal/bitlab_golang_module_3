// internal/db/setup.go
package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // импортируем драйвер Postgres
)

func MustInitDB() *sql.DB {
	// DSN — строка подключения к БД.
	dsn := "postgres://user:password@localhost:5432/myshop?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	// Проверим, что подключение работает
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	log.Println("connected to postgres!")

	return db
}
