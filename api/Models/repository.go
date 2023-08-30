package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/rs/xid"
)

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

func UpdateProduto(db *sql.DB, p Produto) error {
	//executa a query
	sql, err := db.Exec(`UPDATE produto SET nome=?, preco=?, status=? WHERE id=?`, p.Nome, p.Preco, p.Status, p.Id)
	//verifica se houve erro
	if err != nil {
		log.Println(err)
		return err
	}
	//armazena a quantidade de linhas afetadas
	rowsAffected, err := sql.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	//mostra a quantidade de linhas afetadas no log
	log.Printf("%d linhas afetadas\n", rowsAffected)
	return nil
}
