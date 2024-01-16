package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sachinnagesh/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, nameList *pb.NameList) {
	log.Printf("Streaming has started")

	stream, err := client.SayHelloServerStreaming(context.Background(), nameList)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished !!!")

}
