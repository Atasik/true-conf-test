package repository

import (
	"encoding/json"
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

func readJSONFromFile(dir string, v interface{}) error {
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}

func writeJSONToFile(dir string, v interface{}) error {
	file, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(v)
}
