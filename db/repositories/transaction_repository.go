package repositories

import (
	"context"
	"fmt"

	"rinha/backend/api/models"
	"rinha/backend/utils"

	"github.com/jackc/pgx/v5"
)

func getResponseFromTransactionQueryRow(row pgx.Row) (int64, int64, error) {
	var newValue int64
	var currentLimit int64
	var haveError bool
	var msg string

	err := row.Scan(&newValue, &currentLimit, &haveError, &msg)
	if err != nil {
		queryError := utils.QueryError{Code: utils.DbError,
			Msg: "Erro interno",
			Err: err,
		}
		return 0, 0, queryError
	}

	if haveError {
		queryError := utils.QueryError{Code: utils.DbError,
			Msg: "Limite insuficiente",
			Err: err,
		}
		return newValue, currentLimit, &queryError
	}

	return newValue, currentLimit, nil
}

func (tr *TransactionRepository) Credit(transaction *models.Transaction) (int64, int64, error) {
	st := "select * from creditar($1, $2, $3);"
	row := tr.Conn.QueryRow(context.Background(), st, transaction.IdClient, transaction.Value, transaction.Description)
	return getResponseFromTransactionQueryRow(row)
}

func (tr *TransactionRepository) Debit(transaction *models.Transaction) (int64, int64, error) {
	st := "select * from debitar($1, $2, $3);"
	row := tr.Conn.QueryRow(context.Background(), st, transaction.IdClient, transaction.Value, transaction.Description)
	return getResponseFromTransactionQueryRow(row)
}

func (tr *TransactionRepository) GetLast(idClient int, limit int) ([]models.Transaction, error) {
	st := "select * from transacoes where cliente_id = $1 order by realizada_em desc limit $2"
	rows, err := tr.Conn.Query(context.Background(), st, idClient, limit)

	transactions := make([]models.Transaction, 0, 10)

	for rows.Next() {
		transaction := models.Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.IdClient, &transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			fmt.Println("Error while parsing row")
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
