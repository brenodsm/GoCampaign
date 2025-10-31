package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brenodsm/GoCampaign/internal/config"
	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"github.com/brenodsm/GoCampaign/internal/endpoints"
	"github.com/brenodsm/GoCampaign/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	db, err := database.OpenConnection(cfg)

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if cfg.Env == "DEVELOPMENT" {
		err = database.AutoMigrate(db)
		if err != nil {
			log.Fatal("failed to migrations:", err)
		}
	}

	repository := database.CampaignRepository{Db: db}

	campaign_service := campaign.Service{
		Repository: &repository,
	}

	handler := endpoints.Handler{
		CampaignService: &campaign_service,
	}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignsGet)
	r.Get("/campaign/{id}", handler.CampaignGetByID)

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
}
