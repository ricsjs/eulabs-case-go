package models

import (
	"database/sql"

	"github.com/rs/xid"
)

func ProdutoInsert(db *sql.DB, p Produto) error {
	//prepara a consulta ao banco de dados
	sql, err := db.Prepare("INSERT INTO produto (id, nome, preco, status) VALUES (?, ?, ?, ?)")
	//valida se há erros
	if err != nil {
		return err
	}
	//fecha a conexão com banco de dados
	defer sql.Close()
	//utiliza a biblioteca xid para adicionar um id único para o campo Id do Produto
	guid := xid.New()
	p.Id = guid.String()
	//passa os valores
	_, err = sql.Exec(p.Id, p.Nome, p.Preco, p.Status)
	//valida se há erros
	if err != nil {
		return err
	}

	return nil
}
