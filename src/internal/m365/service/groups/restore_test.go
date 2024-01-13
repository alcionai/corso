package groups

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestConsumeRestoreCollections_noErrorOnGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rcc := inject.RestoreConsumerConfig{}
	pth, err := path.Builder{}.
		Append("General").
		ToDataLayerPath(
			"t",
			"g",
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
	require.NoError(t, err, clues.ToCore(err))

	dcs := []data.RestoreCollection{
		mock.Collection{Path: pth},
	}

	_, _, err = NewGroupsHandler(api.Client{}, nil).
		ConsumeRestoreCollections(
			ctx,
			rcc,
			dcs,
			fault.New(false),
			nil)
	assert.NoError(t, err, "Groups Channels restore")
}

type groupsIntegrationSuite struct {
	tester.Suite
	resource string
	tenantID string
	ac       api.Client
}

func TestGroupsIntegrationSuite(t *testing.T) {
	suite.Run(t, &groupsIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *groupsIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.resource = tconfig.M365TeamID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID
}

// test for getSiteName
func (suite *groupsIntegrationSuite) TestGetSiteName() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rootSite, err := suite.ac.Groups().GetRootSite(ctx, suite.resource)
	require.NoError(t, err, clues.ToCore(err))

	// Generate a fake site ID that appears valid to graph API but doesn't actually exist.
	// This "could" be flaky, but highly unlikely
	unavailableSiteID := []rune(ptr.Val(rootSite.GetId()))
	firstIDChar := slices.Index(unavailableSiteID, ',') + 1

	if unavailableSiteID[firstIDChar] != '2' {
		unavailableSiteID[firstIDChar] = '2'
	} else {
		unavailableSiteID[firstIDChar] = '1'
	}

	tests := []struct {
		name              string
		siteID            string
		webURL            string
		siteName          string
		webURLToSiteNames map[string]string
		expectErr         assert.ErrorAssertionFunc
	}{
		{
			name:              "valid",
			siteID:            ptr.Val(rootSite.GetId()),
			webURL:            ptr.Val(rootSite.GetWebUrl()),
			siteName:          *rootSite.GetDisplayName(),
			webURLToSiteNames: map[string]string{},
			expectErr:         assert.NoError,
		},
		{
			name:              "unavailable",
			siteID:            string(unavailableSiteID),
			webURL:            "https://does-not-matter",
			siteName:          "",
			webURLToSiteNames: map[string]string{},
			expectErr:         assert.NoError,
		},
		{
			name:              "previously found",
			siteID:            "random-id",
			webURL:            "https://random-url",
			siteName:          "random-name",
			webURLToSiteNames: map[string]string{"https://random-url": "random-name"},
			expectErr:         assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			siteName, err := getSiteName(
				ctx,
				test.siteID,
				test.webURL,
				suite.ac.Sites(),
				test.webURLToSiteNames)
			require.NoError(t, err, clues.ToCore(err))

			test.expectErr(t, err)
			assert.Equal(t, test.siteName, siteName)
		})
	}
}
