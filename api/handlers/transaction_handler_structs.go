package handlers

import (
	"rinha/backend/db/repositories"
)

type TransactionHandler struct {
	TransactionRepository *repositories.TransactionRepository
	ClientRepository      *repositories.ClientRepository
}

func NewTransactionHandler(repo *repositories.Repositories) *TransactionHandler {
	return &TransactionHandler{TransactionRepository: repo.TransactionRepository,
		ClientRepository: repo.ClientRepository}
}

type TransactionBody struct {
	Id          uint32  `json:"id"`
	Value       float64 `json:"valor"` // Just for the case where user sets a float value
	Type        string  `json:"tipo"`
	Description string  `json:"descricao"`
}

type transactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
