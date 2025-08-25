package endpoints

import (
	"github.com/brenodsm/GoCampaign/internal/dto"
)

type CampaignServiceInterface interface {
	Create(campaignDTO dto.CampaignDTO) (string, error)
}

type Handler struct {
	CampaignService CampaignServiceInterface
}
