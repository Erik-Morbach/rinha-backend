package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"rinha/backend/api/config"
	"rinha/backend/api/handlers"
	"rinha/backend/db"
	"rinha/backend/db/repositories"
)

func main() {
	cfg := config.NewConfig()
	conn, err := db.Connect(cfg)
	if err != nil {
		fmt.Println("Cannot connect with db")
		fmt.Println(err)
		return
	}

	repo := repositories.NewRepositories(conn)
	handlers := handlers.NewHandlers(repo)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/clientes/:id/transacoes", handlers.TransactionHandler.PostTransaction)

	router.GET("/clientes/:id/extrato", handlers.ClientHandler.GetSummary)

	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Error on Gin")
	}
	fmt.Println("Shutdown server!")
}
