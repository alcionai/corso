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
		name                                                                          string
		contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string
		expectIncludeLen                                                              int
	}{
		{
			name:             "no selectors",
			expectIncludeLen: 3,
		},
		{
			name:             "any users",
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "single user",
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "multiple users",
			users:            many,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, any data",
			contacts:         a,
			contactFolders:   a,
			emails:           a,
			emailFolders:     a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, any folders",
			contactFolders:   a,
			emailFolders:     a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "single user, single of each data",
			contacts:         stub,
			contactFolders:   stub,
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "single user, single of each folder",
			contactFolders:   stub,
			emailFolders:     stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, contacts",
			contacts:         a,
			contactFolders:   stub,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, contacts",
			contacts:         stub,
			contactFolders:   stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, emails",
			emails:           a,
			emailFolders:     stub,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, emails",
			emails:           stub,
			emailFolders:     stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, events",
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, events",
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, contacts + email",
			contacts:         a,
			contactFolders:   a,
			emails:           a,
			emailFolders:     a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, contacts + email",
			contacts:         stub,
			contactFolders:   stub,
			emails:           stub,
			emailFolders:     stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "any users, email + event",
			emails:           a,
			emailFolders:     a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, email + event",
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "any users, event + contact",
			contacts:         a,
			contactFolders:   a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, event + contact",
			contacts:         stub,
			contactFolders:   stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "many users, events",
			events:           many,
			eventCalendars:   many,
			users:            many,
			expectIncludeLen: 1,
		},
		{
			name:             "many users, events + contacts",
			contacts:         many,
			contactFolders:   many,
			events:           many,
			eventCalendars:   many,
			users:            many,
			expectIncludeLen: 2,
		},
		{
			name:             "mail, no folder or user",
			emails:           stub,
			expectIncludeLen: 1,
		},
		{
			name:             "contacts, no folder or user",
			contacts:         stub,
			expectIncludeLen: 1,
		},
		{
			name:             "event, no folder or user",
			events:           stub,
			expectIncludeLen: 1,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			utils.IncludeExchangeRestoreDataSelectors(
				sel,
				test.contacts,
				test.contactFolders,
				test.emails,
				test.emailFolders,
				test.events,
				test.eventCalendars,
				test.users)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}

func (suite *ExchangeUtilsSuite) TestFilterExchangeBackupDetailInfoSelectors() {
	stub := "id-stub"

	table := []struct {
		name                                                       string
		contactName                                                string
		after, before, sender, subject                             string
		organizer, recurs, startsAfter, startsBefore, eventSubject string
		expectFilterLen                                            int
	}{
		{
			name:            "no selectors",
			expectFilterLen: 0,
		},
		{
			name:            "contactName",
			contactName:     stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedAfter",
			after:           stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedAfter",
			after:           stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedBefore",
			before:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "sender",
			sender:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "subject",
			subject:         stub,
			expectFilterLen: 1,
		},
		{
			name:            "organizer",
			organizer:       stub,
			expectFilterLen: 1,
		},
		{
			name:            "recurs",
			recurs:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "startsAfter",
			startsAfter:     stub,
			expectFilterLen: 1,
		},
		{
			name:            "startsBefore",
			startsBefore:    stub,
			expectFilterLen: 1,
		},
		{
			name:            "eventSubject",
			eventSubject:    stub,
			expectFilterLen: 1,
		},
		{
			name:            "one of each",
			contactName:     stub,
			after:           stub,
			before:          stub,
			sender:          stub,
			subject:         stub,
			organizer:       stub,
			recurs:          stub,
			startsAfter:     stub,
			startsBefore:    stub,
			eventSubject:    stub,
			expectFilterLen: 10,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			utils.FilterExchangeRestoreInfoSelectors(
				sel,
				test.contactName,
				test.after,
				test.before,
				test.sender,
				test.subject,
				test.organizer,
				test.recurs,
				test.startsAfter,
				test.startsBefore,
				test.eventSubject)
			assert.Equal(t, test.expectFilterLen, len(sel.Filters))
		})
	}
}
