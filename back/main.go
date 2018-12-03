package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Salibert/Gomoku/back/server"
)

func main() {
	listen, err := net.Listen("tcp", server.GrpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	fmt.Println("Server listening Port:", server.GrpcPort)
	if err := server.GrpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
