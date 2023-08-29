package handler

import (
	models "eulabs-case-go/api/Models"
	service "eulabs-case-go/api/Service"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
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
