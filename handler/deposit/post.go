package deposit

import (
	"context"
	account "deposit-grpc/proto"
	"log"
)

func (d DepositHandler) Deposit(context.Context, *account.DepositRequest) (*account.DepositResponse, error) {
	counter++
	log.Printf("counter : %v", counter)
	return &account.DepositResponse{Ok: true}, nil
}
