// lesson2/examples/error_handling/main.go

package main

import "net/http"

func safeGetUserHandler(w http.ResponseWriter, r *http.Request) {
	// допустим, пытаемся достать id пользователя
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// ... ищем пользователя где-то в хранилище
	// допустим, не нашли:
	http.Error(w, "user not found", http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/", safeGetUserHandler)

	http.ListenAndServe(":8080", mux)
}
