package endpoints

import (
	"errors"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/response"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.CampaignService.ListAll()
	if err != nil {
		if errors.Is(err, apperror.ErrInternal) {
			response.ErrorJSON(w, r, http.StatusInternalServerError, err)
			return
		}
		response.ErrorJSON(w, r, http.StatusBadRequest, err)
		return
	}
	response.JSON(w, r, http.StatusOK, "Campaigns successfully returned", campaigns)
}
