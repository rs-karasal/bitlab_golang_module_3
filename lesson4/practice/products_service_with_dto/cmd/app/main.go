package main

import (
	"log"
	"net/http"
	"products_service_with_dto/internal/handlers"
	"products_service_with_dto/internal/repository"
	"products_service_with_dto/internal/service"
)

func main() {
	// 1. Создаём репозиторий (пока in-memory)
	productRepo := repository.NewProductRepository()

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
