package api_test

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
)

type DriveAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func (suite *DriveAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func TestDriveAPIs(t *testing.T) {
	suite.Run(t, &DriveAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *DriveAPIIntgSuite) TestDrives_CreatePagerAndGetPage() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	siteID := tconfig.M365SiteID(t)
	pager := suite.its.ac.Drives().NewSiteDrivePager(siteID, []string{"name"})

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

func (suite *DriveAPIIntgSuite) TestDrives_PostItemInContainer() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rc := testdata.DefaultRestoreConfig("drive_api_post_item")
	acd := suite.its.ac.Drives()

	// generate a parent for the test data
	parent, err := acd.PostItemInContainer(
		ctx,
		suite.its.user.driveID,
		suite.its.user.driveRootFolderID,
		newItem(rc.Location, true),
		control.Replace)
	require.NoError(t, err, clues.ToCore(err))

	// generate a folder to use for collision testing
	folder := newItem("collision", true)
	origFolder, err := acd.PostItemInContainer(
		ctx,
		suite.its.user.driveID,
		ptr.Val(parent.GetId()),
		folder,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	// generate an item to use for collision testing
	file := newItem("collision.txt", false)
	origFile, err := acd.PostItemInContainer(
		ctx,
		suite.its.user.driveID,
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
			i, err := acd.PostItemInContainer(
				ctx,
				suite.its.user.driveID,
				ptr.Val(parent.GetId()),
				test.postItem,
				test.onCollision)

			test.expectErr(t, err)
			test.expectItem(t, i)
		})
	}
}

// purpose: ensure that creating a new folder with "replace" conflict behavior
// makes no changes to the items which exist in that folder.
func (suite *DriveAPIIntgSuite) TestDrives_PostItemInContainer_replaceFolderRegression() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		rc    = testdata.DefaultRestoreConfig("drive_folder_replace_regression")
		acd   = suite.its.ac.Drives()
		files = make([]models.DriveItemable, 0, 5)
	)

	// generate a folder for the test data
	folder, err := acd.PostItemInContainer(
		ctx,
		suite.its.user.driveID,
		suite.its.user.driveRootFolderID,
		newItem(rc.Location, true),
		// skip instead of replace here to get
		// an ErrItemAlreadyExistsConflict, just in case.
		control.Skip)
	require.NoError(t, err, clues.ToCore(err))

	// generate items within that folder
	for i := 0; i < 5; i++ {
		file := newItem(fmt.Sprintf("collision_%d.txt", i), false)
		f, err := acd.PostItemInContainer(
			ctx,
			suite.its.user.driveID,
			ptr.Val(folder.GetId()),
			file,
			control.Copy)
		require.NoError(t, err, clues.ToCore(err))

		files = append(files, f)
	}

	resultFolder, err := acd.PostItemInContainer(
		ctx,
		suite.its.user.driveID,
		ptr.Val(folder.GetParentReference().GetId()),
		newItem(rc.Location, true),
		control.Replace)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, ptr.Val(resultFolder.GetId()))
	require.Equal(t, ptr.Val(folder.GetId()), ptr.Val(resultFolder.GetId()))

	resultFileColl, err := acd.Stable.
		Client().
		Drives().
		ByDriveId(suite.its.user.driveID).
		Items().
		ByDriveItemId(ptr.Val(resultFolder.GetId())).
		Children().
		Get(ctx, nil)
	err = graph.Stack(ctx, err).OrNil()
	require.NoError(t, err, clues.ToCore(err))

	resultFiles := resultFileColl.GetValue()

	// asserting that no file changes have occurred as a result of the
	// "replacement" of the owning folder.
	for _, rf := range resultFiles {
		var (
			rID   = ptr.Val(rf.GetId())
			rName = ptr.Val(rf.GetName())
			rMod  = ptr.Val(rf.GetLastModifiedDateTime())
		)

		check := func(expect models.DriveItemable) bool {
			var (
				eID   = ptr.Val(expect.GetId())
				eName = ptr.Val(expect.GetName())
				eMod  = ptr.Val(expect.GetLastModifiedDateTime())
			)

			return eID == rID && eName == rName && eMod.Equal(rMod)
		}

		assert.True(t, slices.ContainsFunc(files, check))
	}
}
