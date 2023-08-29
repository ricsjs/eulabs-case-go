package models

import "database/sql"

func GetAll(db *sql.DB) ([]Produto, error) {
	//declara a variável produtos do tipo Produto
	produtos := []Produto{}
	//faz a query
	sql, err := db.Query("SELECT * FROM produto")
	//valida se há erro
	if err != nil {
		return nil, err
	}
	//fecha a conexão
	defer sql.Close()
	//faz um for por todos os registros da tabela
	for sql.Next() {
		//declara a variável p do tipo Produto
		var p Produto
		err = sql.Scan(&p.Id, &p.Nome, &p.Preco, &p.Status)
		if err != nil {
			return nil, err
		}
		//adiciona os registros ao slice produtos
		produtos = append(produtos, p)
	}
	return produtos, nil
}
