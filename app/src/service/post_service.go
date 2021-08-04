package service

import (
	"context"

	"com.github/Kinoshita0623/rpc-chat/app/src/core"
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
	pb "com.github/Kinoshita0623/rpc-chat/app/src/service/pb"
)

type PostService struct {
	Core *core.Core
}

func (self *PostService) Create(ctx context.Context, req *pb.CreatePost) (*pb.Post, error) {
	u, e := self.Core.Repository.UserRepository().FindByToken(req.Token)
	if e != nil {
		return nil, e
	}
	post := &entity.Post{
		UserId: u.Id,
		Text:   req.Text,
	}

	post, e = self.Core.Repository.PostRepository().Create(post)
	if e != nil {
		return nil, e
	}
	return &pb.Post{
		Text:   post.Text,
		UserId: post.UserId,
		User: &pb.User{
			Id:        u.Id,
			Name:      u.Name,
			AvatarUrl: u.AvatarUrl,
		},
	}, nil

}

func (self *PostService) FindAll(ctx context.Context, req *pb.RequestPosts) (*pb.Posts, error) {
	_, e := self.Core.Repository.UserRepository().FindByToken(req.Token)
	if e != nil {
		return nil, e
	}
	var untilId *int64
	if req.UntilId == -1 {
		untilId = nil
	} else {
		untilId = &req.UntilId
	}
	posts, e := self.Core.Repository.PostRepository().FindAll(untilId)

	protoPosts := make([]*pb.Post, len(posts))
	for i, v := range posts {
		protoPosts[i] = &pb.Post{
			Id:     v.Id,
			Text:   v.Text,
			UserId: v.UserId,
			User: &pb.User{
				Id:        v.User.Id,
				AvatarUrl: v.User.AvatarUrl,
				Name:      v.User.Name,
			},
		}
	}

	return &pb.Posts{
		Posts: protoPosts,
	}, nil

}
