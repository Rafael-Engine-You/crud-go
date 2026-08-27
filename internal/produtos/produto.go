package produtos

import (
	"context"
	"crud-go/internal/categoria"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Produto struct {
	Id        int
	Nome      string
	Preco     float64
	Categoria categoria.Categoria
}

func AddProduto(db *pgxpool.Pool, nome string, preco float64, cat categoria.Categoria) error {
	sql := `
		INSERT INTO produtos (nome, preco, categoria_id) VALUES ($1,$2,$3)
	`
	_, err := db.Exec(
		context.Background(),
		sql,
		nome,
		preco,
		cat.Id,
	)

	return err
}

func ListarProduto(db *pgxpool.Pool) ([]Produto, error) {
	sql := `
		SELECT p.id, p.nome, p.preco, c.id, c.nome FROM produtos p
		JOIN categorias c ON c.id = p.categoria_id 
		ORDER BY p.id
	`

	linhas, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	produtos := []Produto{}

	for linhas.Next() {
		var produto Produto

		err := linhas.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Preco,
			&produto.Categoria.Id,
			&produto.Categoria.Nome,
		)

		if err != nil {
			return nil, err
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}
