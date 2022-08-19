package onedrive

import (
	"context"
	"io"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/tester"
)

type ItemIntegrationSuite struct {
	suite.Suite
	client  *msgraphsdk.GraphServiceClient
	adapter *msgraphsdk.GraphRequestAdapter
}

func (suite *ItemIntegrationSuite) Client() *msgraphsdk.GraphServiceClient {
	return suite.client
}

func (suite *ItemIntegrationSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return suite.adapter
}

func (suite *ItemIntegrationSuite) ErrPolicy() bool {
	return false
}

func TestItemIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(ItemIntegrationSuite))
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a := tester.NewM365Account(suite.T())

	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	adapter, err := graph.CreateAdapter(m365.TenantID, m365.ClientID, m365.ClientSecret)
	require.NoError(suite.T(), err)
	suite.client = msgraphsdk.NewGraphServiceClient(adapter)
	suite.adapter = adapter
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader() {
	ctx := context.TODO()
	user := tester.M365UserID(suite.T())

	drives, err := drives(ctx, suite, user)
	require.NoError(suite.T(), err)
	// Test Requirement 1: Need a drive
	require.Greaterf(suite.T(), len(drives), 0, "user %s does not have a drive", user)

	// Pick the first drive
	driveID := *drives[0].GetId()

	var driveItemID string
	// This item collector tries to find "a" drive item that is a file to test the reader function
	itemCollector := func(ctx context.Context, driveID string, items []models.DriveItemable) error {
		for _, item := range items {
			if item.GetFile() != nil {
				driveItemID = *item.GetId()
				break
			}
		}
		return nil
	}
	err = collectItems(ctx, suite, driveID, itemCollector)
	require.NoError(suite.T(), err)

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItemID,
		"no file item found for user %s drive %s",
		user,
		driveID,
	)

	// Read data for the file

	name, itemData, err := driveItemReader(ctx, suite, driveID, driveItemID)
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), name)
	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err)
	require.NotZero(suite.T(), size)
	suite.T().Logf("Read %d bytes from file %s.", size, name)
}
