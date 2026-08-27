package errors

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Mensagem string `json:"mensagem"`
	Status   int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Mensagem
}

func ResponderErro(w http.ResponseWriter, err error, statusPadrao int) {
	status := statusPadrao
	mensagem := err.Error()

	// Se for um erro customizado da aplicação, pega o status dele
	if appErr, ok := err.(*AppError); ok {
		status = appErr.Status
		mensagem = appErr.Mensagem
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"erro": mensagem})
}