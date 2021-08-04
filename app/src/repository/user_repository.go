package repository

import (
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
)

type UserRepository interface {

	Find(userId int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(user entity.User) (*entity.User, error)
	Update(user entity.User) (*user.User, error)
}