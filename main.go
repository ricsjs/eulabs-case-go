package main

import (
	"eulabs-case-go/models"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/produtos", models.GetAll)
	e.GET("/produtos/:id", models.GetProduto)
	e.POST("/produtos", models.PostProdutos)
	e.PUT("/produtos/:id", models.PutProduto)
	e.DELETE("/produtos/:id", models.DeleteProduto)
	e.Logger.Fatal(e.Start(":9000"))
}
