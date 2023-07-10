package test_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
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
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type SharePointIntgSuite struct {
	tester.Suite
	its intgTesterSetup
	// the goal of backupInstances is to run a single backup at the start of
	// the suite, and re-use that backup throughout the rest of the suite.
	bi *backupInstance
}

func TestSharePointIntgSuite(t *testing.T) {
	suite.Run(t, &SharePointIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *SharePointIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.its = newIntegrationTesterSetup(suite.T())

	sel := selectors.NewSharePointBackup([]string{suite.its.siteID})
	sel.Include(selTD.SharePointBackupFolderScope(sel))
	sel.DiscreteOwner = suite.its.siteID

	var (
		mb   = evmock.NewBus()
		opts = control.Defaults()
	)

	suite.bi = prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
	suite.bi.runAndCheckBackup(t, ctx, mb, false)
}

func (suite *SharePointIntgSuite) TeardownSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	if suite.bi != nil {
		suite.bi.close(t, ctx)
	}
}

func (suite *SharePointIntgSuite) TestBackup_Run_sharePoint() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		bod     = suite.bi.bod
		sel     = suite.bi.bod.sel
		obo     = suite.bi.obo
		siteID  = suite.its.siteID
		whatSet = deeTD.DriveIDFromRepoRef
	)

	checkBackupIsInManifests(
		t,
		ctx,
		bod,
		obo,
		sel,
		siteID,
		path.LibrariesCategory)

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		obo.Results.BackupID,
		bod.acct.ID(),
		sel,
		path.SharePointService,
		whatSet,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		obo.Results.BackupID,
		whatSet,
		bod.kms,
		bod.sss,
		expectDeets,
		false)
}

func (suite *SharePointIntgSuite) TestBackup_Run_incrementalSharePoint() {
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
		suite.bi,
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

func (suite *SharePointIntgSuite) TestBackup_Run_sharePointExtensions() {
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

	// TODO: use the existing backupInstance for this test
	bi := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
	defer bi.bod.close(t, ctx)

	bi.runAndCheckBackup(t, ctx, mb, false)

	bod := bi.bod
	obo := bi.obo
	bID := obo.Results.BackupID

	checkBackupIsInManifests(
		t,
		ctx,
		bod,
		obo,
		bod.sel,
		suite.its.siteID,
		path.LibrariesCategory)

	deets, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel,
		svc,
		ws,
		bi.bod.kms,
		bi.bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bi.bod.kms,
		bi.bod.sss,
		expectDeets,
		false)

	// Check that the extensions are in the backup
	for _, ent := range deets.Entries {
		if ent.Folder == nil {
			verifyExtensionData(t, ent.ItemInfo, path.SharePointService)
		}
	}
}
