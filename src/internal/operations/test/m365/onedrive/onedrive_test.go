package onedrive_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	m365Ctrl "github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/model"
	. "github.com/alcionai/corso/src/internal/operations/test/m365"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
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
	m365 its.M365IntgTestSetup
}

func TestOneDriveBackupIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDrive() {
	var (
		resourceID = suite.m365.SecondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	RunBasicDriveishBackupTests(
		suite,
		path.OneDriveService,
		control.DefaultOptions(),
		sel.Selector)
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_incrementalOneDrive() {
	runOneDriveIncrementalBackupTests(suite, suite.m365, control.DefaultOptions())
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_extensionsOneDrive() {
	var (
		resourceID = suite.m365.SecondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	RunDriveishBackupWithExtensionsTests(
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
	m365 its.M365IntgTestSetup
}

func TestOneDriveBackupTreeIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupTreeIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupTreeIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeOneDrive() {
	var (
		resourceID = suite.m365.SecondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	RunBasicDriveishBackupTests(
		suite,
		path.OneDriveService,
		opts,
		sel.Selector)
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeIncrementalOneDrive() {
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	runOneDriveIncrementalBackupTests(suite, suite.m365, opts)
}

func (suite *OneDriveBackupTreeIntgSuite) TestBackup_Run_treeExtensionsOneDrive() {
	var (
		resourceID = suite.m365.SecondaryUser.ID
		sel        = selectors.NewOneDriveBackup([]string{resourceID})
		opts       = control.DefaultOptions()
	)

	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	opts.ToggleFeatures.UseDeltaTree = true

	RunDriveishBackupWithExtensionsTests(
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
	m365 its.M365IntgTestSetup,
	opts control.Options,
) {
	sel := selectors.NewOneDriveRestore([]string{m365.User.ID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.Folders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		d, err := m365.AC.Users().GetDefaultDrive(ctx, m365.User.ID)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default user drive").
				With("user", m365.User.ID)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	grh := func(ac api.Client) drive.RestoreHandler {
		return drive.NewUserDriveRestoreHandler(ac)
	}

	RunIncrementalDriveishBackupTest(
		suite,
		opts,
		m365.User.ID,
		m365.User.ID,
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

	ctrl, err := m365Ctrl.NewController(
		ctx,
		acct,
		path.OneDriveService,
		control.DefaultOptions(),
		counter)
	require.NoError(t, err, clues.ToCore(err))

	userable, err := ctrl.AC.Users().GetByID(
		ctx,
		suite.m365.User.ID,
		api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	uid := ptr.Val(userable.GetId())
	uname := ptr.Val(userable.GetUserPrincipalName())

	oldsel := selectors.NewOneDriveBackup([]string{uname})
	oldsel.Include(selTD.OneDriveBackupFolderScope(oldsel))

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, oldsel.Selector, opts, 0, counter)
	defer bod.Close(t, ctx)

	sel := bod.Sel

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
	RunAndCheckBackup(t, ctx, &bo, mb, false)

	newsel := selectors.NewOneDriveBackup([]string{uid})
	newsel.Include(selTD.OneDriveBackupFolderScope(newsel))
	sel = newsel.SetDiscreteOwnerIDName(uid, uname)

	var (
		incMB = evmock.NewBus()
		// the incremental backup op should have a proper user ID for the id.
		incBO = NewTestBackupOp(t, ctx, bod, incMB, opts, counter)
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
	CheckBackupIsInManifests(
		t,
		ctx,
		bod.KW,
		bod.SW,
		&incBO,
		sel,
		uid,
		maps.Keys(categories)...)
	CheckMetadataFilesExist(
		t,
		ctx,
		incBO.Results.BackupID,
		bod.KW,
		bod.KMS,
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

	err = bod.KMS.Get(ctx, model.BackupSchema, bid, bup)
	require.NoError(t, err, clues.ToCore(err))

	var (
		ssid  = bup.StreamStoreID
		deets details.Details
		ss    = streamstore.NewStreamer(bod.KW, creds.AzureTenantID, path.OneDriveService)
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
	m365 its.M365IntgTestSetup
}

func TestOneDriveBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupNightlyIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *OneDriveBackupNightlyIntgSuite) TestBackup_Run_oneDriveVersion9MergeBase() {
	sel := selectors.NewOneDriveBackup([]string{suite.m365.User.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))

	RunMergeBaseGroupsUpdate(suite, sel.Selector, true)
}

//func (suite *OneDriveBackupNightlyIntgSuite) TestBackup_Run_oneDriveVersion9AssistBases() {
//	sel := selectors.NewOneDriveBackup([]string{tconfig.SecondaryM365UserID(suite.T())})
//	sel.Include(selTD.OneDriveBackupFolderScope(sel))
//
//	runDriveAssistBaseGroupsUpdate(suite, sel.Selector, true)
//}

type OneDriveRestoreNightlyIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestOneDriveRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveRestoreNightlyIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *OneDriveRestoreNightlyIntgSuite) TestRestore_Run_onedriveWithAdvancedOptions() {
	sel := selectors.NewOneDriveBackup([]string{suite.m365.User.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))
	sel.DiscreteOwner = suite.m365.User.ID

	RunDriveRestoreWithAdvancedOptions(
		suite.T(),
		suite,
		suite.m365.AC,
		sel.Selector,
		suite.m365.User.DriveID,
		suite.m365.User.DriveRootFolderID)
}

func (suite *OneDriveRestoreNightlyIntgSuite) TestRestore_Run_onedriveAlternateProtectedResource() {
	sel := selectors.NewOneDriveBackup([]string{suite.m365.User.ID})
	sel.Include(selTD.OneDriveBackupFolderScope(sel))
	sel.DiscreteOwner = suite.m365.User.ID

	RunDriveRestoreToAlternateProtectedResource(
		suite.T(),
		suite,
		suite.m365.AC,
		sel.Selector,
		suite.m365.User,
		suite.m365.SecondaryUser,
		suite.m365.SecondaryUser.ID)
}
