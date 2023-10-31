package repository

import (
	"refactoring/internal/domain"
	"sync"
)

type userRepo struct {
	storeDir string
	mutex    *sync.RWMutex
}

func NewUserRepo(storeDir string) *userRepo {
	return &userRepo{
		storeDir: storeDir,
		mutex:    &sync.RWMutex{},
	}
}

func (repo *userRepo) GetAll() (domain.UserList, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	var s domain.UserStore

	if err := readJSONFromFile(repo.storeDir, &s); err != nil {
		return domain.UserList{}, err
	}
	return s.List, nil
}

func (repo *userRepo) Create(user domain.User) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var s domain.UserStore
	if err := readJSONFromFile(repo.storeDir, &s); err != nil {
		return 0, err
	}

	s.Increment++
	s.List[s.Increment] = user

	if err := writeJSONToFile(repo.storeDir, &s); err != nil {
		return 0, err
	}
	return s.Increment, nil
}

func (repo *userRepo) GetByID(id int) (domain.User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	var s domain.UserStore
	if err := readJSONFromFile(repo.storeDir, &s); err != nil {
		return domain.User{}, err
	}

	if _, ok := s.List[id]; !ok {
		return domain.User{}, domain.ErrUserNotFound
	}
	return s.List[id], nil
}

func (repo *userRepo) Update(id int, input domain.UpdateUserInput) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var s domain.UserStore
	if err := readJSONFromFile(repo.storeDir, &s); err != nil {
		return err
	}

	user, ok := s.List[id]
	if !ok {
		return domain.ErrUserNotFound
	}
	user.DisplayName = *input.DisplayName
	s.List[id] = user

	return writeJSONToFile(repo.storeDir, &s)
}

func (repo *userRepo) Delete(id int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var s domain.UserStore
	if err := readJSONFromFile(repo.storeDir, &s); err != nil {
		return err
	}

	if _, ok := s.List[id]; !ok {
		return domain.ErrUserNotFound
	}

	delete(s.List, id)

	return writeJSONToFile(repo.storeDir, &s)
}
