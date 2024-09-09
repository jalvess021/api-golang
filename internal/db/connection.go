package db

import (
	"database/sql"
	"github.com/jalvess021/api-golang/config"
	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	// Carregar variáveis de ambiente
	if err := config.LoadEnv(); err != nil {
		return nil, err
	}

	connStr := config.GetDatabaseURL()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}