package main

import (
	"context"
	"fmt"
	"github.com/lrpinto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Starting Calculator Client")

	// connect to server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect #{err}")
	}

	defer cc.Close()

	// create calculator client
	c := calculatorpb.NewCalculatorServiceClient(cc)

	// Do Unary RPC
	doUnary(c)

	// Do Server Streaming RPC
	doServerStreaming(c)

	// Do Client Streaming RPC
	doClientStreaming(c)
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do ComputeAverage Client Streaming RPC...\n")

	ctx := context.Background()

	requests := []*calculatorpb.ComputeAverageRequest{
		{
			Parcel: 1,
		},
		{
			Parcel: 2,
		},
		{
			Parcel: 3,
		},
		{
			Parcel: 4,
		},
	}

	averageClient, err := c.ComputeAverage(ctx)
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage\n")
	} else {
		for _, req := range requests {
			err := averageClient.Send(req)
			if err != nil {
				log.Fatalf("Error while sending request: %v\n", req)
			} else {
				log.Printf("Sent request: %v\n", req)
			}
		}
		res, err := averageClient.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error while receiving response: %v", res)
		} else {
			log.Printf("Average: %v", res.GetAverage())
		}
	}
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do PrimeNumberComposition Server Streaming RPC")

	ctx := context.Background()

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 120,
	}

	decompositionClient, err := c.PrimeNumberDecomposition(ctx, req)
	if err != nil {
		log.Printf("Failed to do PrimeNumberDecomposition: %v", err)
	} else {
		for {
			res, err := decompositionClient.Recv()
			if err == io.EOF {
				log.Println("No more primes.")
				break
			} else if err != nil {
				log.Printf("Failed to Receive message: %v", err)
			} else {
				log.Println(res.GetResult())
			}
		}
	}

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Sum Unary RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	ctx := context.Background()
	res, err := c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("Failed to request #{err}")
	}

	log.Printf("%d + %d = %d", req.FirstNumber, req.SecondNumber, res.SumResult)
}
