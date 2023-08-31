package repository

import (
	models "eulabs-case-go/api/Models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllProducts(t *testing.T) {
	//criando mock pro banco
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	//cria um slide de produtos que eu espero (nesse caso, eu já tenho esse registro cadastrado no banco de dados)
	expected := []models.Produto{
		{Id: "cjodb0d933lkbrfg74j0", Nome: "Carro", Preco: 300000, Status: "Indisponível"},
	}

	//configuro o mock para retornar os registros esperados
	rows := sqlmock.NewRows([]string{"id", "nome", "preco", "status"}).
		AddRow(expected[0].Id, expected[0].Nome, expected[0].Preco, expected[0].Status)
	mock.ExpectQuery("SELECT \\* FROM produto").WillReturnRows(rows)

	actual, err := GetAllProducts()
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)

	//verifica se o mock foi chamado corretamente
	mock.ExpectationsWereMet()
}

func TestCreateProduto(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	p := models.Produto{
		Nome:   "Produto de Teste",
		Preco:  30.0,
		Status: "Ativo",
	}

	mock.ExpectPrepare("INSERT INTO produto").
		ExpectExec().WithArgs(sqlmock.AnyArg(), p.Nome, p.Preco, p.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = CreateProduto(p)
	assert.NoError(t, err)

	// Verificar se o mock foi chamado corretamente
	mock.ExpectationsWereMet()
}

func TestGetProdutoByID(t *testing.T) {
	// Criar um mock do banco de dados
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Criar um produto que será buscado no banco de dados
	expected := models.Produto{
		Id:     "cjoej167dit4d26vefv0",
		Nome:   "Produto de Teste",
		Preco:  30,
		Status: "Ativo",
	}

	// Configurar o mock para retornar o produto esperado
	rows := sqlmock.NewRows([]string{"id", "nome", "preco", "status"}).
		AddRow(expected.Id, expected.Nome, expected.Preco, expected.Status)
	mock.ExpectPrepare("SELECT \\* FROM produto WHERE id = ?").ExpectQuery().
		WithArgs(expected.Id).
		WillReturnRows(rows)

	// Chamar a função GetProdutoByID e verificar se não há erro
	actual, err := GetProdutoByID(expected.Id)
	assert.NoError(t, err)

	// Verificar se o produto retornado é igual ao esperado
	assert.Equal(t, expected, actual)

	// Verificar se o mock foi chamado corretamente
	mock.ExpectationsWereMet()
}
