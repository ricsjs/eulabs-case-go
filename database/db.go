package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/eulabscase")
	if err != nil {
		return nil, err
	}
	return db, nil
}
