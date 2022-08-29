package deposit

import (
	account "deposit-grpc/proto"
)

var counter = 0

type DepositHandler struct {
	account.UnimplementedDepositServiceServer
}
