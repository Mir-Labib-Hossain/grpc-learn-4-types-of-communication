package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloReq{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the req with name %v", name)
		time.Sleep(2 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatal("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
