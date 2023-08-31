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
