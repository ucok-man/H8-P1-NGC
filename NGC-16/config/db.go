package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/p1_ngc_16")
	if err != nil {
		panic(err.Error())
	}
	DB = db
}
