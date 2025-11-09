package endpoints

import (
	"errors"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CancelCampaign(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.CancelCampaign(id)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrCampaignNotFound):
			response.ErrorJSON(w, r, http.StatusNotFound, apperror.ErrCampaignNotFound)
		case errors.Is(err, apperror.ErrStatusInvalid):
			response.ErrorJSON(w, r, http.StatusConflict, apperror.ErrStatusInvalid)
		default:
			response.ErrorJSON(w, r, http.StatusInternalServerError, apperror.ErrInternal)
		}
		return
	}

	response.JSON(w, r, http.StatusOK, "campaign successfully canceled", nil)
}
