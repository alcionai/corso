package api

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
	siteID string
	creds  account.M365Config
}

func (suite *SharePointPageSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	suite.siteID = tester.M365SiteID(t)
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
	service := createTestBetaService(t, suite.creds)

	pgs, err := FetchPages(ctx, service, suite.siteID)
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
	service := createTestBetaService(t, suite.creds)
	tuples, err := FetchPages(ctx, service, suite.siteID)
	require.NoError(t, err)
	require.NotNil(t, tuples)

	jobs := []string{tuples[0].ID}
	pages, err := GetSitePage(ctx, service, suite.siteID, jobs)
	assert.NoError(t, err)
	assert.NotEmpty(t, pages)
}
