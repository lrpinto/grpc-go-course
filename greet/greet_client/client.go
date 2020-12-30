package main

import (
	"context"
	"fmt"
	"github.com/lrpinto/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Print("Hello I am a client!\n")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// Do Unary gRPC
	doUnary(c)

	// Do Server Streaming gRPC
	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC...")

	ctx := context.Background()

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Luisa",
			LastName:  "Pinto",
		},
	}

	timesClient, err := c.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("Failed to GreetManyTimes %v:", err)
	}

	for {
		res, err := timesClient.Recv()
		if err == io.EOF {
			log.Println("No more greetings.")
			break
		} else if err != nil {
			log.Fatalf("Failed to receive greeting %v:", err)
		} else {
			log.Println("Response from GreetManyTimes: ", res.GetResult())
		}
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Luisa",
			LastName:  "Pinto",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
