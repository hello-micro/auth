package main

import (
	"github.com/hello-micro/auth/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	auth "github.com/hello-micro/auth/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("hello.micro.srv.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Register Struct as Subscriber
	//_ = micro.RegisterSubscriber("hello.micro.srv.auth", service.Server(), new(subscriber.Auth))

	// Register Function as Subscriber
	//_ = micro.RegisterSubscriber("hello.micro.srv.auth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
