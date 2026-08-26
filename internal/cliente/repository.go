package cliente

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db
	}
}

func (r Repository) RegistrarCliente(cliente Cliente) error {

	sql := `
		INSERT INTO clientes
		(nome, email, telefone)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
	)

	return err
}

func (r Repository) CarregarTodosClientes(db *pgxpool.Pool) ([]Cliente, error){
	sql := `
		SELECT id, nome, email, telefone
		FROM clientes
	`
	linhas, err := r.db.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	clientes := []models.Cliente {}

	for linhas.Next() {
		var cliente models.Cliente

		err := linhas.Scan(
			&cliente.Id,
			&cliente.Nome,
			&cliente.Email,
			&cliente.Telefone,
		)

		if err != nil {
			return nil, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (r Repository) CarregarClientePeloId(idCliente int) (Cliente, error){
	
	var cliente models.Cliente
	
	sql := `
		SELECT id, nome, email, telefone
		FROM clientes 
		WHERE id = $1
	`
	err := r.db.QueryRow(
		context.Background(),
		sql,
		idCliente,
	).Scan(
		&cliente.Id,
		&cliente.Nome,
		&cliente.Email,
		&cliente.Telefone,
	)

	return cliente, err
}