package grpc

import (
	"context"
	"log"

	pb "github.com/keivanipchihagh/hello-world-go/api/grpc/proto/golang"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *Server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	log.Println("ListTasks called")
	// For demonstration, return the tasks from the request.
	return &pb.ListTasksResponse{Tasks: req.Tasks}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	log.Printf("GetTask called for id: %s", req.Id)
	// Create a dummy task.
	task := &pb.Task{
		Id:     req.Id,
		Title:  "Dummy Task",
		Author: "Author Name",
	}
	return &pb.GetTaskResponse{Task: task}, nil
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	log.Println("CreateTask called")
	// Return the created task.
	return &pb.CreateTaskResponse{Task: req.Task}, nil
}

func (s *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	log.Println("UpdateTask called")
	return &pb.UpdateTaskResponse{Task: req.Task}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	log.Printf("DeleteTask called for id: %s", req.Id)
	task := &pb.Task{
		Id:    req.Id,
		Title: "Deleted Task",
	}
	return &pb.DeleteTaskResponse{Task: task}, nil
}
