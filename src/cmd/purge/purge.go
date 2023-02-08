package main

import (
	"context"
	"os"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/logger"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge all types of m365 folders",
	RunE:  handleAllFolderPurge,
}

var oneDriveCmd = &cobra.Command{
	Use:   "onedrive",
	Short: "Purges OneDrive folders",
	RunE:  handleOneDriveFolderPurge,
}

var (
	before string
	user   string
	tenant string
	prefix string
)

var ErrPurging = errors.New("not all items were successfully purged")

// ------------------------------------------------------------------------------------------
// CLI command handlers
// ------------------------------------------------------------------------------------------

func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)
	ctx = SetRootCmd(ctx, purgeCmd)

	defer logger.Flush(ctx)

	fs := purgeCmd.PersistentFlags()
	fs.StringVar(&before, "before", "", "folders older than this date are deleted.  (default: now in UTC)")
	fs.StringVar(&user, "user", "", "m365 user id whose folders will be deleted")
	cobra.CheckErr(purgeCmd.MarkPersistentFlagRequired("user"))
	fs.StringVar(&tenant, "tenant", "", "m365 tenant containing the user")
	fs.StringVar(&prefix, "prefix", "", "filters mail folders by displayName prefix")
	cobra.CheckErr(purgeCmd.MarkPersistentFlagRequired("prefix"))

	purgeCmd.AddCommand(oneDriveCmd)

	if err := purgeCmd.ExecuteContext(ctx); err != nil {
		logger.Flush(ctx)
		os.Exit(1)
	}
}

func handleAllFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, t, err := getGCAndBoundaryTime(ctx)
	if err != nil {
		return err
	}

	err = runPurgeForEachUser(
		ctx, gc, t,
		purgeOneDriveFolders,
	)
	if err != nil {
		return Only(ctx, ErrPurging)
	}

	return nil
}

func handleOneDriveFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, t, err := getGCAndBoundaryTime(ctx)
	if err != nil {
		return err
	}

	if err := runPurgeForEachUser(ctx, gc, t, purgeOneDriveFolders); err != nil {
		logger.Ctx(ctx).Error(err)
		return Only(ctx, errors.Wrap(ErrPurging, "OneDrive folders"))
	}

	return nil
}

// ------------------------------------------------------------------------------------------
// Purge Controllers
// ------------------------------------------------------------------------------------------

type purgable interface {
	GetDisplayName() *string
	GetId() *string
}

type purger func(context.Context, *connector.GraphConnector, time.Time, string) error

func runPurgeForEachUser(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	ps ...purger,
) error {
	var errs error

	for pn, uid := range userOrUsers(user, gc.Users) {
		Infof(ctx, "\nUser: %s - %s", pn, uid)

		for _, p := range ps {
			if err := p(ctx, gc, boundary, pn); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	return errs
}

// ----- OneDrive

func purgeOneDriveFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	uid string,
) error {
	getter := func(gs graph.Servicer, uid, prefix string) ([]purgable, error) {
		pager, err := onedrive.PagerForSource(onedrive.OneDriveSource, gs, uid, nil)
		if err != nil {
			return nil, err
		}

		cfs, err := onedrive.GetAllFolders(ctx, gs, pager, prefix)
		if err != nil {
			return nil, err
		}

		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Servicer, uid string, f purgable) error {
		driveFolder, ok := f.(*onedrive.Displayable)
		if !ok {
			return errors.New("non-OneDrive item")
		}

		return onedrive.DeleteItem(
			ctx,
			gs,
			*driveFolder.GetParentReference().GetDriveId(),
			*f.GetId(),
		)
	}

	return purgeFolders(ctx, gc, boundary, "OneDrive Folders", uid, getter, deleter)
}

// ----- controller

func purgeFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	data, uid string,
	getter func(graph.Servicer, string, string) ([]purgable, error),
	deleter func(graph.Servicer, string, purgable) error,
) error {
	Infof(ctx, "Container: %s", data)

	// get them folders
	fs, err := getter(gc.Service, uid, prefix)
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "retrieving %s folders", data))
	}

	if len(fs) == 0 {
		Info(ctx, "None Matched")
		return nil
	}

	var errs error

	// delete any containers that don't pass the boundary
	for _, fld := range fs {
		// compare the folder time to the deletion boundary time first
		displayName := *fld.GetDisplayName()

		dnTime, err := common.ExtractTime(displayName)
		if err != nil && !errors.Is(err, common.ErrNoTimeString) {
			err = errors.Wrapf(err, "!! Error: parsing container named [%s]", displayName)
			errs = multierror.Append(errs, err)
			Info(ctx, err)

			continue
		}

		if !dnTime.Before(boundary) || dnTime == (time.Time{}) {
			continue
		}

		Infof(ctx, "∙ Deleting [%s]", displayName)

		err = deleter(gc.Service, uid, fld)
		if err != nil {
			err = errors.Wrapf(err, "!! Error")
			errs = multierror.Append(errs, err)
			Info(ctx, err)
		}
	}

	return errs
}

// ------------------------------------------------------------------------------------------
// Helpers
// ------------------------------------------------------------------------------------------

func getGC(ctx context.Context) (*connector.GraphConnector, error) {
	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: common.First(tenant, os.Getenv(account.AzureTenantID)),
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, Only(ctx, errors.Wrap(err, "finding m365 account details"))
	}

	// build a graph connector
	gc, err := connector.NewGraphConnector(ctx,
		graph.HTTPClient(
			graph.NoTimeout(),
			graph.RetryOptions(
				graph.RetryHandlerOptions{
					MaxRetries: 3,
				})),
		acct,
		connector.Users)
	if err != nil {
		return nil, Only(ctx, errors.Wrap(err, "connecting to graph api"))
	}

	return gc, nil
}

func getBoundaryTime(ctx context.Context) (time.Time, error) {
	// format the time input
	var (
		err          error
		boundaryTime = time.Now().UTC()
	)

	if len(before) > 0 {
		boundaryTime, err = common.ParseTime(before)
		if err != nil {
			return time.Time{}, Only(ctx, errors.Wrap(err, "parsing before flag to time"))
		}
	}

	return boundaryTime, nil
}

func getGCAndBoundaryTime(ctx context.Context) (*connector.GraphConnector, time.Time, error) {
	gc, err := getGC(ctx)
	if err != nil {
		return nil, time.Time{}, err
	}

	t, err := getBoundaryTime(ctx)
	if err != nil {
		return nil, time.Time{}, err
	}

	return gc, t, nil
}

func userOrUsers(u string, us map[string]string) map[string]string {
	if len(u) == 0 {
		return nil
	}

	if u == "*" {
		return us
	}

	return map[string]string{u: u}
}
