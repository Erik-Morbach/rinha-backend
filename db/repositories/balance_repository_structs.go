package repositories

import (
	"github.com/jackc/pgx/v5"
	"rinha/backend/api/models"
)

type BalanceRepositoryInterface interface {
	FindByClient(int) (*models.Balance, error)
}

type BalanceRepository struct {
	Conn *pgx.Conn
}

func NewBalanceRepository(conn *pgx.Conn) *BalanceRepository {
	return &BalanceRepository{Conn: conn}
}
