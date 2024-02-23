package repositories

import (
	"rinha/backend/api/models"

	"github.com/jackc/pgx/v5"
)

type ClientRepositoryInterface interface {
	FindById(int) (*models.Client, error)
}

type ClientRepository struct {
	Conn *pgx.Conn
}
func NewClientRepository(conn *pgx.Conn) *ClientRepository {
	return &ClientRepository{Conn: conn}
}
