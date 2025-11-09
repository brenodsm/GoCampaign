package database

import (
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := r.Db.Create(campaign)
	return tx.Error
}

func (r *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := r.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (r *CampaignRepository) GetByID(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := r.Db.First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}

func (r *CampaignRepository) UpdateStatus(id, status string) error {
	tx := r.Db.Model(&campaign.Campaign{}).
		Where("id = ?", id).
		Update("status", status)
	return tx.Error
}
