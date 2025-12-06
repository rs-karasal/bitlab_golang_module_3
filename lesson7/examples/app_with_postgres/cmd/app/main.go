package main

import (
	dbCon "app_with_postgres/internal/db"
	"app_with_postgres/internal/handlers"
	"app_with_postgres/internal/repository"
	"app_with_postgres/internal/service"
	"log"
	"net/http"
)

func main() {
	// Пулл соединений с БД Postgres
	db := dbCon.MustInitDB()

	// 1. Создаём репозиторий с подключением к Postgres
	productRepo := repository.NewPostgresProductRepository(db)

	// 2. Создаём сервис и передаём ему репозиторий
	productService := service.NewProductService(productRepo)

	// 3. Создаём хендлер и передаём ему сервис
	productHandler := handlers.NewProductHandler(productService)

	// 4. Настраиваем роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/products", productHandler.HandlerProducts)

	// 5. Запускаем сервер
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
