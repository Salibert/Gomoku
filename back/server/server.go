package server

import (
	"context"
	"fmt"

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

// Server ...
type Server struct{}

func init() {
	pb.RegisterGameServer(GrpcServer, &Server{})
	reflection.Register(GrpcServer)
}

// InitGame ...
func (s *Server) InitGame(context context.Context, in *pb.InitGameRequest) (*pb.InitGameResponse, error) {
	fmt.Println("SALUT TOI")
	fmt.Println(in)
	return &pb.InitGameResponse{Message: "OKOK frere"}, nil
}
