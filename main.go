package main

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"com.github/Kinoshita0623/rpc-chat/app/src/entity"
	"com.github/Kinoshita0623/rpc-chat/app/src/core"
	"com.github/Kinoshita0623/rpc-chat/app/src/repository/impl"
	"com.github/Kinoshita0623/rpc-chat/app/src/service"
	pb "com.github/Kinoshita0623/rpc-chat/app/src/service/pb"
	"net"
	"os"
	"google.golang.org/grpc"


)

func main() {
	fmt.Printf("Hello go")
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "test", "password", "db:3306", "database")

	db, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Post{})

	core := core.Core{
		Repository: &impl.Repository {
			DB: db,
		}
	}

	listen, err := net.Listen("tcp", ":19004")
	if err != nil {
		os.Exit(1)
	}

	server := grpc.NewServer()
	userService := &service.UserService{
		Core: core,
	}


	pb.RegisterUserServiceServer(server, userService)
	server.Serve(listen)

	
}