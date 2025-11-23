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

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/hello", helloHandler)

	// TODO: реализовать свои ручки с параметрами запроса

	log.Printf("Starting server on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// usersHandler - обработчик для ручки /users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		// Method not allowed
		http.Error(w, "Mehtod not allowed", http.StatusMethodNotAllowed)

		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Список пользователей")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Создание пользователя")
}

// homeHandler - обработчик для ручки /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

// helloHandler - обработчик для ручки /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// получаем параметр name из URL
	name := r.URL.Query().Get("name")

	if name == "" {
		// если параметр name не передан, то выводим сообщение "Hello Bitlab!"
		fmt.Fprint(w, "Hello Bitlab!")

		return
	}

	// если параметр name передан, то выводим сообщение "Hello {name}!"
	fmt.Fprintf(w, "Hello %s!", name)
}
