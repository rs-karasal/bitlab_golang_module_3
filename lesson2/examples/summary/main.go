// lesson2/examples/summary/main.go

package main

import (
	"encoding/json"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func helloQueryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Bitlab"
	}

	resp := HelloResponse{
		Message: "Hello, " + name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func helloBodyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	resp := HelloResponse{
		Message: "Hello, " + req.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()

	// GET /hello?name=...
	mux.HandleFunc("/hello", helloQueryHandler)

	// POST /hello-body { "name": "Alice" }
	mux.HandleFunc("/hello-body", helloBodyHandler)

	http.ListenAndServe(":8080", mux)
}
