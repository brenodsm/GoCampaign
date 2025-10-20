package dto

// CampaignDTO holds input data for creating or updating a campaign.
type CampaignDTO struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}

type ResponseCampaignDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
