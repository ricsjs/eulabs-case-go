package models

import (
	"database/sql"

	"github.com/rs/xid"
)

func ProdutoInsert(db *sql.DB, p Produto) error {
	sql, err := db.Prepare("INSERT INTO produto (id, nome, preco, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer sql.Close()
	guid := xid.New()
	p.Id = guid.String()
	_, err = sql.Exec(p.Id, p.Nome, p.Preco, p.Status)

	if err != nil {
		return err
	}

	return nil
}
