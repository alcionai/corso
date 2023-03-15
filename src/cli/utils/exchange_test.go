package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ExchangeUtilsSuite struct {
	tester.Suite
}

func TestExchangeUtilsSuite(t *testing.T) {
	suite.Run(t, &ExchangeUtilsSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeUtilsSuite) TestValidateRestoreFlags() {
	table := []struct {
		name     string
		backupID string
		opts     utils.ExchangeOpts
		expect   assert.ErrorAssertionFunc
	}{
		{
			name:     "with backupid",
			backupID: "bid",
			opts:     utils.ExchangeOpts{},
			expect:   assert.NoError,
		},
		{
			name:   "no backupid",
			opts:   utils.ExchangeOpts{},
			expect: assert.Error,
		},
		{
			name:     "valid time",
			backupID: "bid",
			opts:     utils.ExchangeOpts{EmailReceivedAfter: common.Now()},
			expect:   assert.NoError,
		},
		{
			name:   "invalid time",
			opts:   utils.ExchangeOpts{EmailReceivedAfter: "fnords"},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := utils.ValidateExchangeRestoreFlags(test.backupID, test.opts)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *ExchangeUtilsSuite) TestIncludeExchangeRestoreDataSelectors() {
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
				Contact:       a,
				ContactFolder: a,
				Email:         a,
				EmailFolder:   a,
				Event:         a,
				EventCalendar: a,
				Users:         a,
			},
			expectIncludeLen: 3,
		},
		{
			name: "any users, any folders",
			opts: utils.ExchangeOpts{
				ContactFolder: a,
				EmailFolder:   a,
				EventCalendar: a,
				Users:         a,
			},
			expectIncludeLen: 3,
		},
		{
			name: "single user, single of each data",
			opts: utils.ExchangeOpts{
				Contact:       stub,
				ContactFolder: stub,
				Email:         stub,
				EmailFolder:   stub,
				Event:         stub,
				EventCalendar: stub,
				Users:         stub,
			},
			expectIncludeLen: 3,
		},
		{
			name: "single user, single of each folder",
			opts: utils.ExchangeOpts{
				ContactFolder: stub,
				EmailFolder:   stub,
				EventCalendar: stub,
				Users:         stub,
			},
			expectIncludeLen: 3,
		},
		{
			name: "any users, contacts",
			opts: utils.ExchangeOpts{
				Contact:       a,
				ContactFolder: stub,
				Users:         a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, contacts",
			opts: utils.ExchangeOpts{
				Contact:       stub,
				ContactFolder: stub,
				Users:         stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, emails",
			opts: utils.ExchangeOpts{
				Email:       a,
				EmailFolder: stub,
				Users:       a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, emails",
			opts: utils.ExchangeOpts{
				Email:       stub,
				EmailFolder: stub,
				Users:       stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, events",
			opts: utils.ExchangeOpts{
				Event:         a,
				EventCalendar: a,
				Users:         a,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single user, events",
			opts: utils.ExchangeOpts{
				Event:         stub,
				EventCalendar: stub,
				Users:         stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "any users, contacts + email",
			opts: utils.ExchangeOpts{
				Contact:       a,
				ContactFolder: a,
				Email:         a,
				EmailFolder:   a,
				Users:         a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, contacts + email",
			opts: utils.ExchangeOpts{
				Contact:       stub,
				ContactFolder: stub,
				Email:         stub,
				EmailFolder:   stub,
				Users:         stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "any users, email + event",
			opts: utils.ExchangeOpts{
				Email:         a,
				EmailFolder:   a,
				Event:         a,
				EventCalendar: a,
				Users:         a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, email + event",
			opts: utils.ExchangeOpts{
				Email:         stub,
				EmailFolder:   stub,
				Event:         stub,
				EventCalendar: stub,
				Users:         stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "any users, event + contact",
			opts: utils.ExchangeOpts{
				Contact:       a,
				ContactFolder: a,
				Event:         a,
				EventCalendar: a,
				Users:         a,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single users, event + contact",
			opts: utils.ExchangeOpts{
				Contact:       stub,
				ContactFolder: stub,
				Event:         stub,
				EventCalendar: stub,
				Users:         stub,
			},
			expectIncludeLen: 2,
		},
		{
			name: "many users, events",
			opts: utils.ExchangeOpts{
				Event:         many,
				EventCalendar: many,
				Users:         many,
			},
			expectIncludeLen: 1,
		},
		{
			name: "many users, events + contacts",
			opts: utils.ExchangeOpts{
				Contact:       many,
				ContactFolder: many,
				Event:         many,
				EventCalendar: many,
				Users:         many,
			},
			expectIncludeLen: 2,
		},
		{
			name: "mail, no folder or user",
			opts: utils.ExchangeOpts{
				Email: stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "contacts, no folder or user",
			opts: utils.ExchangeOpts{
				Contact: stub,
			},
			expectIncludeLen: 1,
		},
		{
			name: "event, no folder or user",
			opts: utils.ExchangeOpts{
				Event: stub,
			},
			expectIncludeLen: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := utils.IncludeExchangeRestoreDataSelectors(test.opts)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}

func (suite *ExchangeUtilsSuite) TestAddExchangeInclude() {
	var (
		empty             = []string{}
		single            = []string{"single"}
		multi             = []string{"more", "than", "one"}
		containsOnly      = []string{"contains"}
		prefixOnly        = []string{"/prefix"}
		containsAndPrefix = []string{"contains", "/prefix"}
		eisc              = selectors.NewExchangeRestore(nil).Contacts // type independent, just need the func
	)

	table := []struct {
		name                      string
		resources, folders, items []string
		expectIncludeLen          int
	}{
		{
			name:             "no inputs",
			folders:          empty,
			items:            empty,
			expectIncludeLen: 0,
		},
		{
			name:             "single inputs",
			folders:          single,
			items:            single,
			expectIncludeLen: 1,
		},
		{
			name:             "multi inputs",
			folders:          multi,
			items:            multi,
			expectIncludeLen: 1,
		},
		{
			name:             "folder contains",
			folders:          containsOnly,
			items:            empty,
			expectIncludeLen: 1,
		},
		{
			name:             "folder prefixes",
			folders:          prefixOnly,
			items:            empty,
			expectIncludeLen: 1,
		},
		{
			name:             "folder prefixes and contains",
			folders:          containsAndPrefix,
			items:            empty,
			expectIncludeLen: 2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := selectors.NewExchangeRestore(nil)
			// no return, mutates sel as a side effect
			utils.AddExchangeInclude(sel, test.folders, test.items, eisc)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}

func (suite *ExchangeUtilsSuite) TestFilterExchangeRestoreInfoSelectors() {
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
		suite.Run(test.name, func() {
			sel := selectors.NewExchangeRestore(nil)
			utils.FilterExchangeRestoreInfoSelectors(sel, test.opts)
			assert.Len(suite.T(), sel.Filters, test.expectFilterLen)
		})
	}
}
