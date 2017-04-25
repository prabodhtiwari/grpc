package main

import (
	"log"
	"net"

	"dev/grpc/info"
	pb "dev/grpc/protobuf"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterInfoServer(s, &info.Server{})
	log.Printf("Zen listening on %v\n", lis.Addr())
	s.Serve(lis)
}
