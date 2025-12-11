package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/aishsanal/grpc/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot start client due to error %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	callSayHello(client)
}
