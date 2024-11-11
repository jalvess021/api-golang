package db

import (
	"database/sql"
	"github.com/jalvess021/api-golang/api/config"
	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {

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