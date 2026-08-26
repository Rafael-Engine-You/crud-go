package cliente

type Service struct {
	repository * Repository
}

func NewService(repository *Repository) *Service {
	return &Service {
		repository: repository
	}
}

func (s *Service) ListarClientes() ([] Cliente, error) {
	return s.repository.CarregarTodosClientes()
}

func(s *Service) CadastrarCliente(cliente Cliente) error {

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