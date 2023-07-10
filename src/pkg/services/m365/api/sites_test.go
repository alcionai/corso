package api_test

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SitesUnitSuite struct {
	tester.Suite
}

func TestSitesUnitSuite(t *testing.T) {
	suite.Run(t, &SitesUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SitesUnitSuite) TestValidateSite() {
	site := models.NewSite()
	site.SetWebUrl(ptr.To("sharepoint.com/sites/foo"))
	site.SetDisplayName(ptr.To("testsite"))
	site.SetId(ptr.To("testID"))

	tests := []struct {
		name           string
		args           models.Siteable
		errCheck       assert.ErrorAssertionFunc
		errIsSkippable bool
	}{
		{
			name:     "No ID",
			args:     models.NewSite(),
			errCheck: assert.Error,
		},
		{
			name: "No WebURL",
			args: func() *models.Site {
				s := models.NewSite()
				s.SetId(ptr.To("id"))
				return s
			}(),
			errCheck: assert.Error,
		},
		{
			name: "No name",
			args: func() *models.Site {
				s := models.NewSite()
				s.SetId(ptr.To("id"))
				s.SetWebUrl(ptr.To("sharepoint.com/sites/foo"))
				return s
			}(),
			errCheck: assert.Error,
		},
		{
			name: "Search site",
			args: func() *models.Site {
				s := models.NewSite()
				s.SetId(ptr.To("id"))
				s.SetWebUrl(ptr.To("sharepoint.com/search"))
				return s
			}(),
			errCheck:       assert.Error,
			errIsSkippable: true,
		},
		{
			name: "Personal OneDrive",
			args: func() *models.Site {
				s := models.NewSite()
				s.SetId(ptr.To("id"))
				s.SetWebUrl(ptr.To("https://" + api.PersonalSitePath + "/someone's/onedrive"))
				return s
			}(),
			errCheck:       assert.Error,
			errIsSkippable: true,
		},
		{
			name:     "Valid Site",
			args:     site,
			errCheck: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := api.ValidateSite(test.args)
			test.errCheck(t, err, clues.ToCore(err))

			if test.errIsSkippable {
				assert.ErrorIs(t, err, api.ErrKnownSkippableCase)
			}
		})
	}
}

type SitesIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSitesIntgSuite(t *testing.T) {
	suite.Run(t, &SitesIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SitesIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SitesIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	sites, err := suite.its.ac.
		Sites().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(sites), "must have at least one site")

	for _, site := range sites {
		assert.NotContains(t, ptr.Val(site.GetWebUrl()), api.PersonalSitePath, "must not return onedrive sites")
	}
}

func (suite *SitesIntgSuite) TestSites_GetByID() {
	var (
		t               = suite.T()
		siteID          = tconfig.M365SiteID(t)
		host            = strings.Split(siteID, ",")[0]
		shortID         = strings.TrimPrefix(siteID, host+",")
		siteURL         = tconfig.M365SiteURL(t)
		modifiedSiteURL = siteURL + "foo"
	)

	sitesAPI := suite.its.ac.Sites()

	table := []struct {
		name      string
		id        string
		expectErr func(*testing.T, error)
	}{
		{
			name: "3 part id",
			id:   siteID,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "2 part id",
			id:   shortID,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "malformed id",
			id:   uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "random id",
			id:   uuid.NewString() + "," + uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
		{
			name: "url",
			id:   siteURL,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "host only",
			id:   host,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "malformed url",
			id:   "barunihlda",
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "well formed url, invalid hostname",
			id:   "https://test/sites/testing",
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "well formed url, no sites match",
			id:   modifiedSiteURL,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := sitesAPI.GetByID(ctx, test.id)
			test.expectErr(t, err)
		})
	}
}

func (suite *SitesIntgSuite) TestGetRoot() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	result, err := suite.its.ac.Sites().GetRoot(ctx)
	require.NoError(t, err)
	require.NotNil(t, result, "must find the root site")
	require.NotEmpty(t, ptr.Val(result.GetId()), "must have an id")
}
