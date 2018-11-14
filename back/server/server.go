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

var test int32

// InitGame ...
func (s *Server) InitGame(context context.Context, in *pb.InitGameRequest) (*pb.InitGameResponse, error) {
	fmt.Println(test)
	test++
	// fmt.Println("SALUT TOI")
	// fmt.Println(in)
	return &pb.InitGameResponse{Message: "OKOK frere"}, nil
}

// Played ...
func (s *Server) Played(context context.Context, in *pb.StonePlayed) (*pb.StonePlayed, error) {
	fmt.Println(test)
	test++
	// fmt.Println("SALUT TOI")
	// fmt.Println(in)
	return &pb.StonePlayed{CurrentPlayerMove: &pb.Node{X: in.CurrentPlayerMove.X, Y: in.CurrentPlayerMove.Y + 1, Player: 2}, Message: "OKOK frere"}, nil
}

// Played ...
// func (s *Server) Played(stream pb.Game_PlayedServer) error {
// 	log.Println("Started stream")
// 	ctx := stream.Context()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 		}
// 		req, err := stream.Recv()
// 		log.Println("Received value")
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Println("Got " + req.Message)
// 		resp := pb.StonePlayed{Message: "OKOK"}
// 		if err := stream.Send(&resp); err != nil {
// 			log.Printf("send error %v", err)
// 		}
// 		log.Printf("send new OKOK")
// 	}
// }
