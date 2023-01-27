package onedrive

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type OneDriveSuite struct {
	suite.Suite
	userID string
}

func TestOneDriveDriveSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoOneDriveTests)

	suite.Run(t, new(OneDriveSuite))
}

func (suite *OneDriveSuite) SetupSuite() {
	suite.userID = tester.SecondaryM365UserID(suite.T())
}

func (suite *OneDriveSuite) TestCreateGetDeleteFolder() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	folderIDs := []string{}
	folderName1 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements := []string{folderName1}
	gs := loadTestService(t)

	pager, err := PagerForSource(OneDriveSource, gs, suite.userID, nil)
	require.NoError(t, err)

	drives, err := drives(ctx, pager, true)
	require.NoError(t, err)
	require.NotEmpty(t, drives)

	// TODO: Verify the intended drive
	driveID := *drives[0].GetId()

	defer func() {
		for _, id := range folderIDs {
			err := DeleteItem(ctx, gs, driveID, id)
			if err != nil {
				logger.Ctx(ctx).Warnw("deleting folder", "id", id, "error", err)
			}
		}
	}()

	folderID, err := CreateRestoreFolders(ctx, gs, driveID, folderElements)
	require.NoError(t, err)

	folderIDs = append(folderIDs, folderID)

	folderName2 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements = append(folderElements, folderName2)

	folderID, err = CreateRestoreFolders(ctx, gs, driveID, folderElements)
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
			pager, err := PagerForSource(OneDriveSource, gs, suite.userID, nil)
			require.NoError(t, err)

			allFolders, err := GetAllFolders(ctx, gs, pager, test.prefix)
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

type testFolderMatcher struct {
	scope selectors.OneDriveScope
}

func (fm testFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.OneDriveFolder)
}

func (fm testFolderMatcher) Matches(path string) bool {
	return fm.scope.Matches(selectors.OneDriveFolder, path)
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
				NewOneDriveBackup([]string{test.user}).
				AllData()[0]
			odcs, err := NewCollections(
				graph.LargeItemClient(),
				creds.AzureTenantID,
				test.user,
				OneDriveSource,
				testFolderMatcher{scope},
				service,
				service.updateStatus,
				control.Options{},
			).Get(ctx)
			assert.NoError(t, err)

			for _, entry := range odcs {
				assert.NotEmpty(t, entry.FullPath())
			}
		})
	}
}
