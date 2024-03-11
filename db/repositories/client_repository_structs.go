package repositories

import (
	"rinha/backend/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepositoryInterface interface {
	FindById(uint32) (*models.Client, error)
	Exist(uint32) (bool, error)
}

type ClientRepository struct {
	DbPool *pgxpool.Pool
}

func NewClientRepository(dbPool *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{DbPool: dbPool}
}
