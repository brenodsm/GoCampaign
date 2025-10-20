package database

import (
	"fmt"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}

func (c *CampaignRepository) GetByID(id string) (*campaign.Campaign, error) {
	if len(c.campaigns) == 0 {
		return nil, apperror.ErrCampaignNotFound
	}
	for i, camp := range c.campaigns {
		fmt.Printf("Comparando camp.ID=%q com id=%q\n", camp.ID, id)
		if camp.ID == id {
			return &c.campaigns[i], nil
		}
	}

	return nil, apperror.ErrCampaignNotFound
}
