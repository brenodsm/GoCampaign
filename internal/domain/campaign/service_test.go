package campaign

import (
	"errors"
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

func (r *repositoryMock) GetAll() ([]Campaign, error) {
	return nil, nil
}

func (r *repositoryMock) GetByID(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *repositoryMock) UpdateStatus(id, status string) error {
	args := r.Called(id, status)
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

func TestService_GetByID(t *testing.T) {
	campaignDTO := dto.CampaignDTO{
		Name:    "Campaign X",
		Content: "Body",
		Emails:  []string{"email@gmail.com"},
	}

	campaign, _ := NewCampaign(campaignDTO.Name, campaignDTO.Content, campaignDTO.Emails)

	t.Run("should return campaign successfully", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(repositoryMock)
		service := Service{Repository: mockRepo}

		mockRepo.On("GetByID", mock.MatchedBy(func(id string) bool {
			return id == campaign.ID
		})).Return(campaign, nil)

		campaignReturned, err := service.GetByID(campaign.ID)
		assert.NoError(t, err)
		assert.NotNil(t, campaignReturned)
		assert.Equal(t, campaign.ID, campaignReturned.ID)
		assert.Equal(t, campaign.Name, campaignReturned.Name)
		assert.Equal(t, campaign.Content, campaignReturned.Content)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return internal server error", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(repositoryMock)
		service := Service{Repository: mockRepo}

		mockRepo.On("GetByID", mock.Anything).Return(nil, errors.New("error"))

		campaignReturned, err := service.GetByID(campaign.ID)

		assert.ErrorIs(t, err, apperror.ErrInternal)
		assert.Nil(t, campaignReturned)
		mockRepo.AssertExpectations(t)
	})
}
