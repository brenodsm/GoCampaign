package database

import (
	"errors"

	"github.com/brenodsm/GoCampaign/internal/apperror"
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
	tx := r.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrCampaignNotFound
		}
		return nil, apperror.ErrInternal
	}
	return &campaign, nil
}

func (r *CampaignRepository) UpdateStatus(id, status string) error {
	tx := r.Db.Model(&campaign.Campaign{}).
		Where("id = ?", id).
		Update("status", status)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return apperror.ErrCampaignNotFound
		}
		return apperror.ErrInternal
	}
	return nil
}

func (r *CampaignRepository) Delete(id string) error {
	tx := r.Db.Select("Contacts").Delete(&campaign.Campaign{}, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return apperror.ErrCampaignNotFound
		}
		return apperror.ErrInternal
	}
	return nil
}
