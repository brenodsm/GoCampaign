package campaign

import (
	"testing"
	"time"

	"github.com/brenodsm/GoCampaign/internal/apperror"
	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc    string
		name    string
		content string
		emails  []string
		Status  string
		output  error
	}{
		{
			desc: "should create campaign with valid input", name: "Campaign X", content: "body", emails: []string{"email@gmail.com", "email2@gmail.com"}, Status: StatusPending, output: nil,
		},
		{
			desc: "should return required field error when name is empt", name: "  ", content: "body", emails: []string{"email@gmail.com", "email2@gmail.com"}, Status: StatusPending, output: apperror.ErrRequiredField,
		},
		{
			desc: "should return min value error when name is too short", name: "tt", content: "body", emails: []string{"email@gmail.com"}, Status: StatusPending,
			output: apperror.ErrMinValueNotReached,
		},
		{
			desc: "should return required field error when content is empty", name: "Campaign X", content: " ", emails: []string{"email@gmail.com"}, Status: StatusPending, output: apperror.ErrRequiredField,
		},
		{
			desc: "should return min value error when content is too short", name: "Campaign X", content: "tt", emails: []string{"email@gmail.com"}, Status: StatusPending, output: apperror.ErrMinValueNotReached,
		},
		{
			desc: "should return min value error when no emails are provided", name: "Campaign X", content: "body", emails: []string{}, Status: StatusPending, output: apperror.ErrMinValueNotReached,
		},
		{
			desc: "should return invalid email error when email is malformed", name: "Campaign X", content: "body", emails: []string{"emailgmail.com"}, Status: StatusPending, output: apperror.ErrInvalidEmail,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			tC := tC
			campaign, err := NewCampaign(tC.name, tC.content, tC.emails)

			if tC.output == nil {
				now := time.Now()
				assert.WithinDuration(t, now, campaign.CreatedOn, 1*time.Second)

				assert.NoError(t, err)
				assert.Equal(t, StatusPending, tC.Status)
				assert.NotEmpty(t, campaign.ID)
			} else {
				assert.ErrorIs(t, err, tC.output)
			}
		})
	}
}

func TestEmailsToContacts(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc   string
		input  []string
		output []Contact
	}{
		{
			desc: "nil slice", input: nil, output: []Contact{},
		},
		{
			desc: "empty slice", input: []string{}, output: []Contact{},
		},
		{
			desc: "single email", input: []string{"email@gmail.com"}, output: []Contact{{Email: "email@gmail.com"}},
		},
		{
			desc: "trims whitespace", input: []string{"  email@gmail.com  ", "  email2@gmail.com  "}, output: []Contact{
				{Email: "email@gmail.com"},
				{Email: "email2@gmail.com"}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := emailsToContacts(tc.input)

			var gotEmails, wantEmails []string
			for _, c := range got {
				gotEmails = append(gotEmails, c.Email)
			}
			for _, c := range tc.output {
				wantEmails = append(wantEmails, c.Email)
			}

			assert.Equal(t, gotEmails, wantEmails)
		})
	}
}
