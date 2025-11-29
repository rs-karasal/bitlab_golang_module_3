// lesson2/examples/response_structure/main.go

package main

import (
	"net/http"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK) // 200 OK

	w.Write([]byte("Successfully listed users"))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated) // 201 Created

	w.Write([]byte("Successfully created user"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", listHandler)
	mux.HandleFunc("/users/create", createHandler)

	http.ListenAndServe(":8080", mux)
}
