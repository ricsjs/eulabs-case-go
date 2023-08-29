package handler

import (
	service "eulabs-case-go/api/Service"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, service.Produtos)
}
