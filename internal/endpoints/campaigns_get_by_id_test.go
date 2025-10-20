package endpoints_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/brenodsm/GoCampaign/internal/dto"
	"github.com/brenodsm/GoCampaign/internal/endpoints"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCampaignGetByID(t *testing.T) {
	tests := []struct {
		name           string
		campaignID     string
		mockReturn     *dto.ResponseCampaignDTO
		mockError      error
		expectedStatus int
	}{
		{
			name:       "should return 200 when campaign exists",
			campaignID: "123",
			mockReturn: &dto.ResponseCampaignDTO{
				ID:      "123",
				Name:    "Campanha Teste",
				Content: "Conteúdo",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "should return 404 when campaign not found",
			campaignID:     "not-found",
			mockReturn:     nil,
			mockError:      apperror.ErrCampaignNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "should return 500 when internal error occurs",
			campaignID:     "500",
			mockReturn:     nil,
			mockError:      errors.New("internal error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockService := new(campaignServiceMock)
			handler := endpoints.Handler{CampaignService: mockService}

			mockService.On("GetByID", tt.campaignID).Return(tt.mockReturn, tt.mockError)

			req := httptest.NewRequest(http.MethodGet, "/campaigns/"+tt.campaignID, nil)
			w := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.campaignID)
			req = req.WithContext(
				context.WithValue(req.Context(), chi.RouteCtxKey, rctx),
			)

			// Executa handler
			handler.CampaignGetByID(w, req)

			// Verifica resultado
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}
