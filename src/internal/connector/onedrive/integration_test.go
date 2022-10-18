package onedrive

import (
	"testing"

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

func (suite *OneDriveIntegrationSuite) TestOneDriveDataCollections() {
	ctx, flush := tester.NewContext()
	defer flush()

	user := suite.userID
	scope := selectors.NewOneDriveBackup().Users([]string{user})[0]
	t := suite.T()
	service, err := NewOneDriveService(suite.creds)
	require.NoError(t, err)

	odcs, err := NewCollections(
		suite.creds.AzureTenantID,
		user,
		scope,
		service,
		service.updateStatus,
	).Get(ctx)

	require.NoError(t, err)

	for _, entry := range odcs {
		t.Log(entry.FullPath())
	}

	t.Log(service.status.String())

}
