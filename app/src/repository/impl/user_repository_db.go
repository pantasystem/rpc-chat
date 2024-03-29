package impl

import (
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
)

type UserRepositoryDB struct {
	Repository *Repository
}

func (self *UserRepositoryDB) Find(userId int64) (*entity.User, error) {
	var user entity.User
	if r := self.Repository.DB.First(&user, userId); r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

func (self *UserRepositoryDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if r := self.Repository.DB.Where("email = ?", email).First(&user); r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

func (self *UserRepositoryDB) Create(user entity.User) (*entity.User, error) {
	if r := self.Repository.DB.Create(&user); r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

func (self *UserRepositoryDB) Update(user entity.User) (*entity.User, error) {
	if r := self.Repository.DB.Save(&user); r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

func (self *UserRepositoryDB) FindByToken(token string) (*entity.User, error) {
	var user entity.User
	if result := self.Repository.DB.Where("token = ?", token).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
