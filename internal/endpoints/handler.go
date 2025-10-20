package endpoints

import (
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
)

type Handler struct {
	CampaignService campaign.ServiceInterface
}
