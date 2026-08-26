package models

type Produto struct {
	Id int
	Nome string
	Preco float64
	Categoria Categoria
}