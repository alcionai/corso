package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type OneDriveAPISuite struct {
	tester.Suite
	creds   account.M365Config
	service graph.Servicer
}

func (suite *OneDriveAPISuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
	adpt, err := graph.CreateAdapter(m365.AzureTenantID, m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err)

	suite.service = graph.NewService(adpt)
}

func TestOneDriveAPIs(t *testing.T) {
	suite.Run(t, &OneDriveAPISuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
			tester.CorsoGraphConnectorOneDriveTests),
	})
}

func (suite *OneDriveAPISuite) TestCreatePagerAndGetPage() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	siteID := tester.M365SiteID(t)
	pager := api.NewSiteDrivePager(suite.service, siteID, []string{"name"})
	a, err := pager.GetPage(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func (suite *OneDriveAPISuite) TestGetDriveIDByName() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	siteID := tester.M365SiteID(t)
	pager := api.NewSiteDrivePager(suite.service, siteID, []string{"id", "name"})
	id, err := pager.GetDriveIDByName(ctx, "Documents")
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func (suite *OneDriveAPISuite) TestGetDriveFolderByName() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	siteID := tester.M365SiteID(t)
	pager := api.NewSiteDrivePager(suite.service, siteID, []string{"id", "name"})
	id, err := pager.GetDriveIDByName(ctx, "Documents")
	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = pager.GetFolderIDByName(ctx, id, "folder")
	assert.NoError(t, err)
}
