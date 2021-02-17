package main

import (
	"MY_GO_CODES/Grpc_Basic/Grpc_with_Employee/Emppb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client side active")
	//creating a connection
	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect", err)
	}
	defer conn.Close()
	c := Emppb.NewEmpServiceClient(conn)

	unaryEmp(c)
}
func unaryEmp(c Emppb.EmpServiceClient) {
	fmt.Println("Inside UnaryEmp")

	req := &Emppb.EmpRequest{
		Input: &Emppb.Employee{
			EmployeeId:     124,
			EmployeeName:   "Krishnendu Karmakar",
			EmployeeMail:   "krishkarmakar17@gmail.com",
			EmployeeMobile: "8240903775",
		},
	}

	//sending the request to the server
	res, err := c.Emp(context.Background(), req)
	if err != nil {
		log.Println("Error Occured", err)
	}

	//fmt.Println(res)
	fmt.Println("The ID: ", res.GetOutput().GetEmployeeId())
	fmt.Println("The Name:", res.GetOutput().GetEmployeeName())
	fmt.Println("The MailId:", res.GetOutput().GetEmployeeMail())
	fmt.Println("The Mobile No:", res.GetOutput().GetEmployeeMobile())
}
