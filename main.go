package main

import "github.com/keivanipchihagh/hello-world-go/internal"

func main() {
	internal.Start()
}

// package main

// import (
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"

// 	server "github.com/keivanipchihagh/hello-world-go/api/grpc"
// 	pb "github.com/keivanipchihagh/hello-world-go/api/grpc/proto/golang"
// )

// func main() {
// 	// Listen on TCP port 50051.
// 	lis, err := net.Listen("tcp", "localhost:8080")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	// Create a new gRPC server.
// 	grpcServer := grpc.NewServer()

// 	// Register the TaskService with the gRPC server.
// 	pb.RegisterTaskServiceServer(grpcServer, &server.Server{})

// 	log.Println("gRPC server listening on :50051")
// 	// Start serving requests.
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }
