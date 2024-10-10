package main

import (
	pb "github.com/learn_go_grpc/hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "0.0.0.0:5050"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("connection failed: ", err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	doHello(c)
}
