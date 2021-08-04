package impl

import (
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
)

type PostRepositoryDB struct {
	Repository *Repository
}

func (self *PostRepositoryDB) Create(post *entity.Post) (*entity.Post, error) {
	if result := self.Repository.DB.Create(post); result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (self *PostRepositoryDB) FindAll(untilId *int64) ([]*entity.Post, error) {
	var posts []*entity.Post
	if result := self.Repository.DB.Where("id < ?", untilId).Order("id desc").Limit(100).Find(&posts); result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
