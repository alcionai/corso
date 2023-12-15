package test_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type GroupsBackupIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsBackupIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groups() {
	var (
		resourceID = suite.its.group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	runBasicDriveishBackupTests(
		suite,
		path.GroupsService,
		control.DefaultOptions(),
		sel.Selector)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_incrementalGroups() {
	runGroupsIncrementalBackupTests(suite, suite.its, control.DefaultOptions())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_extensionsGroups() {
	var (
		resourceID = suite.its.group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	runDriveishBackupWithExtensionsTests(
		suite,
		path.GroupsService,
		control.DefaultOptions(),
		sel.Selector)
}

// ---------------------------------------------------------------------------
// test version using the tree-based drive item processor
// ---------------------------------------------------------------------------

type GroupsBackupTreeIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupTreeIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_groups() {
	var (
		resourceID = suite.its.group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runBasicDriveishBackupTests(
		suite,
		path.GroupsService,
		opts,
		sel.Selector)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_incrementalGroups() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	runGroupsIncrementalBackupTests(suite, suite.its, opts)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_extensionsGroups() {
	var (
		resourceID = suite.its.group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	runDriveishBackupWithExtensionsTests(
		suite,
		path.GroupsService,
		opts,
		sel.Selector)
}

// ---------------------------------------------------------------------------
// common backup test wrappers
// ---------------------------------------------------------------------------

func runGroupsIncrementalBackupTests(
	suite tester.Suite,
	its intgTesterSetup,
	opts control.Options,
) {
	sel := selectors.NewGroupsRestore([]string{its.group.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return its.group.RootSite.DriveID
	}

	gtsi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return its.group.RootSite.ID
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewSiteRestoreHandler(ac, path.GroupsService)
	}

	runIncrementalDriveishBackupTest(
		suite,
		opts,
		its.group.ID,
		its.user.ID,
		path.GroupsService,
		path.LibrariesCategory,
		ic,
		gtdi,
		gtsi,
		grh,
		true)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groups9VersionBumpBackup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		sel     = selectors.NewGroupsBackup([]string{suite.its.group.ID})
		opts    = control.DefaultOptions()
		whatSet = deeTD.CategoryFromRepoRef
	)

	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel),
		sel.Conversation(selectors.Any()))

	bo, bod := prepNewTestBackupOp(
		t,
		ctx,
		mb,
		sel.Selector,
		opts,
		version.All8MigrateUserPNToID,
		count.New())
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&bo,
		bod.sel,
		bod.sel.ID(),
		path.ChannelMessagesCategory)

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bo.Results.BackupID,
		bod.acct.ID(),
		bod.sel.ID(),
		path.GroupsService,
		whatSet,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		whatSet,
		bod.kms,
		bod.sss,
		expectDeets,
		false)

	mb = evmock.NewBus()
	forcedFull := newTestBackupOp(
		t,
		ctx,
		bod,
		mb,
		opts,
		count.New())
	forcedFull.BackupVersion = version.Groups9Update

	runAndCheckBackup(t, ctx, &forcedFull, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&forcedFull,
		bod.sel,
		bod.sel.ID(),
		path.ChannelMessagesCategory)

	_, expectDeets = deeTD.GetDeetsInBackup(
		t,
		ctx,
		forcedFull.Results.BackupID,
		bod.acct.ID(),
		bod.sel.ID(),
		path.GroupsService,
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

	// The number of items backed up in the forced full backup should be roughly
	// the same as the number of items in the original backup.
	assert.Equal(
		t,
		bo.Results.Counts[string(count.PersistedNonCachedFiles)],
		forcedFull.Results.Counts[string(count.PersistedNonCachedFiles)],
		"items written")
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groupsBasic() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		sel     = selectors.NewGroupsBackup([]string{suite.its.group.ID})
		opts    = control.DefaultOptions()
		whatSet = deeTD.CategoryFromRepoRef
	)

	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel),
		sel.Conversation(selectors.Any()))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&bo,
		bod.sel,
		bod.sel.ID(),
		path.ChannelMessagesCategory)

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bo.Results.BackupID,
		bod.acct.ID(),
		bod.sel.ID(),
		path.GroupsService,
		whatSet,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		whatSet,
		bod.kms,
		bod.sss,
		expectDeets,
		false)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groupsExtensions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		sel     = selectors.NewGroupsBackup([]string{suite.its.group.ID})
		opts    = control.DefaultOptions()
		tenID   = tconfig.M365TenantID(t)
		svc     = path.GroupsService
		ws      = deeTD.DriveIDFromRepoRef
	)

	opts.ItemExtensionFactory = getTestExtensionFactories()

	// does not apply to channel messages
	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&bo,
		bod.sel,
		bod.sel.ID(),
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
			verifyExtensionData(t, ent.ItemInfo, path.GroupsService)
		}
	}
}

type GroupsBackupNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9MergeBase() {
	sel := selectors.NewGroupsBackup([]string{suite.its.group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel))

	runMergeBaseGroupsUpdate(suite, sel.Selector, false)
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9AssistBases() {
	sel := selectors.NewGroupsBackup([]string{suite.its.group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel))

	runDriveAssistBaseGroupsUpdate(suite, sel.Selector, false)
}

type GroupsRestoreNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsRestoreNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsRestoreNightlyIntgSuite) TestRestore_Run_groupsWithAdvancedOptions() {
	sel := selectors.NewGroupsBackup([]string{suite.its.group.ID})
	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))
	sel.Filter(sel.Library("documents"))
	sel.DiscreteOwner = suite.its.group.ID

	runDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.its.ac,
		sel.Selector,
		suite.its.group.RootSite.DriveID,
		suite.its.group.RootSite.DriveRootFolderID)
}
