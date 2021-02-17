package main

import (
	"MY_GO_CODES/Grpc_Basic/Client_streaming/protopb"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Running server ")

	//setup connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	//setting up grpc
	s := grpc.NewServer()

	protopb.RegisterAverageServer(s, &server{})

	//sending the lis tcp connection to serve
	erv := s.Serve(lis)
	if erv != nil {
		log.Fatalln(erv)
	}
}
func (*server) Average(stream protopb.Average_AverageServer) error {
	fmt.Println("Average invoked")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protopb.Response{
				Avg: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalln(err)
		}
		sum += req.GetNumber()
		count++
	}
}
