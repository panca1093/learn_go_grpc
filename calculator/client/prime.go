package main

import (
	"context"
	pb "github.com/learn_go_grpc/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("Start Do Primes")

	req := &pb.PrimeRequest{
		Number: 120,
	}
	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Println("could not sum: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("End Do Primes")
			break
		}
		if err != nil {
			log.Fatalln("error while stream primes: ", err)
		}
		log.Println("Result: ", res.Result)
	}
}
