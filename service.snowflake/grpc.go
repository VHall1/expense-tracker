package main

import (
	"log"
	"net"

	handler "github.com/vhall1/expense-tracker/service.snowflake/handler/snowflake"
	"github.com/vhall1/expense-tracker/service.snowflake/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	snowflakeService, err := service.NewSnowflakeService(int64(0))
	handler.NewGrpcSnowflakeService(grpcServer, snowflakeService)

	log.Println("[service.snowflake] Starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
