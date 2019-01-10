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
func (s *Server) CDGame(ctx context.Context, in *pb.CDGameRequest) (res *pb.CDGameResponse, err error) {
	if !in.Delete {
		res, err = manegeGame.CurrentGames.AddNewGame(in)
	} else {
		res, err = manegeGame.CurrentGames.DeleteGame(in.GameID)
	}
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		return res, err
	}
}

// Played ...
func (s *Server) Played(ctx context.Context, in *pb.StonePlayed) (res *pb.StonePlayed, err error) {
	res, err = manegeGame.CurrentGames.PlayedIA(in, false)
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		return res, err
	}
}

// Played ...
func (s *Server) PlayedHelp(ctx context.Context, in *pb.StonePlayed) (res *pb.StonePlayed, err error) {
	res, err = manegeGame.CurrentGames.PlayedIA(in, true)
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		return res, err
	}
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
