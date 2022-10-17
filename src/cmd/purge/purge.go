package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge all types of m365 folders",
	RunE:  handleAllFolderPurge,
}

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Purges mail folders",
	RunE:  handleMailFolderPurge,
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Purges calendar event folders",
	RunE:  handleCalendarFolderPurge,
}

var contactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Purges contacts folders",
	RunE:  handleContactsFolderPurge,
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

	purgeCmd.AddCommand(mailCmd)
	purgeCmd.AddCommand(eventsCmd)
	purgeCmd.AddCommand(contactsCmd)
	purgeCmd.AddCommand(oneDriveCmd)

	if err := purgeCmd.ExecuteContext(ctx); err != nil {
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
		purgeMailFolders,
		purgeCalendarFolders,
		purgeContactFolders,
		purgeOneDriveFolders,
	)
	if err != nil {
		return Only(ctx, ErrPurging)
	}

	return nil
}

func handleMailFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, t, err := getGCAndBoundaryTime(ctx)
	if err != nil {
		return err
	}

	if err := runPurgeForEachUser(ctx, gc, t, purgeMailFolders); err != nil {
		return Only(ctx, errors.Wrap(ErrPurging, "mail folders"))
	}

	return nil
}

func handleCalendarFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, t, err := getGCAndBoundaryTime(ctx)
	if err != nil {
		return err
	}

	if err := runPurgeForEachUser(ctx, gc, t, purgeCalendarFolders); err != nil {
		return Only(ctx, errors.Wrap(ErrPurging, "event calendars"))
	}

	return nil
}

func handleContactsFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, t, err := getGCAndBoundaryTime(ctx)
	if err != nil {
		return err
	}

	if err := runPurgeForEachUser(ctx, gc, t, purgeContactFolders); err != nil {
		return Only(ctx, errors.Wrap(ErrPurging, "contact folders"))
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

// ----- mail

func purgeMailFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	uid string,
) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		sel := selectors.NewExchangeBackup()
		sel.Include(sel.MailFolders([]string{uid}, selectors.Any()))

		scope := sel.Scopes()[0]
		params := graph.QueryParams{
			User:     uid,
			Scope:    scope,
			FailFast: false,
			Credentials: account.M365Config{
				M365:     credentials.GetM365(),
				TenantID: common.First(tenant, os.Getenv(account.TenantID)),
			},
		}

		allFolders, err := exchange.GetAllMailFolders(ctx, params, gs)
		if err != nil {
			return nil, err
		}

		mfs := containerFilter(prefix, allFolders)
		purgables := make([]purgable, len(mfs))

		for i, v := range mfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid string, f purgable) error {
		return exchange.DeleteMailFolder(ctx, gs, uid, *f.GetId())
	}

	return purgeFolders(ctx, gc, boundary, "Mail Folders", uid, getter, deleter)
}

// ----- calendars

func purgeCalendarFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	uid string,
) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		sel := selectors.NewExchangeBackup()
		sel.Include(sel.EventCalendars([]string{uid}, selectors.Any()))

		scope := sel.Scopes()[0]
		params := graph.QueryParams{
			User:     uid,
			Scope:    scope,
			FailFast: false,
			Credentials: account.M365Config{
				M365:     credentials.GetM365(),
				TenantID: common.First(tenant, os.Getenv(account.TenantID)),
			},
		}

		allCalendars, err := exchange.GetAllCalendars(ctx, params, gs)
		if err != nil {
			return nil, err
		}

		cfs := containerFilter(prefix, allCalendars)
		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid string, f purgable) error {
		return exchange.DeleteCalendar(ctx, gs, uid, *f.GetId())
	}

	return purgeFolders(ctx, gc, boundary, "Event Calendars", uid, getter, deleter)
}

// ----- contacts

func purgeContactFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	uid string,
) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		sel := selectors.NewExchangeBackup()
		sel.Include(sel.ContactFolders([]string{uid}, selectors.Any()))

		scope := sel.Scopes()[0]
		params := graph.QueryParams{
			User:     uid,
			Scope:    scope,
			FailFast: false,
			Credentials: account.M365Config{
				M365:     credentials.GetM365(),
				TenantID: common.First(tenant, os.Getenv(account.TenantID)),
			},
		}
		allContainers, err := exchange.GetAllContactFolders(ctx, params, gs)
		//, uid, prefix)
		if err != nil {
			return nil, err
		}

		cfs := containerFilter(prefix, allContainers)
		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid string, f purgable) error {
		return exchange.DeleteContactFolder(ctx, gs, uid, *f.GetId())
	}

	return purgeFolders(ctx, gc, boundary, "Contact Folders", uid, getter, deleter)
}

// ----- OneDrive

func purgeOneDriveFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	uid string,
) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		cfs, err := onedrive.GetAllFolders(ctx, gs, uid, prefix)
		if err != nil {
			return nil, err
		}

		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid string, f purgable) error {
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
	getter func(graph.Service, string, string) ([]purgable, error),
	deleter func(graph.Service, string, purgable) error,
) error {
	Infof(ctx, "\nContainer: %s", data)

	// get them folders
	fs, err := getter(gc.Service(), uid, prefix)
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

		Infof(ctx, "Deleting [%s]", displayName)

		err = deleter(gc.Service(), uid, fld)
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
		M365:     credentials.GetM365(),
		TenantID: common.First(tenant, os.Getenv(account.TenantID)),
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, Only(ctx, errors.Wrap(err, "finding m365 account details"))
	}

	// build a graph connector
	gc, err := connector.NewGraphConnector(ctx, acct)
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

// containerFilter filters container list based on prefix
// @returns cachedContainers that meet the requirements for purging.
func containerFilter(nameContains string, containers []graph.CachedContainer) []graph.CachedContainer {
	cacheContainers := make([]graph.CachedContainer, 0)

	for _, folder := range containers {
		include := len(nameContains) == 0 ||
			(len(nameContains) > 0 && strings.Contains(*folder.GetDisplayName(), nameContains))
		if include {
			cacheContainers = append(cacheContainers, folder)
		}
	}

	return cacheContainers
}
