package main

import (
	"context"
	"log"
	"time"

	pb "github.com/prabhatpankaj/GO-grpc-demo/proto"
)

// callSayHello will need to be called from main.go to make the call to the server
func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
