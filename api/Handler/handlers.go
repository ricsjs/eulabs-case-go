package handler

import (
	models "eulabs-case-go/api/Models"
	service "eulabs-case-go/api/Service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	produtos, err := service.GetAllProducts()
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, produtos)
}

func PostProduto(c echo.Context) error {
	produto := models.Produto{}
	err := c.Bind(&produto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	err = service.CreateProduto(produto)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, "Produto inserido com sucesso!")
}

func GetProduto(c echo.Context) error {
	id := c.Param("id")
	produto, err := service.GetProdutoByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	return c.JSON(http.StatusOK, produto)
}

func GetProdutosByPrice(c echo.Context) error {
	price1 := c.Param("price1")
	price2 := c.Param("price2")

	// Converter as strings para float64 usando a base 10 e 32 bits
	price1Float, err := strconv.ParseFloat(price1, 10)
	if err != nil {
		return err
	}
	price2Float, err := strconv.ParseFloat(price2, 10)
	if err != nil {
		return err
	}

	// Converter os float64 para float32
	price1Float32 := float32(price1Float)
	price2Float32 := float32(price2Float)

	// Chamar a função GetProdutosByPrice com os valores convertidos
	produto, err := service.GetProdutosByPrice(price1Float32, price2Float32)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Nenhum produto encontrado")
	}
	return c.JSON(http.StatusOK, produto)
}

func PutProduto(c echo.Context) error {
	var produto models.Produto
	id := c.Param("id")
	err := c.Bind(&produto)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro na requisição")
	}

	produto.Id = id
	err = service.UpdateProduto(produto)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Erro ao atualizar o produto")
	}
	return c.JSON(http.StatusOK, "Produto atualizado com sucesso")
}

func DeleteProduto(c echo.Context) error {
	id := c.Param("id")
	err := service.DeleteProduto(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	return c.JSON(http.StatusOK, "Produto removido com sucesso")
}
