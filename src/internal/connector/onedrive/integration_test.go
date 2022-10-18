package onedrive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type OneDriveIntegrationSuite struct {
	suite.Suite
	userID string
	creds  account.M365Config
}

func TestOneDriveIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoOneDriveTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(OneDriveIntegrationSuite))
}

func (suite *OneDriveIntegrationSuite) SetupSuite() {
	t := suite.T()
	suite.userID = tester.M365UserID(t)
	a := tester.NewM365Account(t)
	credentials, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = credentials
}

func (suite *OneDriveIntegrationSuite) TestOneDriveNewCollections() {
	ctx, flush := tester.NewContext()
	defer flush()

	tests := []struct {
		name, user string
	}{
		{
			name: "Test User w/ Drive",
			user: suite.userID,
		},
		{
			name: "Test User w/out Drive",
			user: "testevents@8qzvrj.onmicrosoft.com",
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			service := loadTestService(t)
			scope := selectors.
				NewOneDriveBackup().
				Users([]string{test.user})[0]
			odcs, err := NewCollections(
				suite.creds.AzureTenantID,
				test.user,
				scope,
				service,
				service.updateStatus,
			).Get(ctx)
			assert.NoError(t, err)

			for _, entry := range odcs {
				assert.NotEmpty(t, entry.FullPath())
			}
		})
	}
}
