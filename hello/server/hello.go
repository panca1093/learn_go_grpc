package main

import (
	"context"
	pb "github.com/learn_go_grpc/hello/proto"
	"log"
)

func (s *Server) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRes, error) {
	log.Println("Hello function was invoked: ", in)

	return &pb.HelloRes{
		Result: "Hello " + in.Name,
	}, nil
}
