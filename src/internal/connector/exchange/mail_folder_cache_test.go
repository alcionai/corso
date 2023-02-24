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
	"github.com/alcionai/corso/src/pkg/fault"
)

const (
	// Need to use a hard-coded ID because GetAllFolderNamesForUser only gets
	// top-level folders right now.
	//nolint:lll
	testFolderID = "AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAABl7AqpAAA="
	//nolint:lll
	topFolderID = "AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEIAAA="
	//nolint:lll
	// Full folder path for the folder above.
	expectedFolderPath = "toplevel/subFolder/subsubfolder"
)

type MailFolderCacheIntegrationSuite struct {
	tester.Suite
	credentials account.M365Config
}

func TestMailFolderCacheIntegrationSuite(t *testing.T) {
	suite.Run(t, &MailFolderCacheIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *MailFolderCacheIntegrationSuite) SetupSuite() {
	t := suite.T()

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
		suite.Run(test.name, func() {
			t := suite.T()

			ac, err := api.NewClient(suite.credentials)
			require.NoError(t, err)

			acm := ac.Mail()

			mfc := mailFolderCache{
				userID: userID,
				enumer: acm,
				getter: acm,
			}

			require.NoError(t, mfc.Populate(ctx, fault.New(true), test.root, test.path...))

			p, l, err := mfc.IDToPath(ctx, testFolderID, true)
			require.NoError(t, err)
			t.Logf("Path: %s\n", p.String())
			t.Logf("Location: %s\n", l.String())

			expectedPath := stdpath.Join(append(test.path, expectedFolderPath)...)
			assert.Equal(t, expectedPath, p.String())
			identifier, ok := mfc.PathInCache(p.String())
			assert.True(t, ok)
			assert.NotEmpty(t, identifier)
		})
	}
}
