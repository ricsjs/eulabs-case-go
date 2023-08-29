package main

import (
	handler "eulabs-case-go/api/Handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/produtos", handler.GetAll)
	e.GET("/produtos/:id", handler.GetProduto)
	e.POST("/produtos", handler.PostProdutos)
	e.Logger.Fatal(e.Start(":9000"))
}