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
		tenantID    = "tid"
		pstExchange = path.ExchangeService
		exchange    = pstExchange.String()
		email       = path.EmailCategory.String()
		contacts    = path.ContactsCategory.String()
	)

	type expect struct {
		tenant            string
		serviceResources  []path.ServiceResource
		category          string
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "timbubba",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "bubba",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "tachoma dhaume",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "vyng vang zoombah",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "fat billie",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "seathane",
						Service:           pstExchange,
					}},
					category:    email,
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
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "perell",
						Service:           pstExchange,
					}},
					category:    email,
					subtreePath: stpFor("perell", email, exchange),
				},
				{
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "perell",
						Service:           pstExchange,
					}},
					category:    contacts,
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
					serviceResources:  r.ServiceResources(),
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
	var tenantID = "tid"

	type expect struct {
		tenant            string
		serviceResources  []path.ServiceResource
		category          string
		subtreePath       string
		subtreePathHadErr bool
	}

	stpFor := func(
		resource string,
		service path.ServiceType,
		category path.CategoryType,
	) string {
		return path.Builder{}.Append(tenantID, service.String(), resource, category.String()).String()
	}

	table := []struct {
		name    string
		sel     func() servicerCategorizerProvider
		useName bool
		expect  []expect
	}{
		{
			name: "exchange",
			sel: func() servicerCategorizerProvider {
				sel := *NewExchangeRestore([]string{"hadrian"})
				sel.Include(sel.MailFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "hadrian",
						Service:           path.ExchangeService,
					}},
					category:    path.EmailCategory.String(),
					subtreePath: stpFor("hadrian", path.ExchangeService, path.EmailCategory),
				},
			},
		},
		{
			name: "onedrive",
			sel: func() servicerCategorizerProvider {
				sel := *NewOneDriveRestore([]string{"hella"})
				sel.Filter(sel.Folders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "hella",
						Service:           path.OneDriveService,
					}},
					category:    path.FilesCategory.String(),
					subtreePath: stpFor("hella", path.OneDriveService, path.FilesCategory),
				},
			},
		},
		{
			name: "sharepoint",
			sel: func() servicerCategorizerProvider {
				sel := *NewSharePointRestore([]string{"lem king"})
				sel.Include(sel.LibraryFolders(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "lem king",
						Service:           path.SharePointService,
					}},
					category:    path.EmailCategory.String(),
					subtreePath: stpFor("lem king", path.SharePointService, path.LibrariesCategory),
				},
			},
		},
		{
			name: "groups",
			sel: func() servicerCategorizerProvider {
				sel := *NewGroupsRestore([]string{"fero feritas"})
				sel.Include(sel.TODO(Any()))
				return sel
			},
			expect: []expect{
				{
					tenant: tenantID,
					serviceResources: []path.ServiceResource{{
						ProtectedResource: "fero feritas",
						Service:           path.GroupsService,
					}},
					category:    path.EmailCategory.String(),
					subtreePath: stpFor("fero feritas", path.GroupsService, path.LibrariesCategory),
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
					serviceResources:  r.ServiceResources(),
					category:          r.Category().String(),
					subtreePath:       stpStr,
					subtreePathHadErr: err != nil,
				})
			}

			assert.ElementsMatch(t, test.expect, results)
		})
	}
}
