package test_test

import (
	"context"
	"io"
	"sync/atomic"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type OneDriveBackupIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveBackupIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDrive() {
	var (
		resourceID = suite.its.secondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	runBasicDriveishBackupTests(
		suite,
		path.OneDriveService,
		control.DefaultOptions(),
		sel.Selector)
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_incrementalOneDrive() {
	runOneDriveIncrementalBackupTests(suite, suite.its, control.DefaultOptions())
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_extensionsOneDrive() {
	var (
		resourceID = suite.its.secondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	runDriveishBackupWithExtensionsTests(
		suite,
		path.OneDriveService,
		control.DefaultOptions(),
		sel.Selector)
}

// ---------------------------------------------------------------------------
// test version using the tree-based drive item processor
// ---------------------------------------------------------------------------

type OneDriveBackupTreeIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupTreeIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeOneDrive() {
	var (
		resourceID = suite.its.secondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runBasicDriveishBackupTests(
		suite,
		path.OneDriveService,
		opts,
		sel.Selector)
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeIncrementalOneDrive() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	runOneDriveIncrementalBackupTests(suite, suite.its, opts)
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeExtensionsOneDrive() {
	var (
		resourceID = suite.its.secondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runDriveishBackupWithExtensionsTests(
		suite,
		path.OneDriveService,
		opts,
		sel.Selector)
}

// ---------------------------------------------------------------------------
// common backup test wrappers
// ---------------------------------------------------------------------------

func runOneDriveIncrementalBackupTests(
	suite tester.Suite,
	its intgTesterSetup,
	opts control.Options,
) {
	sel := selectors.NewOneDriveRestore([]string{its.user.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.Folders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		d, err := its.ac.Users().GetDefaultDrive(ctx, its.user.ID)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default user drive").
				With("user", its.user.ID)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewUserDriveRestoreHandler(ac)
	}

	runIncrementalDriveishBackupTest(
		suite,
		opts,
		its.user.ID,
		its.user.ID,
		path.OneDriveService,
		path.FilesCategory,
		ic,
		gtdi,
		nil,
		grh,
		false)
}

// ---------------------------------------------------------------------------
// other drive tests
// ---------------------------------------------------------------------------

var (
	_ io.ReadCloser                    = &failFirstRead{}
	_ extensions.CreateItemExtensioner = &createFailFirstRead{}
)

// failFirstRead fails the first read on a file being uploaded during a
// snapshot. Only one file is failed during the snapshot even if it the snapshot
// contains multiple files.
type failFirstRead struct {
	firstFile *atomic.Bool
	io.ReadCloser
}

func (e *failFirstRead) Read(p []byte) (int, error) {
	if e.firstFile.CompareAndSwap(true, false) {
		// This is the first file being read, return an error for it.
		return 0, clues.New("injected error for testing")
	}

	return e.ReadCloser.Read(p)
}

func newCreateSingleFileFailExtension() *createFailFirstRead {
	firstItem := &atomic.Bool{}
	firstItem.Store(true)

	return &createFailFirstRead{
		firstItem: firstItem,
	}
}

type createFailFirstRead struct {
	firstItem *atomic.Bool
}

func (ce *createFailFirstRead) CreateItemExtension(
	_ context.Context,
	r io.ReadCloser,
	_ details.ItemInfo,
	_ *details.ExtensionData,
) (io.ReadCloser, error) {
	return &failFirstRead{
		firstFile:  ce.firstItem,
		ReadCloser: r,
	}, nil
}

func runDriveAssistBaseGroupsUpdate(
	suite tester.Suite,
	sel selectors.Selector,
	expectCached bool,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		whatSet = deeTD.CategoryFromRepoRef
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	opts.ToggleFeatures.UseDeltaTree = true
	opts.ItemExtensionFactory = []extensions.CreateItemExtensioner{
		newCreateSingleFileFailExtension(),
	}

	// Creating out here so bod lasts for full test and isn't closed until the
	// test is compltely done.
	bo, bod := prepNewTestBackupOp(
		t,
		ctx,
		mb,
		sel,
		opts,
		version.All8MigrateUserPNToID,
		counter)
	defer bod.close(t, ctx)

	suite.Run("makeAssistBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		// Need to run manually cause runAndCheckBackup assumes success for the most
		// part.
		err := bo.Run(ctx)
		assert.Error(t, err, clues.ToCore(err))
		assert.NotEmpty(t, bo.Results, "backup had non-zero results")
		assert.NotEmpty(t, bo.Results.BackupID, "backup generated an ID")
		assert.NotZero(t, bo.Results.ItemsWritten)

		// TODO(ashmrtn): Check that the base is marked as an assist base.
		t.Logf("base error: %v\n", err)
	})

	// Don't run the below if we've already failed since it won't make sense
	// anymore.
	if suite.T().Failed() {
		return
	}

	suite.Run("makeIncrementalBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		var (
			mb      = evmock.NewBus()
			counter = count.New()
			opts    = control.DefaultOptions()
		)

		forcedFull := newTestBackupOp(
			t,
			ctx,
			bod,
			mb,
			opts,
			counter)
		forcedFull.BackupVersion = version.Groups9Update

		runAndCheckBackup(t, ctx, &forcedFull, mb, false)

		reasons, err := bod.sel.Reasons(bod.acct.ID(), false)
		require.NoError(t, err, clues.ToCore(err))

		for _, reason := range reasons {
			checkBackupIsInManifests(
				t,
				ctx,
				bod.kw,
				bod.sw,
				&forcedFull,
				bod.sel,
				bod.sel.ID(),
				reason.Category())
		}

		_, expectDeets := deeTD.GetDeetsInBackup(
			t,
			ctx,
			forcedFull.Results.BackupID,
			bod.acct.ID(),
			bod.sel.ID(),
			bod.sel.PathService(),
			whatSet,
			bod.kms,
			bod.sss)
		deeTD.CheckBackupDetails(
			t,
			ctx,
			forcedFull.Results.BackupID,
			whatSet,
			bod.kms,
			bod.sss,
			expectDeets,
			false)

		// For groups the forced full backup shouldn't have any cached items. For
		// OneDrive and SharePoint it should since they shouldn't be forcing full
		// backups.
		cachedCheck := assert.NotZero
		if !expectCached {
			cachedCheck = assert.Zero
		}

		cachedCheck(
			t,
			forcedFull.Results.Counts[string(count.PersistedCachedFiles)],
			"kopia cached items")
	})
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDriveOwnerMigration() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct    = tconfig.NewM365Account(t)
		opts    = control.DefaultOptions()
		mb      = evmock.NewBus()
		counter = count.New()

		categories = map[path.CategoryType][][]string{
			path.FilesCategory: {{bupMD.DeltaURLsFileName}, {bupMD.PreviousPathFileName}},
		}
	)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ctrl, err := m365.NewController(
		ctx,
		acct,
		path.OneDriveService,
		control.DefaultOptions(),
		counter)
	require.NoError(t, err, clues.ToCore(err))

	userable, err := ctrl.AC.Users().GetByID(
		ctx,
		suite.its.user.ID,
		api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	uid := ptr.Val(userable.GetId())
	uname := ptr.Val(userable.GetUserPrincipalName())

	oldsel := selectors.NewOneDriveBackup([]string{uname})
	oldsel.Include(selTD.OneDriveBackupFolderScope(oldsel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, oldsel.Selector, opts, 0, counter)
	defer bod.close(t, ctx)

	sel := bod.sel

	// ensure the initial owner uses name in both cases
	bo.ResourceOwner = sel.SetDiscreteOwnerIDName(uname, uname)
	// required, otherwise we don't run the migration
	bo.BackupVersion = version.All8MigrateUserPNToID - 1

	require.Equalf(
		t,
		bo.ResourceOwner.Name(),
		bo.ResourceOwner.ID(),
		"historical representation of user id [%s] should match pn [%s]",
		bo.ResourceOwner.ID(),
		bo.ResourceOwner.Name())

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	newsel := selectors.NewOneDriveBackup([]string{uid})
	newsel.Include(selTD.OneDriveBackupFolderScope(newsel))
	sel = newsel.SetDiscreteOwnerIDName(uid, uname)

	var (
		incMB = evmock.NewBus()
		// the incremental backup op should have a proper user ID for the id.
		incBO = newTestBackupOp(t, ctx, bod, incMB, opts, counter)
	)

	require.NotEqualf(
		t,
		incBO.ResourceOwner.Name(),
		incBO.ResourceOwner.ID(),
		"current representation of user: id [%s] should differ from PN [%s]",
		incBO.ResourceOwner.ID(),
		incBO.ResourceOwner.Name())

	err = incBO.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&incBO,
		sel,
		uid,
		maps.Keys(categories)...)
	checkMetadataFilesExist(
		t,
		ctx,
		incBO.Results.BackupID,
		bod.kw,
		bod.kms,
		creds.AzureTenantID,
		uid,
		path.OneDriveService,
		categories)

	// 2 on read/writes to account for metadata: 1 delta and 1 path.
	assert.LessOrEqual(t, 2, incBO.Results.ItemsWritten, "items written")
	assert.LessOrEqual(t, 1, incBO.Results.NonMetaItemsWritten, "non meta items written")
	assert.LessOrEqual(t, 2, incBO.Results.ItemsRead, "items read")
	assert.NoError(t, incBO.Errors.Failure(), "non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
	assert.Empty(t, incBO.Errors.Recovered(), "recoverable/iteration errors")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "backup-end events")

	bid := incBO.Results.BackupID
	bup := &backup.Backup{}

	err = bod.kms.Get(ctx, model.BackupSchema, bid, bup)
	require.NoError(t, err, clues.ToCore(err))

	var (
		ssid  = bup.StreamStoreID
		deets details.Details
		ss    = streamstore.NewStreamer(bod.kw, creds.AzureTenantID, path.OneDriveService)
	)

	err = ss.Read(ctx, ssid, streamstore.DetailsReader(details.UnmarshalTo(&deets)), fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	for _, ent := range deets.Entries {
		// 46 is the tenant uuid + "onedrive" + two slashes
		if len(ent.RepoRef) > 46 {
			assert.Contains(t, ent.RepoRef, uid)
		}
	}
}

type OneDriveBackupNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveBackupNightlyIntgSuite) TestBackup_Run_oneDriveVersion9MergeBase() {
	sel := selectors.NewOneDriveBackup([]string{suite.its.user.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	runMergeBaseGroupsUpdate(suite, sel.Selector, true)
}

//func (suite *OneDriveBackupNightlyIntgSuite) TestBackup_Run_oneDriveVersion9AssistBases() {
//	sel := selectors.NewOneDriveBackup([]string{tconfig.SecondaryM365UserID(suite.T())})
//	sel.Include(selTD.OneDriveBackupFolderScope(sel))
//
//	runDriveAssistBaseGroupsUpdate(suite, sel.Selector, true)
//}

type OneDriveRestoreNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveRestoreNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveRestoreNightlyIntgSuite) TestRestore_Run_onedriveWithAdvancedOptions() {
	sel := selectors.NewOneDriveBackup([]string{suite.its.user.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))
	sel.DiscreteOwner = suite.its.user.ID

	runDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.its.ac,
		sel.Selector,
		suite.its.user.DriveID,
		suite.its.user.DriveRootFolderID)
}

func runDriveRestoreWithAdvancedOptions(
	t *testing.T,
	suite tester.Suite,
	ac api.Client,
	sel selectors.Selector, // both Restore and Backup types work.
	driveID, rootFolderID string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	var (
		restoreCfg          = ctrlTD.DefaultRestoreConfig("drive_adv_restore")
		containerID         string
		countItemsInRestore int
		collKeys            = map[string]api.DriveItemIDType{}
		fileIDs             map[string]api.DriveItemIDType
		acd                 = ac.Drives()
	)

	// initial restore

	suite.Run("baseline", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		contGC, err := acd.GetFolderByName(ctx, driveID, rootFolderID, restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		// the folder containing the files is a child of the folder created by the restore.
		contGC, err = acd.GetFolderByName(ctx, driveID, ptr.Val(contGC.GetId()), selTD.TestFolderName)
		require.NoError(t, err, clues.ToCore(err))

		containerID = ptr.Val(contGC.GetId())

		collKeys, err = acd.GetItemsInContainerByCollisionKey(
			ctx,
			driveID,
			containerID)
		require.NoError(t, err, clues.ToCore(err))

		countItemsInRestore = len(collKeys)

		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)

		fileIDs, err = acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))
	})

	// skip restore

	suite.Run("skip collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Skip

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)

		checkRestoreCounts(t, ctr, countItemsInRestore, 0, 0)
		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")

		// get all files in folder, use these as the base
		// set of files to compare against.

		result := filterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, 0, "no new items should get added")

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, fileIDs, currentFileIDs, "ids are equal")
	})

	// replace restore

	suite.Run("replace collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Replace

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		checkRestoreCounts(t, ctr, 0, countItemsInRestore, 0)
		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been replaced")

		result := filterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, 0, "all items should have been replaced")

		for k, v := range result {
			assert.NotEqual(t, v, collKeys[k], "replaced items should have new IDs")
		}

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, len(fileIDs), len(currentFileIDs), "count of ids ids are equal")
		for orig := range fileIDs {
			assert.NotContains(t, currentFileIDs, orig, "original item should not exist after replacement")
		}

		fileIDs = currentFileIDs
	})

	// copy restore

	suite.Run("copy collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)
		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been copied")

		result := filterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, len(collKeys), "all items should have been added as copies")

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, 2*len(fileIDs), len(currentFileIDs), "count of ids should be double from before")
		assert.Subset(t, maps.Keys(currentFileIDs), maps.Keys(fileIDs), "original item should exist after copy")
	})
}

func (suite *OneDriveRestoreNightlyIntgSuite) TestRestore_Run_onedriveAlternateProtectedResource() {
	sel := selectors.NewOneDriveBackup([]string{suite.its.user.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))
	sel.DiscreteOwner = suite.its.user.ID

	runDriveRestoreToAlternateProtectedResource(
		suite.T(),
		suite,
		suite.its.ac,
		sel.Selector,
		suite.its.user,
		suite.its.secondaryUser,
		suite.its.secondaryUser.ID)
}

func runDriveRestoreToAlternateProtectedResource(
	t *testing.T,
	suite tester.Suite,
	ac api.Client,
	sel selectors.Selector, // owner should match 'from', both Restore and Backup types work.
	driveFrom, driveTo ids,
	toResource string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	var (
		restoreCfg        = ctrlTD.DefaultRestoreConfig("drive_restore_to_resource")
		fromCollisionKeys map[string]api.DriveItemIDType
		fromItemIDs       map[string]api.DriveItemIDType
		acd               = ac.Drives()
	)

	// first restore to the 'from' resource

	suite.Run("restore original resource", func() {
		mb = evmock.NewBus()
		fromCtr := count.New()
		driveID := driveFrom.DriveID
		rootFolderID := driveFrom.DriveRootFolderID
		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			fromCtr,
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		fromItemIDs, fromCollisionKeys = getDriveCollKeysAndItemIDs(
			t,
			ctx,
			acd,
			driveID,
			rootFolderID,
			restoreCfg.Location,
			selTD.TestFolderName)
	})

	// then restore to the 'to' resource
	var (
		toCollisionKeys map[string]api.DriveItemIDType
		toItemIDs       map[string]api.DriveItemIDType
	)

	suite.Run("restore to alternate resource", func() {
		mb = evmock.NewBus()
		toCtr := count.New()
		driveID := driveTo.DriveID
		rootFolderID := driveTo.DriveRootFolderID
		restoreCfg.ProtectedResource = toResource

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			toCtr,
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		toItemIDs, toCollisionKeys = getDriveCollKeysAndItemIDs(
			t,
			ctx,
			acd,
			driveID,
			rootFolderID,
			restoreCfg.Location,
			selTD.TestFolderName)
	})

	// compare restore results
	assert.Equal(t, len(fromItemIDs), len(toItemIDs))
	assert.ElementsMatch(t, maps.Keys(fromCollisionKeys), maps.Keys(toCollisionKeys))
}

type GetItemsKeysAndFolderByNameer interface {
	GetItemIDsInContainer(
		ctx context.Context,
		driveID, containerID string,
	) (map[string]api.DriveItemIDType, error)
	GetFolderByName(
		ctx context.Context,
		driveID, parentFolderID, folderName string,
	) (models.DriveItemable, error)
	GetItemsInContainerByCollisionKey(
		ctx context.Context,
		driveID, containerID string,
	) (map[string]api.DriveItemIDType, error)
}

func getDriveCollKeysAndItemIDs(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	gikafbn GetItemsKeysAndFolderByNameer,
	driveID, parentContainerID string,
	containerNames ...string,
) (map[string]api.DriveItemIDType, map[string]api.DriveItemIDType) {
	var (
		c   models.DriveItemable
		err error
		cID string
	)

	for _, cn := range containerNames {
		pcid := parentContainerID

		if len(cID) != 0 {
			pcid = cID
		}

		c, err = gikafbn.GetFolderByName(ctx, driveID, pcid, cn)
		require.NoError(t, err, clues.ToCore(err))

		cID = ptr.Val(c.GetId())
	}

	itemIDs, err := gikafbn.GetItemIDsInContainer(ctx, driveID, cID)
	require.NoError(t, err, clues.ToCore(err))

	collisionKeys, err := gikafbn.GetItemsInContainerByCollisionKey(ctx, driveID, cID)
	require.NoError(t, err, clues.ToCore(err))

	return itemIDs, collisionKeys
}
