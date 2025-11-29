package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Всё в одном файле и в одном пакете main.
// Здесь и сущность, и "хранилище", и бизнес-логика, и HTTP-обработка.

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var (
	notes  = []Note{}
	nextID = 1
)

func notesHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь свалено всё:
	// - разбор метода (GET/POST/DELETE)
	// - валидация
	// - работа с "БД" (in-memory слайс)
	// - формирование JSON-ответа и кодов статуса

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// GET /notes — вернуть список заметок
		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, "failed to encode notes", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		// POST /notes — создать новую заметку
		var req struct {
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if req.Text == "" {
			http.Error(w, "text is required", http.StatusBadRequest)
			return
		}

		note := Note{
			ID:   nextID,
			Text: req.Text,
		}
		nextID++

		notes = append(notes, note)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(note); err != nil {
			http.Error(w, "failed to encode note", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		// DELETE /notes?id=1 — удалить заметку по id
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id must be an integer", http.StatusBadRequest)
			return
		}

		// ищем и удаляем заметку из слайса
		found := false
		newNotes := make([]Note, 0, len(notes))
		for _, n := range notes {
			if n.ID == id {
				found = true
				continue
			}
			newNotes = append(newNotes, n)
		}
		notes = newNotes

		if !found {
			http.Error(w, fmt.Sprintf("note with id=%d not found", id), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/notes", notesHandler)

	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
