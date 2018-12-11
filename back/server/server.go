package server

import (
	"context"
	"fmt"

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
		return manegeGame.CurrentGames.AddNewGame(in)
	}
	return manegeGame.CurrentGames.DeleteGame(in.GameID)
}

// Played ...
func (s *Server) Played(context context.Context, in *pb.StonePlayed) (*pb.StonePlayed, error) {
	defer fmt.Println("END PLAYED")
	fmt.Println("START PLAYED")
	return manegeGame.CurrentGames.PlayedIA(in)
}

// CheckRules ...
func (s *Server) CheckRules(context context.Context, in *pb.StonePlayed) (*pb.CheckRulesResponse, error) {
	return manegeGame.CurrentGames.ProccessRules(in)
}
