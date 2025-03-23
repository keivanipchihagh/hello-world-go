package grpc_test

import (
	"context"
	"net"
	"testing"
	"time"

	server "github.com/keivanipchihagh/hello-world-go/api/grpc"
	pb "github.com/keivanipchihagh/hello-world-go/api/grpc/proto/golang"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

// startTestGRPCServer starts an in-process gRPC server for testing.
func startTestGRPCServer(t *testing.T) (addr string, cleanup func()) {
	t.Helper()

	// Listen on an ephemeral port on localhost.
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server.
	grpcServer := grpc.NewServer()
	// Register the TaskService server implementation.
	pb.RegisterTaskServiceServer(grpcServer, &server.Server{})

	// Start the server in a separate goroutine.
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// Log server error; note that the error may occur during shutdown.
			t.Logf("gRPC server error: %v", err)
		}
	}()

	// Return the server address and a cleanup function for shutting down the server.
	return lis.Addr().String(), func() {
		grpcServer.Stop()
		lis.Close()
	}
}

func TestGetTask(t *testing.T) {
	// Start an in-process gRPC server for testing.
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	// Dial the gRPC server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the TaskService.
	client := pb.NewTaskServiceClient(conn)

	// Create a context with timeout to avoid hanging tests.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Build the GetTask request with a test task ID.
	req := &pb.GetTaskRequest{Id: "123"}
	// Call the GetTask RPC.
	resp, err := client.GetTask(ctx, req)
	if err != nil {
		t.Fatalf("GetTask failed: %v", err)
	}

	// Verify that the Task in the response is not nil.
	if resp.Task == nil {
		t.Fatalf("expected non-nil Task, got nil")
	}
	// Verify that the returned Task has the expected ID.
	if resp.Task.Id != req.Id {
		t.Errorf("expected task ID '%s', got %s", req.Id, resp.Task.Id)
	}
}

func TestCreateTask(t *testing.T) {
	// Start an in-process gRPC server for testing.
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	// Dial the gRPC server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the TaskService.
	client := pb.NewTaskServiceClient(conn)

	// Create a context with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Define a dummy task to be created.
	task := &pb.Task{
		Id:     "123",
		Title:  "New Task",
		Author: "Tester",
	}

	// Build the CreateTask request.
	req := &pb.CreateTaskRequest{Task: task}
	// Call the CreateTask RPC.
	resp, err := client.CreateTask(ctx, req)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	// Compare the created task with the one sent in the request.
	if !proto.Equal(resp.Task, task) {
		t.Errorf("expected task %+v, got %+v", task, resp.Task)
	}
}

func TestUpdateTask(t *testing.T) {
	// Start an in-process gRPC server for testing.
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	// Dial the gRPC server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the TaskService.
	client := pb.NewTaskServiceClient(conn)

	// Create a context with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Define the task with updated information.
	task := &pb.Task{
		Id:     "123",
		Title:  "Updated Task",
		Author: "Updater",
	}

	// Build the UpdateTask request.
	req := &pb.UpdateTaskRequest{Task: task}
	// Call the UpdateTask RPC.
	resp, err := client.UpdateTask(ctx, req)
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	// Compare the updated task with the one sent in the request.
	if !proto.Equal(resp.Task, task) {
		t.Errorf("expected task %+v, got %+v", task, resp.Task)
	}
}

func TestDeleteTask(t *testing.T) {
	// Start an in-process gRPC server for testing.
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	// Dial the gRPC server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the TaskService.
	client := pb.NewTaskServiceClient(conn)

	// Create a context with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Build the DeleteTask request with the task ID to delete.
	req := &pb.DeleteTaskRequest{Id: "101"}
	// Call the DeleteTask RPC.
	resp, err := client.DeleteTask(ctx, req)
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}

	// Verify that the response includes a Task.
	if resp.Task == nil {
		t.Fatalf("expected non-nil Task in response")
	}
	// Check that the returned Task has the expected ID.
	if resp.Task.Id != req.Id {
		t.Errorf("expected task ID '%v', got %s", req.Id, resp.Task.Id)
	}
}
