package main

import (
	"context"
	"fmt"
	"github.com/lrpinto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
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
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC")

	req := &calculatorpb.CalculatorRequest{
		Int1: 3,
		Int2: 10,
	}

	ctx := context.Background()
	res, err := c.Calculator(ctx, req)
	if err != nil {
		log.Fatalf("Failed to request #{err}")
	}

	log.Printf("%d + %d = %d", req.Int1, req.Int2, res.Sum)
}
