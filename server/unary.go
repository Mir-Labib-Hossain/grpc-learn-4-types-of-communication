package main

import (
	"context"

	pb "github.com/Mir-Labib-Hossain/grpc-4-types/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloRes, error) {
	return &pb.HelloRes{
		Message: "Hello LabibNoob",
	}, nil
}
