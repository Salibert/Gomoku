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
func (s *Server) CDGame(context context.Context, in *pb.CDGameRequest) (res *pb.CDGameResponse, err error) {
	if !in.Delete {
		res, err = manegeGame.CurrentGames.AddNewGame(in)
	}
	res, err = manegeGame.CurrentGames.DeleteGame(in.GameID)
	return
}

// Played ...
func (s *Server) Played(context context.Context, in *pb.StonePlayed) (res *pb.StonePlayed, err error) {
	manegeGame.CurrentGames.PlayedIA(in)
	return
}

// CheckRules ...
func (s *Server) CheckRules(ctx context.Context, in *pb.StonePlayed) (*pb.CheckRulesResponse, error) {
	res, err := manegeGame.CurrentGames.ProccessRules(in)
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		return res, err
	}
}
