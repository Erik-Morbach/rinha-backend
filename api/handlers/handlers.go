package handlers

import (
	"rinha/backend/db/repositories"
)

type Handlers struct {
	ClientHandler      *ClientHandler
	TransactionHandler *TransactionHandler
}

func NewHandlers(repo *repositories.Repositories) *Handlers {
	return &Handlers{
		ClientHandler:      NewClientHandler(repo),
		TransactionHandler: NewTransactionHandler(repo),
	}
}
