package repository

import (
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
)
type Post interface {

	Create(post *entity.Post) (*entity.Post, error)
	FindAll(sinceId int64) ([]*entity.Post, error)
}