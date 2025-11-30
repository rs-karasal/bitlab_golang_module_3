// internal/service/user_service.go
package service

import "users_service_with_dto/internal/entities"

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	List() []entities.User
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(u entities.User) (entities.User, error) {
	if err := u.Validate(); err != nil {
		return entities.User{}, err
	}

	return s.repo.Create(u)
}

func (s *UserService) ListUsers() []entities.User {
	return s.repo.List()
}
