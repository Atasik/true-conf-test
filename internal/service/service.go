package service

import "refactoring/internal/domain"

type Services struct {
	User
}

type User interface {
	GetAll() (domain.UserList, error)
	Create(user domain.User) (int, error)
	GetByID(id int) (domain.User, error)
	Update(id int, input domain.UpdateUserInput) error
	Delete(id int) error
}

func NewServices(userService User) *Services {
	return &Services{
		User: userService,
	}
}
