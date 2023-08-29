package models

import (
	"database/sql"
)

func ProdutoInsert(db *sql.DB, p Produto) error {
	sql, err := db.Prepare("INSERT INTO produto (nome, preco, marca, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer sql.Close()
	_, err = sql.Exec(p.Nome, p.Preco, p.Marca, p.Status)

	if err != nil {
		return err
	}

	return nil
}
