package services

import (
	"codingDojo/cmd/domain/entity"
	"codingDojo/cmd/domain/interface"
)

// UserService represents a services that interacts with UserRepository.
type UserService struct {
	repository _interface.UserRepositoryInterface
}

// NewUserService creates a new instance of UserService with the given UserRepository dependency.
func NewUserService(repository _interface.UserRepositoryInterface) *UserService {
	return &UserService{
		repository: repository,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user entity.User) error {
	return s.repository.Create(user)
}

// GetUser retrieves a user based on ID.
func (s *UserService) GetUser(id string) (entity.User, error) {
	return s.repository.Read(id)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(user entity.User) error {
	return s.repository.Update(user)
}

// DeleteUser deletes a user based on ID.
func (s *UserService) DeleteUser(id string) error {
	return s.repository.Delete(id)
}
