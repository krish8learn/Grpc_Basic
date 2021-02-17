package main

import (
	"MY_GO_CODES/Grpc_Basic/Client_streaming/protopb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect", err)
	}
	defer conn.Close()
	c := protopb.NewAverageClient(conn)
	//fmt.Printf("created client %f\n", c)

	doclientstream(c)
}

func doclientstream(c protopb.AverageClient) {
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	numbers := []int32{12, 32, 21}
	for _, num := range numbers {
		fmt.Println("Number streams: ", num)
		stream.Send(&protopb.Request{
			Number: num,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The Average: ", res.GetAvg())

}
