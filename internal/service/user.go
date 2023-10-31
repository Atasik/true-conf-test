package service

import (
	"refactoring/internal/domain"
	"refactoring/internal/repository"
)

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAll() (domain.UserList, error) {
	return s.userRepo.GetAll()
}

func (s *userService) Create(user domain.User) (int, error) {
	return s.userRepo.Create(user)
}

func (s *userService) GetByID(id int) (domain.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) Update(id int, input domain.UpdateUserInput) error {
	return s.userRepo.Update(id, input)
}

func (s *userService) Delete(id int) error {
	return s.userRepo.Delete(id)
}
