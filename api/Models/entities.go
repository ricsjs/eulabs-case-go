package models

type Produtos []Produto

var produtos Produtos

type Produto struct {
	Id     string  `json:"string"`
	Nome   string  `json:"nome"`
	Preco  float64 `json:"preco"`
	Status string  `json:"status"`
}
