package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/learn_go_grpc/hello/proto"
)

var addr = "0.0.0.0:5000"

type Server struct {
	pb.HelloServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}
	defer lis.Close()
	log.Println("listening on", addr)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}
