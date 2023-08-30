package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewConnection(username, password, dbname string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
