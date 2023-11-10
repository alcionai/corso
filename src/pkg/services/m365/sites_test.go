package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type siteIntegrationSuite struct {
	tester.Suite
}

func TestSiteIntegrationSuite(t *testing.T) {
	suite.Run(t, &siteIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *siteIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)
}

func (suite *siteIntegrationSuite) TestSites() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tconfig.NewM365Account(t)

	sites, err := Sites(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)

	for _, s := range sites {
		suite.Run("site_"+s.ID, func() {
			t := suite.T()
			assert.NotEmpty(t, s.WebURL)
			assert.NotEmpty(t, s.ID)
		})
	}
}

func (suite *siteIntegrationSuite) TestSites_GetByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tconfig.NewM365Account(t)

	sites, err := Sites(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)

	for _, s := range sites {
		suite.Run("site_"+s.ID, func() {
			t := suite.T()
			site, err := SiteByID(ctx, acct, s.ID)
			assert.NoError(t, err, clues.ToCore(err))
			assert.NotEmpty(t, site.WebURL)
			assert.NotEmpty(t, site.ID)
			assert.NotEmpty(t, site.OwnerType)
		})
	}
}

func (suite *siteIntegrationSuite) TestSites_InvalidCredentials() {
	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "data",
					})
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
		{
			name: "Empty Credentials",
			acct: func(t *testing.T) account.Account {
				// intentionally swallowing the error here
				a, _ := account.NewAccount(account.ProviderM365)
				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sites, err := Sites(ctx, test.acct(t), fault.New(true))
			assert.Empty(t, sites, "returned some sites")
			assert.NotNil(t, err)
		})
	}
}

// ---------------------------------------------------------------------------
// Unit
// ---------------------------------------------------------------------------

type siteUnitSuite struct {
	tester.Suite
}

func TestSiteUnitSuite(t *testing.T) {
	suite.Run(t, &siteUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockGASites struct {
	response []models.Siteable
	err      error
}

func (m mockGASites) GetAll(context.Context, *fault.Bus) ([]models.Siteable, error) {
	return m.response, m.err
}

func (m mockGASites) GetByID(context.Context, string, api.CallConfig) (models.Siteable, error) {
	if len(m.response) == 0 {
		return nil, m.err
	}

	return m.response[0], m.err
}

func (suite *siteUnitSuite) TestGetAllSites() {
	table := []struct {
		name      string
		mock      func(context.Context) getAller[models.Siteable]
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getAller[models.Siteable] {
				return mockGASites{[]models.Siteable{}, nil}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getAller[models.Siteable] {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.NoSPLicense)))
				odErr.SetErrorEscaped(merr)

				return mockGASites{nil, graph.Stack(ctx, odErr)}
			},
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrServiceNotEnabled, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getAller[models.Siteable] {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockGASites{nil, graph.Stack(ctx, odErr)}
			},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gas := test.mock(ctx)

			_, err := getAllSites(ctx, gas)
			test.expectErr(t, err)
		})
	}
}

func (suite *siteUnitSuite) TestGetSites() {
	table := []struct {
		name       string
		mock       func(context.Context) api.GetByIDer[models.Siteable]
		expectErr  assert.ErrorAssertionFunc
		expectSite func(*testing.T, *Site)
	}{
		{
			name: "ok - no owner",
			mock: func(ctx context.Context) api.GetByIDer[models.Siteable] {
				return mockGASites{[]models.Siteable{
					mock.DummySite(nil),
				}, nil}
			},
			expectErr: assert.NoError,
			expectSite: func(t *testing.T, site *Site) {
				assert.NotEmpty(t, site.ID)
				assert.NotEmpty(t, site.WebURL)
				assert.Empty(t, site.OwnerID)
			},
		},
		{
			name: "ok - owner user",
			mock: func(ctx context.Context) api.GetByIDer[models.Siteable] {
				return mockGASites{[]models.Siteable{
					mock.DummySite(mock.UserIdentity("userid", "useremail")),
				}, nil}
			},
			expectErr: assert.NoError,
			expectSite: func(t *testing.T, site *Site) {
				assert.NotEmpty(t, site.ID)
				assert.NotEmpty(t, site.WebURL)
				assert.Equal(t, site.OwnerID, "userid")
				assert.Equal(t, site.OwnerEmail, "useremail")
				assert.Equal(t, site.OwnerType, SiteOwnerUser)
			},
		},
		{
			name: "ok - group user with ID and email",
			mock: func(ctx context.Context) api.GetByIDer[models.Siteable] {
				return mockGASites{[]models.Siteable{
					mock.DummySite(mock.GroupIdentitySet("groupid", "groupemail")),
				}, nil}
			},
			expectErr: assert.NoError,
			expectSite: func(t *testing.T, site *Site) {
				assert.NotEmpty(t, site.ID)
				assert.NotEmpty(t, site.WebURL)
				assert.Equal(t, SiteOwnerGroup, site.OwnerType)
				assert.Equal(t, "groupid", site.OwnerID)
				assert.Equal(t, "groupemail", site.OwnerEmail)
			},
		},
		{
			name: "ok - group user with no ID but email",
			mock: func(ctx context.Context) api.GetByIDer[models.Siteable] {
				return mockGASites{[]models.Siteable{
					mock.DummySite(mock.GroupIdentitySet("", "groupemail")),
				}, nil}
			},
			expectErr: assert.NoError,
			expectSite: func(t *testing.T, site *Site) {
				assert.NotEmpty(t, site.ID)
				assert.NotEmpty(t, site.WebURL)
				assert.Equal(t, SiteOwnerGroup, site.OwnerType)
				assert.Equal(t, "", site.OwnerID)
				assert.Equal(t, "groupemail", site.OwnerEmail)
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gas := test.mock(ctx)

			site, err := getSiteByID(ctx, gas, "id", api.CallConfig{})
			test.expectSite(t, site)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
