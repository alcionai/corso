package test_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type SharePointBackupIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSharePointBackupIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointBackupIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SharePointBackupIntgSuite) TestBackup_Run_sharePoint() {
	var (
		resourceID = suite.its.site.ID
		sel        = selectors.NewSharePointBackup([]string{resourceID})
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	runBasicDriveishBackupTests(
		suite,
		path.SharePointService,
		control.DefaultOptions(),
		sel.Selector)
}

func (suite *SharePointBackupIntgSuite) TestBackup_Run_incrementalSharePoint() {
	runSharePointIncrementalBackupTests(suite, suite.its, control.DefaultOptions())
}

func (suite *SharePointBackupIntgSuite) TestBackup_Run_extensionsSharePoint() {
	var (
		resourceID = suite.its.site.ID
		sel        = selectors.NewSharePointBackup([]string{resourceID})
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	runDriveishBackupWithExtensionsTests(
		suite,
		path.SharePointService,
		control.DefaultOptions(),
		sel.Selector)
}

// ---------------------------------------------------------------------------
// test version using the tree-based drive item processor
// ---------------------------------------------------------------------------

type SharePointBackupTreeIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSharePointBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointBackupTreeIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SharePointBackupTreeIntgSuite) TestBackup_Run_sharePoint() {
	var (
		resourceID = suite.its.site.ID
		sel        = selectors.NewSharePointBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runBasicDriveishBackupTests(
		suite,
		path.SharePointService,
		opts,
		sel.Selector)
}

func (suite *SharePointBackupTreeIntgSuite) TestBackup_Run_incrementalSharePoint() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	runSharePointIncrementalBackupTests(suite, suite.its, opts)
}

func (suite *SharePointBackupTreeIntgSuite) TestBackup_Run_extensionsSharePoint() {
	var (
		resourceID = suite.its.site.ID
		sel        = selectors.NewSharePointBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runDriveishBackupWithExtensionsTests(
		suite,
		path.SharePointService,
		opts,
		sel.Selector)
}

// ---------------------------------------------------------------------------
// common backup test wrappers
// ---------------------------------------------------------------------------

func runSharePointIncrementalBackupTests(
	suite tester.Suite,
	its intgTesterSetup,
	opts control.Options,
) {
	sel := selectors.NewSharePointRestore([]string{its.site.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		d, err := its.ac.Sites().GetDefaultDrive(ctx, its.site.ID)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default site drive").
				With("site", its.site.ID)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewSiteRestoreHandler(ac, path.SharePointService)
	}

	runIncrementalDriveishBackupTest(
		suite,
		opts,
		its.site.ID,
		its.user.ID,
		path.SharePointService,
		path.LibrariesCategory,
		ic,
		gtdi,
		nil,
		grh,
		true)
}

type SharePointBackupNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSharePointBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointBackupNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SharePointBackupNightlyIntgSuite) TestBackup_Run_sharePointVersion9MergeBase() {
	sel := selectors.NewSharePointBackup([]string{suite.its.site.ID})
	sel.Include(selTD.SharePointBackupFolderScope(sel))

	runMergeBaseGroupsUpdate(suite, sel.Selector, true)
}

func (suite *SharePointBackupNightlyIntgSuite) TestBackup_Run_sharePointVersion9AssistBases() {
	sel := selectors.NewSharePointBackup([]string{suite.its.site.ID})
	sel.Include(selTD.SharePointBackupFolderScope(sel))

	runDriveAssistBaseGroupsUpdate(suite, sel.Selector, true)
}

type SharePointRestoreNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSharePointRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointRestoreNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SharePointRestoreNightlyIntgSuite) TestRestore_Run_sharepointWithAdvancedOptions() {
	sel := selectors.NewSharePointBackup([]string{suite.its.site.ID})
	sel.Include(selTD.SharePointBackupFolderScope(sel))
	sel.Filter(sel.Library("documents"))
	sel.DiscreteOwner = suite.its.site.ID

	runDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.its.ac,
		sel.Selector,
		suite.its.site.DriveID,
		suite.its.site.DriveRootFolderID)
}

func (suite *SharePointRestoreNightlyIntgSuite) TestRestore_Run_sharepointAlternateProtectedResource() {
	sel := selectors.NewSharePointBackup([]string{suite.its.site.ID})
	sel.Include(selTD.SharePointBackupFolderScope(sel))
	sel.Filter(sel.Library("documents"))
	sel.DiscreteOwner = suite.its.site.ID

	runDriveRestoreToAlternateProtectedResource(
		suite.T(),
		suite,
		suite.its.ac,
		sel.Selector,
		suite.its.site,
		suite.its.secondarySite,
		suite.its.secondarySite.ID)
}

func (suite *SharePointRestoreNightlyIntgSuite) TestRestore_Run_sharepointDeletedDrives() {
	t := suite.T()

	// despite the client having a method for drive.Patch and drive.Delete, both only return
	// the error code and message `invalidRequest`.
	t.Skip("graph api doesn't allow patch or delete on drives, so we cannot run any conditions")

	ctx, flush := tester.NewContext(t)
	defer flush()

	rc := ctrlTD.DefaultRestoreConfig("restore_deleted_drives")
	rc.OnCollision = control.Copy

	// create a new drive
	md, err := suite.its.ac.Lists().PostDrive(ctx, suite.its.site.ID, rc.Location)
	require.NoError(t, err, clues.ToCore(err))

	driveID := ptr.Val(md.GetId())

	// get the root folder
	mdi, err := suite.its.ac.Drives().GetRootFolder(ctx, driveID)
	require.NoError(t, err, clues.ToCore(err))

	rootFolderID := ptr.Val(mdi.GetId())

	// add an item to it
	itemName := uuid.NewString()

	item := models.NewDriveItem()
	item.SetName(ptr.To(itemName + ".txt"))

	file := models.NewFile()
	item.SetFile(file)

	_, err = suite.its.ac.Drives().PostItemInContainer(
		ctx,
		driveID,
		rootFolderID,
		item,
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))

	// run a backup
	var (
		mb          = evmock.NewBus()
		counter     = count.New()
		opts        = control.DefaultOptions()
		graphClient = suite.its.ac.Stable.Client()
	)

	bsel := selectors.NewSharePointBackup([]string{suite.its.site.ID})
	bsel.Include(selTD.SharePointBackupFolderScope(bsel))
	bsel.Filter(bsel.Library(rc.Location))
	bsel.DiscreteOwner = suite.its.site.ID

	bo, bod := prepNewTestBackupOp(t, ctx, mb, bsel.Selector, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	// test cases:

	// first test, we take the current drive and rename it.
	// the restore should find the drive by id and restore items
	// into it like normal.  Due to collision handling, this should
	// create a copy of the current item.
	suite.Run("renamed drive", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		patchBody := models.NewDrive()
		patchBody.SetName(ptr.To("some other name"))

		md, err = graphClient.
			Drives().
			ByDriveId(driveID).
			Patch(ctx, patchBody, nil)
		require.NoError(t, err, clues.ToCore(graph.Stack(ctx, err)))

		var (
			mb  = evmock.NewBus()
			ctr = count.New()
		)

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			bod.sel,
			opts,
			rc)

		runAndCheckRestore(t, ctx, &ro, mb, false)
		assert.Equal(t, 1, ctr.Get(count.NewItemCreated), "restored an item")

		resp, err := graphClient.
			Drives().
			ByDriveId(driveID).
			Items().
			ByDriveItemId(rootFolderID).
			Children().
			Get(ctx, nil)
		require.NoError(t, err, clues.ToCore(graph.Stack(ctx, err)))

		items := resp.GetValue()
		assert.Len(t, items, 2)

		for _, item := range items {
			assert.Contains(t, ptr.Val(item.GetName()), itemName)
		}
	})

	// second test, we delete the drive altogether.  the restore should find
	// no existing drives, but it should have the old drive's name and attempt
	// to recreate that drive by name.
	suite.Run("deleted drive", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		err = graphClient.
			Drives().
			ByDriveId(driveID).
			Delete(ctx, nil)
		require.NoError(t, err, clues.ToCore(graph.Stack(ctx, err)))

		var (
			mb  = evmock.NewBus()
			ctr = count.New()
		)

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			bod.sel,
			opts,
			rc)

		runAndCheckRestore(t, ctx, &ro, mb, false)
		assert.Equal(t, 1, ctr.Get(count.NewItemCreated), "restored an item")

		pgr := suite.its.ac.
			Drives().
			NewSiteDrivePager(suite.its.site.ID, []string{"id", "name"})

		drives, err := api.GetAllDrives(ctx, pgr)
		require.NoError(t, err, clues.ToCore(err))

		var created models.Driveable

		for _, drive := range drives {
			if ptr.Val(drive.GetName()) == ptr.Val(created.GetName()) &&
				ptr.Val(drive.GetId()) != driveID {
				created = drive
				break
			}
		}

		require.NotNil(t, created, "found the restored drive by name")
		md = created
		driveID = ptr.Val(md.GetId())

		mdi, err := suite.its.ac.Drives().GetRootFolder(ctx, driveID)
		require.NoError(t, err, clues.ToCore(err))

		rootFolderID = ptr.Val(mdi.GetId())

		resp, err := graphClient.
			Drives().
			ByDriveId(driveID).
			Items().
			ByDriveItemId(rootFolderID).
			Children().
			Get(ctx, nil)
		require.NoError(t, err, clues.ToCore(graph.Stack(ctx, err)))

		items := resp.GetValue()
		assert.Len(t, items, 1)

		assert.Equal(t, ptr.Val(items[0].GetName()), itemName+".txt")
	})

	// final test, run a follow-up restore.  This should match the
	// drive we created in the prior test by name, but not by ID.
	suite.Run("different drive - same name", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		var (
			mb  = evmock.NewBus()
			ctr = count.New()
		)

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			bod.sel,
			opts,
			rc)

		runAndCheckRestore(t, ctx, &ro, mb, false)

		assert.Equal(t, 1, ctr.Get(count.NewItemCreated), "restored an item")

		resp, err := graphClient.
			Drives().
			ByDriveId(driveID).
			Items().
			ByDriveItemId(rootFolderID).
			Children().
			Get(ctx, nil)
		require.NoError(t, err, clues.ToCore(graph.Stack(ctx, err)))

		items := resp.GetValue()
		assert.Len(t, items, 2)

		for _, item := range items {
			assert.Contains(t, ptr.Val(item.GetName()), itemName)
		}
	})
}
