// lesson2/examples/simple_router/main.go

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from /hello")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}

func main() {
	// создаем новый ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	// запускаем сервер с использованием ServeMux
	http.ListenAndServe(":8080", mux)
}
