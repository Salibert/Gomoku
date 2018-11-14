package server

import (
	"context"
	"fmt"
	"io"
	"log"

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

// PlayedAI ...
func (s *Server) PlayedAI(stream pb.Game_PlayedAIServer) error {
	log.Println("Started stream")
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Got " + req.Message)
		resp := pb.PlayedAIResponse{Message: "OKOK"}
		if err := stream.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new OKOK")
	}
}
