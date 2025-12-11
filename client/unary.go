package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aishsanal/grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error while greeting: %v", err)
	}
	log.Printf("%s", response.Message)

}