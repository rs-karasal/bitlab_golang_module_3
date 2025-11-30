// internal/entities/user.go
package entities

import "errors"

type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Password string
}

// Validate проверяет, что пользователь имеет все необходимые поля
func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Age < 0 {
		return errors.New("age must be positive")
	}
	return nil
}

// IsAdult проверяет, является ли пользователь взрослым
func (u User) IsAdult() bool {
	return u.Age >= 18
}

// CanAccessRestrictedArea проверяет, может ли пользователь получить доступ к ограниченной области
func (u User) CanAccessRestrictedArea() bool {
	return u.IsAdult() && u.Email != ""
}
