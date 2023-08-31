package repository

import (
	"errors"
	models "eulabs-case-go/api/Models"
	"eulabs-case-go/database"
	"log"

	"github.com/rs/xid"
)

func GetAllProducts() ([]models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return []models.Produto{}, err
	}
	defer db.Close()

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

func CreateProduto(p models.Produto) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

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

func GetProdutoByID(id string) (models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return models.Produto{}, err
	}
	defer db.Close()

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

func UpdateProduto(p models.Produto) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

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

func DeleteProduto(id string) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

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
