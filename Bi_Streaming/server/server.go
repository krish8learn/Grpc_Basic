package main

import (
	"MY_GO_CODES/Grpc_Basic/Bi_Streaming/protopb"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Server running")
	//setup connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	//setting up grpc
	s := grpc.NewServer()

	protopb.RegisterFindmaxServer(s, &server{})

	//sending the lis tcp connection to serve
	erv := s.Serve(lis)
	if erv != nil {
		log.Fatalln(erv)
	}
}
func (*server) Findmax(stream protopb.Findmax_FindmaxServer) error {
	fmt.Println("Findmax invoked")
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		num := req.GetInput()
		if num > max {
			max = num
			senderr := stream.Send(&protopb.Response{
				Output: max,
			})
			if senderr != nil {
				log.Fatalln("Error while sending")
				return senderr
			}
		} else if num < max {
			senderr := stream.Send(&protopb.Response{
				Output: max,
			})
			if senderr != nil {
				log.Fatalln("Error while sending")
				return senderr
			}
		}
		//max = num

	}
}
