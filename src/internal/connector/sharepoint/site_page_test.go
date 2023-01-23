package sharepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type SharePointPageSuite struct {
	suite.Suite

	creds account.M365Config
}

func (suite *SharePointPageSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
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
	siteID := tester.M365SiteID(t)
	service, err := createTestService(suite.creds)
	require.NoError(t, err)

	pgs, err := fetchPages(ctx, *service, siteID)
	assert.NoError(t, err)
	require.NotNil(t, pgs)
	assert.NotZero(t, len(pgs))

	for _, entry := range pgs {
		t.Logf("id: %s\t name: %s\n", entry.id, entry.name)
	}
}

func (suite *SharePointPageSuite) TestGetSitePage() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	siteID := tester.M365SiteID(t)

	service, err := createTestService(suite.creds)
	require.NoError(t, err)
	tuples, err := fetchPages(ctx, *service, siteID)
	require.NoError(t, err)
	require.NotNil(t, tuples)

	jobs := []string{tuples[0].id}
	pages, err := GetSitePage(ctx, service, siteID, jobs)
	assert.NoError(t, err)
	assert.NotEmpty(t, pages)
}
