package entity

type Post struct {
	Id int64	`gorm:"primaryKey;"`
	Text string
	UserId int64
	User User	`gorm:"foreignKey:UserId;references:Id;constant:OnUpdateOnUpdate:CASCADE, OnDelete:CASCADE;"`
}