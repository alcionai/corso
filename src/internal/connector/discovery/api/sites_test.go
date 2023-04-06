package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
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
		args           any
		want           models.Siteable
		errCheck       assert.ErrorAssertionFunc
		errIsSkippable bool
	}{
		{
			name:     "Invalid type",
			args:     string("invalid type"),
			errCheck: assert.Error,
		},
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
				s.SetWebUrl(ptr.To("https://" + personalSitePath + "/someone's/onedrive"))
				return s
			}(),
			errCheck:       assert.Error,
			errIsSkippable: true,
		},
		{
			name:     "Valid Site",
			args:     site,
			want:     site,
			errCheck: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			got, err := validateSite(test.args)
			test.errCheck(t, err, clues.ToCore(err))

			if test.errIsSkippable {
				assert.ErrorIs(t, err, errKnownSkippableCase)
			}

			assert.Equal(t, test.want, got)
		})
	}
}

type SitesIntgSuite struct {
	tester.Suite

	creds account.M365Config
}

func TestSitesIntgSuite(t *testing.T) {
	suite.Run(t, &SitesIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs, tester.AWSStorageCredEnvs}),
	})
}

func (suite *SitesIntgSuite) SetupSuite() {
	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
	)

	m365, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365
}

func (suite *SitesIntgSuite) TestGetAll() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	cli, err := NewClient(suite.creds)
	require.NoError(t, err, clues.ToCore(err))

	sites, err := cli.Sites().GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(sites), "must have at least one site")

	for _, site := range sites {
		assert.NotContains(t, ptr.Val(site.GetWebUrl()), personalSitePath, "must not return onedrive sites")
	}
}
