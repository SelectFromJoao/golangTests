package repository

import (
	"codingDojo/cmd/domain/entity"
	"errors"
)

// InMemoryUserRepository is an in-memory implementation of UserRepository.
type InMemoryUserRepository struct {
	users map[string]entity.User
}

// NewInMemoryUserRepository creates a new instance of InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]entity.User),
	}
}

// Read get a user in the repository.
func (repo *InMemoryUserRepository) Read(id string) (entity.User, error) {
	user := &entity.User{}
	if _, exists := repo.users[id]; !exists {
		return *user, errors.New("user not found")
	}

	return repo.users[id], nil
}

// Create a user in the repository.
func (repo *InMemoryUserRepository) Create(user entity.User) error {
	repo.users[user.ID] = user
	return nil
}

// Update updates an existing user in the repository.
func (repo *InMemoryUserRepository) Update(user entity.User) error {
	if _, exists := repo.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	repo.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository based on ID.
func (repo *InMemoryUserRepository) Delete(id string) error {
	if _, exists := repo.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(repo.users, id)
	return nil
}
