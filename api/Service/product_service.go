package service

import (
	models "eulabs-case-go/api/Models"

	"github.com/rs/xid"
)

var marcas map[string]string // Declara um mapa de nomes de marcas para ids

func init() {
	marcas = make(map[string]string) // Inicializa o mapa
}

var Produtos []models.Produto

func Save(produto models.Produto) error {
	guid := xid.New()
	produto.Id = guid.String()
	// Verifica se a marca já existe no mapa
	if id, ok := marcas[produto.Marca.Nome]; ok {
		// Se sim, usa o id existente
		produto.Marca.Id = id
	} else {
		// Se não, gera um novo id e salva no mapa
		produto.Marca.Id = guid.String()
		marcas[produto.Marca.Nome] = produto.Marca.Id
	}

	// Usa append para adicionar o produto à slice produtos
	Produtos = append(Produtos, produto)

	return nil
}
