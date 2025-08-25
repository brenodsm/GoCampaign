package endpoints_test

import (
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
