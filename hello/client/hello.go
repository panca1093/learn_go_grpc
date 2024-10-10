package main

import (
	"context"
	pb "github.com/learn_go_grpc/hello/proto"
	"log"
)

func doHello(c pb.HelloServiceClient) {
	log.Println("do hello was invoked")

	res, err := c.Hello(context.Background(), &pb.HelloReq{
		Name: "Isaac",
	})
	if err != nil {
		log.Println("could not greet: ", err)
	}
	log.Println(res.Result)
}
