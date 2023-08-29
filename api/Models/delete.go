package models

import (
	"database/sql"
	"errors"
)

func DeleteProduto(db *sql.DB, id string) error {
	//prepara a consulta SQL usando o parâmetro id
	sql, err := db.Prepare("DELETE FROM produto WHERE id = ?")
	if err != nil {
		return err
	}
	defer sql.Close()
	//executa a consulta e verifica se alguma linha foi afetada
	result, err := sql.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	//se nenhuma linha foi afetada, significa que o id não existe na tabela
	if rowsAffected == 0 {
		return errors.New("Produto não encontrado")
	}
	//se alguma linha foi afetada, significa que o produto foi deletado com sucesso
	return nil
}
