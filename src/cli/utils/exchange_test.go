package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ExchangeUtilsSuite struct {
	suite.Suite
}

func TestExchangeUtilsSuite(t *testing.T) {
	suite.Run(t, new(ExchangeUtilsSuite))
}

func (suite *ExchangeUtilsSuite) TestValidateBackupDetailFlags() {
	table := []struct {
		name     string
		backupID string
		expect   assert.ErrorAssertionFunc
	}{
		{
			name:     "with backupid",
			backupID: "bid",
			expect:   assert.NoError,
		},
		{
			name:   "no backupid",
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, utils.ValidateExchangeRestoreFlags(test.backupID))
		})
	}
}

func (suite *ExchangeUtilsSuite) TestIncludeExchangeBackupDetailDataSelectors() {
	stub := []string{"id-stub"}
	many := []string{"fnord", "smarf"}
	a := []string{utils.Wildcard}

	table := []struct {
		name             string
		opts             utils.ExchangeOpts
		expectIncludeLen int
	}{
		{
			name:             "no selectors",
			expectIncludeLen: 3,
		},
		{
			name: "any users",
			opts: utils.ExchangeOpts{
				Users: a,
			},
			expectIncludeLen: 3,
		},
		{
			name: "single user",
			opts: utils.ExchangeOpts{
				Users: stub,
			},
			expectIncludeLen: 3,
		},
		{
			name: "multiple users",
			opts: utils.ExchangeOpts{
				Users: many,
			},
			expectIncludeLen: 3,
		},
		{
			name: "any users, any data",
			opts: utils.ExchangeOpts{
				Contacts:       a,
				ContactFolders: a,
				Emails:         a,
				EmailFolders:   a,
				Events:         a,
				EventCalendars: a,
				Users:          a,
			},
			expectIncludeLen: 3,
		},
		{
			name: "any users, any folders",
			opts: utils.ExchangeOpts{
				ContactFolders: a,
				EmailFolders:   a,
				EventCalendars: a,
				Users:          a,
			},
			expectIncludeLen: 3,
		},
		{
			name: "single user, single of each data",
			opts: utils.ExchangeOpts{
				Contacts:       stub,
				ContactFolders: stub,
				Emails:         stub,
				EmailFolders:   stub,
				Events:         stub,
				EventCalendars: stub,
				Users:          stub,
			},
			expectIncludeLen: 3,
		},
		{
			name: "single user, single of each folder",
			opts: utils.ExchangeOpts{
				ContactFolders: stub,
				EmailFolders:   stub,
				EventCalendars: stub,
				Users:          stub,
			},
			expectIncludeLen: 3,
		},
		{
			name: "any users, contacts",
			opts: utils.ExchangeOpts{
				Contacts:       a,
				ContactFolders: stub,
				Users:          a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, contacts",
			opts: utils.ExchangeOpts{
				Contacts:       stub,
				ContactFolders: stub,
				Users:          stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, emails",
			opts: utils.ExchangeOpts{
				Emails:       a,
				EmailFolders: stub,
				Users:        a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, emails",
			opts: utils.ExchangeOpts{
				Emails:       stub,
				EmailFolders: stub,
				Users:        stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, events",
			opts: utils.ExchangeOpts{
				Events:         a,
				EventCalendars: a,
				Users:          a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, events",
			opts: utils.ExchangeOpts{
				Events:         stub,
				EventCalendars: stub,
				Users:          stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, contacts + email",
			opts: utils.ExchangeOpts{
				Contacts:       a,
				ContactFolders: a,
				Emails:         a,
				EmailFolders:   a,
				Users:          a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, contacts + email",
			opts: utils.ExchangeOpts{
				Contacts:       stub,
				ContactFolders: stub,
				Emails:         stub,
				EmailFolders:   stub,
				Users:          stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "any users, email + event",
			opts: utils.ExchangeOpts{
				Emails:         a,
				EmailFolders:   a,
				Events:         a,
				EventCalendars: a,
				Users:          a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, email + event",
			opts: utils.ExchangeOpts{
				Emails:         stub,
				EmailFolders:   stub,
				Events:         stub,
				EventCalendars: stub,
				Users:          stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "any users, event + contact",
			opts: utils.ExchangeOpts{
				Contacts:       a,
				ContactFolders: a,
				Events:         a,
				EventCalendars: a,
				Users:          a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, event + contact",
			opts: utils.ExchangeOpts{
				Contacts:       stub,
				ContactFolders: stub,
				Events:         stub,
				EventCalendars: stub,
				Users:          stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "many users, events",
			opts: utils.ExchangeOpts{
				Events:         many,
				EventCalendars: many,
				Users:          many,
			},
			expectIncludeLen: 1,
		},
		{
			name: "many users, events + contacts",
			opts: utils.ExchangeOpts{
				Contacts:       many,
				ContactFolders: many,
				Events:         many,
				EventCalendars: many,
				Users:          many,
			},
			expectIncludeLen: 2,
		},
		{
			name: "mail, no folder or user",
			opts: utils.ExchangeOpts{
				Emails: stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "contacts, no folder or user",
			opts: utils.ExchangeOpts{
				Contacts: stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "event, no folder or user",
			opts: utils.ExchangeOpts{
				Events: stub,
			},
			expectIncludeLen: 1,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			utils.IncludeExchangeRestoreDataSelectors(sel, test.opts)
			assert.Len(t, sel.Includes, test.expectIncludeLen)
		})
	}
}

func (suite *ExchangeUtilsSuite) TestFilterExchangeBackupDetailInfoSelectors() {
	stub := "id-stub"

	table := []struct {
		name            string
		opts            utils.ExchangeOpts
		expectFilterLen int
	}{
		{
			name:            "no selectors",
			expectFilterLen: 0,
		},
		{
			name: "contactName",
			opts: utils.ExchangeOpts{
				ContactName: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "receivedAfter",
			opts: utils.ExchangeOpts{
				EmailReceivedAfter: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "receivedAfter",
			opts: utils.ExchangeOpts{
				EmailReceivedAfter: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "receivedBefore",
			opts: utils.ExchangeOpts{
				EmailReceivedBefore: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "sender",
			opts: utils.ExchangeOpts{
				EmailSender: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "subject",
			opts: utils.ExchangeOpts{
				EmailSubject: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "organizer",
			opts: utils.ExchangeOpts{
				EventOrganizer: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "recurs",
			opts: utils.ExchangeOpts{
				EventRecurs: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "startsAfter",
			opts: utils.ExchangeOpts{
				EventStartsAfter: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "startsBefore",
			opts: utils.ExchangeOpts{
				EventStartsBefore: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "eventSubject",
			opts: utils.ExchangeOpts{
				EventSubject: stub,
			},
			expectFilterLen: 1,
		},
		{
			name: "one of each",
			opts: utils.ExchangeOpts{
				ContactName:         stub,
				EmailReceivedAfter:  stub,
				EmailReceivedBefore: stub,
				EmailSender:         stub,
				EmailSubject:        stub,
				EventOrganizer:      stub,
				EventRecurs:         stub,
				EventStartsAfter:    stub,
				EventStartsBefore:   stub,
				EventSubject:        stub,
			},
			expectFilterLen: 10,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			utils.FilterExchangeRestoreInfoSelectors(sel, test.opts)
			assert.Len(t, sel.Filters, test.expectFilterLen)
		})
	}
}
