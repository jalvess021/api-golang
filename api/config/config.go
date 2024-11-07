package config

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	envPath := filepath.Join("..", ".env") 
	return godotenv.Load(envPath)
}

func GetDatabaseURL() string {
	// Lendo variáveis de ambiente
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Configuração da string de conexão
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
}
