package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"
	"crud-go/internal/clientes"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/cruddb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		panic("Erro ao conectar")
		// log.Fatal("Erro ao conectar ", err )
	}

	defer db.Close()

	// Factory
	repository := clientes.NewRepository(db)
	service    := clientes.NewService(repository)
	handler    := clientes.NewHandler(service)

	router := chi.NewRouter()

	router.Get("/clientes", handler.ListarTodosClientes)

	log.Println(
        "Servidor executando em http://localhost:8081",
    )

	err := http.ListenAndServe(
		":8081",
		router
	)


	if err != nil {
		log.Fatal(err)
	}
}	