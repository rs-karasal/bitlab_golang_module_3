// internal/http/user_handler.go
package handlers

import (
	"bitlab_layered_service_example/internal/entities"
	"encoding/json"
	"net/http"
)

type UserService interface {
	GetUserByID(id int) (entities.User, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// здесь мы бы достали id из path или query
	id := 1 // пока захардкодим для примера

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
