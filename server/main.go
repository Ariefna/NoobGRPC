package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/yourusername/NoobGRPC/proto" // Replace with the actual import path
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sapa(ctx context.Context, req *proto.SalamRequest) (*proto.SalamResponse, error) {
	pesan := fmt.Sprintf("Halo, %s!", req.Nama)
	return &proto.SalamResponse{Pesan: pesan}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	proto.RegisterSalamServiceServer(grpcServer, &server{})

	fmt.Println("Server is running on port 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
