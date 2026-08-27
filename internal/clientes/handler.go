package clientes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CadastrarCliente godoc
// @Summary Cadastra um novo cliente
// @Description Cria um novo cliente no sistema
// @Tags clientes
// @Accept json
// @Produce json
// @Param cliente body clientes.Cliente true "Dados do Cliente"
// @Security BearerAuth
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Dados inválidos"
// @Router /clientes [post]
func (h *Handler) CadastrarCliente(
	response http.ResponseWriter,
	request *http.Request,
) {
	var cliente Cliente

	err := json.NewDecoder(request.Body).Decode(&cliente)
	if err != nil {
		http.Error(response, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.service.CadastrarCliente(cliente)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{"mensagem": "Cliente cadastrado com sucesso!"})
}

// BuscarClientePorId godoc
// @Summary Busca um cliente pelo ID
// @Description Retorna os dados de um cliente específico com base no ID
// @Tags clientes
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} clientes.Cliente
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Cliente não encontrado"
// @Router /clientes/{id} [get]
func (h *Handler) BuscarClientePorId(response http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	cliente, err := h.service.BuscarClientePorId(id)
	if err != nil {
		http.Error(response, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(cliente)
}

// AtualizarCliente godoc
// @Summary Atualiza um cliente existente
// @Description Altera os dados de um cliente com base no ID
// @Tags clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Param cliente body clientes.Cliente true "Novos dados do Cliente"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "ID inválido ou dados inválidos"
// @Router /clientes/{id} [put]
func (h *Handler) AtualizarCliente(response http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	var cliente Cliente
	err = json.NewDecoder(request.Body).Decode(&cliente)
	if err != nil {
		http.Error(response, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.service.AtualizarCliente(id, cliente)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{"mensagem": "Cliente atualizado com sucesso!"})
}

// DeletarCliente godoc
// @Summary Remove um cliente
// @Description Exclui um cliente do sistema com base no ID
// @Tags clientes
// @Produce json
// @Param id path int true "ID do Cliente"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Cliente não encontrado"
// @Router /clientes/{id} [delete]
func (h *Handler) DeletarCliente(response http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.service.DeletarCliente(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{"mensagem": "Cliente deletado com sucesso!"})
}

// ListarTodosClientes godoc
// @Summary Lista todos os clientes
// @Description Retorna uma lista paginada de clientes
// @Tags clientes
// @Produce json
// @Param limite query int false "Limite por página" default(10)
// @Param pagina query int false "Número da página" default(1)
// @Success 200 {array} clientes.Cliente
// @Failure 500 {string} string "Erro ao buscar clientes"
// @Router /clientes [get]
func (h *Handler) ListarTodosClientes(response http.ResponseWriter, request *http.Request) {
	limiteStr := request.URL.Query().Get("limite")
	paginaStr := request.URL.Query().Get("pagina")

	limite := 10
	pagina := 1

	if l, err := strconv.Atoi(limiteStr); err == nil && l > 0 {
		limite = l
	}

	if p, err := strconv.Atoi(paginaStr); err == nil && p > 0 {
		pagina = p
	}

	clientes, err := h.service.ListarClientes(limite, pagina)
	if err != nil {
		http.Error(response, "Erro ao buscar clientes", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(clientes)
}
