package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// создаем сервер
	server := &http.Server{
		Addr: "localhost:8080",
	}

	// создаем ручку
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bitlab!")
	})

	// запускаем сервер
	log.Printf("Starting server on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		// если ошибка, то выводим ошибку
		log.Fatal("Error starting server:", err)
	}
}
