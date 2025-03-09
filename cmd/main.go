package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/equaerdist/simple-api/internal/domain/consts"
	"github.com/equaerdist/simple-api/internal/handlers"
	"github.com/equaerdist/simple-api/internal/infrastructure/database/connections"
	"github.com/equaerdist/simple-api/pkg/github.com/equaerdist/simple-api/pkg/car"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	port := os.Getenv(consts.PORT)
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool := connections.GetDBPool()
	defer pool.Close()

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	car.RegisterCarServiceServer(grpcServer, &handlers.CarGrpcService{})

	fmt.Println("Server is running on port " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
