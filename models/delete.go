package models

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeleteProduto(c echo.Context) error {
	id := c.Param("id")
	for i := range produtos {
		if produtos[i].Id == id {
			produtos = append(produtos[:i], produtos[i+1:]...)
			c.JSON(http.StatusOK, produtos)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
