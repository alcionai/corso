package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type ItemIntegrationSuite struct {
	tester.Suite
	user        string
	userDriveID string
	service     graph.Servicer
}

func TestItemIntegrationSuite(t *testing.T) {
	suite.Run(t, &ItemIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	suite.service = loadTestService(t)
	suite.user = tester.SecondaryM365UserID(t)

	pager, err := PagerForSource(OneDriveSource, suite.service, suite.user, nil)
	require.NoError(t, err)

	odDrives, err := drives(ctx, pager, true)
	require.NoError(t, err)
	// Test Requirement 1: Need a drive
	require.Greaterf(t, len(odDrives), 0, "user %s does not have a drive", suite.user)

	// Pick the first drive
	suite.userDriveID = ptr.Val(odDrives[0].GetId())
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader_oneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var driveItem models.DriveItemable
	// This item collector tries to find "a" drive item that is a file to test the reader function
	itemCollector := func(
		_ context.Context,
		_, _ string,
		items []models.DriveItemable,
		_ map[string]string,
		_ map[string]string,
		_ map[string]struct{},
		_ map[string]map[string]string,
		_ bool,
		_ *fault.Bus,
	) error {
		for _, item := range items {
			if item.GetFile() != nil {
				driveItem = item
				break
			}
		}

		return nil
	}
	_, _, _, err := collectItems(
		ctx,
		defaultItemPager(
			suite.service,
			suite.userDriveID,
			""),
		suite.userDriveID,
		"General",
		itemCollector,
		map[string]string{},
		"",
		fault.New(true))
	require.NoError(suite.T(), err)

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItem,
		"no file item found for user %s drive %s",
		suite.user,
		suite.userDriveID,
	)

	// Read data for the file
	itemInfo, itemData, err := oneDriveItemReader(ctx, graph.HTTPClient(graph.NoTimeout()), driveItem)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), itemInfo.OneDrive)
	require.NotEmpty(suite.T(), itemInfo.OneDrive.ItemName)

	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err)
	require.NotZero(suite.T(), size)
	require.Equal(suite.T(), size, itemInfo.OneDrive.Size)
	suite.T().Logf("Read %d bytes from file %s.", size, itemInfo.OneDrive.ItemName)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new `testfolder_<timestamp` folder with a new
// testitem_<timestamp> item and writes data to it
func (suite *ItemIntegrationSuite) TestItemWriter() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			srv := suite.service

			root, err := srv.Client().DrivesById(test.driveID).Root().Get(ctx, nil)
			require.NoError(t, err)

			// Test Requirement 2: "Test Folder" should exist
			folder, err := getFolder(ctx, srv, test.driveID, ptr.Val(root.GetId()), "Test Folder")
			require.NoError(t, err)

			newFolderName := "testfolder_" + common.FormatNow(common.SimpleTimeTesting)
			t.Logf("Test will create folder %s", newFolderName)

			newFolder, err := CreateItem(
				ctx,
				srv,
				test.driveID,
				ptr.Val(folder.GetId()),
				newItem(newFolderName, true))
			require.NoError(t, err)
			require.NotNil(t, newFolder.GetId())

			newItemName := "testItem_" + common.FormatNow(common.SimpleTimeTesting)
			t.Logf("Test will create item %s", newItemName)

			newItem, err := CreateItem(
				ctx,
				srv,
				test.driveID,
				ptr.Val(newFolder.GetId()),
				newItem(newItemName, false))
			require.NoError(t, err)
			require.NotNil(t, newItem.GetId())

			// HACK: Leveraging this to test getFolder behavior for a file. `getFolder()` on the
			// newly created item should fail because it's a file not a folder
			_, err = getFolder(ctx, srv, test.driveID, ptr.Val(newFolder.GetId()), newItemName)
			require.ErrorIs(t, err, errFolderNotFound)

			// Initialize a 100KB mockDataProvider
			td, writeSize := mockDataReader(int64(100 * 1024))

			w, err := driveItemWriter(ctx, srv, test.driveID, ptr.Val(newItem.GetId()), writeSize)
			require.NoError(t, err)

			// Using a 32 KB buffer for the copy allows us to validate the
			// multi-part upload. `io.CopyBuffer` will only write 32 KB at
			// a time
			copyBuffer := make([]byte, 32*1024)

			size, err := io.CopyBuffer(w, td, copyBuffer)
			require.NoError(t, err)

			require.Equal(t, writeSize, size)
		})
	}
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return bytes.NewReader(data), size
}

func (suite *ItemIntegrationSuite) TestDriveGetFolder() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "oneDrive",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			srv := suite.service

			root, err := srv.Client().DrivesById(test.driveID).Root().Get(ctx, nil)
			require.NoError(t, err)

			// Lookup a folder that doesn't exist
			_, err = getFolder(ctx, srv, test.driveID, ptr.Val(root.GetId()), "FolderDoesNotExist")
			require.ErrorIs(t, err, errFolderNotFound)

			// Lookup a folder that does exist
			_, err = getFolder(ctx, srv, test.driveID, ptr.Val(root.GetId()), "")
			require.NoError(t, err)
		})
	}
}

func getPermsUperms(permID, userID, entity string, scopes []string) (models.Permissionable, UserPermission) {
	identity := models.NewIdentity()
	identity.SetId(&userID)
	identity.SetAdditionalData(map[string]any{"email": &userID})

	sharepointIdentity := models.NewSharePointIdentitySet()

	switch entity {
	case "user":
		sharepointIdentity.SetUser(identity)
	case "group":
		sharepointIdentity.SetGroup(identity)
	case "application":
		sharepointIdentity.SetApplication(identity)
	case "device":
		sharepointIdentity.SetDevice(identity)
	}

	perm := models.NewPermission()
	perm.SetId(&permID)
	perm.SetRoles([]string{"read"})
	perm.SetGrantedToV2(sharepointIdentity)

	uperm := UserPermission{
		ID:       permID,
		Roles:    []string{"read"},
		EntityID: userID,
	}

	return perm, uperm
}

type ItemUnitTestSuite struct {
	tester.Suite
}

func TestItemUnitTestSuite(t *testing.T) {
	suite.Run(t, &ItemUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitTestSuite) TestOneDrivePermissionsFilter() {
	permID := "fakePermId"
	userID := "fakeuser@provider.com"
	userID2 := "fakeuser2@provider.com"

	userReadPerm, userReadUperm := getPermsUperms(permID, userID, "user", []string{"read"})
	userReadWritePerm, userReadWriteUperm := getPermsUperms(permID, userID2, "user", []string{"read", "write"})

	groupReadPerm, groupReadUperm := getPermsUperms(permID, userID, "group", []string{"read"})
	groupReadWritePerm, groupReadWriteUperm := getPermsUperms(permID, userID2, "group", []string{"read", "write"})

	noPerm, _ := getPermsUperms(permID, userID, "user", []string{"read"})
	noPerm.SetGrantedToV2(nil) // eg: link shares

	cases := []struct {
		name              string
		graphPermissions  []models.Permissionable
		parsedPermissions []UserPermission
	}{
		{
			name:              "no perms",
			graphPermissions:  []models.Permissionable{},
			parsedPermissions: []UserPermission{},
		},
		{
			name:              "no user bound to perms",
			graphPermissions:  []models.Permissionable{noPerm},
			parsedPermissions: []UserPermission{},
		},

		// user
		{
			name:              "user with read permissions",
			graphPermissions:  []models.Permissionable{userReadPerm},
			parsedPermissions: []UserPermission{userReadUperm},
		},
		{
			name:              "user with read and write permissions",
			graphPermissions:  []models.Permissionable{userReadWritePerm},
			parsedPermissions: []UserPermission{userReadWriteUperm},
		},
		{
			name:              "multiple users with separate permissions",
			graphPermissions:  []models.Permissionable{userReadPerm, userReadWritePerm},
			parsedPermissions: []UserPermission{userReadUperm, userReadWriteUperm},
		},

		// group
		{
			name:              "group with read permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm},
			parsedPermissions: []UserPermission{groupReadUperm},
		},
		{
			name:              "group with read and write permissions",
			graphPermissions:  []models.Permissionable{groupReadWritePerm},
			parsedPermissions: []UserPermission{groupReadWriteUperm},
		},
		{
			name:              "multiple groups with separate permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm, groupReadWritePerm},
			parsedPermissions: []UserPermission{groupReadUperm, groupReadWriteUperm},
		},
	}
	for _, tc := range cases {
		suite.Run(tc.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			actual := filterUserPermissions(ctx, tc.graphPermissions)
			assert.ElementsMatch(suite.T(), tc.parsedPermissions, actual)
		})
	}
}
