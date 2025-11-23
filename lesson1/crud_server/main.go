package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}

	// GET localhost:8080
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// GET localhost:8080/users
			fmt.Fprint(w, "Получение пользователя методом GET прошло успешно!")
		case http.MethodPost:
			// POST localhost:8080/users
			fmt.Fprint(w, "Создание пользователя методом POST прошло успешно!")
		case http.MethodPut:
			// PUT localhost:8080/users
			fmt.Fprint(w, "Обновление пользователя методом PUT прошло успешно!")
		case http.MethodPatch:
			// PATCH localhost:8080/users
			fmt.Fprint(w, "Частичное обновление пользователя методом PATCH прошло успешно!")
		case http.MethodDelete:
			// DELETE localhost:8080/users
			fmt.Fprint(w, "Удаление пользователя методом DELETE прошло успешно!")
		default:
			fmt.Fprintf(w, "method %s is forbidden", r.Method)
		}
	})

	// TODO: реализовать свои ручки

	log.Printf("Starting server on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
