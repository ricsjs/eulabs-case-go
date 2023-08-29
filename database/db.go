package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// declara a função OpenConnection
func OpenConnection() (*sql.DB, error) {
	//string de conexão do banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/eulabscase")
	//valida se há erros na conexão
	if err != nil {
		return nil, err
	}
	//retorna a string de conexão
	return db, nil
}
