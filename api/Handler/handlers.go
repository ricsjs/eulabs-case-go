package handler

import (
	models "eulabs-case-go/api/Models"
	service "eulabs-case-go/api/Service"
	"log"
	"net/http"

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
