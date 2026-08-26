package services

import (
	"crud-go/models"
	"crud-go/repositories"
)

type ClienteService struct {
	ClienteRepository repositories.ClienteRepository
}

func (clienteRepository  ClienteRepository) addCliente() error {
	clienteRepository.RegistrarCliente()
}