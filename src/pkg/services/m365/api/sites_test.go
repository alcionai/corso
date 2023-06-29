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
	"github.com/alcionai/corso/src/internal/tester"
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
			[][]string{tester.M365AcctCredEnvs, tester.AWSStorageCredEnvs}),
	})
}

func (suite *SitesIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SitesIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	sites, err := suite.its.ac.Sites().GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(sites), "must have at least one site")

	for _, site := range sites {
		assert.NotContains(t, ptr.Val(site.GetWebUrl()), api.PersonalSitePath, "must not return onedrive sites")
	}
}

func (suite *SitesIntgSuite) TestSites_GetByID() {
	var (
		t       = suite.T()
		siteID  = tester.M365SiteID(t)
		host    = strings.Split(siteID, ",")[0]
		shortID = strings.TrimPrefix(siteID, host+",")
		siteURL = tester.M365SiteURL(t)
	)

	sitesAPI := suite.its.ac.Sites()

	table := []struct {
		name      string
		id        string
		expectErr assert.ErrorAssertionFunc
	}{
		{"3 part id", siteID, assert.NoError},
		{"2 part id", shortID, assert.NoError},
		{"malformed id", uuid.NewString(), assert.Error},
		{"random id", uuid.NewString() + "," + uuid.NewString(), assert.Error},
		{"url", siteURL, assert.NoError},
		{"host only", host, assert.NoError},
		{"malformed url", "barunihlda", assert.Error},
		{"non-matching url", "https://test/sites/testing", assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(t)
			defer flush()

			t := suite.T()

			_, err := sitesAPI.GetByID(ctx, test.id)
			test.expectErr(t, err, clues.ToCore(err))
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
