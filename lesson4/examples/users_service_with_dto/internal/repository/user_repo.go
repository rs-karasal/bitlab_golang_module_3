// internal/repository/user_repo.go
package repository

import (
	"users_service_with_dto/internal/entities"
)

type UserRepository struct {
	data   map[int]entities.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: map[int]entities.User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
		nextID: 3,
	}
}

func (r *UserRepository) Create(user entities.User) (entities.User, error) {
	user.ID = r.nextID // генерируем новый ID
	r.data[user.ID] = user
	r.nextID++

	return user, nil
}

func (r *UserRepository) List() []entities.User {
	users := make([]entities.User, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users
}
