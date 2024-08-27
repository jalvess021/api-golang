package db

import (
	"log"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

// Configura as variáveis de ambiente e a string de conexão
func GetDBConnection() (*sql.DB, error) {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Lendo variáveis de ambiente
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Configuração da string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Conectando ao banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verifica a conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDb() error {
	_, err := GetDBConnection()
	if err != nil {
			log.Fatalf("\r\x1b[KErro ao conectar ao banco de dados! %v\n", err)
			return err
	}
	return nil
}
		