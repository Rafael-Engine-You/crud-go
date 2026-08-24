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

	categoria := p.Categoria {
		Nome: "Tecnológia",
	}

	err = p.AddCategoria(db, categoria)
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Produto cadastrado com sucesso!")
}	