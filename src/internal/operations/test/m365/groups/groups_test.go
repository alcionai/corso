package groups_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	. "github.com/alcionai/corso/src/internal/operations/test/m365"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
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
	m365 its.M365IntgTestSetup
}

func TestGroupsBackupIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groups() {
	var (
		resourceID = suite.m365.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	RunBasicDriveishBackupTests(
		suite,
		path.GroupsService,
		control.DefaultOptions(),
		sel.Selector)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_incrementalGroups() {
	runGroupsIncrementalBackupTests(suite, suite.m365, control.DefaultOptions())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_extensionsGroups() {
	var (
		resourceID = suite.m365.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
	)

	// This test does not apply to channel messages or conversations.
	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	RunDriveishBackupWithExtensionsTests(
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
	m365 its.M365IntgTestSetup
}

func TestGroupsBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupTreeIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeGroups() {
	var (
		resourceID = suite.m365.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	RunBasicDriveishBackupTests(
		suite,
		path.GroupsService,
		opts,
		sel.Selector)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeIncrementalGroups() {
	opts := control.DefaultOptions()
	runGroupsIncrementalBackupTests(suite, suite.m365, opts)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeExtensionsGroups() {
	var (
		resourceID = suite.m365.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	RunDriveishBackupWithExtensionsTests(
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
	m365 its.M365IntgTestSetup,
	opts control.Options,
) {
	sel := selectors.NewGroupsRestore([]string{m365.Group.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return m365.Group.RootSite.DriveID
	}

	gtsi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return m365.Group.RootSite.ID
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewSiteRestoreHandler(ac, path.GroupsService)
	}

	RunIncrementalDriveishBackupTest(
		suite,
		opts,
		m365.Group.ID,
		m365.SecondaryGroup.ID, // more reliable than user
		path.GroupsService,
		path.LibrariesCategory,
		ic,
		gtdi,
		gtsi,
		grh,
		true)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groupsBasic() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		sel     = selectors.NewGroupsBackup([]string{suite.m365.Group.ID})
		opts    = control.DefaultOptions()
		whatSet = deeTD.CategoryFromRepoRef
	)

	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel))

	// TODO(pandeyabs): CorsoCITeam group mailbox backup is currently broken because of invalid
	// odata.NextLink which causes an infinite loop during paging. Disabling conversations tests while
	// we go fix the group mailbox.
	// selTD.GroupsBackupConversationScope(sel))

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	reasons, err := bod.Sel.Reasons(bod.Acct.ID(), false)
	require.NoError(t, err, clues.ToCore(err))

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	for _, reason := range reasons {
		CheckBackupIsInManifests(
			t,
			ctx,
			bod.KW,
			bod.SW,
			&bo,
			bod.Sel,
			bod.Sel.ID(),
			reason.Category())
	}

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bo.Results.BackupID,
		bod.Acct.ID(),
		bod.Sel.ID(),
		path.GroupsService,
		whatSet,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		whatSet,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)

	// Basic, happy path incremental test.  No changes are dictated or expected.
	// This only tests that an incremental backup is runnable at all, and that it
	// produces fewer results than the last backup.
	//
	// Incremental testing for conversations is limited because of API restrictions.
	// Since graph doesn't provide us a way to programmatically delete conversations,
	// or create new conversations without a delegated token, we can't do incremental
	// testing with newly added items.
	incMB := evmock.NewBus()
	incBO := NewTestBackupOp(
		t,
		ctx,
		bod,
		incMB,
		opts,
		count.New())

	RunAndCheckBackup(t, ctx, &incBO, incMB, true)

	for _, reason := range reasons {
		CheckBackupIsInManifests(
			t,
			ctx,
			bod.KW,
			bod.SW,
			&incBO,
			bod.Sel,
			bod.Sel.ID(),
			reason.Category())
	}

	_, expectDeets = deeTD.GetDeetsInBackup(
		t,
		ctx,
		incBO.Results.BackupID,
		bod.Acct.ID(),
		bod.Sel.ID(),
		bod.Sel.PathService(),
		whatSet,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		incBO.Results.BackupID,
		whatSet,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)

	assert.NotZero(
		t,
		incBO.Results.Counts[string(count.PersistedCachedFiles)],
		"cached items")
	assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
	assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
	assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
}

type GroupsBackupNightlyIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestGroupsBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupNightlyIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9MergeBase() {
	sel := selectors.NewGroupsBackup([]string{suite.m365.Group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel))

	// TODO(pandeyabs): CorsoCITeam group mailbox backup is currently broken because of invalid
	// odata.NextLink which causes an infinite loop during paging. Disabling conv backups while
	// we go fix the group mailbox.
	// selTD.GroupsBackupConversationScope(sel))

	RunMergeBaseGroupsUpdate(suite, sel.Selector, false)
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9AssistBases() {
	sel := selectors.NewGroupsBackup([]string{suite.m365.Group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel))

	// TODO(pandeyabs): CorsoCITeam group mailbox backup is currently broken because of invalid
	// odata.NextLink which causes an infinite loop during paging. Disabling conv backups while
	// we go fix the group mailbox.
	// selTD.GroupsBackupConversationScope(sel))

	RunDriveAssistBaseGroupsUpdate(suite, sel.Selector, false)
}

type GroupsRestoreNightlyIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestGroupsRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsRestoreNightlyIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *GroupsRestoreNightlyIntgSuite) TestRestore_Run_groupsWithAdvancedOptions() {
	sel := selectors.NewGroupsBackup([]string{suite.m365.Group.ID})
	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))
	sel.Filter(sel.Library("documents"))
	sel.DiscreteOwner = suite.m365.Group.ID

	RunDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.m365.AC,
		sel.Selector,
		suite.m365.Group.RootSite.DriveID,
		suite.m365.Group.RootSite.DriveRootFolderID)
}
