package models

import "sistema-estoque/database"

// Produto representa o modelo do produto
type Produto struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Nome       string  `json:"nome" binding:"required"`
	Preco      float64 `json:"preco" binding:"required"`
	Quantidade int     `json:"quantidade" binding:"required"`
}

// CriarProduto salva um novo produto no banco
func CriarProduto(produto *Produto) error {
	return database.DB.Create(produto).Error
}

// ListarProdutos retorna todos os produtos
func ListarProdutos() ([]Produto, error) {
	var produtos []Produto
	err := database.DB.Find(&produtos).Error
	return produtos, err
}

// AtualizarProduto atualiza um produto existente
func AtualizarProduto(id uint, novosDados *Produto) error {
	var produto Produto
	if err := database.DB.First(&produto, id).Error; err != nil {
		return err
	}
	produto.Nome = novosDados.Nome
	produto.Preco = novosDados.Preco
	produto.Quantidade = novosDados.Quantidade
	return database.DB.Save(&produto).Error
}

// DeletarProduto deleta um produto pelo ID
func DeletarProduto(id uint) error {
	return database.DB.Delete(&Produto{}, id).Error
}

// MigrarTabelaProduto cria a tabela se não existir
func MigrarTabelaProduto() {
	database.DB.AutoMigrate(&Produto{})
}
