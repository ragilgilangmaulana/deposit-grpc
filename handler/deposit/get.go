package deposit

import (
	"context"
	account "deposit-grpc/proto"
	"log"
)

func (d DepositHandler) GetDeposit(context.Context, *account.GetDepositRequest) (*account.GetDepositResponse, error) {
	log.Printf("counter : %v", counter)
	return &account.GetDepositResponse{TotalDeposit: float32(counter)}, nil
}
