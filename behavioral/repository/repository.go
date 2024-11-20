package repository

import "errors"

// UserRepository определяет интерфейс для доступа к пользователям.
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
	Delete(id int) error
}

// InMemoryUserRepository реализует UserRepository в памяти.
type InMemoryUserRepository struct {
	users  map[int]*User
	nextID int
}

// NewInMemoryUserRepository создает новый экземпляр InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// FindByID ищет пользователя по ID.
func (repo *InMemoryUserRepository) FindByID(id int) (*User, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Save сохраняет пользователя.
func (repo *InMemoryUserRepository) Save(user *User) error {
	if user.ID == 0 {
		user.ID = repo.nextID
		repo.nextID++
	}
	repo.users[user.ID] = user
	return nil
}

// Delete удаляет пользователя по ID.
func (repo *InMemoryUserRepository) Delete(id int) error {
	if _, exists := repo.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(repo.users, id)
	return nil
}
