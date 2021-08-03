package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int64
	Name              string
	AvatarUrl         string
	Token             string `json:"-"`
	Email             string `json:"-"`
	EncryptedPassword string `json:"-"`
}


func (self *User) ApplyPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	self.EncryptedPassword = hash
	return nil
}

func (self *User) CheckRawPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(self.EncryptedPassword), []byte(password)) == nil
}
