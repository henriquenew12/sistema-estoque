package controllers

import (
	"fmt"
	"net/http"
	"sistema-estoque/models"

	"github.com/gin-gonic/gin"
)

func AtualizarProduto(c *gin.Context) {
	var novosDados models.Produto
	if err := c.ShouldBindJSON(&novosDados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	// Pegando ID da URL
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	// Validações
	if novosDados.Preco <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "O preço deve ser maior que 0"})
		return
	}
	if novosDados.Quantidade < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "A quantidade não pode ser negativa"})
		return
	}

	// Atualização
	if err := models.AtualizarProduto(id, &novosDados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Produto atualizado com sucesso"})
}
