package api_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *api.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	return api.NewBetaService(adapter)
}

type SharePointPageSuite struct {
	tester.Suite
	siteID  string
	creds   account.M365Config
	service *api.BetaService
}

func (suite *SharePointPageSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.siteID = tconfig.M365SiteID(t)
	a := tconfig.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365
	suite.service = createTestBetaService(t, suite.creds)
}

func TestSharePointPageSuite(t *testing.T) {
	suite.Run(t, &SharePointPageSuite{
		Suite: tester.NewIntegrationSuite(t, [][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointPageSuite) TestFetchPages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	pgs, err := api.FetchPages(ctx, suite.service, suite.siteID)
	assert.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, pgs)
	assert.NotZero(t, len(pgs))

	for _, entry := range pgs {
		t.Logf("id: %s\t name: %s\n", entry.ID, entry.Name)
	}
}

func (suite *SharePointPageSuite) TestGetSitePages() {
	t := suite.T()
	t.Skip("skipping until code is maintained again")

	ctx, flush := tester.NewContext(t)
	defer flush()

	tuples, err := api.FetchPages(ctx, suite.service, suite.siteID)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, tuples)

	jobs := []string{tuples[0].ID}
	pages, err := api.GetSitePages(ctx, suite.service, suite.siteID, jobs, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, pages)
}

func (suite *SharePointPageSuite) TestRestoreSinglePage() {
	t := suite.T()
	t.Skip("skipping until code is maintained again")

	ctx, flush := tester.NewContext(t)
	defer flush()

	destName := testdata.DefaultRestoreConfig("").Location
	testName := "MockPage"

	// Create Test Page
	//nolint:lll
	byteArray := spMock.Page("Byte Test")

	pageData, err := data.NewPrefetchedItem(
		io.NopCloser(bytes.NewReader(byteArray)),
		testName,
		time.Now())
	require.NoError(t, err, clues.ToCore(err))

	info, err := api.RestoreSitePage(
		ctx,
		suite.service,
		pageData,
		suite.siteID,
		destName)

	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, info)

	// Clean Up
	pageID := info.SharePoint.ParentPath

	err = api.DeleteSitePage(ctx, suite.service, suite.siteID, pageID)
	assert.NoError(t, err, clues.ToCore(err))
}
