package exchange

import (
	stdpath "path"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/its"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
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

type MailFolderCacheIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestMailFolderCacheIntegrationSuite(t *testing.T) {
	suite.Run(t, &MailFolderCacheIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MailFolderCacheIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *MailFolderCacheIntgSuite) TestDeltaFetch() {
	suite.T().Skipf("Test depends on hardcoded folder names. Skipping till that is fixed")

	tests := []struct {
		name string
		root string
		path []string
	}{
		{
			name: "Default Root",
			root: api.MsgFolderRoot,
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

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			acm := suite.m365.AC.Mail()

			mfc := mailContainerCache{
				userID: suite.m365.User.ID,
				enumer: acm,
				getter: acm,
			}

			err := mfc.Populate(ctx, fault.New(true), test.root, test.path...)
			require.NoError(t, err, clues.ToCore(err))

			p, l, err := mfc.IDToPath(ctx, testFolderID)
			require.NoError(t, err, clues.ToCore(err))
			t.Logf("Path: %s\n", p.String())
			t.Logf("Location: %s\n", l.String())

			expectedPath := stdpath.Join(append(test.path, expectedFolderPath)...)
			assert.Equal(t, expectedPath, p.String())
			identifier, ok := mfc.LocationInCache(p.String())
			assert.True(t, ok)
			assert.NotEmpty(t, identifier)
		})
	}
}
