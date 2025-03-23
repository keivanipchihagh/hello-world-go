// package main

// import "github.com/keivanipchihagh/hello-world-go/internal"

// func main() {
// 	internal.Start()
// }

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/keivanipchihagh/hello-world-go/api/grpc/proto/golang"
)

// server implements the TaskServiceServer interface.
type server struct {
	pb.UnimplementedTaskServiceServer
}

// ListTasks returns the list of tasks.
// Here we simply echo back the tasks provided in the request.
func (s *server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	log.Println("ListTasks called")
	// For demonstration, return the tasks from the request.
	return &pb.ListTasksResponse{Tasks: req.Tasks}, nil
}

// GetTask returns a task for the provided id.
// In this dummy implementation, a new Task is created using the id.
func (s *server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	log.Printf("GetTask called for id: %s", req.Id)
	// Create a dummy task.
	task := &pb.Task{
		Id:     req.Id,
		Title:  "Dummy Task",
		Author: "Author Name",
	}
	return &pb.GetTaskResponse{Task: task}, nil
}

// CreateTask creates a new task.
// This dummy implementation simply returns the provided task.
func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	log.Println("CreateTask called")
	return &pb.CreateTaskResponse{Task: req.Task}, nil
}

// UpdateTask updates an existing task.
// Here, we simply return the updated task from the request.
func (s *server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	log.Println("UpdateTask called")
	return &pb.UpdateTaskResponse{Task: req.Task}, nil
}

// DeleteTask deletes a task by id.
// For this example, a dummy task with the given id is returned as confirmation.
func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	log.Printf("DeleteTask called for id: %s", req.Id)
	task := &pb.Task{
		Id:    req.Id,
		Title: "Deleted Task",
	}
	return &pb.DeleteTaskResponse{Task: task}, nil
}

func main() {
	// Listen on TCP port 50051.
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server.
	grpcServer := grpc.NewServer()

	// Register the TaskService with the gRPC server.
	pb.RegisterTaskServiceServer(grpcServer, &server{})

	log.Println("gRPC server listening on :50051")
	// Start serving requests.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
