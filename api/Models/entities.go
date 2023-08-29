package models

type Produtos []Produto

var produtos Produtos

type Produto struct {
	Id     string
	Nome   string  `json:"nome"`
	Preco  float64 `json:"preco"`
	Marca  Marca   `json:"marca"`
	Status string  `json:"status"`
}

type Marca struct {
	Id   string
	Nome string `json:"nome"`
}
