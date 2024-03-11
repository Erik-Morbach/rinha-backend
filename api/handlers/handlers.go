package handlers

import (
	"rinha/backend/db/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	ClientHandler      *ClientHandler
	TransactionHandler *TransactionHandler
	ErrorHandler       *ErrorHandler
}

func NewHandlers(repo *repositories.Repositories) *Handlers {
	return &Handlers{
		ClientHandler:      NewClientHandler(repo),
		TransactionHandler: NewTransactionHandler(repo),
		ErrorHandler:       &ErrorHandler{},
	}
}

// Just an utils that is used across handlers
func GetIdFromRequest(ctx *fiber.Ctx) (uint32, error) {
	parsedValue, err := strconv.ParseInt(ctx.Params("id"), 10, 32)
	return uint32(parsedValue), err
}
