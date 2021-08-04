package impl

import (
	"gorm.io/gorm"
	"com.github/Kinoshita0623/rpc-chat/app/src/repository"
)

type Repository struct {
	DB *gorm.DB
}

func (self *Repository) UserRepository() repository.UserRepository {
	return &UserRepositoryDB{self}
}

func (self *Repository) PostRepository() repository.PostRepository {
	return &PostRepositoryDB{self}
}