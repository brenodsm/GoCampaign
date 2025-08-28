package endpoints

import (
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"github.com/brenodsm/GoCampaign/internal/dto"
)

type CampaignServiceInterface interface {
	Create(campaignDTO dto.CampaignDTO) (string, error)
	ListAll() ([]campaign.Campaign, error)
}

type Handler struct {
	CampaignService CampaignServiceInterface
}
