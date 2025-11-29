// lesson2/examples/parameters/main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CreateUserRequest структура для запроса на создание пользователя
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// createUserHandler обрабатывает запрос на создание пользователя с телом запроса
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "method not allowed")
		return
	}

	var req CreateUserRequest
	// Декодируем JSON из тела запроса в структуру
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprint(w, "invalid json")
		return
	}

	// Простая проверка
	if req.Name == "" || req.Email == "" {
		fmt.Fprint(w, "name and email are required")
		return
	}

	// Здесь мы бы вызывали сервис / сохраняли в БД
	// Пока просто отвечаем
	fmt.Fprintf(w, "user %s with email %s created", req.Name, req.Email)
}

// searchHandler обрабатывает запрос с query-параметрами
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")

	if query == "" {
		fmt.Fprint(w, "query is required")
		return
	}

	if page == "" {
		page = "1"
	}

	fmt.Fprintf(w, "Ищем '%s', страница %s\n", query, page)
}

// userHandler обрабатывает запрос с параметром в пути
func userHandler(w http.ResponseWriter, r *http.Request) {
	// ожидаем путь вида /users/10
	path := r.URL.Path                // например, "/users/10"
	parts := strings.Split(path, "/") // ["", "users", "10"]

	if len(parts) != 3 {
		fmt.Fprint(w, "invalid path")
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprint(w, "invalid user id")
		return
	}

	fmt.Fprintf(w, "Пользователь с id = %d\n", id)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/users/", userHandler)
	mux.HandleFunc("/users/create", createUserHandler)

	http.ListenAndServe(":8080", mux)
}
