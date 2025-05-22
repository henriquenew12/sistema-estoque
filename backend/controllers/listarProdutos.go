package controllers

import (
	"net/http"
	"sistema-estoque/models"

	"github.com/gin-gonic/gin"
)

func ListarProdutos(c *gin.Context) {
	produtos, err := models.ListarProdutos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produtos)
}
