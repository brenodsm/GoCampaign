package endpoints_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"github.com/brenodsm/GoCampaign/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestCampaignGet(t *testing.T) {
	testCases := []struct {
		desc               string
		mockReturnData     []campaign.Campaign
		mockReturnErr      error
		expectedStatusCode int
	}{
		{
			desc: "should returns 200 OK with a lista of campaigns when repository suceesds",
			mockReturnData: []campaign.Campaign{
				{
					ID:      "1",
					Name:    "Campaign X",
					Content: "Body",
					Contacts: []campaign.Contact{
						{Email: "test@gmail.com"},
					},
				},
			},
			mockReturnErr:      nil,
			expectedStatusCode: http.StatusOK,
		}, {
			desc:               "should returns 500 Internal Server Error when repository returns ErrInternal",
			mockReturnData:     []campaign.Campaign{},
			mockReturnErr:      apperror.ErrInternal,
			expectedStatusCode: http.StatusInternalServerError,
		}, {
			desc:               "shoul returns 400 Bad Request when repository returns unknown error",
			mockReturnData:     []campaign.Campaign{},
			mockReturnErr:      errors.New("bad request"),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tC := range testCases {
		t.Parallel()
		tC := tC
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/campaings", nil)

		mockService := new(campaignServiceMock)
		handler := endpoints.Handler{CampaignService: mockService}
		mockService.On("ListAll").Return(tC.mockReturnData, tC.mockReturnErr)

		handler.CampaignsGet(w, req)

		assert.Equal(t, tC.expectedStatusCode, w.Code)

	}

}
