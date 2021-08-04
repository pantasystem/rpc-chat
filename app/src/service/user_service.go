package service

import (
	pb "com.github.Kinoshita0623/rpc-chat/app/src/service/pb"
	"context"
	"com.github/Kinoshita0623/rpc-chat/app/src/core"
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
	"fmt"
)

type UserService struct {
	Core core.Core
}

func (self *UserService) Register(ctx context.Context, req *pb.RegisterUserRequest) (*pb.TokenResponse, error) {
	u := entity.User{
		Email: req.Email,
		Name: req.Name,
	}
	u.GenerateToken()
	u.ApplyPassword(req.Password)
	r, err := self.Core.Repository.UserRepository().Create(u)
	if err != nil {
		return nil, err
	}

	return &pb.TokenResponse {
		Token: r.Token,
	}, nil
}

func (self *UserService) Login(ctx context.Context,req  *pb.RegisterUserRequest) (*pb.TokenResponse, error) {
	u, err := self.Core.Repository().UserRepository().FindByEmail(pb.Email)
	if err != nil {
		return nil, err
	}
	if u.CheckRawPassword(req.Password) {
		return &pb.TokenResponse {
			Token: u.Token,
		}, nil
	}else{
		return nil, fmt.Errorf("password missmatch")
	}
}

func (self *UserService) Find(ctx context.Context, req *pb.FindUser) (*pb.User, error) {
	u, err := self.Core.Repository().UserRepository().Find(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User {
		Email: u.Email,
		Name: u.Token,
		Id: u.Token,
	}, nil
}