package handler

import (
	"fmt"
	"context"
	"encoding/json"

	"go_grpc/model/proto/user"
	"go_grpc/model/proto/transaction"
)

func (h *handlerService) GreetUser(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	// bisa diakses oleh public
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!",
		req.Salutation, req.Name)
	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func (h *handlerService) ApproveTransactions(ctx context.Context, req *transaction.TransactionRequest) (*transaction.TransactionResponse, error) {
	// bisa diakses oleh user yg sudah login
	greeting := user.GreetingResponse{GreetingMessage: "Hari sudah malam!"}
	exampleBytes, err := json.Marshal(greeting)
	if err != nil {
		return &transaction.TransactionResponse{}, err
	}

	return &transaction.TransactionResponse{
		Message:    fmt.Sprint("Total harga: ", req.TotalPrice),
		HttpStatus: 200,
		Data: &transaction.DataResponse{
			Data: exampleBytes,
		},
	}, nil
}
