package main

import (
	"context"
	pb "github.com/learn_go_grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked with: ", in)

	return &pb.SumResponse{
		Result: in.Number + in.Number_1,
	}, nil
}
