package api

import (
	"fmt"
	"testing"
	"time"

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
	"github.com/alcionai/corso/src/internal/tester/tsetup"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
)

type DriveAPIIntgSuite struct {
	tester.Suite
	its tsetup.M365
}

func (suite *DriveAPIIntgSuite) SetupSuite() {
	suite.its = tsetup.NewM365IntegrationTester(suite.T())
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
	pager := suite.its.AC.Drives().NewSiteDrivePager(siteID, []string{"name"})

	a, err := pager.GetPage(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, a)
}

func (suite *DriveAPIIntgSuite) TestDrives_PostItemInContainer() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rc := testdata.DefaultRestoreConfig("drive_api_post_item")
	acd := suite.its.AC.Drives()

	// generate a parent for the test data
	parent, err := acd.PostItemInContainer(
		ctx,
		suite.its.User.DriveID,
		suite.its.User.DriveRootFolderID,
		NewDriveItem(rc.Location, true),
		control.Replace)
	require.NoError(t, err, clues.ToCore(err))

	// generate a folder to use for collision testing
	folder := NewDriveItem("collision", true)
	origFolder, err := acd.PostItemInContainer(
		ctx,
		suite.its.User.DriveID,
		ptr.Val(parent.GetId()),
		folder,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	// generate an item to use for collision testing
	file := NewDriveItem("collision.txt", false)
	origFile, err := acd.PostItemInContainer(
		ctx,
		suite.its.User.DriveID,
		ptr.Val(parent.GetId()),
		file,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	// ensure we don't bucket the mod time within a second
	time.Sleep(2 * time.Second)

	updatedFile := models.NewDriveItem()
	updatedFile.SetAdditionalData(origFile.GetAdditionalData())
	updatedFile.SetCreatedBy(origFile.GetCreatedBy())
	updatedFile.SetCreatedDateTime(origFile.GetCreatedDateTime())
	updatedFile.SetDescription(origFile.GetDescription())
	updatedFile.SetFile(origFile.GetFile())
	updatedFile.SetName(ptr.To("updated" + ptr.Val(origFile.GetName())))
	updatedFile.SetSize(origFile.GetSize())
	updatedFile.SetWebUrl(origFile.GetWebUrl())

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
		// Note: this currently behaves the same as folder collision, but there used to be a
		// bug or a deviation in graph api behavior that prevented it from succeeding.
		// No response on the ticket below, so this test code is being kept around to showcase
		// that prior behavior while we're evaluating the permanence of the fix.
		// See open ticket: https://github.com/OneDrive/onedrive-api-docs/issues/1702
		// {
		// 	name:        "replace file",
		// 	onCollision: control.Replace,
		// 	postItem:    file,
		// 	expectErr: func(t *testing.T, err error) {
		// 		assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
		// 	},
		// 	expectItem: func(t *testing.T, i models.DriveItemable) {
		// 		assert.Nil(t, i)
		// 	},
		// },
		{
			name:        "replace file",
			onCollision: control.Replace,
			postItem:    updatedFile,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectItem: func(t *testing.T, i models.DriveItemable) {
				// the name was updated
				assert.Equal(
					t,
					"updated"+ptr.Val(origFile.GetName()),
					ptr.Val(i.GetName()),
					"replaced item should have the updated name")
				// the mod time automatically updates
				assert.True(
					t,
					ptr.Val(origFile.GetLastModifiedDateTime()).Before(ptr.Val(i.GetLastModifiedDateTime())),
					"replaced item should have a later mod time")
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			i, err := acd.PostItemInContainer(
				ctx,
				suite.its.User.DriveID,
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
		acd   = suite.its.AC.Drives()
		files = make([]models.DriveItemable, 0, 5)
	)

	// generate a folder for the test data
	folder, err := acd.PostItemInContainer(
		ctx,
		suite.its.User.DriveID,
		suite.its.User.DriveRootFolderID,
		NewDriveItem(rc.Location, true),
		// skip instead of replace here to get
		// an ErrItemAlreadyExistsConflict, just in case.
		control.Skip)
	require.NoError(t, err, clues.ToCore(err))

	// generate items within that folder
	for i := 0; i < 5; i++ {
		file := NewDriveItem(fmt.Sprintf("collision_%d.txt", i), false)
		f, err := acd.PostItemInContainer(
			ctx,
			suite.its.User.DriveID,
			ptr.Val(folder.GetId()),
			file,
			control.Copy)
		require.NoError(t, err, clues.ToCore(err))

		files = append(files, f)
	}

	resultFolder, err := acd.PostItemInContainer(
		ctx,
		suite.its.User.DriveID,
		ptr.Val(folder.GetParentReference().GetId()),
		NewDriveItem(rc.Location, true),
		control.Replace)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, ptr.Val(resultFolder.GetId()))
	require.Equal(t, ptr.Val(folder.GetId()), ptr.Val(resultFolder.GetId()))

	resultFileColl, err := acd.Stable.
		Client().
		Drives().
		ByDriveId(suite.its.User.DriveID).
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
