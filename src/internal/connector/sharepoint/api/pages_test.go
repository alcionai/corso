package api_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/sharepoint/api"
	spMock "github.com/alcionai/corso/src/internal/connector/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	m365api "github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SharePointPageSuite struct {
	tester.Suite
	siteID  string
	creds   account.M365Config
	service *m365api.BetaService
}

func (suite *SharePointPageSuite) SetupSuite() {
	t := suite.T()

	suite.siteID = tester.M365SiteID(t)
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365
	suite.service = createTestBetaService(t, suite.creds)
}

func TestSharePointPageSuite(t *testing.T) {
	suite.Run(t, &SharePointPageSuite{
		Suite: tester.NewIntegrationSuite(t, [][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *SharePointPageSuite) TestFetchPages() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	pgs, err := api.FetchPages(ctx, suite.service, suite.siteID)
	assert.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, pgs)
	assert.NotZero(t, len(pgs))

	for _, entry := range pgs {
		t.Logf("id: %s\t name: %s\n", entry.ID, entry.Name)
	}
}

func (suite *SharePointPageSuite) TestGetSitePages() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	tuples, err := api.FetchPages(ctx, suite.service, suite.siteID)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, tuples)

	jobs := []string{tuples[0].ID}
	pages, err := api.GetSitePages(ctx, suite.service, suite.siteID, jobs, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, pages)
}

func (suite *SharePointPageSuite) TestRestoreSinglePage() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	destName := tester.DefaultTestRestoreDestination("").ContainerName
	testName := "MockPage"

	// Create Test Page
	//nolint:lll
	byteArray := spMock.Page("Byte Test")

	pageData := sharepoint.NewItem(
		testName,
		io.NopCloser(bytes.NewReader(byteArray)),
	)

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
