package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lrpinto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest,
	decompositionServer calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	log.Println("Invoked Function PrimeNumberDecomposition")

	number := request.Number

	res := &calculatorpb.PrimeNumberDecompositionResponse{
		Result: -1,
	}

	// Decompose number in prime factors
	k := int64(2)
	for number > 1 {
		if number%k == 0 {
			res.Result = k
			err := decompositionServer.Send(res)
			if err != nil {
				log.Fatalf("Failed to send response: %v", res)
			}
			number /= k
		} else {
			k++
		}
	}

	return nil
}

func (s server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	res := &calculatorpb.SumResponse{
		SumResult: int64(request.FirstNumber + request.SecondNumber),
	}

	return res, nil
}

func main() {
	fmt.Println("Starting Calculator Server.")

	// write a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen #{err}")
	}

	// write the grpc server
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// serve the listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve #{err}")
	}
}
