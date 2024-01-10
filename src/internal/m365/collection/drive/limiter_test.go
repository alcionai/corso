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
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type LimiterUnitSuite struct {
	tester.Suite
}

func TestLimiterUnitSuite(t *testing.T) {
	suite.Run(t, &LimiterUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type backupLimitTest struct {
	name       string
	limits     control.PreviewItemLimits
	enumerator enumerateDriveItemsDelta
	// Collection name -> set of item IDs. We can't check item data because
	// that's not mocked out. Metadata is checked separately.
	expectedItemIDsInCollection map[string][]string
	// Collection name -> set of item IDs. We can't check item data because
	// that's not mocked out. Metadata is checked separately.
	// the tree version has some different (more accurate) expectations
	// for success
	expectedItemIDsInCollectionTree map[string][]string
}

func backupLimitTable(t *testing.T, d1, d2 *deltaDrive) []backupLimitTest {
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileWSizeAt(7, root, "f1"),
							d1.fileWSizeAt(1, root, "f2"),
							d1.fileWSizeAt(1, root, "f3"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t): {fileID("f2"), fileID("f3")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileWSizeAt(1, root, "f1"),
							d1.fileWSizeAt(2, root, "f2"),
							d1.fileWSizeAt(1, root, "f3"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t): {fileID("f1"), fileID("f2")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileWSizeAt(1, root, "f1"),
							d1.folderAt(root),
							d1.fileWSizeAt(2, folder, "f2"),
							d1.fileWSizeAt(1, folder, "f3"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1")},
				d1.strPath(t, folderName()): {folderID(), fileID("f2")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3"),
							d1.fileAt(root, "f4"),
							d1.fileAt(root, "f5"),
							d1.fileAt(root, "f6"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t): {fileID("f1"), fileID("f2"), fileID("f3")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2")),
						aPage(
							// Repeated items shouldn't count against the limit.
							d1.fileAt(root, "f1"),
							d1.folderAt(root),
							d1.fileAt(folder, "f3"),
							d1.fileAt(folder, "f4"),
							d1.fileAt(folder, "f5"),
							d1.fileAt(folder, "f6"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2")},
				d1.strPath(t, folderName()): {folderID(), fileID("f3")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f3"),
							d1.fileAt(folder, "f4"),
							d1.fileAt(folder, "f5"),
							d1.fileAt(folder, "f6"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t): {fileID("f1"), fileID("f2")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f4"),
							d1.fileAt(folder, "f5"))))),
			expectedItemIDsInCollection: map[string][]string{
				// Root has an additional item. It's hard to fix that in the code though.
				d1.strPath(t):               {fileID("f1"), fileID("f2")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4")},
			},
			expectedItemIDsInCollectionTree: map[string][]string{
				// the tree version doesn't have this problem.
				d1.strPath(t):               {fileID("f1")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f1"),
							d1.fileAt(folder, "f2")),
						aPage(
							d1.folderAt(root),
							// Updated item that shouldn't count against the limit a second time.
							d1.fileAt(folder, "f2"),
							d1.fileAt(folder, "f3"),
							d1.fileAt(folder, "f4"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {},
				d1.strPath(t, folderName()): {folderID(), fileID("f1"), fileID("f2"), fileID("f3")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							// Put root/folder at limit.
							d1.folderAt(root),
							d1.fileAt(folder, "f3"),
							d1.fileAt(folder, "f4")),
						aPage(
							d1.folderAt(root),
							// Try to move item from root to folder 0 which is already at the limit.
							d1.fileAt(folder, "f1"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2")},
				d1.strPath(t, folderName()): {folderID(), fileID("f3"), fileID("f4")},
			},
			expectedItemIDsInCollectionTree: map[string][]string{
				d1.strPath(t): {fileID("f2")},
				// note that the tree version allows f1 to get moved.
				// we've already committed to backing up the file as part of the preview,
				// it doesn't seem rational to prevent its movement
				d1.strPath(t, folderName()): {folderID(), fileID("f3"), fileID("f4"), fileID("f1")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f4")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f5"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2"), fileID("f3")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4"), fileID("f5")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f4"),
							d1.fileAt(folder, "f5"),
							// This container shouldn't be returned.
							d1.folderAt(root, 2),
							d1.fileAt(2, "f7"),
							d1.fileAt(2, "f8"),
							d1.fileAt(2, "f9"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2"), fileID("f3")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4"), fileID("f5")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f4"),
							d1.fileAt(folder, "f5")),
						aPage(
							// This container shouldn't be returned.
							d1.folderAt(root, 2),
							d1.fileAt(2, "f7"),
							d1.fileAt(2, "f8"),
							d1.fileAt(2, "f9"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2"), fileID("f3")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4"), fileID("f5")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3"),
							d1.fileAt(root, "f4"),
							d1.fileAt(root, "f5")))),
				d2.newEnumer().with(
					delta(nil).with(
						aPage(
							d2.fileAt(root, "f1"),
							d2.fileAt(root, "f2"),
							d2.fileAt(root, "f3"),
							d2.fileAt(root, "f4"),
							d2.fileAt(root, "f5"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t): {fileID("f1"), fileID("f2"), fileID("f3")},
				d2.strPath(t): {fileID("f1"), fileID("f2"), fileID("f3")},
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
			enumerator: driveEnumerator(
				d1.newEnumer().with(
					delta(nil).with(
						aPage(
							d1.fileAt(root, "f1"),
							d1.fileAt(root, "f2"),
							d1.fileAt(root, "f3")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f4")),
						aPage(
							d1.folderAt(root),
							d1.fileAt(folder, "f5"))))),
			expectedItemIDsInCollection: map[string][]string{
				d1.strPath(t):               {fileID("f1"), fileID("f2"), fileID("f3")},
				d1.strPath(t, folderName()): {folderID(), fileID("f4"), fileID("f5")},
			},
		},
	}
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_noTree() {
	iterGetPreviewLimitsTests(suite, control.DefaultOptions(), control.DefaultBackupConfig())
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_tree() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	iterGetPreviewLimitsTests(suite, opts, control.DefaultBackupConfig())
}

func iterGetPreviewLimitsTests(
	suite *LimiterUnitSuite,
	opts control.Options,
	backupOpts control.BackupConfig,
) {
	d1, d2 := drive(), drive(2)

	for _, test := range backupLimitTable(suite.T(), d1, d2) {
		suite.Run(test.name, func() {
			runGetPreviewLimits(
				suite.T(),
				test,
				d1, d2,
				opts,
				backupOpts)
		})
	}
}

func runGetPreviewLimits(
	t *testing.T,
	test backupLimitTest,
	drive1, drive2 *deltaDrive,
	opts control.Options,
	backupOpts control.BackupConfig,
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

	backupOpts.PreviewLimits = test.limits

	var (
		mbh       = defaultDriveBHWith(user, test.enumerator)
		c         = collWithMBHAndOpts(mbh, opts, backupOpts)
		errs      = fault.New(true)
		delList   = prefixmatcher.NewStringSetBuilder()
		collPaths = []string{}
	)

	cols, canUsePreviousBackup, err := c.Get(ctx, nil, delList, errs)
	require.NoError(t, err, clues.ToCore(err))
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

		expectItemIDs := test.expectedItemIDsInCollection[folderPath]

		if opts.ToggleFeatures.UseDeltaTree && test.expectedItemIDsInCollectionTree != nil {
			expectItemIDs = test.expectedItemIDsInCollectionTree[folderPath]
		}

		assert.ElementsMatchf(
			t,
			expectItemIDs,
			itemIDs,
			"item IDs in collection with path:\n\t%q",
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
	// the tree handling behavior may deviate under certain conditions
	// since it allows one file to slightly step over the byte limit
	numItemsTreePadding int
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
				numItemsTreePadding:  1,
			},
		},
	}
}

// TestGet_PreviewLimits_Defaults checks that default values are used when
// making a preview backup if the user didn't provide some options.
// These tests run a reduced set of checks that really just look for item counts
// and such. Other tests are expected to provide more comprehensive checks.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_defaultsNoTree() {
	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				control.DefaultOptions(),
				control.DefaultBackupConfig())
		})
	}
}

// TestGet_PreviewLimits_Defaults checks that default values are used when
// making a preview backup if the user didn't provide some options.
// These tests run a reduced set of checks that really just look for item counts
// and such. Other tests are expected to provide more comprehensive checks.
func (suite *LimiterUnitSuite) TestGet_PreviewLimits_defaultsWithTree() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	for _, test := range defaultLimitsTable() {
		suite.Run(test.name, func() {
			runGetPreviewLimitsDefaults(
				suite.T(),
				test,
				opts,
				control.DefaultBackupConfig())
		})
	}
}

func runGetPreviewLimitsDefaults(
	t *testing.T,
	test defaultLimitTest,
	opts control.Options,
	backupOpts control.BackupConfig,
) {
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

	ctx, flush := tester.NewContext(t)
	defer flush()

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making metadata path", clues.ToCore(err))

	d := drive()
	pages := make([]nextPage, 0, test.numContainers)

	for containerIdx := 0; containerIdx < test.numContainers; containerIdx++ {
		page := nextPage{
			Items: []models.DriveItemable{
				rootFolder(),
				driveItem(
					folderID(containerIdx),
					folderName(containerIdx),
					d.dir(),
					rootID,
					isFolder),
			},
		}

		for itemIdx := 0; itemIdx < test.numItemsPerContainer; itemIdx++ {
			itemSuffix := fmt.Sprintf("%d-%d", containerIdx, itemIdx)

			page.Items = append(page.Items, driveItemWSize(
				fileID(itemSuffix),
				fileName(itemSuffix),
				d.dir(folderName(containerIdx)),
				folderID(containerIdx),
				test.itemSize,
				isFile))
		}

		pages = append(pages, page)
	}

	backupOpts.PreviewLimits = test.limits

	var (
		mockEnumerator = driveEnumerator(
			d.newEnumer().with(
				delta(nil).with(pages...)))
		mbh           = defaultDriveBHWith(user, mockEnumerator)
		c             = collWithMBHAndOpts(mbh, opts, backupOpts)
		errs          = fault.New(true)
		delList       = prefixmatcher.NewStringSetBuilder()
		numContainers int
		numItems      int
	)

	cols, canUsePreviousBackup, err := c.Get(ctx, nil, delList, errs)
	require.NoError(t, err, clues.ToCore(err))
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
		expected := test.expect.numItemsPerContainer + 1
		if opts.ToggleFeatures.UseDeltaTree {
			expected += test.expect.numItemsTreePadding
		}

		assert.Len(
			t,
			col.driveItems,
			expected,
			"number of items in collection at:\n\t%+v",
			col.FullPath())
	}

	assert.Equal(
		t,
		test.expect.numContainers,
		numContainers,
		"total count of collections")

	// Add one to account for the folder permissions item.
	expected := test.expect.numItems + test.expect.numContainers
	if opts.ToggleFeatures.UseDeltaTree {
		expected += test.expect.numItemsTreePadding
	}

	// Each container also gets an item so account for that here.
	assert.Equal(
		t,
		expected,
		numItems,
		"total sum of item counts in all collections")
}
