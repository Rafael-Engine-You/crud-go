package cliente

import(
	"encoding/json"
    "net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handle {
	return &Handler{
		service: service
	}
}

func(h *Handler) ListarTodosClientes(
	response http.ResponseWriter, 
	request http.Request){

	clientes, err := h.service.ListarClientes()	

	if err != nil {
		http.Error(
			response,
			"Error ao buscar cliente",
			http.StatusInternalServerError,
		)

		return
	}

	response.Header().Set(
		"Content-Type",
		"application/json"
	)

	json.NewEncoder(response).Encode(clientes)
}
