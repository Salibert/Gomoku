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

// Server ...
type Server struct{}

func init() {
	pb.RegisterGameServer(GrpcServer, &Server{})
	reflection.Register(GrpcServer)
}

var test int32

// InitGame ...
func (s *Server) InitGame(context context.Context, in *pb.InitGameRequest) (*pb.InitGameResponse, error) {
	return &pb.InitGameResponse{Message: "OKOK frere"}, nil
}

// Played ...
func (s *Server) Played(context context.Context, in *pb.StonePlayed) (*pb.StonePlayed, error) {
	return &pb.StonePlayed{CurrentPlayerMove: &pb.Node{X: in.CurrentPlayerMove.X, Y: in.CurrentPlayerMove.Y + 1, Player: 2}, Message: "OKOK frere"}, nil
}
