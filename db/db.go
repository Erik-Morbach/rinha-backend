package db

import (
	"context"
	"rinha/backend/api/config"
	"rinha/backend/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

// This returns an pgx.Rows with Next() already called, just to verify
// the existence of errors
func ExecuteQuery(dbPool *pgxpool.Pool, query string, args ...any) (pgx.Rows, error) {
	errorMsg := "Error on query " + query + ": "
	rows, err := dbPool.Query(context.Background(), query, args...)
	if err != nil {
		errorMsg += "on calling query:"
		return nil, utils.NewError(utils.DbError, errorMsg + err.Error())
	}
	hasRows := rows.Next()

	if hasRows {
		return rows, nil
	}
	rows.Close()

	err = rows.Err()
	if err != nil {
		errorMsg += "after calling Next:"
		return nil, utils.NewError(utils.DbError, errorMsg + err.Error())
	}

	return nil, utils.NewError(utils.EmptyResultError, "Query returned empty set")
}
