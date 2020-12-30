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

	doUnary(c)

	doServerStreaming(c)
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
