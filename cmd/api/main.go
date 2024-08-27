package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jalvess021/api-golang/db" // Atualize o caminho conforme a sua estrutura de diretórios
)

func main() {
	db.ConnectDb()

	// Inicializa o roteador Gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8282")
}
