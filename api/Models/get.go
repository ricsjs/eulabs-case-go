package models

import "database/sql"

func GetById(db *sql.DB, id string) (Produto, error) {
	//variável que armazena o produto com o id encontrado
	var produto Produto

	sql, err := db.Prepare("SELECT * FROM produto WHERE id = ?")
	if err != nil {
		return produto, err
	}
	defer sql.Close()
	//consulta e armazena o resultado na variável produto
	err = sql.QueryRow(id).Scan(&produto.Id, &produto.Nome, &produto.Preco, &produto.Status)
	if err != nil {
		return produto, err
	}
	return produto, nil
}
