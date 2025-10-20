package endpoints

import (
	"errors"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	campaing, err := h.CampaignService.GetByID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrCampaignNotFound) {
			response.ErrorJSON(w, r, 404, apperror.ErrCampaignNotFound)
			return
		}
		response.ErrorJSON(w, r, 500, apperror.ErrInternal)
		return
	}
	response.JSON(w, r, 200, "success", campaing)
}
