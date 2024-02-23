package handlers

import (
	"rinha/backend/api/models"
	"rinha/backend/db/repositories"
)

type ClientHandler struct {
	ClientRepository      *repositories.ClientRepository
	BalanceRepository     *repositories.BalanceRepository
	TransactionRepository *repositories.TransactionRepository
}

func NewClientHandler(repo *repositories.Repositories) *ClientHandler {
	return &ClientHandler{
		ClientRepository:      repo.ClientRepository,
		BalanceRepository:     repo.BalanceRepository,
		TransactionRepository: repo.TransactionRepository,
	}
}

type BalanceSummary struct {
	Total     int64  `json:"total"`
	Limit     int64  `json:"limite"`
	CreatedAt string `json:"data_extrato"`
}
type ClientSummary struct {
	Balance          BalanceSummary       `json:"saldo"`
	LastTransactions []models.Transaction `json:"ultimas_transacoes"`
}
