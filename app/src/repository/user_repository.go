package repository

import (
	"entity"
)

type UserRepository interface {

	Find(userId int64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}