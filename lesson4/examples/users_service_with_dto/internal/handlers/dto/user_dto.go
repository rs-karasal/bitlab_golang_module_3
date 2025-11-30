// internal/handlers/dto/user_dto.go
package dto

import "users_service_with_dto/internal/entities"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// NewUserFromCreateRequest создаёт новый пользователь из запроса на создание
func NewUserFromCreateRequest(req CreateUserRequest) entities.User {
	return entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Password: req.Password,
	}
}

// NewUserResponse создаёт новый ответ на запрос пользователя
func NewUserResponse(u entities.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}
