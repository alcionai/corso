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
		tenID = tester.M365TenantID(t)
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
