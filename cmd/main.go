package main

import (
	"eulabs-case-go/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
