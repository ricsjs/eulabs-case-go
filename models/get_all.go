package models

import (
	"net/http"

	"github.com/labstack/echo"
)

type Produtos []Produto

var produtos Produtos

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, produtos)
}
