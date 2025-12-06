package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const migrationsDir = "migrations"

func main() {
	// Флаги командной строки
	command := flag.String("command", "up", "goose command: up, down, status, reset, version")
	flag.Parse()

	// Получаем DSN из переменной окружения
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/migration_example?sslmode=disable"
	}

	// Подключаемся к БД
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	// Выполняем команду goose
	if err := runGooseCommand(db, *command); err != nil {
		log.Fatalf("goose %s: %v", *command, err)
	}
}

func runGooseCommand(db *sql.DB, command string) error {
	switch command {
	case "up":
		// Применить все миграции
		return goose.Up(db, migrationsDir)

	case "up-one":
		// Применить только одну следующую миграцию
		return goose.UpByOne(db, migrationsDir)

	case "down":
		// Откатить последнюю миграцию
		return goose.Down(db, migrationsDir)

	case "reset":
		// Откатить все миграции
		return goose.Reset(db, migrationsDir)

	case "status":
		// Показать статус миграций
		return goose.Status(db, migrationsDir)

	case "version":
		// Показать текущую версию
		version, err := goose.GetDBVersion(db)
		if err != nil {
			return err
		}
		fmt.Printf("Current DB version: %d\n", version)
		return nil

	default:
		return fmt.Errorf("unknown command: %s. Available: up, up-one, down, reset, status, version", command)
	}
}
