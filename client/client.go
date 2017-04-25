package main

import (
	"context"
	pb "dev/grpc/protobuf"
	"flag"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewInfoClient(conn)
	msg := &pb.PingRequest{}
	res, _ := client.Ping(context.Background(), msg)
	fmt.Println("Server said", res.Message)
}
