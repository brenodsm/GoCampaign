package response

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	statusSuccess = "success"
	statusError   = "error"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, message string, data any) {
	respondJSON(w, r, statusCode, statusSuccess, message, data)
}

func ErrorJSON(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	respondJSON(w, r, statusCode, statusError, err.Error(), nil)
}

func respondJSON(w http.ResponseWriter, r *http.Request, statusCode int, status string, message string, data any) {
	render.Status(r, statusCode)
	render.JSON(w, r, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
