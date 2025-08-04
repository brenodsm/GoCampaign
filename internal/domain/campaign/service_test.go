package campaign

import (
	"testing"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaign) error {
	args := r.Called(campaing)
	return args.Error(0)
}

func TestCreateCampaign(t *testing.T) {
	t.Parallel()
	campaign := dto.CampaignDTO{
		Name:    "Campaign X",
		Content: "Body",
		Emails:  []string{"email@gmail.com", "email2@gmail.com"},
	}

	mockRepo := new(repositoryMock)
	mockRepo.On("Save", mock.AnythingOfType("*campaign.Campaign")).Return(nil)

	service := Service{Repository: mockRepo}
	id, err := service.Create(campaign)

	assert.NotEmpty(t, id)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateCampaign_Save(t *testing.T) {
	t.Parallel()
	campaign := dto.CampaignDTO{
		Name:    "Campaign X",
		Content: "Body",
		Emails:  []string{"email@gmail.com"},
	}

	mockRepo := new(repositoryMock)
	mockRepo.On("Save", mock.MatchedBy(func(c *Campaign) bool {
		if c.Name != campaign.Name {
			return false
		} else if c.Content != campaign.Content {
			return false
		} else if len(c.Contacts) != len(campaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service := Service{Repository: mockRepo}

	id, err := service.Create(campaign)

	assert.NotEmpty(t, id)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreate_ValidateRepositorySave(t *testing.T) {
	t.Parallel()
	campaign := dto.CampaignDTO{
		Name:    "Campaign X",
		Content: "Body",
		Emails:  []string{"email@gmail.com"},
	}

	mockRepo := new(repositoryMock)
	service := Service{mockRepo}
	mockRepo.On("Save", mock.Anything).Return(apperror.ErrInternal)

	_, err := service.Create(campaign)

	assert.Error(t, err)
	assert.ErrorIs(t, err, apperror.ErrInternal)

}
