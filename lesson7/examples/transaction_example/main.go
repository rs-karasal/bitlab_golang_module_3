package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "postgres://user:password@localhost:5432/myshop?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	log.Println("connected to postgres!")

	// Пример транзакции: перевод товара со склада
	err = transferProduct(db, 1, 5)
	if err != nil {
		log.Printf("transfer failed: %v", err)
	} else {
		log.Println("transfer completed successfully")
	}
}

// transferProduct уменьшает quantity продукта и записывает лог операции
func transferProduct(db *sql.DB, productID int, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// важно: если что-то пойдёт не так — откатить
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Проверяем наличие товара
	var currentQty int
	err = tx.QueryRow("SELECT quantity FROM products WHERE id = $1", productID).Scan(&currentQty)
	if err != nil {
		return fmt.Errorf("failed to get product quantity: %w", err)
	}

	if currentQty < amount {
		return fmt.Errorf("insufficient quantity: have %d, need %d", currentQty, amount)
	}

	// Уменьшаем количество товара
	_, err = tx.Exec("UPDATE products SET quantity = quantity - $1 WHERE id = $2", amount, productID)
	if err != nil {
		return fmt.Errorf("failed to update product quantity: %w", err)
	}

	log.Printf("reduced product %d quantity by %d (was: %d, now: %d)", productID, amount, currentQty, currentQty-amount)

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
