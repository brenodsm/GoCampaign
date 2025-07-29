package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc    string
		name    string
		content string
		emails  []string
		output  error
	}{
		{
			desc: "valid input", name: "Campaign X", content: "body", emails: []string{"email@gmail.com", "email2@gmail.com"}, output: nil,
		},
		{
			desc: "empty name", name: "  ", content: "Body", emails: []string{"email@gmail.com", "email2@gmail.com"}, output: errNameEmpty,
		},
		{
			desc: "empty content", name: "Campaign X", content: " ", emails: []string{"email@gmail.com"}, output: errContentEmpty,
		},
		{
			desc: "empty emails", name: "Campaign X", content: "body", emails: []string{}, output: errContactsEmpty,
		},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			campaign, err := NewCampaign(tC.name, tC.content, tC.emails)

			if tC.output == nil {
				now := time.Now()
				assert.WithinDuration(t, now, campaign.CreatedOn, 1*time.Second)

				assert.NoError(t, err)
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
		{desc: "nil slice", input: nil, output: []Contact{}},
		{desc: "empty slice", input: []string{}, output: []Contact{}},
		{desc: "single email", input: []string{"email@gmail.com"}, output: []Contact{{Email: "email@gmail.com"}}},
		{desc: "trims whitespace", input: []string{"  email@gmail.com  ", "  email2@gmail.com  "}, output: []Contact{
			{Email: "email@gmail.com"},
			{Email: "email2@gmail.com"},
		}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := emailsToContacts(tc.input)
			assert.Equal(t, tc.output, got)
		})
	}
}
