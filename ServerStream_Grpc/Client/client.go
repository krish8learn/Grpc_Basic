package main

import (
	"MY_GO_CODES/Grpc_Basic/ServerStream_Grpc/Greetpb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello, client here")
	//creating a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect", err)
	}
	defer conn.Close()
	c := Greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("created client %f\n", c)

	//dounary(c)
	doserverstreaming(c)
}

func dounary(c Greetpb.GreetServiceClient) {
	fmt.Println("IN UNARY")
	req := &Greetpb.GreetRequest{
		Greeting: &Greetpb.Greeting{
			FirstName: "krish",
			LastName:  "karmakar",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Println("ERROR OCURRED in Greet function", err)
	}
	log.Printf("RESPONSE FROM GREET: %v", res)

}

func doserverstreaming(c Greetpb.GreetServiceClient) {
	fmt.Println("server streaming here")
	req := &Greetpb.GreetManyTimesRequest{
		Greeting: &Greetpb.Greeting{
			FirstName: "Toni",
			LastName:  "Kroos",
		},
	}
	resstream, erb := c.GreetManyTimes(context.Background(), req)
	if erb != nil {
		log.Fatalln("erb error", erb)
	}
	for {
		msg, err := resstream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv error", err)
		}
		log.Printf("%v", msg.GetResult())
	}

}
