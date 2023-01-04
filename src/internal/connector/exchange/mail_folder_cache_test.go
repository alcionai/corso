package exchange

import (
	stdpath "path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

const (
	// Need to use a hard-coded ID because GetAllFolderNamesForUser only gets
	// top-level folders right now.
	//nolint:lll
	testFolderID = "AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAABl7AqpAAA="

	//nolint:lll
	topFolderID = "AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEIAAA="
	// Full folder path for the folder above.
	expectedFolderPath = "toplevel/subFolder/subsubfolder"
)

type MailFolderCacheIntegrationSuite struct {
	suite.Suite
	credentials account.M365Config
}

func TestMailFolderCacheIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(MailFolderCacheIntegrationSuite))
}

func (suite *MailFolderCacheIntegrationSuite) SetupSuite() {
	t := suite.T()

	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
}

func (suite *MailFolderCacheIntegrationSuite) TestDeltaFetch() {
	suite.T().Skipf("Test depends on hardcoded folder names. Skipping till that is fixed")

	ctx, flush := tester.NewContext()
	defer flush()

	tests := []struct {
		name string
		root string
		path []string
	}{
		{
			name: "Default Root",
			root: rootFolderAlias,
		},
		{
			name: "Node Root",
			root: topFolderID,
		},
		{
			name: "Node Root Non-empty Path",
			root: topFolderID,
			path: []string{"some", "leading", "path"},
		},
	}
	userID := tester.M365UserID(suite.T())

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			ac, err := api.NewClient(suite.credentials)
			require.NoError(t, err)

			mfc := mailFolderCache{
				userID: userID,
				ac:     ac,
			}

			require.NoError(t, mfc.Populate(ctx, test.root, test.path...))

			p, err := mfc.IDToPath(ctx, testFolderID)
			require.NoError(t, err)
			t.Logf("Path: %s\n", p.String())

			expectedPath := stdpath.Join(append(test.path, expectedFolderPath)...)
			assert.Equal(t, expectedPath, p.String())
			identifier, ok := mfc.PathInCache(p.String())
			assert.True(t, ok)
			assert.NotEmpty(t, identifier)
		})
	}
}
