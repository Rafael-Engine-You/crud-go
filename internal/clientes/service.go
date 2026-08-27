package clientes

import (
	"errors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository, // vírgula obrigatória aqui
	}
}

func (s *Service) BuscarClientePorId(id int) (Cliente, error) {
	return s.repository.CarregarClientePeloId(id)
}

func (s *Service) CadastrarCliente(cliente Cliente) error {
	if cliente.Nome == "" {
		return errors.New("Nome é obrigatório")
	}

	if cliente.Email == "" {
		return errors.New("E-mail é obrigatório")
	}

	if cliente.Telefone == "" {
		return errors.New("Telefone é obrigatório")
	}

	return s.repository.RegistrarCliente(cliente)
}

func (s *Service) AtualizarCliente(id int, cliente Cliente) error {
	if cliente.Nome == "" {
		return errors.New("Nome é obrigatório")
	}
	
	if cliente.Email == "" {
		return errors.New("E-mail é obrigatório")
	}
	
	if cliente.Telefone == "" {
		return errors.New("Telefone é obrigatório")
	}

	return s.repository.AtualizarCliente(id, cliente)
}

func (s *Service) DeletarCliente(id int) error {
	return s.repository.DeletarCliente(id)
}

func (s *Service) ListarClientes(limite int, pagina int) ([]Cliente, error) {
	if limite <= 0 {
		limite = 10 // Padrão de 10 itens por página
	}
	if pagina <= 0 {
		pagina = 1
	}

	offset := (pagina - 1) * limite

	return s.repository.ListarClientes(limite, offset)
}