package main

import (
	"log"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client didn't connect %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{Names: []string{"Labib", "Fayek", "Ayon"}}
	// callSayHello(client)
	// callSayHelloServerStreaming(client,names)
	// callSayHelloClientStreaming(client,names)
	callSayHelloBiDirectionalStreaming(client, names)
}
