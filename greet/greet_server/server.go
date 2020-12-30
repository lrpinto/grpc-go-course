package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lrpinto/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GreetManyTimes(request *greetpb.GreetManyTimesRequest,
	timesServer greetpb.GreetService_GreetManyTimesServer) error {

	log.Printf("Invoked Function GreetManyTimes")

	firstName := request.Greeting.FirstName
	result := "Hello " + firstName

	res := &greetpb.GreetManyTimesResponse{
		Result: result,
	}

	for i := 0; i < 10; i++ {
		timesServer.Send(res)
		log.Printf("Sent Greeting Number: %v", i+1)
	}

	return nil
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Print("Hello World")

	// write a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
