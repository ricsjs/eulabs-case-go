package repository

import (
	"database/sql"
	"errors"
	models "eulabs-case-go/api/Models"
	"log"

	"github.com/rs/xid"
)

func GetAllProducts(db *sql.DB) ([]models.Produto, error) {
	produtos := []models.Produto{}
	sql, err := db.Query("SELECT * FROM produto")
	if err != nil {
		return produtos, err
	}

	defer sql.Close()

	for sql.Next() {
		var p models.Produto
		err = sql.Scan(&p.Id, &p.Nome, &p.Preco, &p.Status)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return produtos, nil
}

func CreateProduto(db *sql.DB, p models.Produto) error {
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

func GetProdutoByID(db *sql.DB, id string) (models.Produto, error) {
	var produto models.Produto

	sql, err := db.Prepare("SELECT * FROM produto WHERE id = ?")
	if err != nil {
		return produto, err
	}
	defer sql.Close()

	err = sql.QueryRow(id).Scan(&produto.Id, &produto.Nome, &produto.Preco, &produto.Status)
	if err != nil {
		return produto, err
	}
	return produto, nil
}

func DeleteProduto(db *sql.DB, id string) error {
	sql, err := db.Prepare("DELETE FROM produto WHERE id = ?")
	if err != nil {
		return err
	}

	defer sql.Close()

	result, err := sql.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Produto não encontrado")
	}

	log.Printf("%d linhas afetadas\n", rowsAffected)
	return nil
}

func UpdateProduto(db *sql.DB, p models.Produto) error {
	sql, err := db.Exec(`UPDATE produto SET nome=?, preco=?, status=? WHERE id=?`, p.Nome, p.Preco, p.Status, p.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := sql.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("%d linhas afetadas\n", rowsAffected)
	return nil
}
