package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type OneDriveAPISuite struct {
	tester.Suite
	creds account.M365Config
	ac    api.Client
}

func (suite *OneDriveAPISuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = creds
	suite.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))
}

func TestOneDriveAPIs(t *testing.T) {
	suite.Run(t, &OneDriveAPISuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveAPISuite) TestCreatePagerAndGetPage() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	siteID := tester.M365SiteID(t)
	pager := suite.ac.Drives().NewSiteDrivePager(siteID, []string{"name"})

	a, err := pager.GetPage(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, a)
}
