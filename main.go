package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"rinha/backend/api/config"
	"rinha/backend/api/handlers"
	"rinha/backend/db"
	"rinha/backend/db/repositories"
)

func main() {
	cfg := config.NewConfig()

	dbPool, err := db.Connect(cfg)
	if err != nil {
		fmt.Println("Cannot connect with db")
		fmt.Println(err)
		return
	}
	defer dbPool.Close()

	repo := repositories.NewRepositories(dbPool)
	handlers := handlers.NewHandlers(repo)

	api := fiber.New(fiber.Config{
		DisableKeepalive: true,
		Prefork: true,
		ErrorHandler: handlers.ErrorHandler.OnError,
	})

	api.Post("/clientes/:id/transacoes", handlers.TransactionHandler.PostTransaction)

	api.Get("/clientes/:id/extrato", handlers.ClientHandler.GetSummary)

	err = api.Listen(cfg.ServerPort)
	if err != nil {
		fmt.Println("Error on Fiber")
		fmt.Println(err.Error())
	}
	fmt.Println("Shutdown server!")
}
