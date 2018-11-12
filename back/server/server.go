package server

import (
	"context"

	pb "./pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	//GrpcPort ...
	GrpcPort = ":50051"
)

var (
	// GrpcServer ...
	GrpcServer = grpc.NewServer()
)

func init() {
	reflection.Register(GrpcServer)
}

func newServer(srv pb.GameServer) {
	pb.RegisterGameServer(GrpcServer, srv)
}

// Server ...
type Server struct{}

// InitGame ...
func (s *Server) InitGame(context context.Context, in *pb.InitGameRequest) (*pb.InitGameResponse, error) {
	return &pb.InitGameResponse{Message: "OKOK frere"}, nil
}
