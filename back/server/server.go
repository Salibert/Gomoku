package server

import (
	"context"

	"github.com/Salibert/Gomoku/back/manegeGame"
	pb "github.com/Salibert/Gomoku/back/server/pb"
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

// CDGame ...
func (s *Server) CDGame(context context.Context, in *pb.CDGameRequest) (*pb.CDGameResponse, error) {
	if !in.Delete {
		return manegeGame.CurrentGames.AddNewGame(in.GameID)
	}
	return manegeGame.CurrentGames.DeleteGame(in.GameID)
}

// Played ...
func (s *Server) Played(context context.Context, in *pb.StonePlayed) (*pb.StonePlayed, error) {
	return nil, manegeGame.CurrentGames.PlayedAI(context, in)
}
