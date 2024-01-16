package main

import (
	"log"

	pb "github.com/sachinnagesh/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connnect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)

	name := &pb.NameList{
		Names: []string{"Scorpio", "King", "Yo"},
	}
	//callSayHelloServerStream(client, name)

	//callSayHelloClientStream(client, name)
	callBidirectionalStream(client, name)

}
