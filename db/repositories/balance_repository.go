package repositories

import (
	"context"
	"rinha/backend/api/models"
)

func (sr *BalanceRepository) FindByClient(idClient int) (*models.Balance, error) {
	st := "SELECT id, cliente_id, valor from saldos where cliente_id = $1;"

	row := sr.Conn.QueryRow(context.Background(), st, idClient)

	saldo := &models.Balance{}
	err := row.Scan(&saldo.Id, &saldo.IdClient, &saldo.Value)

	if err != nil {
		return nil, err
	}
	return saldo, nil
}
