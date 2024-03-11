package repositories

import (
	"rinha/backend/api/models"
	"rinha/backend/db"
	"rinha/backend/utils"

	"github.com/jackc/pgx/v5"
)

func getResponseFromTransactionQueryRow(rows pgx.Rows) (int64, int64, error) {
	var newValue int64
	var currentLimit int64
	var haveError bool
	var msg string

	err := rows.Scan(&newValue, &currentLimit, &haveError, &msg)
	if err != nil {
		return 0, 0, utils.NewError(utils.DbError, "Error scanning row " + err.Error())
	}

	if haveError {
		return newValue, currentLimit, utils.NewError(utils.PaymentError, "Not enought balance")
	}

	return newValue, currentLimit, nil
}

func (tr *TransactionRepository) credit(transaction *models.Transaction) (int64, int64, error) {
	st := "select * from creditar($1, $2, $3);"
	rows, err := db.ExecuteQuery(tr.DbPool, st, transaction.IdClient, transaction.Value, transaction.Description)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()
	return getResponseFromTransactionQueryRow(rows)
}

func (tr *TransactionRepository) debit(transaction *models.Transaction) (int64, int64, error) {
	st := "select * from debitar($1, $2, $3);"
	rows, err := db.ExecuteQuery(tr.DbPool, st, transaction.IdClient, transaction.Value, transaction.Description)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()
	return getResponseFromTransactionQueryRow(rows)
}

func (tr *TransactionRepository) GetLast(idClient uint32, limit int) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0, limit)
	st := "select * from transacoes where cliente_id = $1 order by realizada_em desc limit $2"
	rows, err := db.ExecuteQuery(tr.DbPool, st, idClient, limit)

	if err != nil {
		if utils.VerifyErrorCode(err) == utils.EmptyResultError {
			return transactions, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transaction := models.Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.IdClient, &transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreatedAt)
		if err != nil {
			return transactions, utils.NewError(utils.DbError, "Error while parsing row: " + err.Error())
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (tr *TransactionRepository) ExecuteTransaction(transaction *models.Transaction) (int64, int64, error) {
	switch transaction.Type[0] {
	case 'c':
		return tr.credit(transaction)
	case 'd':
		return tr.debit(transaction)
	default:
		return 0, 0, utils.NewError(utils.ValidationError, "Unsuported transaction")
	}
}
