package campaign

import (
	"errors"
	"strings"
	"time"

	"github.com/rs/xid"
)

// Contact represents an email contact for a campaign.
type Contact struct {
	Email string
}

// Campaign represents an email campaign with a name, content, contacts, and creation timestamp.
type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

// NewCampaign creates a new Campaign with the given name, content, and a list of emails.
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	err := validate(name, content, emails)
	if err != nil {
		return nil, err
	}
	contacts := emailsToContacts(emails)
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
	}, nil
}

var (
	// errNameEmpty is returned when the campaign name is empty or whitespace only.
	errNameEmpty error = errors.New("name cannot be empty")

	// errContentEmpty is returned when the campaign content is empty or whitespace only.
	errContentEmpty error = errors.New("content cannot be empty")

	// errContactsEmpty is returned when the list of emails/contacts is empty.
	errContactsEmpty error = errors.New("contacts list cannot be empty")
)

// validate checks if the campaign parameters are valid.
// It returns a specific error if the name, content, or email list is empty.
func validate(name string, content string, emails []string) error {
	if strings.TrimSpace(name) == "" {
		return errNameEmpty
	}
	if strings.TrimSpace(content) == "" {
		return errContentEmpty
	}
	if len(emails) == 0 {
		return errContactsEmpty
	}
	return nil
}

func emailsToContacts(emails []string) (contacts []Contact) {
	contacts = make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = strings.TrimSpace(email)
	}
	return contacts
}
