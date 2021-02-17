package main

import (
	"MY_GO_CODES/Grpc_Basic/Bi_Streaming/protopb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running")

	//connecting to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect", err)
	}
	defer conn.Close()
	c := protopb.NewFindmaxClient(conn)

	dobistream(c)
}
func dobistream(c protopb.FindmaxClient) {
	stream, err := c.Findmax(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	//creating channel
	waitc := make(chan struct{})

	//to send
	go func() {
		numbers := []int32{12, 2, 21}
		for _, num := range numbers {
			fmt.Println("sending number:", num)
			stream.Send(&protopb.Request{
				Input: num,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		close(waitc)
	}()
	//to receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
				break
			}
			clientmax := res.GetOutput()
			fmt.Println("The max value: ", clientmax)
		}
		close(waitc)
	}()
	<-waitc
}
