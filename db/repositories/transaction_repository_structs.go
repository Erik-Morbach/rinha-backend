package repositories

import (
	"github.com/jackc/pgx/v5"
	"rinha/backend/api/models"
)

type TransactionRespositoryInterface interface {
	Credit(*models.Transaction) (int64, int64, error)
	Debit(*models.Transaction) (int64, int64, error)
	GetLast(int) ([]models.Transaction, error)
}

type TransactionRepository struct {
	Conn *pgx.Conn
}

func NewTransactionRepository(conn *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{Conn: conn}
}
