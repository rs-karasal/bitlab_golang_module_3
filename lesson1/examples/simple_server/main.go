package main

import (
	"fmt"
	"net/http"
)

func main() {
	// создаем ручку
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bitlab!")
	})

	// запускаем сервер
	http.ListenAndServe(":8080", nil)
}
