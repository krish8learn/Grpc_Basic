package main

import (
	"MY_GO_CODES/Grpc_Basic/Unary_Grpc/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Client running")

	//trying to connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Problem occured while connecting to the server")
	}
	//conn will end at the last
	defer conn.Close()
	//to serve the client
	c := proto.NewSquarerootClient(conn)

	//he we request server to do the opeartion
	dosquareroot(c)
}
func dosquareroot(c proto.SquarerootClient) {
	//num := 4
	var num int
	fmt.Println("Inside dosquareroot, Enter the Input")
	fmt.Scanf("%d", &num)
	req := &proto.Request{
		Input: int32(num),
	}
	res, err := c.Squareroot(context.Background(), req)
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			//error from GRPC
			fmt.Println("Error from GRPC")
			fmt.Println(resErr.Message())
			//fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("Hey -ve Input")
			}
		} else {
			log.Fatalln("Error occured while performing sqaureroot", err)
		}
	}
	fmt.Println("The square root is :", res.GetOutput())
}
