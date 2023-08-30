package handler

import (
	models "eulabs-case-go/api/Models"
	repository "eulabs-case-go/api/Repository"
	"eulabs-case-go/database"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	db, err := database.OpenConnection()

	if err != nil {
		log.Println(err)
		return err
	}

	defer db.Close()
	produtos, err := repository.GetAll(db)

	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, produtos)
}

func PostProdutos(c echo.Context) error {

	produto := models.Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	db, err := database.OpenConnection()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	err = repository.ProdutoInsert(db, produto)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, "Produto inserido com sucesso!")
}

func GetProduto(c echo.Context) error {

	id := c.Param("id")
	db, err := database.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	produto, err := repository.GetById(db, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}

	return c.JSON(http.StatusOK, produto)
}

func PutProduto(c echo.Context) error {

	var p models.Produto
	id := c.Param("id")
	err := c.Bind(&p)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro na requisição")
	}

	p.Id = id
	db, err := database.OpenConnection()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	err = repository.UpdateProduto(db, p)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, "Produto atualizado com sucesso")
}

func DeleteProduto(c echo.Context) error {

	id := c.Param("id")
	db, err := database.OpenConnection()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	err = repository.DeleteProduto(db, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	return c.JSON(http.StatusOK, "Produto removido com sucesso")
}
