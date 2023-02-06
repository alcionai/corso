package api_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/sharepoint/api"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type SharePointPageSuite struct {
	suite.Suite
	siteID  string
	creds   account.M365Config
	service *discover.BetaService
}

func (suite *SharePointPageSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	suite.siteID = tester.M365SiteID(t)
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
	suite.service = createTestBetaService(t, suite.creds)
}

func TestSharePointPageSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorSharePointTests)
	suite.Run(t, new(SharePointPageSuite))
}

func (suite *SharePointPageSuite) TestFetchPages() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	pgs, err := api.FetchPages(ctx, suite.service, suite.siteID)
	assert.NoError(t, err)
	require.NotNil(t, pgs)
	assert.NotZero(t, len(pgs))

	for _, entry := range pgs {
		t.Logf("id: %s\t name: %s\n", entry.ID, entry.Name)
	}
}

func (suite *SharePointPageSuite) TestGetSitePage() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	tuples, err := api.FetchPages(ctx, suite.service, suite.siteID)
	require.NoError(t, err)
	require.NotNil(t, tuples)

	jobs := []string{tuples[0].ID}
	pages, err := api.GetSitePage(ctx, suite.service, suite.siteID, jobs)
	assert.NoError(t, err)
	assert.NotEmpty(t, pages)
}

func (suite *SharePointPageSuite) TestRestoreSinglePage() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	destName := "Corso_Restore_" + common.FormatNow(common.SimpleTimeTesting)
	testName := "MockPage"

	// Create Test Page
	//nolint:lll
	byteArray := mockconnector.GetMockPage("Byte Test")

	pageData := sharepoint.NewItem(
		testName,
		io.NopCloser(bytes.NewReader(byteArray)),
	)

	info, err := api.RestoreSitePage(
		ctx,
		suite.service,
		pageData,
		suite.siteID,
		destName,
	)

	require.NoError(t, err)
	require.NotNil(t, info)

	// Clean Up
	pageID := info.SharePoint.ParentPath
	err = api.DeleteSitePage(ctx, suite.service, suite.siteID, pageID)
	assert.NoError(t, err)
}
