package dto

// CampaignDTO holds input data for creating or updating a campaign.
type CampaignDTO struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Emails  []string `json:"emails"`
}
