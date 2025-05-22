package main

import (
	"sistema-estoque/database"
	"sistema-estoque/models"
	"sistema-estoque/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Conectar()
	models.MigrarTabelaProduto()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	routes.ConfigurarRotas(r)

	r.Run(":8080")
}
