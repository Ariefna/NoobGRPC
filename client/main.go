package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yourusername/NoobGRPC/proto" // Replace with the actual import path
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewSalamServiceClient(conn)

	req := &proto.SalamRequest{Nama: "John"}
	res, err := client.Sapa(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Sapa: %v", err)
	}

	fmt.Println(res.Pesan)
}
