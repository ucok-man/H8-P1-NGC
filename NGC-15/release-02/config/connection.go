package db

import (
	"database/sql"
	"fmt"
)

func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error establishing connection: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error ping connection: %v", err)
	}
	return db, nil
}
