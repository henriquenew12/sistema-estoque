package main

import (
	"fmt"
	"net/http"
	"sistema-estoque/database"
	"sistema-estoque/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	database.Conectar()

	// Migrar tabela Produto
	models.MigrarTabelaProduto()

	// Criar router Gin
	r := gin.Default()

	// Rota para criar produto
	r.POST("/produtos", func(c *gin.Context) {
		var produto models.Produto

		// Bind JSON
		if err := c.ShouldBindJSON(&produto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		// Validações
		if produto.Preco <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "O preço deve ser maior que 0"})
			return
		}
		if produto.Quantidade < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "A quantidade não pode ser negativa"})
			return
		}

		// Create
		if err := models.CriarProduto(&produto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, produto)
	})

	// Read
	r.GET("/produtos", func(c *gin.Context) {
		produtos, err := models.ListarProdutos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, produtos)
	})

	// Update
	r.PUT("/produtos/:id", func(c *gin.Context) {
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

		if err := models.AtualizarProduto(id, &novosDados); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"mensagem": "Produto atualizado com sucesso"})
	})

	// Delete
	r.DELETE("/produtos/:id", func(c *gin.Context) {
		// Pegando ID da URL
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
	})

	// Rodar servidor
	r.Run(":8080")
}
