package endpoints_test

import (
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"github.com/brenodsm/GoCampaign/internal/dto"
	"github.com/stretchr/testify/mock"
)

type campaignServiceMock struct {
	mock.Mock
}

func (m *campaignServiceMock) Create(campaignDTO dto.CampaignDTO) (string, error) {
	args := m.Called(campaignDTO)
	return args.String(0), args.Error(1)
}

func (m *campaignServiceMock) ListAll() ([]campaign.Campaign, error) {
	args := m.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}
