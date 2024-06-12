package main

import (
	"log"
	"time"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got requwst with names %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloRes{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
