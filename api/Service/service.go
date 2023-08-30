package service

import (
	models "eulabs-case-go/api/Models"
	repository "eulabs-case-go/api/Repository"
)

func GetAllProducts() ([]models.Produto, error) {
	return repository.GetAllProducts()
}

func CreateProduto(produto models.Produto) error {
	return repository.CreateProduto(produto)
}

func GetProdutoByID(id string) (models.Produto, error) {

	produto, err := repository.GetProdutoByID(id)

	if err != nil {
		return models.Produto{}, err
	}

	return produto, nil
}

func UpdateProduto(produto models.Produto) error {
	return repository.UpdateProduto(produto)
}

func DeleteProduto(id string) error {
	return repository.DeleteProduto(id)
}
