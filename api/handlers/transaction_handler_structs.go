package handlers

import (
	"rinha/backend/db/repositories"
)

type TransactionHandler struct {
	TransactionRepository *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.Repositories) *TransactionHandler {
	return &TransactionHandler{TransactionRepository: repo.TransactionRepository}
}

type TransactionBody struct {
	Id          uint32 `json:"id"`
	Value       int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type transactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
