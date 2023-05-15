package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
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
	require.NoError(t, err, clues.ToCore(err))

	odDrives, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	require.NoError(t, err, clues.ToCore(err))
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
	require.NoError(suite.T(), err, clues.ToCore(err))

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItem,
		"no file item found for user %s drive %s",
		suite.user,
		suite.userDriveID,
	)

	// Read data for the file
	itemInfo, itemData, err := oneDriveItemReader(ctx, graph.NewNoTimeoutHTTPWrapper(), driveItem)

	require.NoError(suite.T(), err, clues.ToCore(err))
	require.NotNil(suite.T(), itemInfo.OneDrive)
	require.NotEmpty(suite.T(), itemInfo.OneDrive.ItemName)

	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err, clues.ToCore(err))
	require.NotZero(suite.T(), size)
	require.Equal(suite.T(), size, itemInfo.OneDrive.Size)

	suite.T().Logf("Read %d bytes from file %s.", size, itemInfo.OneDrive.ItemName)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new folder with a new item and writes data to it
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
			require.NoError(t, err, clues.ToCore(err))

			newFolderName := tester.DefaultTestRestoreDestination("folder").ContainerName
			t.Logf("creating folder %s", newFolderName)

			newFolder, err := CreateItem(
				ctx,
				srv,
				test.driveID,
				ptr.Val(root.GetId()),
				newItem(newFolderName, true))
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, newFolder.GetId())

			newItemName := "testItem_" + dttm.FormatNow(dttm.SafeForTesting)
			t.Logf("creating item %s", newItemName)

			newItem, err := CreateItem(
				ctx,
				srv,
				test.driveID,
				ptr.Val(newFolder.GetId()),
				newItem(newItemName, false))
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, newItem.GetId())

			// HACK: Leveraging this to test getFolder behavior for a file. `getFolder()` on the
			// newly created item should fail because it's a file not a folder
			_, err = api.GetFolderByName(ctx, srv, test.driveID, ptr.Val(newFolder.GetId()), newItemName)
			require.ErrorIs(t, err, api.ErrFolderNotFound, clues.ToCore(err))

			// Initialize a 100KB mockDataProvider
			td, writeSize := mockDataReader(int64(100 * 1024))

			w, err := driveItemWriter(ctx, srv, test.driveID, ptr.Val(newItem.GetId()), writeSize)
			require.NoError(t, err, clues.ToCore(err))

			// Using a 32 KB buffer for the copy allows us to validate the
			// multi-part upload. `io.CopyBuffer` will only write 32 KB at
			// a time
			copyBuffer := make([]byte, 32*1024)

			size, err := io.CopyBuffer(w, td, copyBuffer)
			require.NoError(t, err, clues.ToCore(err))

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
			require.NoError(t, err, clues.ToCore(err))

			// Lookup a folder that doesn't exist
			_, err = api.GetFolderByName(ctx, srv, test.driveID, ptr.Val(root.GetId()), "FolderDoesNotExist")
			require.ErrorIs(t, err, api.ErrFolderNotFound, clues.ToCore(err))

			// Lookup a folder that does exist
			_, err = api.GetFolderByName(ctx, srv, test.driveID, ptr.Val(root.GetId()), "")
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

func getPermsAndResourceOwnerPerms(
	permID, resourceOwner string,
	gv2t metadata.GV2Type,
	scopes []string,
) (models.Permissionable, metadata.Permission) {
	sharepointIdentitySet := models.NewSharePointIdentitySet()

	switch gv2t {
	case metadata.GV2App, metadata.GV2Device, metadata.GV2Group, metadata.GV2User:
		identity := models.NewIdentity()
		identity.SetId(&resourceOwner)
		identity.SetAdditionalData(map[string]any{"email": &resourceOwner})

		switch gv2t {
		case metadata.GV2User:
			sharepointIdentitySet.SetUser(identity)
		case metadata.GV2Group:
			sharepointIdentitySet.SetGroup(identity)
		case metadata.GV2App:
			sharepointIdentitySet.SetApplication(identity)
		case metadata.GV2Device:
			sharepointIdentitySet.SetDevice(identity)
		}

	case metadata.GV2SiteUser, metadata.GV2SiteGroup:
		spIdentity := models.NewSharePointIdentity()
		spIdentity.SetId(&resourceOwner)
		spIdentity.SetAdditionalData(map[string]any{"email": &resourceOwner})

		switch gv2t {
		case metadata.GV2SiteUser:
			sharepointIdentitySet.SetSiteUser(spIdentity)
		case metadata.GV2SiteGroup:
			sharepointIdentitySet.SetSiteGroup(spIdentity)
		}
	}

	perm := models.NewPermission()
	perm.SetId(&permID)
	perm.SetRoles([]string{"read"})
	perm.SetGrantedToV2(sharepointIdentitySet)

	ownersPerm := metadata.Permission{
		ID:         permID,
		Roles:      []string{"read"},
		EntityID:   resourceOwner,
		EntityType: gv2t,
	}

	return perm, ownersPerm
}

type ItemUnitTestSuite struct {
	tester.Suite
}

func TestItemUnitTestSuite(t *testing.T) {
	suite.Run(t, &ItemUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitTestSuite) TestDrivePermissionsFilter() {
	var (
		pID  = "fakePermId"
		uID  = "fakeuser@provider.com"
		uID2 = "fakeuser2@provider.com"
		own  = []string{"owner"}
		r    = []string{"read"}
		rw   = []string{"read", "write"}
	)

	userOwnerPerm, userOwnerROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2User, own)
	userReadPerm, userReadROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2User, r)
	userReadWritePerm, userReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, metadata.GV2User, rw)
	siteUserOwnerPerm, siteUserOwnerROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2SiteUser, own)
	siteUserReadPerm, siteUserReadROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2SiteUser, r)
	siteUserReadWritePerm, siteUserReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, metadata.GV2SiteUser, rw)

	groupReadPerm, groupReadROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2Group, r)
	groupReadWritePerm, groupReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, metadata.GV2Group, rw)
	siteGroupReadPerm, siteGroupReadROperm := getPermsAndResourceOwnerPerms(pID, uID, metadata.GV2SiteGroup, r)
	siteGroupReadWritePerm, siteGroupReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, metadata.GV2SiteGroup, rw)

	noPerm, _ := getPermsAndResourceOwnerPerms(pID, uID, "user", []string{"read"})
	noPerm.SetGrantedToV2(nil) // eg: link shares

	cases := []struct {
		name              string
		graphPermissions  []models.Permissionable
		parsedPermissions []metadata.Permission
	}{
		{
			name:              "no perms",
			graphPermissions:  []models.Permissionable{},
			parsedPermissions: []metadata.Permission{},
		},
		{
			name:              "no user bound to perms",
			graphPermissions:  []models.Permissionable{noPerm},
			parsedPermissions: []metadata.Permission{},
		},

		// user
		{
			name:              "user with read permissions",
			graphPermissions:  []models.Permissionable{userReadPerm},
			parsedPermissions: []metadata.Permission{userReadROperm},
		},
		{
			name:              "user with owner permissions",
			graphPermissions:  []models.Permissionable{userOwnerPerm},
			parsedPermissions: []metadata.Permission{userOwnerROperm},
		},
		{
			name:              "user with read and write permissions",
			graphPermissions:  []models.Permissionable{userReadWritePerm},
			parsedPermissions: []metadata.Permission{userReadWriteROperm},
		},
		{
			name:              "multiple users with separate permissions",
			graphPermissions:  []models.Permissionable{userReadPerm, userReadWritePerm},
			parsedPermissions: []metadata.Permission{userReadROperm, userReadWriteROperm},
		},

		// site-user
		{
			name:              "site user with read permissions",
			graphPermissions:  []models.Permissionable{siteUserReadPerm},
			parsedPermissions: []metadata.Permission{siteUserReadROperm},
		},
		{
			name:              "site user with owner permissions",
			graphPermissions:  []models.Permissionable{siteUserOwnerPerm},
			parsedPermissions: []metadata.Permission{siteUserOwnerROperm},
		},
		{
			name:              "site user with read and write permissions",
			graphPermissions:  []models.Permissionable{siteUserReadWritePerm},
			parsedPermissions: []metadata.Permission{siteUserReadWriteROperm},
		},
		{
			name:              "multiple site users with separate permissions",
			graphPermissions:  []models.Permissionable{siteUserReadPerm, siteUserReadWritePerm},
			parsedPermissions: []metadata.Permission{siteUserReadROperm, siteUserReadWriteROperm},
		},

		// group
		{
			name:              "group with read permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm},
			parsedPermissions: []metadata.Permission{groupReadROperm},
		},
		{
			name:              "group with read and write permissions",
			graphPermissions:  []models.Permissionable{groupReadWritePerm},
			parsedPermissions: []metadata.Permission{groupReadWriteROperm},
		},
		{
			name:              "multiple groups with separate permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm, groupReadWritePerm},
			parsedPermissions: []metadata.Permission{groupReadROperm, groupReadWriteROperm},
		},

		// site-group
		{
			name:              "site group with read permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadPerm},
			parsedPermissions: []metadata.Permission{siteGroupReadROperm},
		},
		{
			name:              "site group with read and write permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadWritePerm},
			parsedPermissions: []metadata.Permission{siteGroupReadWriteROperm},
		},
		{
			name:              "multiple site groups with separate permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadPerm, siteGroupReadWritePerm},
			parsedPermissions: []metadata.Permission{siteGroupReadROperm, siteGroupReadWriteROperm},
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
