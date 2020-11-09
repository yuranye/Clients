// Package main implements a server for client service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	empty "github.com/golang/protobuf/ptypes/empty"

	pb "github.com/showcase/clients/build/gen/clients"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement client service
type server struct {
	pb.UnimplementedClientServiceServer
}

func (s *server) CreateClient(ctx context.Context, in *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	log.Printf("Received create for client uuid: %v", in.Client.ClientUuid)
	return &pb.CreateClientResponse{
		Client: &pb.Client{
			ClientUuid: uuid.New().String(),
			FirstName:  in.Client.FirstName,
			SecondName: in.Client.SecondName,
			Balance:    in.Client.Balance,
		},
	}, nil
}

func (s *server) UpdateClient(ctx context.Context, in *pb.UpdateClientRequest) (*pb.UpdateClientResponse, error) {
	log.Printf("Received update for client uuid: %v", in.Client.ClientUuid)
	return &pb.UpdateClientResponse{
		Client: &pb.Client{
			ClientUuid: in.Client.ClientUuid,
			FirstName:  in.Client.FirstName,
			SecondName: in.Client.SecondName,
			Balance:    in.Client.Balance,
		},
	}, nil
}

func (s *server) GetClient(ctx context.Context, in *pb.GetClientRequest) (*pb.GetClientResponse, error) {
	log.Printf("Received get for client uuid: %v", in.ClientUuid)
	return &pb.GetClientResponse{
		Client: &pb.Client{
			ClientUuid: uuid.New().String(),
			FirstName:  "SomeFirstName",
			SecondName: "SomeSecondName",
			Balance:    0,
		},
	}, nil
}

func (s *server) ListClients(ctx context.Context, in *pb.ListClientsRequest) (*pb.ListClientsResponse, error) {
	for _, ch := range in.ClientUuids {
		log.Printf("Received list for client uuid: %v", ch)
	}

	return &pb.ListClientsResponse{
		Clients: []*pb.Client{
			{
				ClientUuid: uuid.New().String(),
				FirstName:  "SomeFirstName",
				SecondName: "SomeSecondName",
				Balance:    0,
			},
		},
	}, nil
}

func (s *server) DeleteClient (ctx context.Context, in *pb.DeleteClientRequest) (*empty.Empty, error) {
	log.Printf("Received delete for client uuid: %v", in.ClientUuid)
	//simply return empty
	return &empty.Empty{}, nil
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
