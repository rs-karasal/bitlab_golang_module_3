// internal/repository/user_repo.go
package repository

import (
	"bitlab_layered_service_example/internal/entities"
	"errors"
)

type UserRepository struct {
	data map[int]entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: map[int]entities.User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

func (r *UserRepository) FindByID(id int) (entities.User, error) {
	user, ok := r.data[id]
	if !ok {
		return entities.User{}, errors.New("not found")
	}
	return user, nil
}
