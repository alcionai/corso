package selectors

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type ReasonsUnitSuite struct {
	tester.Suite
}

func TestReasonsUnitSuite(t *testing.T) {
	suite.Run(t, &ReasonsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ReasonsUnitSuite) TestReasonsFor_thorough() {
	var (
		tenantID = "tid"
		exchange = path.ExchangeService.String()
		email    = path.EmailCategory.String()
		contacts = path.ContactsCategory.String()
	)

	type expect struct {
		tenant            string
		resource          string
		category          string
		service           string
		subtreePath       string
		subtreePathHadErr bool
	}

	stpFor := func(resource, category, service string) string {
		return path.Builder{}.Append(tenantID, service, resource, category).String()
	}

	table := []struct {
		name    string
		sel     func() ExchangeRestore
		useName bool
		expect  []expect
	}{
		{
			name: "no scopes",
			sel: func() ExchangeRestore {
				return *NewExchangeRestore([]string{"timbo"})
			},
			expect: []expect{},
		},
		{
			name: "use name",
			sel: func() ExchangeRestore {
				sel := NewExchangeRestore([]string{"timbo"})
				sel.Include(sel.MailFolders(Any()))
				plainSel := sel.SetDiscreteOwnerIDName("timbo", "timbubba")

				sel, err := plainSel.ToExchangeRestore()
				require.NoError(suite.T(), err, clues.ToCore(err))

				return *sel
			},
			useName: true,
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "timbubba",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("timbubba", email, exchange),
				},
			},
		},
		{
			name: "only includes",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"bubba"})
				sel.Include(sel.MailFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "bubba",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("bubba", email, exchange),
				},
			},
		},
		{
			name: "only filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"tachoma dhaume"})
				sel.Filter(sel.MailFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "tachoma dhaume",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("tachoma dhaume", email, exchange),
				},
			},
		},
		{
			name: "duplicate includes and filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"vyng vang zoombah"})
				sel.Include(sel.MailFolders(Any()))
				sel.Filter(sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "vyng vang zoombah",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("vyng vang zoombah", email, exchange),
				},
			},
		},
		{
			name: "duplicate includes",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"fat billie"})
				sel.Include(sel.MailFolders(Any()), sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "fat billie",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("fat billie", email, exchange),
				},
			},
		},
		{
			name: "duplicate filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"seathane"})
				sel.Filter(sel.MailFolders(Any()), sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "seathane",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("seathane", email, exchange),
				},
			},
		},
		{
			name: "no duplicates",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"perell"})
				sel.Include(sel.MailFolders(Any()), sel.ContactFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "perell",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("perell", email, exchange),
				},
				{
					tenant:      tenantID,
					resource:    "perell",
					category:    contacts,
					service:     exchange,
					subtreePath: stpFor("perell", contacts, exchange),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			results := []expect{}
			rs := reasonsFor(test.sel(), tenantID, test.useName)

			for _, r := range rs {
				stp, err := r.SubtreePath()

				t.Log("stp err", err)

				stpStr := ""
				if stp != nil {
					stpStr = stp.String()
				}

				results = append(results, expect{
					tenant:            r.Tenant(),
					resource:          r.ProtectedResource(),
					service:           r.Service().String(),
					category:          r.Category().String(),
					subtreePath:       stpStr,
					subtreePathHadErr: err != nil,
				})
			}

			assert.ElementsMatch(t, test.expect, results)
		})
	}
}

func (suite *ReasonsUnitSuite) TestReasonsFor_serviceChecks() {
	var (
		tenantID = "tid"
		exchange = path.ExchangeService.String()
		email    = path.EmailCategory.String()
		contacts = path.ContactsCategory.String()
	)

	type expect struct {
		tenant            string
		resource          string
		category          string
		service           string
		subtreePath       string
		subtreePathHadErr bool
	}

	stpFor := func(resource, category, service string) string {
		return path.Builder{}.Append(tenantID, service, resource, category).String()
	}

	table := []struct {
		name    string
		sel     func() ExchangeRestore
		useName bool
		expect  []expect
	}{
		{
			name: "no scopes",
			sel: func() ExchangeRestore {
				return *NewExchangeRestore([]string{"timbo"})
			},
			expect: []expect{},
		},
		{
			name: "only includes",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"bubba"})
				sel.Include(sel.MailFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "bubba",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("bubba", email, exchange),
				},
			},
		},
		{
			name: "only filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"tachoma dhaume"})
				sel.Filter(sel.MailFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "tachoma dhaume",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("tachoma dhaume", email, exchange),
				},
			},
		},
		{
			name: "duplicate includes and filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"vyng vang zoombah"})
				sel.Include(sel.MailFolders(Any()))
				sel.Filter(sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "vyng vang zoombah",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("vyng vang zoombah", email, exchange),
				},
			},
		},
		{
			name: "duplicate includes",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"fat billie"})
				sel.Include(sel.MailFolders(Any()), sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "fat billie",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("fat billie", email, exchange),
				},
			},
		},
		{
			name: "duplicate filters",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"seathane"})
				sel.Filter(sel.MailFolders(Any()), sel.MailFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "seathane",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("seathane", email, exchange),
				},
			},
		},
		{
			name: "no duplicates",
			sel: func() ExchangeRestore {
				sel := *NewExchangeRestore([]string{"perell"})
				sel.Include(sel.MailFolders(Any()), sel.ContactFolders(Any()))

				return sel
			},
			expect: []expect{
				{
					tenant:      tenantID,
					resource:    "perell",
					category:    email,
					service:     exchange,
					subtreePath: stpFor("perell", email, exchange),
				},
				{
					tenant:      tenantID,
					resource:    "perell",
					category:    contacts,
					service:     exchange,
					subtreePath: stpFor("perell", contacts, exchange),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			results := []expect{}
			rs := reasonsFor(test.sel(), tenantID, test.useName)

			for _, r := range rs {
				stp, err := r.SubtreePath()

				t.Log("stp err", err)

				stpStr := ""
				if stp != nil {
					stpStr = stp.String()
				}

				results = append(results, expect{
					tenant:            r.Tenant(),
					resource:          r.ProtectedResource(),
					service:           r.Service().String(),
					category:          r.Category().String(),
					subtreePath:       stpStr,
					subtreePathHadErr: err != nil,
				})
			}

			assert.ElementsMatch(t, test.expect, results)
		})
	}
}
