package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "path_to_generated_code/greet"

	"google.golang.org/grpc"
)

// реализация сервиса GreetService
type server struct {
	pb.UnimplementedGreetServiceServer
}

// реализация  метода SayHello
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}

func main() {
	// настройка сервера
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
