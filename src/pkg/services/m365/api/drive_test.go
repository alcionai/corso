package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type DriveAPISuite struct {
	tester.Suite
	creds        account.M365Config
	ac           api.Client
	driveID      string
	rootFolderID string
}

func (suite *DriveAPISuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	userID := tester.M365UserID(t)
	a := tester.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = creds
	suite.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	drive, err := suite.ac.Users().GetDefaultDrive(ctx, userID)
	require.NoError(t, err, clues.ToCore(err))

	suite.driveID = ptr.Val(drive.GetId())

	rootFolder, err := suite.ac.Drives().GetRootFolder(ctx, suite.driveID)
	require.NoError(t, err, clues.ToCore(err))

	suite.rootFolderID = ptr.Val(rootFolder.GetId())
}

func TestDriveAPIs(t *testing.T) {
	suite.Run(t, &DriveAPISuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DriveAPISuite) TestDrives_CreatePagerAndGetPage() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	siteID := tester.M365SiteID(t)
	pager := suite.ac.Drives().NewSiteDrivePager(siteID, []string{"name"})

	a, err := pager.GetPage(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, a)
}

// newItem initializes a `models.DriveItemable` that can be used as input to `createItem`
func newItem(name string, folder bool) *models.DriveItem {
	itemToCreate := models.NewDriveItem()
	itemToCreate.SetName(&name)

	if folder {
		itemToCreate.SetFolder(models.NewFolder())
	} else {
		itemToCreate.SetFile(models.NewFile())
	}

	return itemToCreate
}

func (suite *DriveAPISuite) TestDrives_PostItemInContainer() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rc := testdata.DefaultRestoreConfig("drive_api_post_item")

	// generate a parent for the test data
	parent, err := suite.ac.Drives().PostItemInContainer(
		ctx,
		suite.driveID,
		suite.rootFolderID,
		newItem(rc.Location, true),
		control.Replace)
	require.NoError(t, err, clues.ToCore(err))

	// generate a folder to use for collision testing
	folder := newItem("collision", true)
	origFolder, err := suite.ac.Drives().PostItemInContainer(
		ctx,
		suite.driveID,
		ptr.Val(parent.GetId()),
		folder,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	// generate an item to use for collision testing
	file := newItem("collision.txt", false)
	origFile, err := suite.ac.Drives().PostItemInContainer(
		ctx,
		suite.driveID,
		ptr.Val(parent.GetId()),
		file,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name        string
		onCollision control.CollisionPolicy
		postItem    models.DriveItemable
		expectErr   func(t *testing.T, err error)
		expectItem  func(t *testing.T, i models.DriveItemable)
	}{
		{
			name:        "fail folder",
			onCollision: control.Skip,
			postItem:    folder,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.Nil(t, i)
			},
		},
		{
			name:        "rename folder",
			onCollision: control.Copy,
			postItem:    folder,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.NotEqual(
					t,
					ptr.Val(origFolder.GetId()),
					ptr.Val(i.GetId()),
					"renamed item should have a different id")
				assert.NotEqual(
					t,
					ptr.Val(origFolder.GetName()),
					ptr.Val(i.GetName()),
					"renamed item should have a different name")
			},
		},
		{
			name:        "replace folder",
			onCollision: control.Replace,
			postItem:    folder,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.Equal(
					t,
					ptr.Val(origFolder.GetId()),
					ptr.Val(i.GetId()),
					"replaced item should have the same id")
				assert.Equal(
					t,
					ptr.Val(origFolder.GetName()),
					ptr.Val(i.GetName()),
					"replaced item should have the same name")
			},
		},
		{
			name:        "fail file",
			onCollision: control.Skip,
			postItem:    file,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.Nil(t, i)
			},
		},
		{
			name:        "rename file",
			onCollision: control.Copy,
			postItem:    file,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.NotEqual(
					t,
					ptr.Val(origFile.GetId()),
					ptr.Val(i.GetId()),
					"renamed item should have a different id")
				assert.NotEqual(
					t,
					ptr.Val(origFolder.GetName()),
					ptr.Val(i.GetName()),
					"renamed item should have a different name")
			},
		},
		// FIXME: this *should* behave the same as folder collision, but there's either a
		// bug or a deviation in graph api behavior.
		// See open ticket: https://github.com/OneDrive/onedrive-api-docs/issues/1702
		{
			name:        "replace file",
			onCollision: control.Replace,
			postItem:    file,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				assert.Nil(t, i)
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			i, err := suite.ac.Drives().PostItemInContainer(
				ctx,
				suite.driveID,
				ptr.Val(parent.GetId()),
				test.postItem,
				test.onCollision)

			test.expectErr(t, err)
			test.expectItem(t, i)
		})
	}
}
