package drive

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func minimumLimitOpts() control.Options {
	minLimitOpts := control.DefaultOptions()
	minLimitOpts.PreviewLimits.Enabled = true
	minLimitOpts.PreviewLimits.MaxBytes = 1
	minLimitOpts.PreviewLimits.MaxContainers = 1
	minLimitOpts.PreviewLimits.MaxItems = 1
	minLimitOpts.PreviewLimits.MaxItemsPerContainer = 1
	minLimitOpts.PreviewLimits.MaxPages = 1

	return minLimitOpts
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type LimiterUnitSuite struct {
	tester.Suite
}

func TestLimiterUnitSuite(t *testing.T) {
	suite.Run(t, &LimiterUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type backupLimitTest struct {
	name       string
	limits     control.PreviewItemLimits
	drives     []models.Driveable
	enumerator mock.EnumerateItemsDeltaByDrive
	// Collection name -> set of item IDs. We can't check item data because
	// that's not mocked out. Metadata is checked separately.
	expectedItemIDsInCollection map[string][]string
}

func backupLimitTable(drive1, drive2 models.Driveable) []backupLimitTest {
	return []backupLimitTest{
		{
			name: "OneDrive SinglePage ExcludeItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             5,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(pageItems(
							driveItemWithSize(idx(file, 1), namex(file, 1), parentDir(), rootID, 7, isFile),
							driveItemWithSize(idx(file, 2), namex(file, 2), parentDir(), rootID, 1, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parentDir(), rootID, 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath(): {idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive SinglePage SingleFolder ExcludeCombinedItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             3,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(pageItems(
							driveItemWithSize(idx(file, 1), namex(file, 1), parentDir(), rootID, 1, isFile),
							driveItemWithSize(idx(file, 2), namex(file, 2), parentDir(), rootID, 2, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parentDir(), rootID, 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath(): {idx(file, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive SinglePage MultipleFolders ExcludeCombinedItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             3,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(pageItems(
							driveItemWithSize(idx(file, 1), namex(file, 1), parentDir(), rootID, 1, isFile),
							driveItemWithSize(idx(folder, 1), namex(folder, 1), parentDir(), rootID, 1, isFolder),
							driveItemWithSize(idx(file, 2), namex(file, 2), parentDir(namex(folder, 1)), idx(folder, 1), 2, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parentDir(namex(folder, 1)), idx(folder, 1), 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive SinglePage SingleFolder ItemLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(pageItems(
							driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), parentDir(), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), parentDir(), rootID, isFile),
							driveItem(idx(file, 6), namex(file, 6), parentDir(), rootID, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath(): {idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages MultipleFolders ItemLimit WithRepeatedItem",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile)),
							pageItems(
								// Repeated items shouldn't count against the limit.
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 6), namex(file, 6), parentDir(namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1), idx(file, 2)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages PageLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             1,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 6), namex(file, 6), parentDir(namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath(): {idx(file, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 1,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				// Root has an additional item. It's hard to fix that in the code
				// though.
				fullPath():                 {idx(file, 1), idx(file, 2)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 4)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit ItemUpdated",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 3,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(file, 1), namex(file, 1), parentDir(name(folder)), id(folder), isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(name(folder)), id(folder), isFile)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								// Updated item that shouldn't count against the limit a second time.
								driveItem(idx(file, 2), namex(file, 2), parentDir(name(folder)), id(folder), isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(name(folder)), id(folder), isFile),
								driveItem(idx(file, 4), namex(file, 4), parentDir(name(folder)), id(folder), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():             {},
				fullPath(name(folder)): {id(folder), idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit MoveItemBetweenFolders",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 2,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								// Put folder 0 at limit.
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parentDir(name(folder)), id(folder), isFile),
								driveItem(idx(file, 4), namex(file, 4), parentDir(name(folder)), id(folder), isFile)),
							pageItems(
								driveItem(id(folder), name(folder), parentDir(), rootID, isFolder),
								// Try to move item from root to folder 0 which is already at the limit.
								driveItem(idx(file, 1), namex(file, 1), parentDir(name(folder)), id(folder), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():             {idx(file, 1), idx(file, 2)},
				fullPath(name(folder)): {id(folder), idx(file, 3), idx(file, 4)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit LastContainerSplitAcrossPages",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit NextContainerOnSamePage",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								// This container shouldn't be returned.
								driveItem(idx(folder, 2), namex(folder, 2), parentDir(), rootID, isFolder),
								driveItem(idx(file, 7), namex(file, 7), parentDir(namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 8), namex(file, 8), parentDir(namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 9), namex(file, 9), parentDir(namex(folder, 2)), idx(folder, 2), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit NextContainerOnNextPage",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile)),
							pageItems(
								// This container shouldn't be returned.
								driveItem(idx(folder, 2), namex(folder, 2), parentDir(), rootID, isFolder),
								driveItem(idx(file, 7), namex(file, 7), parentDir(namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 8), namex(file, 8), parentDir(namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 9), namex(file, 9), parentDir(namex(folder, 2)), idx(folder, 2), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "TwoDrives SeparateLimitAccounting",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1, drive2},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(pageItems(
							driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), parentDir(), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), parentDir(), rootID, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
					idx(drive, 2): {
						Pages: pagesOf(pageItems(
							driveItem(idx(file, 1), namex(file, 1), driveParentDir(2), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), driveParentDir(2), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), driveParentDir(2), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), driveParentDir(2), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), driveParentDir(2), rootID, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():       {idx(file, 1), idx(file, 2), idx(file, 3)},
				driveFullPath(2): {idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive PreviewDisabled MinimumLimitsIgnored",
			limits: control.PreviewItemLimits{
				MaxItems:             1,
				MaxItemsPerContainer: 1,
				MaxContainers:        1,
				MaxBytes:             1,
				MaxPages:             1,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					id(drive): {
						Pages: pagesOf(
							pageItems(
								driveItem(idx(file, 1), namex(file, 1), parentDir(), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parentDir(), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parentDir(), rootID, isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parentDir(namex(folder, 1)), idx(folder, 1), isFile)),
							pageItems(
								driveItem(idx(folder, 1), namex(folder, 1), parentDir(), rootID, isFolder),
								driveItem(idx(file, 5), namex(file, 5), parentDir(namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedItemIDsInCollection: map[string][]string{
				fullPath():                 {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
	}
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_noTree() {
	t := suite.T()

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(id(drive)))
	drive1.SetName(ptr.To(name(drive)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	for _, test := range backupLimitTable(drive1, drive2) {
		suite.Run(test.name, func() {
			runGetPreviewLimits(
				suite.T(),
				test,
				metadataPath,
				control.DefaultOptions())
		})
	}
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_tree() {
	t := suite.T()

	t.Skip("TODO: unskip when tree produces collections")

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(id(drive)))
	drive1.SetName(ptr.To(name(drive)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	for _, test := range backupLimitTable(drive1, drive2) {
		suite.Run(test.name, func() {
			runGetPreviewLimits(
				suite.T(),
				test,
				metadataPath,
				opts)
		})
	}
}

func runGetPreviewLimits(
	t *testing.T,
	test backupLimitTest,
	metadata path.Path,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	opts.PreviewLimits = test.limits

	var (
		mockDrivePager = &apiMock.Pager[models.Driveable]{
			ToReturn: []apiMock.PagerResult[models.Driveable]{
				{Values: test.drives},
			},
		}
		mbh       = mock.DefaultDriveBHWith(user, mockDrivePager, test.enumerator)
		c         = collWithMBHAndOpts(mbh, opts)
		errs      = fault.New(true)
		delList   = prefixmatcher.NewStringSetBuilder()
		collPaths = []string{}
	)

	cols, canUsePreviousBackup, err := c.Get(ctx, nil, delList, errs)

	if opts.ToggleFeatures.UseDeltaTree {
		require.ErrorIs(t, err, errGetTreeNotImplemented, clues.ToCore(err))
	} else {
		require.NoError(t, err, clues.ToCore(err))
	}

	assert.True(t, canUsePreviousBackup, "can use previous backup")
	assert.Empty(t, errs.Skipped())

	for _, baseCol := range cols {
		// There shouldn't be any deleted collections.
		if !assert.NotEqual(
			t,
			data.DeletedState,
			baseCol.State(),
			"collection marked deleted") {
			continue
		}

		folderPath := baseCol.FullPath().String()

		if folderPath == metadata.String() {
			continue
		}

		collPaths = append(collPaths, folderPath)

		// TODO: We should really be getting items in the collection
		// via the Items() channel. The lack of that makes this check a bit more
		// bittle since internal details can change.  The wiring to support
		// mocked GetItems is available.  We just haven't plugged it in yet.
		col, ok := baseCol.(*Collection)
		require.True(t, ok, "getting onedrive.Collection handle")

		itemIDs := make([]string, 0, len(col.driveItems))

		for id := range col.driveItems {
			itemIDs = append(itemIDs, id)
		}

		assert.ElementsMatchf(
			t,
			test.expectedItemIDsInCollection[folderPath],
			itemIDs,
			"item IDs in collection with path %q",
			folderPath)
	}

	assert.ElementsMatch(
		t,
		maps.Keys(test.expectedItemIDsInCollection),
		collPaths,
		"collection paths")
}

// The number of pages returned can be indirectly tested by checking how many
// containers/items were returned.
type defaultLimitTestExpects struct {
	numItems             int
	numContainers        int
	numItemsPerContainer int
}

type defaultLimitTest struct {
	name                 string
	numContainers        int
	numItemsPerContainer int
	itemSize             int64
	limits               control.PreviewItemLimits
	expect               defaultLimitTestExpects
}

func defaultLimitsTable() []defaultLimitTest {
	return []defaultLimitTest{
		{
			name:                 "DefaultNumItems",
			numContainers:        1,
			numItemsPerContainer: defaultPreviewMaxItems + 1,
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItemsPerContainer: 99999999,
				MaxContainers:        99999999,
				MaxBytes:             99999999,
				MaxPages:             99999999,
			},
			expect: defaultLimitTestExpects{
				numItems:             defaultPreviewMaxItems,
				numContainers:        1,
				numItemsPerContainer: defaultPreviewMaxItems,
			},
		},
		{
			name:                 "DefaultNumContainers",
			numContainers:        defaultPreviewMaxContainers + 1,
			numItemsPerContainer: 1,
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             99999999,
				MaxItemsPerContainer: 99999999,
				MaxBytes:             99999999,
				MaxPages:             99999999,
			},
			expect: defaultLimitTestExpects{
				// Root is counted as a container in the code but won't be counted or
				// have items in the test.
				numItems:             defaultPreviewMaxContainers - 1,
				numContainers:        defaultPreviewMaxContainers - 1,
				numItemsPerContainer: 1,
			},
		},
		{
			name:                 "DefaultNumItemsPerContainer",
			numContainers:        1,
			numItemsPerContainer: defaultPreviewMaxItemsPerContainer + 1,
			limits: control.PreviewItemLimits{
				Enabled:       true,
				MaxItems:      99999999,
				MaxContainers: 99999999,
				MaxBytes:      99999999,
				MaxPages:      99999999,
			},
			expect: defaultLimitTestExpects{
				numItems:             defaultPreviewMaxItemsPerContainer,
				numContainers:        1,
				numItemsPerContainer: defaultPreviewMaxItemsPerContainer,
			},
		},
		{
			name:                 "DefaultNumPages",
			numContainers:        defaultPreviewMaxPages + 1,
			numItemsPerContainer: 1,
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             99999999,
				MaxContainers:        99999999,
				MaxItemsPerContainer: 99999999,
				MaxBytes:             99999999,
			},
			expect: defaultLimitTestExpects{
				numItems:             defaultPreviewMaxPages,
				numContainers:        defaultPreviewMaxPages,
				numItemsPerContainer: 1,
			},
		},
		{
			name:                 "DefaultNumBytes",
			numContainers:        1,
			numItemsPerContainer: int(defaultPreviewMaxBytes/1024/1024) + 1,
			itemSize:             1024 * 1024,
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             99999999,
				MaxContainers:        99999999,
				MaxItemsPerContainer: 99999999,
				MaxPages:             99999999,
			},
			expect: defaultLimitTestExpects{
				numItems:             int(defaultPreviewMaxBytes) / 1024 / 1024,
				numContainers:        1,
				numItemsPerContainer: int(defaultPreviewMaxBytes) / 1024 / 1024,
			},
		},
	}
}

// TestGet_PreviewLimits_Defaults checks that default values are used when
// making a preview backup if the user didn't provide some options.
// These tests run a reduced set of checks that really just look for item counts
// and such. Other tests are expected to provide more comprehensive checks.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_defaultsNoTree() {
	t := suite.T()

	// Add a check that will fail if we make the default smaller than expected.
	require.LessOrEqual(
		t,
		int64(1024*1024),
		defaultPreviewMaxBytes,
		"default number of bytes changed; DefaultNumBytes test case may need updating!")
	require.Zero(
		t,
		defaultPreviewMaxBytes%(1024*1024),
		"default number of bytes isn't divisible by 1MB; DefaultNumBytes test case may need updating!")

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(id(drive)))
	drive1.SetName(ptr.To(name(drive)))

	// The number of pages the test generates can be controlled by setting the
	// number of containers. The test will put one (non-root) container per page.

	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				drive1,
				metadataPath,
				control.DefaultOptions())
		})
	}
}

// TestGet_PreviewLimits_Defaults checks that default values are used when
// making a preview backup if the user didn't provide some options.
// These tests run a reduced set of checks that really just look for item counts
// and such. Other tests are expected to provide more comprehensive checks.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_defaultsWithTree() {
	t := suite.T()

	t.Skip("TODO: unskip when tree produces collections")

	// Add a check that will fail if we make the default smaller than expected.
	require.LessOrEqual(
		t,
		int64(1024*1024),
		defaultPreviewMaxBytes,
		"default number of bytes changed; DefaultNumBytes test case may need updating!")
	require.Zero(
		t,
		defaultPreviewMaxBytes%(1024*1024),
		"default number of bytes isn't divisible by 1MB; DefaultNumBytes test case may need updating!")

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(id(drive)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	// The number of pages the test generates can be controlled by setting the
	// number of containers. The test will put one (non-root) container per page.
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				drive1,
				metadataPath,
				opts)
		})
	}
}

func runGetPreviewLimitsDefaults(
	t *testing.T,
	test defaultLimitTest,
	drv models.Driveable,
	metadata path.Path,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	pages := make([]mock.NextPage, 0, test.numContainers)

	for containerIdx := 0; containerIdx < test.numContainers; containerIdx++ {
		page := mock.NextPage{
			Items: []models.DriveItemable{
				driveRootItem(),
				driveItem(
					idx(folder, containerIdx),
					namex(folder, containerIdx),
					parentDir(),
					rootID,
					isFolder),
			},
		}

		for itemIdx := 0; itemIdx < test.numItemsPerContainer; itemIdx++ {
			itemSuffix := fmt.Sprintf("%d-%d", containerIdx, itemIdx)

			page.Items = append(page.Items, driveItemWithSize(
				idx(file, itemSuffix),
				namex(file, itemSuffix),
				parentDir(namex(folder, containerIdx)),
				idx(folder, containerIdx),
				test.itemSize,
				isFile))
		}

		pages = append(pages, page)
	}

	opts.PreviewLimits = test.limits

	var (
		mockDrivePager = &apiMock.Pager[models.Driveable]{
			ToReturn: []apiMock.PagerResult[models.Driveable]{
				{Values: []models.Driveable{drv}},
			},
		}
		mockEnumerator = mock.EnumerateItemsDeltaByDrive{
			DrivePagers: map[string]*mock.DriveItemsDeltaPager{
				id(drive): {
					Pages:       pages,
					DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
				},
			},
		}
		mbh           = mock.DefaultDriveBHWith(user, mockDrivePager, mockEnumerator)
		c             = collWithMBHAndOpts(mbh, opts)
		errs          = fault.New(true)
		delList       = prefixmatcher.NewStringSetBuilder()
		numContainers int
		numItems      int
	)

	cols, canUsePreviousBackup, err := c.Get(ctx, nil, delList, errs)

	if opts.ToggleFeatures.UseDeltaTree {
		require.ErrorIs(t, err, errGetTreeNotImplemented, clues.ToCore(err))
	} else {
		require.NoError(t, err, clues.ToCore(err))
	}

	assert.True(t, canUsePreviousBackup, "can use previous backup")
	assert.Empty(t, errs.Skipped())

	for _, baseCol := range cols {
		require.NotEqual(
			t,
			data.DeletedState,
			baseCol.State(),
			"no collections should be marked deleted")

		folderPath := baseCol.FullPath().String()

		if folderPath == metadata.String() {
			continue
		}

		// Skip the root container and don't count it because we don't put
		// anything in it.
		dp, err := path.ToDrivePath(baseCol.FullPath())
		require.NoError(t, err, clues.ToCore(err))

		if len(dp.Folders) == 0 {
			continue
		}

		numContainers++

		// TODO: We should really be getting items in the collection
		// via the Items() channel. The lack of that makes this check a bit more
		// bittle since internal details can change.  The wiring to support
		// mocked GetItems is available.  We just haven't plugged it in yet.
		col, ok := baseCol.(*Collection)
		require.True(t, ok, "baseCol must be type *Collection")

		numItems += len(col.driveItems)

		// Add one to account for the folder permissions item.
		assert.Len(
			t,
			col.driveItems,
			test.expect.numItemsPerContainer+1,
			"items in container %v",
			col.FullPath())
	}

	assert.Equal(
		t,
		test.expect.numContainers,
		numContainers,
		"total containers")

	// Each container also gets an item so account for that here.
	assert.Equal(
		t,
		test.expect.numItems+test.expect.numContainers,
		numItems,
		"total items across all containers")
}
