package database

import (
	"database/sql"
	"eulabs-case-go/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection(c *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.GetDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
