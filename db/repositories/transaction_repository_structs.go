package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"rinha/backend/api/models"
)

type TransactionRespositoryInterface interface {
	Credit(*models.Transaction) (int64, int64, error)
	Debit(*models.Transaction) (int64, int64, error)
	GetLast(uint32) ([]models.Transaction, error)
}

type TransactionRepository struct {
	DbPool *pgxpool.Pool
}

func NewTransactionRepository(dbPool *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{DbPool: dbPool}
}
