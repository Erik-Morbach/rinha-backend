package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ch *ClientHandler) GetSummary(ctx *gin.Context) {
	parsedValue, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Id Error")
		return
	}
	clientId := int(parsedValue)

	balance, err := ch.BalanceRepository.FindByClient(clientId)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := ch.ClientRepository.FindById(clientId)
	if err != nil {
		fmt.Println(err)
		return
	}

	transactions, err := ch.TransactionRepository.GetLast(clientId, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	timeStampFormated := time.Now().Format(time.RFC3339)

	balanceSummary := BalanceSummary{
		Total:     balance.Value,
		Limit:     client.Limit,
		CreatedAt: timeStampFormated,
	}
	clientSummary := ClientSummary{
		Balance:          balanceSummary,
		LastTransactions: transactions,
	}
	ctx.JSON(http.StatusOK, clientSummary)
}
