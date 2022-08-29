package main

import (
	"context"
	account "deposit-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50050"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect server %v", err)
	}
	defer conn.Close()
	c := account.NewDepositServiceClient(conn)

	depositReq := account.DepositRequest{}
	Deposit(c, depositReq)

	getDepositReq := account.GetDepositRequest{}
	GetDeposit(c, getDepositReq)

}

func Deposit(conn account.DepositServiceClient, req account.DepositRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := conn.Deposit(ctx, &req)
	if err != nil {
		log.Fatalf("error while get deposit  %v", err)
	}

	log.Printf("response : %v", res)
}

func GetDeposit(conn account.DepositServiceClient, req account.GetDepositRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := conn.GetDeposit(ctx, &req)
	if err != nil {
		log.Fatalf("error while get deposit  %v", err)
	}

	log.Printf("response : %v", res)
}