package main

import (
	"fmt"
	"net/http"
)

func main() {
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

	http.ListenAndServe(":8080", nil)
}
