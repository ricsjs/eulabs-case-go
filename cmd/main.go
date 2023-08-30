package main

import (
	"eulabs-case-go/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
