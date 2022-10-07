package onedrive

import (
	"context"
	"strings"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

// TODO(ashmrtn): Merge with similar structs in graph and exchange packages.
type testService struct {
	adapter     msgraphsdk.GraphRequestAdapter
	client      msgraphsdk.GraphServiceClient
	failFast    bool
	credentials account.M365Config
}

func (ts *testService) Client() *msgraphsdk.GraphServiceClient {
	return &ts.client
}

func (ts *testService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &ts.adapter
}

func (ts *testService) ErrPolicy() bool {
	return ts.failFast
}

// TODO(ashmrtn): Merge with similar functions in connector and exchange
// packages.
func loadService(t *testing.T) *testService {
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	adapter, err := graph.CreateAdapter(
		m365.TenantID,
		m365.ClientID,
		m365.ClientSecret,
	)
	require.NoError(t, err)

	service := &testService{
		adapter:     *adapter,
		client:      *msgraphsdk.NewGraphServiceClient(adapter),
		failFast:    false,
		credentials: m365,
	}

	return service
}

type OneDriveDriveSuite struct {
	suite.Suite
	userID string
}

func TestOneDriveDriveSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoOneDriveTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(OneDriveDriveSuite))
}

func (suite *OneDriveDriveSuite) SetupSuite() {
	suite.userID = tester.M365UserID(suite.T())
}

func (suite *OneDriveDriveSuite) TestCreateGetDeleteFolder() {
	t := suite.T()
	ctx := context.Background()
	folderIDs := []string{}
	folderName1 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements := []string{folderName1}
	gs := loadService(t)

	drives, err := drives(ctx, gs, suite.userID)
	require.NoError(t, err)
	require.NotEmpty(t, drives)

	driveID := *drives[0].GetId()

	folderID, err := createRestoreFolders(ctx, gs, driveID, folderElements)
	require.NoError(t, err)

	folderIDs = append(folderIDs, folderID)

	defer func() {
		assert.NoError(t, DeleteItem(ctx, gs, driveID, folderIDs[0]))
	}()

	folderName2 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements = append(folderElements, folderName2)

	folderID, err = createRestoreFolders(ctx, gs, driveID, folderElements)
	require.NoError(t, err)

	folderIDs = append(folderIDs, folderID)

	table := []struct {
		name   string
		prefix string
	}{
		{
			name:   "NoPrefix",
			prefix: "",
		},
		{
			name:   "Prefix",
			prefix: "Corso_Folder_Test",
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			allFolders, err := GetAllFolders(ctx, gs, suite.userID, test.prefix)
			require.NoError(t, err)

			foundFolderIDs := []string{}

			for _, f := range allFolders {
				if *f.GetName() == folderName1 || *f.GetName() == folderName2 {
					foundFolderIDs = append(foundFolderIDs, *f.GetId())
				}

				assert.True(t, strings.HasPrefix(*f.GetName(), test.prefix), "folder prefix")
			}

			assert.ElementsMatch(t, folderIDs, foundFolderIDs)
		})
	}
}
