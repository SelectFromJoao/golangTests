package main

import (
	"codingDojo/cmd/domain/entity"
	"codingDojo/cmd/domain/repository"
	"codingDojo/cmd/domain/services"
	"testing"
)

// TestUserCreate is a test suite for create a user.
func TestUserCreate(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := services.NewUserService(repo)

	// Create and retrieve a user
	user := entity.User{ID: "1", Name: "John Doe"}
	err := service.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
}

// TestUserCreateAndRead is a test suite for create a user and get.
func TestUserCreateAndRead(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := services.NewUserService(repo)

	user := entity.User{ID: "1", Name: "John Doe"}
	err := service.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	retrievedUser, err := service.GetUser("1")
	if err != nil {
		t.Errorf("Failed to retrieve user: %v", err)
	} else if retrievedUser.Name != user.Name {
		t.Errorf("Retrieved user name doesn't match: expected %s, got %s", user.Name, retrievedUser.Name)
	}
}

// TestUserCreateAndUpdate is a test suite for create a user and update.
func TestUserCreateAndUpdate(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := services.NewUserService(repo)

	user := entity.User{ID: "1", Name: "John Doe"}
	err := service.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	// Update and verify the updated user
	updatedUser := entity.User{ID: "1", Name: "Jane Doe"}
	err = service.UpdateUser(updatedUser)
	if err != nil {
		t.Errorf("Failed to update user: %v", err)
	}
}

// TestUserDeleteAndGet is a test suite for delete a user and try get it.
func TestUserDeleteAndGet(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := services.NewUserService(repo)

	user := entity.User{ID: "1", Name: "John Doe"}
	err := service.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	// Delete the user and verify it's deleted
	err = service.DeleteUser("1")
	if err != nil {
		t.Errorf("Failed to delete user: %v", err)
	}

	_, err = service.GetUser("1")
	if err == nil {
		t.Error("Retrieved user after deletion")
	}
}
