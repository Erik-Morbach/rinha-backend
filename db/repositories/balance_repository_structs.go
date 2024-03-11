package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"rinha/backend/api/models"
)

type BalanceRepositoryInterface interface {
	FindByClient(uint32) (*models.Balance, error)
}

type BalanceRepository struct {
	DbPool *pgxpool.Pool
}

func NewBalanceRepository(dbPool *pgxpool.Pool) *BalanceRepository {
	return &BalanceRepository{DbPool: dbPool}
}
