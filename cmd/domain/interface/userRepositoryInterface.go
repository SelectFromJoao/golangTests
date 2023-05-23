package _interface

import (
	"codingDojo/cmd/domain/entity"
)

type UserRepositoryInterface interface {
	Create(user entity.User) error
	Read(id string) (entity.User, error)
	Update(user entity.User) error
	Delete(id string) error
}
