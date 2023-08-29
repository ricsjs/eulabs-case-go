package handler

import (
	models "eulabs-case-go/api/Models"
	service "eulabs-case-go/api/Service"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	//retorna o status 201 e o JSON de produtos
	return c.JSON(http.StatusOK, service.Produtos)
}

func PostProdutos(c echo.Context) error {
	produto := models.Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	err = service.Save(produto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, "Produto inserido com sucesso!")
}

func GetProduto(c echo.Context) error {
	id := c.Param("id")
	for _, produto := range service.Produtos {
		if produto.Id == id {
			c.JSON(http.StatusOK, produto)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func PutProduto(c echo.Context) error {
	id := c.Param("id")
	produto := models.Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	for i, p := range service.Produtos {
		if p.Id == id {
			service.Produtos[i] = produto
			return c.JSON(http.StatusOK, "Produto atualizado com sucesso!")
		}
	}
	return c.JSON(http.StatusNotFound, "Produto não encontrado!")
}

func DeleteProduto(c echo.Context) error {
	id := c.Param("id")
	for i, produto := range service.Produtos {
		if produto.Id == id {
			service.Produtos = append(service.Produtos[:i], service.Produtos[i+1:]...)
			return c.JSON(http.StatusOK, "Produto removido com sucesso!")
		}
	}
	return c.JSON(http.StatusNotFound, "Produto não encontrado!")
}
