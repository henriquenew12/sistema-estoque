package controllers

import (
	"net/http"
	"sistema-estoque/models"

	"github.com/gin-gonic/gin"
)

func CriarProdutoHandler(c *gin.Context) {
	var produto models.Produto

	// Bind JSON
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if produto.Preco <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "O preço deve ser maior que 0"})
		return
	}
	if produto.Quantidade < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "A quantidade não pode ser negativa"})
		return
	}

	// Criar produto no banco
	if err := models.CriarProduto(&produto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, produto)
}
