package clientes

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db, // vírgula obrigatória
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

func (r Repository) CarregarTodosClientes() ([]Cliente, error) {
	sql := `
		SELECT id, nome, email, telefone
		FROM clientes
	`
	linhas, err := r.db.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var clientes []Cliente

	for linhas.Next() {
		var cliente Cliente

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

func (r Repository) CarregarClientePeloId(idCliente int) (Cliente, error) {
	var cliente Cliente

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

func (r Repository) AtualizarCliente(id int, cliente Cliente) error {
	sql := `
		UPDATE clientes 
		SET nome = $1, email = $2, telefone = $3 
		WHERE id = $4
	`
	_, err := r.db.Exec(
		context.Background(),
		sql,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
		id,
	)

	return err
}

func (r Repository) DeletarCliente(id int) error {
	sql := `DELETE FROM clientes WHERE id = $1`
	
	result, err := r.db.Exec(
		context.Background(),
		sql,
		id,
	)
	
	if err != nil {
		return err
	}

	// Opcional: verificar se realmente apagou algum registro
	if result.RowsAffected() == 0 {
		return errors.New("cliente não encontrado")
	}

	return nil
}

func (r Repository) ListarClientes(limite int, offset int) ([]Cliente, error) {
	sql := `
		SELECT id, nome, email, telefone 
		FROM clientes 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(context.Background(), sql, limite, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.Nome, &c.Email, &c.Telefone)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}

	return clientes, nil
}