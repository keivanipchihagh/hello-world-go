package grpc_test

import (
	"context"
	"net"
	"testing"
	"time"

	server "github.com/keivanipchihagh/hello-world-go/api/grpc"
	pb "github.com/keivanipchihagh/hello-world-go/api/grpc/proto/golang"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// startTestGRPCServer starts an in-process gRPC server for testing.
func startTestGRPCServer(t *testing.T) (addr string, cleanup func()) {
	t.Helper()

	// Listen on an ephemeral port.
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server and register our service implementation.
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &server.Server{}) // our server implementation

	// Start the server in a goroutine.
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Logf("gRPC server error: %v", err)
		}
	}()

	// Return the address and a cleanup function.
	return lis.Addr().String(), func() {
		grpcServer.Stop()
		lis.Close()
	}
}

func TestGetTask(t *testing.T) {
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	// Create a context with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetTaskRequest{Id: "123"}
	resp, err := client.GetTask(ctx, req)
	if err != nil {
		t.Fatalf("GetTask failed: %v", err)
	}

	if resp.Task == nil {
		t.Fatalf("expected non-nil Task, got nil")
	}
	if resp.Task.Id != "123" {
		t.Errorf("expected task ID '123', got %s", resp.Task.Id)
	}
}

func TestCreateTask(t *testing.T) {
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a dummy task
	task := &pb.Task{
		Id:     "456",
		Title:  "New Task",
		Author: "Tester",
	}

	req := &pb.CreateTaskRequest{Task: task}
	resp, err := client.CreateTask(ctx, req)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	if !proto.Equal(resp.Task, task) {
		t.Errorf("expected task %+v, got %+v", task, resp.Task)
	}
}

func TestUpdateTask(t *testing.T) {
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	task := &pb.Task{
		Id:     "789",
		Title:  "Updated Task",
		Author: "Updater",
	}

	req := &pb.UpdateTaskRequest{Task: task}
	resp, err := client.UpdateTask(ctx, req)
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	if !proto.Equal(resp.Task, task) {
		t.Errorf("expected task %+v, got %+v", task, resp.Task)
	}
}

func TestDeleteTask(t *testing.T) {
	addr, cleanup := startTestGRPCServer(t)
	defer cleanup()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteTaskRequest{Id: "101"}
	resp, err := client.DeleteTask(ctx, req)
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}

	if resp.Task == nil {
		t.Fatalf("expected non-nil Task in response")
	}
	if resp.Task.Id != "101" {
		t.Errorf("expected task ID '101', got %s", resp.Task.Id)
	}
}
