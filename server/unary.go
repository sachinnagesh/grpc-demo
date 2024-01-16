package main

import (
	"context"

	pb "github.com/sachinnagesh/grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, re *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil

}
