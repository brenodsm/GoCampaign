package database

import (
	"log"

	"github.com/brenodsm/GoCampaign/internal/domain/campaign"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	if err != nil {
		return err
	}

	log.Println("migrates executada com sucesso")
	return nil
}
