package repositories

import "github.com/jackc/pgx/v5"

type Repositories struct {
	ClientRepository      *ClientRepository
	TransactionRepository *TransactionRepository
	BalanceRepository     *BalanceRepository
}

func NewRepositories(conn *pgx.Conn) *Repositories {
	return &Repositories{
		ClientRepository:      NewClientRepository(conn),
		TransactionRepository: NewTransactionRepository(conn),
		BalanceRepository:     NewBalanceRepository(conn),
	}
}
