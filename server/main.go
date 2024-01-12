package main

import (
	"log"
	"net"

	pb "github.com/sachinnagesh/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(gRPCServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
