package endpoints

import (
	"errors"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
	"github.com/brenodsm/GoCampaign/internal/response"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request dto.CampaignDTO
	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		response.ErrorJSON(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := h.CampaignService.Create(request)
	if err != nil {
		if errors.Is(err, apperror.ErrInternal) {
			response.ErrorJSON(w, r, http.StatusInternalServerError, err)
			return
		}
		response.ErrorJSON(w, r, http.StatusBadRequest, err)
		return
	}
	response.JSON(w, r, http.StatusCreated, "Campaign created successfully", map[string]string{"id": id})
}
