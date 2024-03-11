package repositories

import (
	"rinha/backend/api/models"
	"rinha/backend/db"
	"rinha/backend/utils"
)

func (cr *ClientRepository) FindById(id uint32) (*models.Client, error) {
	st := "SELECT id, nome, limite from clientes where id = $1;"
	rows, err := db.ExecuteQuery(cr.DbPool, st, id)
	if err != nil {
		if utils.VerifyErrorCode(err) == utils.EmptyResultError {
			return nil, utils.NewError(utils.UserNotFoundError, "User not found")
		}
		return nil, err
	}
	defer rows.Close()

	cliente := &models.Client{}
	err = rows.Scan(&cliente.Id, &cliente.Name, &cliente.Limit)

	if err != nil {
		return cliente, utils.NewError(utils.DbError, "Cannot scan user")
	}

	return cliente, nil
}

func (cr *ClientRepository) Exist(id uint32) (bool, error) {
	st := "SELECT id from clientes where id = $1;"
	rows, err := db.ExecuteQuery(cr.DbPool, st, id)
	if err != nil {
		if utils.VerifyErrorCode(err) == utils.EmptyResultError {
			return false, nil
		}
		return false, err
	}
	defer rows.Close()

	var clientId uint32
	err = rows.Scan(&clientId)

	if err != nil{
		return false, utils.NewError(utils.DbError, "Could not parse result")
	}
	return clientId == id, nil
}
