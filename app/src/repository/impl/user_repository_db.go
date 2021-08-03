package impl

import (
	"fmt"
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
)

type UserRepositoryDB struct {
	Repository *Repository
}

func (self *UserRepositoryDB) Find(userId int64) (*entity.User, error) {
	return nil, fmt.Errorf("not impl")
}

func (self *UserRepositoryDB) FindByEmail(email string) (*entity.User, error) {
	return nil, fmt.Errorf("not impl")
}

func (self *UserRepositoryDB) Create(user entity.User) (*entity.User, error) {
	return nil, fmt.Errorf("not impl")
}