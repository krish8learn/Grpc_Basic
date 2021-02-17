package main

import (
	"MY_GO_CODES/Grpc_Basic/Grpc_with_Employee/Emppb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Inside server main")

	//setting up TCP connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	//setting up grpc server
	s := grpc.NewServer()

	//registering the grpc server to implement service from Emp.pb.go
	Emppb.RegisterEmpServiceServer(s, &server{})

	//sending the lis connection to serve
	erv := s.Serve(lis)
	if erv != nil {
		log.Fatalln(erv)
	}
}

func (*server) Emp(ctx context.Context, req *Emppb.EmpRequest) (*Emppb.EmpResponse, error) {
	fmt.Println("Emp function inveoked", req)
	empid := req.GetInput().GetEmployeeId()
	empname := req.GetInput().GetEmployeeName()
	empmail := req.GetInput().GetEmployeeMail()
	empmobile := req.GetInput().GetEmployeeMobile()

	res := &Emppb.EmpResponse{
		Output: &Emppb.Employee{
			EmployeeId:     empid,
			EmployeeName:   empname,
			EmployeeMail:   empmail,
			EmployeeMobile: empmobile,
		},
	}
	return res, nil
}
