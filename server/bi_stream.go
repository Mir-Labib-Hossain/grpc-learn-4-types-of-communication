package main

import (
	"io"
	"log"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &pb.HelloRes{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

	}
}
