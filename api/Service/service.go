package service

import (
	models "eulabs-case-go/api/Models"
	repository "eulabs-case-go/api/Repository"
	"eulabs-case-go/database"
)

func GetAllProducts() ([]models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return repository.GetAllProducts(db)
}

func CreateProduto(produto models.Produto) error {
	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	return repository.CreateProduto(db, produto)
}

func GetProdutoByID(id string) (models.Produto, error) {
	db, err := database.OpenConnection()
	if err != nil {
		return models.Produto{}, err
	}
	defer db.Close()

	return repository.GetProdutoByID(db, id)
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
