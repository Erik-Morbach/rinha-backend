package repositories

import (
	"context"
	"rinha/backend/api/models"
)

func (cr *ClientRepository) FindById(id int) (*models.Client, error) {
	st := "SELECT id, nome, limite from clientes where id = $1;"
	row := cr.Conn.QueryRow(context.Background(), st, id)

	cliente := &models.Client{}
	err := row.Scan(&cliente.Id, &cliente.Name, &cliente.Limit)

	if err != nil {
		return nil, err
	}

	return cliente, nil
}
