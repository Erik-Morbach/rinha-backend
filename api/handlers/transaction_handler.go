package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"rinha/backend/api/models"
	"rinha/backend/utils"
)

func (th *TransactionHandler) PostTransaction(ctx *gin.Context) {
	stamp := time.Now()
	parsedValue, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Error")
		return
	}

	var clientId uint32 = uint32(parsedValue)

	transactionJson := TransactionBody{}

	err = ctx.ShouldBind(&transactionJson)
	if err != nil {
		ctx.String(422, "Error")
		return
	}
	transactionJson.Id = clientId

	transaction := models.Transaction{
		Id:          0,
		IdClient:    transactionJson.Id,
		Value:       transactionJson.Value,
		Type:        transactionJson.Type,
		Description: transactionJson.Description,
		CreatedAt:   stamp,
	}

	var limit int64
	var balance int64
	var unparsedError error

	if transaction.Type[0] == 'c' {
		balance, limit, unparsedError = th.TransactionRepository.Credit(&transaction)
	} else if transaction.Type[0] == 'd' {
		balance, limit, unparsedError = th.TransactionRepository.Debit(&transaction)
	}

	if unparsedError == nil {
		ctx.JSON(http.StatusOK, transactionResponse{Limit: limit, Balance: balance})
		return
	}

	queryError := unparsedError.(*utils.QueryError)

	switch queryError.Code {
	case utils.DbError:
		ctx.String(404, err.Error())
	case utils.InsuficientLimitError:
		ctx.JSON(402, transactionResponse{Limit: limit, Balance: balance})
	}
}
