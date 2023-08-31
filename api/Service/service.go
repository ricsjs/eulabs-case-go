package service

import (
	models "eulabs-case-go/api/Models"
	repository "eulabs-case-go/api/Repository"
)

func GetAllProducts() ([]models.Produto, error) {
	produtos, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func CreateProduto(produto models.Produto) error {
	err := repository.CreateProduto(produto)
	if err != nil {
		return err
	}
	return nil
}

func GetProdutoByID(id string) (models.Produto, error) {
	produto, err := repository.GetProdutoByID(id)
	if err != nil {
		return models.Produto{}, err
	}
	return produto, nil
}

func GetProdutosByPrice(price1 float32, price2 float32) ([]models.Produto, error) {
	produtos, err := repository.GetProdutosByPrice(price1, price2)
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func UpdateProduto(produto models.Produto) error {
	err := repository.UpdateProduto(produto)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduto(id string) error {
	err := repository.DeleteProduto(id)
	if err != nil {
		return err
	}
	return nil
}
