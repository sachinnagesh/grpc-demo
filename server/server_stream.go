package main

import (
	"log"
	"time"

	pb "github.com/sachinnagesh/grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
