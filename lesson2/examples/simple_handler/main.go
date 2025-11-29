// lesson2/examples/simple_handler/main.go

package main

import (
	"fmt"
	"net/http"
)

// helloHandler обрабатывает запросы на /hello и отправляет ответ "Hello Bitlab!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bitlab!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}
