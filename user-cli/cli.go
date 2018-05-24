package main

import (
	"log"
	pb "shippy/user-service/proto/user"
	"golang.org/x/net/context"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.user"))
	service.Init()
	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", service.Client())

	r, err := client.Create(context.Background(), &pb.User{
		Name:     "233",
		Email:    "233",
		Password: "233",
		Company:  "233",
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
}
