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
	expectedCollections map[string][]string
}

func backupLimitTable() (models.Driveable, models.Driveable, []backupLimitTest) {
	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	tbl := []backupLimitTest{
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
					idx(drive, 1): {
						Pages: pagesOf(rootAnd(
							driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 7, isFile),
							driveItemWithSize(idx(file, 2), namex(file, 2), parent(1), rootID, 1, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parent(1), rootID, 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 2), idx(file, 3)},
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
					idx(drive, 1): {
						Pages: pagesOf(rootAnd(
							driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 1, isFile),
							driveItemWithSize(idx(file, 2), namex(file, 2), parent(1), rootID, 2, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parent(1), rootID, 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2)},
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
					idx(drive, 1): {
						Pages: pagesOf(rootAnd(
							driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 1, isFile),
							driveItemWithSize(idx(folder, 1), namex(folder, 1), parent(1), rootID, 1, isFolder),
							driveItemWithSize(idx(file, 2), namex(file, 2), parent(1, namex(folder, 1)), idx(folder, 1), 2, isFile),
							driveItemWithSize(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), 1, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 2)},
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
					idx(drive, 1): {
						Pages: pagesOf(rootAnd(
							driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), parent(1), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), parent(1), rootID, isFile),
							driveItem(idx(file, 6), namex(file, 6), parent(1), rootID, isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2), idx(file, 3)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile)),
							rootAnd(
								// Repeated items shouldn't count against the limit.
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 6), namex(file, 6), parent(1, namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 3)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 6), namex(file, 6), parent(1, namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				// Root has an additional item. It's hard to fix that in the code
				// though.
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
								driveItem(idx(file, 1), namex(file, 1), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 0)), idx(folder, 0), isFile)),
							rootAnd(
								driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
								// Updated item that shouldn't count against the limit a second time.
								driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 0)), idx(folder, 0), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {},
				fullPath(1, namex(folder, 0)): {idx(folder, 0), idx(file, 1), idx(file, 2), idx(file, 3)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								// Put folder 0 at limit.
								driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
								driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 0)), idx(folder, 0), isFile)),
							rootAnd(
								driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
								// Try to move item from root to folder 0 which is already at the limit.
								driveItem(idx(file, 1), namex(file, 1), parent(1, namex(folder, 0)), idx(folder, 0), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 0)): {idx(folder, 0), idx(file, 3), idx(file, 4)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								// This container shouldn't be returned.
								driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
								driveItem(idx(file, 7), namex(file, 7), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 8), namex(file, 8), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 9), namex(file, 9), parent(1, namex(folder, 2)), idx(folder, 2), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile)),
							rootAnd(
								// This container shouldn't be returned.
								driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
								driveItem(idx(file, 7), namex(file, 7), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 8), namex(file, 8), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								driveItem(idx(file, 9), namex(file, 9), parent(1, namex(folder, 2)), idx(folder, 2), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
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
					idx(drive, 1): {
						Pages: pagesOf(rootAnd(
							driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), parent(1), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), parent(1), rootID, isFile),
						)),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
					idx(drive, 2): {
						Pages: pagesOf(rootAnd(
							driveItem(idx(file, 1), namex(file, 1), parent(2), rootID, isFile),
							driveItem(idx(file, 2), namex(file, 2), parent(2), rootID, isFile),
							driveItem(idx(file, 3), namex(file, 3), parent(2), rootID, isFile),
							driveItem(idx(file, 4), namex(file, 4), parent(2), rootID, isFile),
							driveItem(idx(file, 5), namex(file, 5), parent(2), rootID, isFile),
						)),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(2): {idx(file, 1), idx(file, 2), idx(file, 3)},
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
					idx(drive, 1): {
						Pages: pagesOf(
							rootAnd(
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile)),
							rootAnd(
								driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
								driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile))),
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
	}

	return drive1, drive2, tbl
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_noTree() {
	_, _, tbl := backupLimitTable()

	for _, test := range tbl {
		suite.Run(test.name, func() {
			runGetPreviewLimits(
				suite.T(),
				test,
				control.DefaultOptions())
		})
	}
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_tree() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	_, _, tbl := backupLimitTable()

	for _, test := range tbl {
		suite.Run(test.name, func() {
			runGetPreviewLimits(
				suite.T(),
				test,
				opts)
		})
	}
}

func runGetPreviewLimits(
	t *testing.T,
	test backupLimitTest,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

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

		if folderPath == metadataPath.String() {
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
			test.expectedCollections[folderPath],
			itemIDs,
			"expected elements to match in collection with path %q",
			folderPath)
	}

	assert.ElementsMatch(
		t,
		maps.Keys(test.expectedCollections),
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
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_Defaults_noTree() {
	// Add a check that will fail if we make the default smaller than expected.
	require.LessOrEqual(
		suite.T(),
		int64(1024*1024),
		defaultPreviewMaxBytes,
		"default number of bytes changed; DefaultNumBytes test case may need updating!")
	require.Zero(
		suite.T(),
		defaultPreviewMaxBytes%(1024*1024),
		"default number of bytes isn't divisible by 1MB; DefaultNumBytes test case may need updating!")

	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				control.DefaultOptions())
		})
	}
}

// TestGet_PreviewLimits_Defaults checks that default values are used when
// making a preview backup if the user didn't provide some options.
// These tests run a reduced set of checks that really just look for item counts
// and such. Other tests are expected to provide more comprehensive checks.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_Defaults_tree() {
	// Add a check that will fail if we make the default smaller than expected.
	require.LessOrEqual(
		suite.T(),
		int64(1024*1024),
		defaultPreviewMaxBytes,
		"default number of bytes changed; DefaultNumBytes test case may need updating!")
	require.Zero(
		suite.T(),
		defaultPreviewMaxBytes%(1024*1024),
		"default number of bytes isn't divisible by 1MB; DefaultNumBytes test case may need updating!")

	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				opts)
		})
	}
}

func runGetPreviewLimitsDefaults(
	t *testing.T,
	test defaultLimitTest,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	drv := models.NewDrive()
	drv.SetId(ptr.To(idx(drive, 1)))
	drv.SetName(ptr.To(namex(drive, 1)))

	pages := make([]mock.NextPage, 0, test.numContainers)

	for containerIdx := 0; containerIdx < test.numContainers; containerIdx++ {
		page := mock.NextPage{
			Items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(
					idx(folder, containerIdx),
					namex(folder, containerIdx),
					parent(1),
					rootID,
					isFolder),
			},
		}

		for itemIdx := 0; itemIdx < test.numItemsPerContainer; itemIdx++ {
			itemSuffix := fmt.Sprintf("%d-%d", containerIdx, itemIdx)

			page.Items = append(page.Items, driveItemWithSize(
				idx(file, itemSuffix),
				namex(file, itemSuffix),
				parent(1, namex(folder, containerIdx)),
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
				idx(drive, 1): {
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

		if folderPath == metadataPath.String() {
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
