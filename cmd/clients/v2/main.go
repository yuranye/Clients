// Package main implements a server v2 for client service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"

	pb "github.com/showcase/clients/build/gen/clients/v2"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement client service v2
type server struct {
	pb.UnimplementedClientServiceServer
}

func (s *server) CreateManyClients(ctx context.Context, in *pb.CreateBatchClientsRequest) (*pb.CreateBatchClientsResponse, error) {
	log.Printf("Received v2 batch client create: %v", in.Client.ClientUuid)
	log.Printf("Description: %v", in.Client.Description)
	return &pb.CreateBatchClientsResponse{
		Client: &pb.Client{
			ClientUuid: uuid.New().String(),
			Description: "Some description",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClientServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
