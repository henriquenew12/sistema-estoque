package routes

import (
	"sistema-estoque/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine) {
	r.POST("/produtos", controllers.CriarProdutoHandler)
	r.GET("/produtos", controllers.ListarProdutos)
	r.PUT("/produtos/:id", controllers.AtualizarProduto)
	r.DELETE("/produtos/:id", controllers.DeletarProduto)
}
