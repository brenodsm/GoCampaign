package campaign

import (
	"errors"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
)

type ServiceInterface interface {
	Create(campaignDTO dto.CampaignDTO) (string, error)
	ListAll() ([]Campaign, error)
	GetByID(id string) (*dto.ResponseCampaignDTO, error)
}

type Service struct {
	Repository Repository
}

func (s *Service) Create(campaignDTO dto.CampaignDTO) (string, error) {
	campaign, err := NewCampaign(campaignDTO.Name, campaignDTO.Content, campaignDTO.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", apperror.ErrInternal
	}
	return campaign.ID, nil
}

func (s *Service) ListAll() ([]Campaign, error) {
	return s.Repository.GetAll()
}

func (s *Service) GetByID(id string) (*dto.ResponseCampaignDTO, error) {
	campaign, err := s.Repository.GetByID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrCampaignNotFound) {
			return nil, apperror.ErrCampaignNotFound
		}
		return nil, apperror.ErrInternal
	}

	return &dto.ResponseCampaignDTO{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, err
}
