package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Server streamiing started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while server streaming %v", err)
		}
		log.Println(message)

	}
	log.Printf("Server stream finished")
}
