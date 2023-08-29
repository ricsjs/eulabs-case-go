package handler

import (
	service "eulabs-case-go/api/Service"
	"net/http"

	"github.com/labstack/echo"
)

func GetProduto(c echo.Context) error {
	id := c.Param("id")
	for _, produto := range service.Produtos {
		if produto.Id == id {
			c.JSON(http.StatusOK, produto)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
