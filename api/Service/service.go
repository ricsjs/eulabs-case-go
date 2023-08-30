package service

import (
	models "eulabs-case-go/api/Models"
	repository "eulabs-case-go/api/Repository"
	"eulabs-case-go/database"
)

func GetAll() ([]models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return repository.GetAll(db)
}

func CreateProduto(produto models.Produto) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	return repository.ProdutoInsert(db, produto)
}

func GetProdutoByID(id string) (models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return models.Produto{}, err
	}
	defer db.Close()

	return repository.GetById(db, id)
}

func UpdateProduto(produto models.Produto) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	return repository.UpdateProduto(db, produto)
}

func DeleteProduto(id string) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	return repository.DeleteProduto(db, id)
}
