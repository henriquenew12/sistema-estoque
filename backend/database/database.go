package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	var err error
	DB, err = gorm.Open(sqlite.Open("produtos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
}
