package main

import (
	"MY_GO_CODES/Grpc_Basic/ServerStream_Grpc/Greetpb"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Inside server main")

	//setup connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	//setting up grpc
	s := grpc.NewServer()

	//registering the grpc to implement service from greetpb.go
	Greetpb.RegisterGreetServiceServer(s, &server{})

	//sending the lis tcp connection to serve
	erv := s.Serve(lis)
	if erv != nil {
		log.Fatalln(erv)
	}
}

func (*server) Greet(ctx context.Context, req *Greetpb.GreetRequest) (*Greetpb.GreetResponse, error) {
	fmt.Printf("Greet function invoked %v", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	result := "Hello " + firstname + lastname
	res := &Greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *Greetpb.GreetManyTimesRequest, stream Greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManTimes has invoked")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number: " + strconv.Itoa(i)
		res := &Greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
