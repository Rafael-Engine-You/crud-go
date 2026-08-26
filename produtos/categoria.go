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

func CarregarTodasCategorias(db *pgxpool.Pool) ([]Categoria, error) {
	sql := `
		SELECT id, nome FROM categorias
	`

	linhas, err := db.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	categorias := []Categoria {}

	for linhas.Next() {
		var categoria Categoria

		err := linhas.Scan(
			&categoria.Id,
			&categoria.Nome,
		)

		if err != nil {
			return nil, err
		} 

		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

func AtualizarCategoria(db *pgxpool.Pool, novoNome string, idCategoria int) error {
	sql := `
		UPDATE categorias SET nome = $1 WHERE
		id = $2 
	` 
	_, err := db.Exec(
		context.Background(),
		sql,
		novoNome,
		idCategoria,
	)

	return err
}

func BuscarCategoriarPeloId(db *pgxpool.Pool,idCategoria int)(Categoria, error) {
	sql := `
		SELECT id, nome WHERE id = $1
	`
	var categoria Categoria

	err := db.QueryRow(
		context.Background(),
		sql,
		idCategoria,
	).Scan(
		&categoria.Id,
		&categoria.Nome,
	)

	return categoria, err 
}

func RemoverCategoria(db *pgxpool.Pool, idCategoria int) error {
	sql := `
		DELETE FROM categorias WHERE id = $ 
	` 
	_, err := db.Exec(
		context.Background(),
		sql,
		idCategoria,
	)

	return err
}