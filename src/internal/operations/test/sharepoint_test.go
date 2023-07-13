package test_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/m365/sharepoint"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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

func (suite *SharePointBackupIntgSuite) TestBackup_Run_incrementalSharePoint() {
	sel := selectors.NewSharePointRestore([]string{suite.its.siteID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		d, err := suite.its.ac.Sites().GetDefaultDrive(ctx, suite.its.siteID)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default site drive").
				With("site", suite.its.siteID)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	grh := func(ac api.Client) onedrive.RestoreHandler {
		return sharepoint.NewRestoreHandler(ac)
	}

	runDriveIncrementalTest(
		suite,
		suite.its.siteID,
		suite.its.userID,
		resource.Sites,
		path.SharePointService,
		path.LibrariesCategory,
		ic,
		gtdi,
		grh,
		true)
}

func (suite *SharePointBackupIntgSuite) TestBackup_Run_sharePoint() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb   = evmock.NewBus()
		sel  = selectors.NewSharePointBackup([]string{suite.its.siteID})
		opts = control.Defaults()
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&bo,
		bod.sel,
		suite.its.siteID,
		path.LibrariesCategory)
}

func (suite *SharePointBackupIntgSuite) TestBackup_Run_sharePointExtensions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb    = evmock.NewBus()
		sel   = selectors.NewSharePointBackup([]string{suite.its.siteID})
		opts  = control.Defaults()
		tenID = tconfig.M365TenantID(t)
		svc   = path.SharePointService
		ws    = deeTD.DriveIDFromRepoRef
	)

	opts.ItemExtensionFactory = getTestExtensionFactories()

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&bo,
		bod.sel,
		suite.its.siteID,
		path.LibrariesCategory)

	bID := bo.Results.BackupID

	deets, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel.ID(),
		svc,
		ws,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bod.kms,
		bod.sss,
		expectDeets,
		false)

	// Check that the extensions are in the backup
	for _, ent := range deets.Entries {
		if ent.Folder == nil {
			verifyExtensionData(t, ent.ItemInfo, path.SharePointService)
		}
	}
}

type SharePointRestoreIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestSharePointRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointRestoreIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *SharePointRestoreIntgSuite) TestRestore_Run_sharepointWithAdvancedOptions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	baseSel := selectors.NewSharePointBackup([]string{suite.its.siteID})
	baseSel.Include(selTD.SharePointBackupFolderScope(baseSel))
	baseSel.Filter(baseSel.Library("documents"))

	baseSel.DiscreteOwner = suite.its.siteID

	var (
		mb   = evmock.NewBus()
		opts = control.Defaults()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, baseSel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	rsel, err := bod.sel.ToSharePointRestore()
	require.NoError(t, err, clues.ToCore(err))

	var (
		restoreCfg          = ctrlTD.DefaultRestoreConfig("sharepoint_adv_restore")
		sel                 = rsel.Selector
		siteDriveID         = suite.its.siteDriveID
		containerID         string
		countItemsInRestore int
		collKeys            = map[string]api.DriveCollisionItem{}
		acd                 = suite.its.ac.Drives()
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
		contGC, err := acd.GetFolderByName(ctx, siteDriveID, suite.its.siteDriveRootFolderID, restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		// the restored items are in the child of the newly created folder
		contGC, err = acd.GetFolderByName(ctx, siteDriveID, ptr.Val(contGC.GetId()), selTD.TestFolderName)
		require.NoError(t, err, clues.ToCore(err))

		containerID = ptr.Val(contGC.GetId())

		collKeys, err = acd.GetItemsInContainerByCollisionKey(
			ctx,
			siteDriveID,
			containerID)
		require.NoError(t, err, clues.ToCore(err))

		countItemsInRestore = len(collKeys)

		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)
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

		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")
		checkRestoreCounts(t, ctr, countItemsInRestore, 0, 0)

		// get all files in folder, use these as the base
		// set of files to compare against.

		result := filterCollisionKeyResults(t, ctx,
			siteDriveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, 0, "no new items should get added")
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

		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been replaced")
		checkRestoreCounts(t, ctr, 0, countItemsInRestore, 0)

		result := filterCollisionKeyResults(
			t,
			ctx,
			siteDriveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, 0, "all items should have been replaced")

		for k, v := range result {
			assert.NotEqual(t, v, collKeys[k], "replaced items should have new IDs")
		}
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

		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been copied")
		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)

		result := filterCollisionKeyResults(
			t,
			ctx,
			siteDriveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, len(collKeys), "all items should have been added as copies")
	})
}
