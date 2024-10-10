package main

import (
	pb "github.com/learn_go_grpc/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Println("Sum function was invoked with: ", in)

	no := in.Number
	div := int64(2)

	for no > 1 {
		if no%div == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: div,
			})
			no /= div
		} else {
			div++
		}
	}

	return nil
}
