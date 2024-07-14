package server

import (
    "log"
    "net"
    "volta_blockchain/handlers"
    "github.com/EmekaIwuagwu/volta_blockchain/proto"
    "google.golang.org/grpc"
)

func RunServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterVoltaBlockchainServer(s, &handlers.Server{})

    log.Println("Server is running on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
