package main

import (
	"eulabs-case-go/models"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/produtos", models.GetAll)
	e.POST("/produtos", models.PostProdutos)
	e.Logger.Fatal(e.Start(":9000"))
}
