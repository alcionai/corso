package backup_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/print"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlRepo "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/storage/testdata"
)

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type ids struct {
	ID                string
	DriveID           string
	DriveRootFolderID string
}

type intgTesterSetup struct {
	acct   account.Account
	ac     api.Client
	gockAC api.Client
	user   ids
	site   ids
	group  ids
	team   ids
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	its.acct = tconfig.NewM365Account(t)
	creds, err := its.acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	uids := ids{}

	uids.ID = tconfig.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, uids.ID)
	require.NoError(t, err, clues.ToCore(err))

	uids.DriveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, uids.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	uids.DriveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	its.user = uids

	// site

	sids := ids{}

	sids.ID = tconfig.M365SiteID(t)

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, sids.ID)
	require.NoError(t, err, clues.ToCore(err))

	sids.DriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, sids.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	sids.DriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	its.site = sids

	// group

	gids := ids{}

	// use of the TeamID is intentional here, so that we are assured
	// the group has full usage of the teams api.
	gids.ID = tconfig.M365TeamID(t)

	its.group = gids

	// team

	tids := ids{}
	tids.ID = tconfig.M365TeamID(t)
	its.team = tids

	return its
}

type dependencies struct {
	st             storage.Storage
	repo           repository.Repository
	vpr            *viper.Viper
	recorder       strings.Builder
	configFilePath string
}

func prepM365Test(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) dependencies {
	var (
		acct     = tconfig.NewM365Account(t)
		st       = testdata.NewPrefixedS3Storage(t)
		recorder = strings.Builder{}
	)

	cfg, err := st.S3Config()
	require.NoError(t, err, clues.ToCore(err))

	force := map[string]string{
		tconfig.TestCfgAccountProvider: "M365",
		tconfig.TestCfgStorageProvider: "S3",
		tconfig.TestCfgPrefix:          cfg.Prefix,
	}

	vpr, cfgFP := tconfig.MakeTempTestConfigClone(t, force)
	ctx = config.SetViper(ctx, vpr)

	repo, err := repository.Initialize(
		ctx,
		acct,
		st,
		control.DefaultOptions(),
		ctrlRepo.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	return dependencies{
		st:             st,
		repo:           repo,
		vpr:            vpr,
		recorder:       recorder,
		configFilePath: cfgFP,
	}
}

// ---------------------------------------------------------------------------
// funcs
// ---------------------------------------------------------------------------

func buildExchangeBackupCmd(
	ctx context.Context,
	configFile, user, category string,
	recorder *strings.Builder,
) (*cobra.Command, context.Context) {
	cmd := cliTD.StubRootCmd(
		"backup", "create", "exchange",
		"--config-file", configFile,
		"--"+flags.UserFN, user,
		"--"+flags.CategoryDataFN, category)
	cli.BuildCommandTree(cmd)
	cmd.SetOut(recorder)

	return cmd, print.SetRootCmd(ctx, cmd)
}
