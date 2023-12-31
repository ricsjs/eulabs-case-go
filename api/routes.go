package api

import (
	handler "eulabs-case-go/api/Handler"

	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/produtos", handler.GetAll)
	e.GET("/produtos/:id", handler.GetProduto)
	e.GET("/produtos/preco/:price1/:price2", handler.GetProdutosByPrice)
	e.POST("/produtos", handler.PostProduto)
	e.PUT("/produtos/:id", handler.PutProduto)
	e.DELETE("/produtos/:id", handler.DeleteProduto)
}
