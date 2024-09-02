package db

import (
	"database/sql"
	"log"

	"github.com/jalvess021/api-golang/config"
	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
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


func ConnectDb() error {
	_, err := GetDBConnection()
	if err != nil {
			log.Fatalf("\r\x1b[KErro ao conectar ao banco de dados! %v\n", err)
			return err
	}
	return nil
}
		