package campaign

import (
	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
)

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
