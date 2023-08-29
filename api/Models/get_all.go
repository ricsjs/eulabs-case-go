package models

import "database/sql"

func GetAll(db *sql.DB) ([]Produto, error) {
	produtos := []Produto{}
	rows, err := db.Query("SELECT * FROM produto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Produto
		err = rows.Scan(&p.Id, &p.Nome, &p.Preco, &p.Marca.Nome, &p.Status)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return produtos, nil
}
