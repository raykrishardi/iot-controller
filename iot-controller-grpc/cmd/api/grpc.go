package main

import (
	"fmt"
	"log"
	"net"

	"github.com/raykrishardi/iot-controller-grpc/internal/pkg/iot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IotServer struct {
	iot.UnimplementedIotServiceServer
}

func grpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s) // for debugging purpose

	iot.RegisterIotServiceServer(s, &IotServer{})

	log.Printf("gRPC server started on port %s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
}
