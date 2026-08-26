package main

import (
	"context"
	"fmt"
	// "log"
	"github.com/jackc/pgx/v5/pgxpool"
	p "crud-go/produtos"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/cruddb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()

	err = p.AddProduto(
		db, 
		"Monitor DEll S272 22 polegadas", 
		99.9,
		p.Categoria{
			Id: 2,
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Produto cadastrado!") 

	produtos, err := p.ListarProduto(db)

	if err != nil {
		fmt.Println(err)
	}

	for _, produto := range produtos {
		fmt.Printf("%d - %s - %s\n", produto.Id, produto.Nome, produto.Categoria.Nome)
	} 

	
}	