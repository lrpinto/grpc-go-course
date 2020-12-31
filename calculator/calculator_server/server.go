package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/lrpinto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) ComputeAverage(averageServer calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Printf("ComputeAverage Invoked...")

	var n int32
	var sum int64
	var avg float32
	for {
		req, err := averageServer.Recv()
		if err == io.EOF {
			// finished receiving parcels
			if n == 0 {
				avg = 0
			} else {
				avg = float32(sum) / float32(n)
			}
			res := &calculatorpb.ComputeAverageResponse{
				Average: avg,
			}
			return averageServer.SendAndClose(res)
		} else if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		} else {
			sum += int64(req.GetParcel())
			n++
		}
	}

	return nil
}

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
