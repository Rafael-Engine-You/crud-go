package produtos

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Categoria struct {
	Id int
	Nome string
}

func AddCategoria(db *pgxpool.Pool, categoria Categoria) error {

	sql := `
		INSERT INTO categorias(nome) VALUES ($1)
	`
	_, err := db.Exec(
		context.Background(),
		sql,
		categoria.Nome,
	)

	return err
}