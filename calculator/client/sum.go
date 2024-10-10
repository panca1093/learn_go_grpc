package main

import (
	"context"
	pb "github.com/learn_go_grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("start sum")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number:   10,
		Number_1: 3,
	})
	if err != nil {
		log.Println("could not sum: ", err)
	}
	log.Println("Result: ", res.Result)
}
