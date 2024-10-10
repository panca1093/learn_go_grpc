package main

import (
	pb "github.com/learn_go_grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:5051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}
	defer lis.Close()
	log.Println("listening on", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}
