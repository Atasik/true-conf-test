package repository

import (
	"encoding/json"
	"io/fs"
	"os"
	"refactoring/internal/domain"
)

type Repository struct {
	UserRepo
}

type UserRepo interface {
	GetAll() (domain.UserList, error)
	Create(domain.User) (int, error)
	GetByID(id int) (domain.User, error)
	Update(id int, input domain.UpdateUserInput) error
	Delete(id int) error
}

func NewRepository(userRepo UserRepo) *Repository {
	return &Repository{
		UserRepo: userRepo,
	}
}

func readJSONFromFile(dir string, v any) error {
	f, err := os.ReadFile(dir)
	if err != nil {
		return err
	}
	return json.Unmarshal(f, v)
}

func writeJSONToFile(dir string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(dir, b, fs.ModePerm)
}
