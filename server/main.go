package main

import (
	"deposit-grpc/handler/deposit"
	account "deposit-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50050"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	account.RegisterDepositServiceServer(s, &deposit.DepositHandler{})
	log.Printf("Server listening at : %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
