package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/login", loginHandler)

	log.Printf("Starting server on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// loginHandler - обработчик для ручки /login
//
// Получает информацию о пользователе из тела запроса в формате JSON
// и проверяет его почту и пароль
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})

		return
	}

	if user.Email == "admin@example.com" && user.Password == "123456" {
		fmt.Fprintf(w, "Login successful")
	} else {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	}
}
