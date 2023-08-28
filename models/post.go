package models

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/xid"
)

func PostProdutos(c echo.Context) error {
	produto := Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	guid := xid.New()
	produto.Id = guid.String()

	produtos = append(produtos, produto)
	return c.JSON(http.StatusCreated, produtos)
}
