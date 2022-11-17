package onedrive

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type OneDriveSuite struct {
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

	suite.Run(t, new(OneDriveSuite))
}

func (suite *OneDriveSuite) SetupSuite() {
	suite.userID = tester.SecondaryM365UserID(suite.T())
}

func (suite *OneDriveSuite) TestCreateGetDeleteFolder() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	gs := loadTestService(t)

	drives, err := drives(ctx, gs, suite.userID)
	require.NoError(t, err)
	require.NotEmpty(t, drives)

	driveID := *drives[0].GetId()

	names := []string{}
	ids := []string{}

	for i := 0; i < 2; i++ {
		name := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
		folderElements := []string{name}
		id, err := createRestoreFolders(ctx, gs, driveID, folderElements)
		require.NoError(t, err)

		ids = append(ids, id)
		names = append(names, name)

		defer func(fid string) {
			assert.NoError(t, DeleteItem(ctx, gs, driveID, fid))
		}(id)
	}

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
				for _, name := range names {
					if *f.GetName() == name {
						foundFolderIDs = append(foundFolderIDs, *f.GetId())
					}

					assert.True(t, strings.HasPrefix(*f.GetName(), test.prefix), "folder prefix")
				}
			}

			assert.ElementsMatch(t, ids, foundFolderIDs)
		})
	}
}

func (suite *OneDriveSuite) TestOneDriveNewCollections() {
	ctx, flush := tester.NewContext()
	defer flush()

	creds, err := tester.NewM365Account(suite.T()).M365Config()
	require.NoError(suite.T(), err)

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
				creds.AzureTenantID,
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
