// internal/service/user_service.go
package service

import "bitlab_layered_service_example/internal/entities"

type UserRepository interface {
	FindByID(id int) (entities.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByID(id int) (entities.User, error) {
	// тут можно добавить проверки, логику, кеш и т.д.
	return s.repo.FindByID(id)
}
