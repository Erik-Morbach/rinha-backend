package handlers

import (
	"rinha/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
}

type ErrorBaseJson struct {
	Msg string
	Err error
}

func (eh *ErrorHandler) OnError(ctx *fiber.Ctx, err error) error {
	baseJson := ErrorBaseJson{Msg: "error not identified by server", Err: err}

	predeterminedBody := ctx.Locals("body")
	var nCtx *fiber.Ctx = nil

	switch utils.VerifyErrorCode(err) {
	case utils.DbError:
		baseJson.Msg = "Db error"
		nCtx = ctx.Status(fiber.StatusInternalServerError)
	case utils.EmptyResultError:
		baseJson.Msg = "Empty result set"
		nCtx = ctx.Status(fiber.StatusInternalServerError)
	case utils.RequestError:
		baseJson.Msg = "Error on request"
		nCtx = ctx.Status(fiber.StatusBadRequest)
	case utils.ValidationError:
		baseJson.Msg = "validation error on request"
		nCtx = ctx.Status(fiber.StatusUnprocessableEntity)
	case utils.PaymentError:
		baseJson.Msg = "Payment error"
		nCtx = ctx.Status(fiber.StatusUnprocessableEntity)
	case utils.UserNotFoundError:
		baseJson.Msg = "User not found"
		nCtx = ctx.Status(fiber.StatusNotFound)
	default:
		baseJson.Msg = "Error not treated"
		nCtx = ctx.Status(fiber.StatusBadRequest)
	}
	if predeterminedBody != nil{
		return nCtx.JSON(predeterminedBody)
	} else {
		return nCtx.JSON(baseJson)
	}
}
