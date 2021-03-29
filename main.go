package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/hello-micro/auth/handler"

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
