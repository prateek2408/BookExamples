package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/prateek2408/BookExamples/gRPC/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyFunc(ctx context.Context, input *pb.Request) (*pb.StringMessage, error) {
	fmt.Print("Inside The actual caller rpc FUnction")
	return &pb.StringMessage{Reply: "Hey There"}, nil
}

func main() {
	plistener, err := net.Listen("tcp", ":2408")
	if err != nil {
		panic("Failed to bind to port")
	}

	//Creating a New gRPC server
	gServ := grpc.NewServer()
	//Binding the stub function with the func we created
	pb.RegisterMyServiceServer(gServ, &server{})
	if err := gServ.Serve(plistener); err != nil {
		panic("Unable to start gRPC server")
	}
}
