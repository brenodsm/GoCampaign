package endpoints_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
	"github.com/brenodsm/GoCampaign/internal/endpoints"
	"github.com/brenodsm/GoCampaign/internal/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCampaignPost(t *testing.T) {
	testCases := []struct {
		desc                string
		mockCreatedCampaign dto.CampaignDTO
		rawRequestBody      string
		mockReturnID        string
		mockReturnErr       error
		expectedStatusCode  int
		expectedStatus      string
		expectedMessage     string
	}{
		{
			desc: "should returns 201 Created",
			mockCreatedCampaign: dto.CampaignDTO{
				Name:    "Campaign X",
				Content: "Body",
				Emails:  []string{"test@gmail.com"},
			},
			rawRequestBody:     "",
			mockReturnID:       "1",
			mockReturnErr:      nil,
			expectedStatusCode: http.StatusCreated,
			expectedStatus:     "success",
			expectedMessage:    "Campaign created successfully",
		},
		{
			desc:                "should returns 500 Internal Server Error",
			mockCreatedCampaign: dto.CampaignDTO{},
			rawRequestBody:      "",
			mockReturnID:        "",
			mockReturnErr:       apperror.ErrInternal,
			expectedStatusCode:  http.StatusInternalServerError,
			expectedStatus:      "error",
			expectedMessage:     apperror.ErrInternal.Error(),
		},
		{
			desc: "should returns 400 Bad Request on service error",
			mockCreatedCampaign: dto.CampaignDTO{
				Name:    "",
				Content: "Body",
				Emails:  []string{"test@gmail.com"},
			},
			rawRequestBody:     "",
			mockReturnID:       "1",
			mockReturnErr:      errors.New("Bad Request"),
			expectedStatusCode: http.StatusBadRequest,
			expectedStatus:     "error",
			expectedMessage:    "Bad Request",
		},
		{
			desc:                "should return 400 Bad Request on JSON decode error",
			mockCreatedCampaign: dto.CampaignDTO{},
			rawRequestBody:      `"name": "test", "content": "body"`,
			expectedStatusCode:  http.StatusBadRequest,
			expectedStatus:      "error",
			expectedMessage:     "json: cannot unmarshal string into Go value of type dto.CampaignDTO",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			var b bytes.Buffer

			if tC.rawRequestBody != "" {
				b.WriteString(tC.rawRequestBody)
			} else {
				err := json.NewEncoder(&b).Encode(tC.mockCreatedCampaign)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/campaigns", &b)
			w := httptest.NewRecorder()

			mockService := new(campaignServiceMock)
			handler := endpoints.Handler{
				CampaignService: mockService,
			}

			if tC.rawRequestBody == "" {
				mockService.On("Create", mock.AnythingOfType("dto.CampaignDTO")).Return(tC.mockReturnID, tC.mockReturnErr)
			}

			handler.CampaignPost(w, req)
			assert.Equal(t, tC.expectedStatusCode, w.Code)

			var responseBody response.Response

			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err, "Failed to decode body")

			assert.Equal(t, tC.expectedStatus, responseBody.Status)

			if tC.rawRequestBody != "" {
				assert.Contains(t, responseBody.Message, tC.expectedMessage)
			} else {
				assert.Equal(t, tC.expectedMessage, responseBody.Message)
			}

			if tC.expectedStatus != "error" {
				dataMap, ok := responseBody.Data.(map[string]any)
				assert.True(t, ok, "Data field is not a map expected")

				assert.NotEmpty(t, dataMap["id"])
				assert.Equal(t, tC.mockReturnID, dataMap["id"])
			}

			mockService.AssertExpectations(t)

		})
	}
}
