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
