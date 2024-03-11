package handlers

import (
	"net/http"
	"rinha/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (ch *ClientHandler) GetSummary(ctx *fiber.Ctx) error {
	clientId, err := GetIdFromRequest(ctx)
	if err != nil {
		return utils.NewError(utils.ValidationError, "Requested Id cannot be converted to an integer")
	}

	exists, err := ch.ClientRepository.Exist(clientId)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(utils.UserNotFoundError, "User not found")
	}

	client, err := ch.ClientRepository.FindById(clientId)
	if err != nil {
		return err
		//return ctx.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	balance, err := ch.BalanceRepository.FindByClient(clientId)
	if err != nil {
		return err
		//return ctx.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	transactions, err := ch.TransactionRepository.GetLast(clientId, 10)
	if err != nil {
		return err
		//return ctx.Status(http.StatusUnprocessableEntity).SendString(err.Error())
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
	return ctx.Status(http.StatusOK).JSON(clientSummary)
}
