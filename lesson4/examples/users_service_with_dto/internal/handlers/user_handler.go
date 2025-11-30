// internal/handlers/user_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"users_service_with_dto/internal/handlers/dto"
	"users_service_with_dto/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createUser(w, r)
	case http.MethodGet:
		h.listUsers(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user := dto.NewUserFromCreateRequest(req)
	created, err := h.service.CreateUser(user)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}
	createdResp := dto.NewUserResponse(created)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdResp)
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.ListUsers()

	usersResp := make([]dto.UserResponse, len(users))
	for i, user := range users {
		usersResp[i] = dto.NewUserResponse(user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersResp)
}
