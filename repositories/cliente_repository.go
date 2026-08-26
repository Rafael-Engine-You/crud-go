package repositories

import (
	"context"
	"crud-go/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClienteRepository struct {
	Db *pgxpool.Pool
}

func (c ClienteRepository) CadastrarCliente(cliente Cliente) error {

	sql := `
		INSERT INTO clientes
		(nome, email, telefone)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	_, err := db.Exec(
		context.Background(),
		sql,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
	)

	return err
}

func (c ClienteRepository) CarregarTodosClientes(db *pgxpool.Pool) ([]Cliente, error){
	sql := `
		SELECT id, nome, email, telefone
		FROM clientes
	`
	linhas, err := db.Query(context.Background(), sql)

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

func (c ClienteRepository) func CarregarClientePeloId(idCliente int) (Cliente, error){
	
	var cliente models.Cliente
	
	sql := `
		SELECT id, nome, email, telefone
		FROM clientes 
		WHERE id = $1
	`
	err := db.QueryRow(
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