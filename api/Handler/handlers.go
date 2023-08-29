package handler

import (
	models "eulabs-case-go/api/Models"
	service "eulabs-case-go/api/Service"
	"eulabs-case-go/database"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	//abre a conexão
	db, err := database.OpenConnection()
	//se tiver erro retorna o erro
	if err != nil {
		log.Fatal(err)
	}
	//fecha a conxão
	defer db.Close()
	produtos, err := models.GetAll(db)
	//se tiver erro retorna o erro
	if err != nil {
		log.Fatal(err)
	}
	//percorre a quantidade de registros
	for _, p := range produtos {
		//retorna os produtos em json
		c.JSON(http.StatusOK, p)
	}
	return nil
}

func PostProdutos(c echo.Context) error {
	produto := models.Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = models.ProdutoInsert(db, produto)
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, "Produto inserido com sucesso!")
}

func GetProduto(c echo.Context) error {
	// Obtém o valor do parâmetro id da URL
	id := c.Param("id")
	// Abre a conexão com o banco de dados
	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Chama a função GetById para buscar o produto com o id informado
	produto, err := models.GetById(db, id)
	if err != nil {
		// Se houver erro, retorna o status 400 e uma mensagem de erro
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	// Se não houver erro, retorna o status 200 e o produto em formato JSON
	return c.JSON(http.StatusOK, produto)
}

func PutProduto(c echo.Context) error {
	id := c.Param("id")
	produto := models.Produto{}
	err := c.Bind(&produto)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	for i, p := range service.Produtos {
		if p.Id == id {
			service.Produtos[i] = produto
			return c.JSON(http.StatusOK, "Produto atualizado com sucesso!")
		}
	}
	return c.JSON(http.StatusNotFound, "Produto não encontrado!")
}

func DeleteProduto(c echo.Context) error {
	id := c.Param("id")
	for i, produto := range service.Produtos {
		if produto.Id == id {
			service.Produtos = append(service.Produtos[:i], service.Produtos[i+1:]...)
			return c.JSON(http.StatusOK, "Produto removido com sucesso!")
		}
	}
	return c.JSON(http.StatusNotFound, "Produto não encontrado!")
}
