package handler

import (
	models "eulabs-case-go/api/Models"
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
		log.Println(err)
	}
	//fecha a conxão
	defer db.Close()
	produtos, err := models.GetAll(db)
	//se tiver erro retorna o erro
	if err != nil {
		log.Println(err)
	}
	//percorre a quantidade de registros
	for _, p := range produtos {
		//retorna os produtos em json
		c.JSON(http.StatusOK, p)
	}
	return nil
}

func PostProdutos(c echo.Context) error {
	//declara a variável produto do tipo Produto
	produto := models.Produto{}
	//obter os dados da requisição
	err := c.Bind(&produto)
	//validacão de erro se existir
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	//abre a conexão com banco de dados
	db, err := database.OpenConnection()
	//validacão de erro se existir
	if err != nil {
		log.Println(err)
	}
	//fecha a conexão com banco de dados
	defer db.Close()
	//chama a função ProdutoInsert e passa como parâmetro a conexão com bd e o produto
	err = models.ProdutoInsert(db, produto)
	//validacão de erro se existir
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	defer db.Close()
	// Chama a função GetById para buscar o produto com o id informado
	produto, err := models.GetById(db, id)
	if err != nil {
		// Se houver erro, retorna o status 400 e uma mensagem de erro
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	// Se não houver erro, retorna o status 200 e o produto em formato JSON
	return c.JSON(http.StatusOK, produto)
}

func PutProduto(c echo.Context) error {
	// cria uma variável para armazenar o produto
	var p models.Produto
	// obtém o id do produto da rota
	id := c.Param("id")

	// vincula o corpo do pedido à estrutura do produto
	err := c.Bind(&p)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro na requisição")
	}
	// valida se há erros na vinculação
	if err != nil {
		log.Println(err)
	}

	p.Id = id

	// abre a conexão com o banco de dados
	db, err := database.OpenConnection()
	// valida se há erros na abertura da conexão
	if err != nil {
		log.Println(err)
	}
	// fecha a conexão ao final da função
	defer db.Close()
	// chama a função UpdateProduto para executar a query de update
	err = models.UpdateProduto(db, p)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, "Produto atualizado com sucesso")
}

func DeleteProduto(c echo.Context) error {
	//recebe um id como parâmetro
	id := c.Param("id")
	//abre a conexão com o banco de dados
	db, err := database.OpenConnection()
	//validacão de erro se existir
	if err != nil {
		log.Println(err)
	}
	//fecha o banco de dados após a execução da função
	defer db.Close()
	//chama a função DeleteProduto e passa a conexão com o banco de dados e o id do produto
	err = models.DeleteProduto(db, id)
	//validacão de erro se existir
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Produto não encontrado")
	}
	return c.JSON(http.StatusOK, "Produto removido com sucesso")
}
