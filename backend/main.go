package main

import (
	"sistema-estoque/database"
	"sistema-estoque/models"
	"sistema-estoque/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Conectar()
	models.MigrarTabelaProduto()

	r := gin.Default()
	routes.ConfigurarRotas(r)

	r.Run(":8080")
}
