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

func (s server) Calculator(ctx context.Context, request *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	res := &calculatorpb.CalculatorResponse{
		Sum: int64(request.Int1 + request.Int2),
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
