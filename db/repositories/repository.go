package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	ClientRepository      *ClientRepository
	TransactionRepository *TransactionRepository
	BalanceRepository     *BalanceRepository
}

func NewRepositories(dbPool *pgxpool.Pool) *Repositories {
	return &Repositories{
		ClientRepository:      NewClientRepository(dbPool),
		TransactionRepository: NewTransactionRepository(dbPool),
		BalanceRepository:     NewBalanceRepository(dbPool),
	}
}
