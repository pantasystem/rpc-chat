package entity

type User struct {
	Id                int64
	Name              string
	AvatarUrl         string
	Token             string `json:"-"`
	Email             string `json:"-"`
	EncryptedPassword string `json:"-"`
}
