package main

import (
	"net"

	pb "hello-world"

	"google.golang.org/grpc"
)

type server struct {
	pb.Unimplemented
}

//func (s *server)

func main() {
	plistener, err := net.Listen("tcp:2408")
	if err != nil {
		panic("Failed to bind to port: %v", err)
	}

	//Creating a New gRPC server
	gServ := grpc.NewServer()
	//Binding the stub function with the func we created
	pb.RegisterMyFunc(gServ, &server{})
	if err := gServ.Serve(plistener); err != nil {
		panic("Unable to start gRPC server: %v", err)
	}
}
