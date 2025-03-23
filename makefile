proto-go:
	protoc -I. -I./x/googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./api/grpc/proto/*.proto

tests:
	go test ./...