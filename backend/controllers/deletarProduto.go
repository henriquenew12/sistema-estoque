package controllers

import (
	"fmt"
	"net/http"
	"sistema-estoque/models"

	"github.com/gin-gonic/gin"
)

func DeletarProduto(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := models.DeletarProduto(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Produto deletado com sucesso"})
}
