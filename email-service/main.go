package main

import (
	userPb "shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	"encoding/json"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)
	srv.Init()

	pubSub := srv.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}

	// 订阅消息
	_, err := pubSub.Subscribe(topic, func(pub broker.Publication) error {
		var user *userPb.User
		if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
			return err
		}
		log.Printf("[CREATE USER]: %v\n", user)
		go senEmail(user)
		return nil
	})

	if err != nil {
		log.Printf("sub error: %v\n", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

func senEmail(user *userPb.User) error {
	log.Printf("[SENDING A EMAIL TO %s...]", user.Name)
	return nil
}
