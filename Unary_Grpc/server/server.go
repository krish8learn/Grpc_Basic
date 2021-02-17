package main

import (
	"MY_GO_CODES/Grpc_Basic/Unary_Grpc/proto"
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func main() {
	fmt.Println("Server started")

	//setting up tcp connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("Error occured during setting tcp connection", err)
	}

	//setting the connection as GRPC
	grps := grpc.NewServer()
	proto.RegisterSquarerootServer(grps, &server{})

	// the connection (lis) will serve
	go func() {
		fmt.Println("started inidvidual go routine for server")
		erv := grps.Serve(lis)
		if erv != nil {
			log.Fatalln("Error occured while serving", erv)
		}
	}()

	//wait for channel c to exist
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	//block untill the signal is received
	<-channel
	fmt.Println("Stoping server")
	grps.Stop()
	fmt.Println("Closing the server")
	lis.Close()
	fmt.Println("End of service")

}

func (*server) Squareroot(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	fmt.Println("Performing Squareroot")
	number := req.GetInput()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintln("Received -ve Input", number),
		)
	}
	result := math.Sqrt(float64(number))
	res := &proto.Response{
		Output: int32(result),
	}
	return res, nil
}
