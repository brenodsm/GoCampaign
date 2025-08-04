package dto

// CampaignDTO holds input data for creating or updating a campaign.
type CampaignDTO struct {
	Name    string
	Content string
	Emails  []string
}
