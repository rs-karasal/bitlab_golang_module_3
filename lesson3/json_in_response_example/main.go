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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/profile", profileHandler)

	log.Printf("Starting server on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// homeHandler - обработчик для ручки /
//
// отправляет ответ в формате text/plain
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Hello Bitlab!")
}

// profileHandler - обработчик для ручки /profile
//
// отправляет ответ в формате JSON с информацией о профиле пользователя
func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"name":       "Гарри",
		"surname":    "Поттер",
		"patronymic": "Джеймсович",
		"age":        "14",
		"email":      "harry.potter@example.com",
	})
}
