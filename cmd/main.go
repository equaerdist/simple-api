package main

import (
	"fmt"
	"log"
	"net"

	"github.com/equaerdist/simple-api/internal/handlers"
	"github.com/equaerdist/simple-api/internal/infrastructure/database/connections"
	"github.com/equaerdist/simple-api/pkg/github.com/equaerdist/simple-api/pkg/car"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool := connections.GetDBPool()
	defer pool.Close()

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	car.RegisterCarServiceServer(grpcServer, &handlers.CarGrpcService{})

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
