package endpoints

import (
	"errors"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.CampaignService.DeleteCampaign(id)
	if err != nil {
		if errors.Is(err, apperror.ErrCampaignNotFound) {
			response.ErrorJSON(w, r, http.StatusNotFound, err)
			return
		}
		response.ErrorJSON(w, r, http.StatusInternalServerError, apperror.ErrInternal)
		return
	}

	response.JSON(w, r, http.StatusOK, "client successfully deleted", nil)
}
