package models

import (
	"net/http"

	"github.com/labstack/echo"
)

func PutProdutos(c echo.Context) error {
	id := c.Param("id")
	for i := range produtos {
		if produtos[i].Id == id {
			produtos[i].Status = false
			c.JSON(http.StatusOK, produtos)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
