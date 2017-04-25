package info

import (
	pb "dev/grpc/protobuf"

	"golang.org/x/net/context"
)

// Server implements the info service
type Server struct {
}

// Ping responses with pong
func (s *Server) Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	res := &pb.PingResponse{
		Message: "pong",
	}
	return res, nil
}
