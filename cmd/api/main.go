package main

import (
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"github.com/brenodsm/GoCampaign/internal/endpoints"
	"github.com/brenodsm/GoCampaign/internal/infrastructure/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaign_service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: &campaign_service,
	}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignsGet)

	http.ListenAndServe(":5000", r)
}
