package main

import (
	"log"
	pb "shippy/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	// 连接到数据库
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	repo := &UserRepository{db}

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	s.Init()

	// 获取 broker 实例
	pubSub := s.Server().Options().Broker
	t := TokenService{repo}
	pb.RegisterUserServiceHandler(s.Server(), &handler{repo, &t, pubSub})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}
