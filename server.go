package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc_web_todo/todo"
)

const ADDRESS = ":50051"

func main() {
	lis, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}
	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}
