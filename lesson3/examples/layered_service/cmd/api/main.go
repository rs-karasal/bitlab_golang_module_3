// cmd/api/main.go
package main

import (
	"log"
	"net/http"

	"bitlab_layered_service_example/internal/handlers"
	"bitlab_layered_service_example/internal/repository"
	"bitlab_layered_service_example/internal/service"
)

func main() {
	// 1. Создаём репозиторий (пока in-memory)
	userRepo := repository.NewUserRepository()

	// 2. Создаём сервис и передаём ему репозиторий
	userService := service.NewUserService(userRepo)

	// 3. Создаём хендлер и передаём ему сервис
	userHandler := handlers.NewUserHandler(userService)

	// 4. Настраиваем роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/users/me", userHandler.GetUser)

	// 5. Запускаем сервер
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
