package campaign

import (
	"strings"
	"time"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/rs/xid"
)

const (
	StatusPending  = "pending"
	StatusStarting = "starting"
	StatusDone     = "done"
)

// Contact represents an email contact for a campaign.
type Contact struct {
	ID         string `gorm:"size:50"`
	CampaignID string
	Email      string `gorm:"size:100" validate:"required,email"`
}

// Campaign represents an email campaign with a name, content, contacts, and creation timestamp.
type Campaign struct {
	ID        string    `gorm:"size:50" validate:"required"`
	Name      string    `gorm:"size:50" validate:"required,min=3,max=50"`
	Content   string    `gorm:"size:1024" validate:"required,min=3,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
	CreatedOn time.Time
}

// NewCampaign creates a new Campaign with the given name, content, and a list of emails.
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	name, content = validate(name, content)
	contacts := emailsToContacts(emails)
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		Status:    StatusPending,
		CreatedOn: time.Now(),
	}

	err := apperror.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}

// validate checks if the campaign parameters are valid.
// It returns a specific error if the name, content, or email list is empty.
func validate(name string, content string) (string, string) {
	name = strings.TrimSpace(name)
	content = strings.TrimSpace(content)
	return name, content
}

func emailsToContacts(emails []string) (contacts []Contact) {
	contacts = make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].ID = xid.New().String()
		contacts[i].Email = strings.TrimSpace(email)
	}
	return contacts
}
