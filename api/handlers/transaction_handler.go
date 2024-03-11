package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"rinha/backend/api/models"
	"rinha/backend/utils"
)

func getTransactionFromRequest(ctx *fiber.Ctx, clientId uint32) (*models.Transaction, error) {
	transactionJson := TransactionBody{}

	err := ctx.BodyParser(&transactionJson)
	if err != nil {
		return nil, utils.NewError(utils.RequestError, "json cannot be decoded"+err.Error())
	}

	// This should be done in an validation layer
	if len(transactionJson.Description) > 10 || len(transactionJson.Description) == 0 {
		return nil, utils.NewError(utils.ValidationError, "Description too long")
	}
	if transactionJson.Value != math.Floor(transactionJson.Value) {
		return nil, utils.NewError(utils.ValidationError, "Value is float")
	}

	intValue := int64(math.Floor(transactionJson.Value))

	return &models.Transaction{
		Id:          0,
		IdClient:    clientId,
		Value:       intValue,
		Type:        transactionJson.Type,
		Description: transactionJson.Description,
		CreatedAt:   time.Now(),
	}, nil
}

func (th *TransactionHandler) PostTransaction(ctx *fiber.Ctx) error {
	responseBody := transactionResponse{Limit: 0, Balance: 0}

	clientId, err := GetIdFromRequest(ctx)
	if err != nil {
		ctx.Locals("body", responseBody)
		return utils.NewError(utils.ValidationError, "Requested Id cannot be converted to an integer")
	}

	exists, err := th.ClientRepository.Exist(clientId)
	if err != nil {
		ctx.Locals("body", responseBody)
		return err
	}
	if !exists {
		ctx.Locals("body", responseBody)
		return utils.NewError(utils.UserNotFoundError, "User not found")
	}

	transaction, err := getTransactionFromRequest(ctx, clientId)
	if err != nil {
		ctx.Locals("body", responseBody)
		return err
	}

	balance, limit, err := th.TransactionRepository.ExecuteTransaction(transaction)

	responseBody.Limit = limit
	responseBody.Balance = balance

	if err != nil {
		ctx.Locals("body", responseBody)
		return err
	}

	return ctx.Status(http.StatusOK).JSON(responseBody)
}
