package groups_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	. "github.com/alcionai/corso/src/internal/operations/test/m365"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type GroupsBackupIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestGroupsBackupIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_groups() {
	var (
		resourceID = suite.its.Group.ID
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
	runGroupsIncrementalBackupTests(suite, suite.its, control.DefaultOptions())
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_extensionsGroups() {
	var (
		resourceID = suite.its.Group.ID
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
	its IntgTesterSetup
}

func TestGroupsBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupTreeIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeGroups() {
	var (
		resourceID = suite.its.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	RunBasicDriveishBackupTests(
		suite,
		path.GroupsService,
		opts,
		sel.Selector)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeIncrementalGroups() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	runGroupsIncrementalBackupTests(suite, suite.its, opts)
}

func (suite *GroupsBackupTreeIntgSuite) TestBackup_Run_treeExtensionsGroups() {
	var (
		resourceID = suite.its.Group.ID
		sel        = selectors.NewGroupsBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

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
	its IntgTesterSetup,
	opts control.Options,
) {
	sel := selectors.NewGroupsRestore([]string{its.Group.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return its.Group.RootSite.DriveID
	}

	gtsi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		return its.Group.RootSite.ID
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewSiteRestoreHandler(ac, path.GroupsService)
	}

	RunIncrementalDriveishBackupTest(
		suite,
		opts,
		its.Group.ID,
		its.User.ID,
		path.GroupsService,
		path.LibrariesCategory,
		ic,
		gtdi,
		gtsi,
		grh,
		true)
}

func (suite *GroupsBackupIntgSuite) TestBackup_Run_basicBackup() {
	sel := selectors.NewGroupsBackup([]string{suite.its.Group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel),
		selTD.GroupsBackupConversationScope(sel))

	RunBasicBackupTest(suite, sel.Selector)
}

type GroupsBackupNightlyIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestGroupsBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsBackupNightlyIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9MergeBase() {
	sel := selectors.NewGroupsBackup([]string{suite.its.Group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel),
		selTD.GroupsBackupConversationScope(sel))

	RunMergeBaseGroupsUpdate(suite, sel.Selector, false)
}

func (suite *GroupsBackupNightlyIntgSuite) TestBackup_Run_groupsVersion9AssistBases() {
	sel := selectors.NewGroupsBackup([]string{suite.its.Group.ID})
	sel.Include(
		selTD.GroupsBackupLibraryFolderScope(sel),
		selTD.GroupsBackupChannelScope(sel),
		selTD.GroupsBackupConversationScope(sel))

	RunDriveAssistBaseGroupsUpdate(suite, sel.Selector, false)
}

type GroupsRestoreNightlyIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestGroupsRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *GroupsRestoreNightlyIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *GroupsRestoreNightlyIntgSuite) TestRestore_Run_groupsWithAdvancedOptions() {
	sel := selectors.NewGroupsBackup([]string{suite.its.Group.ID})
	sel.Include(selTD.GroupsBackupLibraryFolderScope(sel))
	sel.Filter(sel.Library("documents"))
	sel.DiscreteOwner = suite.its.Group.ID

	RunDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.its.AC,
		sel.Selector,
		suite.its.Group.RootSite.DriveID,
		suite.its.Group.RootSite.DriveRootFolderID)
}
