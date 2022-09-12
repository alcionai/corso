package main

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
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

var (
	before string
	user   string
	tenant string
	prefix string
)

// ------------------------------------------------------------------------------------------
// CLI command handlers
// ------------------------------------------------------------------------------------------

func main() {
	ctx := SetRootCmd(context.Background(), purgeCmd)
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

	if err := purgeCmd.ExecuteContext(ctx); err != nil {
		Info(purgeCmd.Context(), "Error: ", err.Error())
		os.Exit(1)
	}
}

func handleAllFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, err := getGC(ctx)
	if err != nil {
		return err
	}

	t, err := getBoundaryTime(ctx)
	if err != nil {
		return err
	}

	err = purgeMailFolders(ctx, gc, t)
	if err != nil {
		return errors.Wrap(err, "purging mail folders")
	}

	err = purgeCalendarFolders(ctx, gc, t)
	if err != nil {
		return errors.Wrap(err, "purging calendar folders")
	}

	err = purgeContactFolders(ctx, gc, t)
	if err != nil {
		return errors.Wrap(err, "purging contacts folders")
	}

	return nil
}

func handleMailFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, err := getGC(ctx)
	if err != nil {
		return err
	}

	t, err := getBoundaryTime(ctx)
	if err != nil {
		return err
	}

	return purgeMailFolders(ctx, gc, t)
}

func handleCalendarFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	gc, err := getGC(ctx)
	if err != nil {
		return err
	}

	t, err := getBoundaryTime(ctx)
	if err != nil {
		return err
	}

	return purgeCalendarFolders(ctx, gc, t)
}

func handleContactsFolderPurge(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	gc, err := getGC(ctx)
	if err != nil {
		return err
	}

	t, err := getBoundaryTime(ctx)
	if err != nil {
		return err
	}

	return purgeContactFolders(ctx, gc, t)
}

// ------------------------------------------------------------------------------------------
// Purge Controllers
// ------------------------------------------------------------------------------------------

type purgable interface {
	GetDisplayName() *string
	GetId() *string
}

// ----- mail

func purgeMailFolders(ctx context.Context, gc *connector.GraphConnector, boundary time.Time) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		mfs, err := exchange.GetAllMailFolders(gs, uid, prefix)
		if err != nil {
			return nil, err
		}

		purgables := make([]purgable, len(mfs))

		for i, v := range mfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid, fid string) error {
		return exchange.DeleteMailFolder(gs, uid, fid)
	}

	return purgeFolders(ctx, gc, boundary, "mail", getter, deleter)
}

// ----- calendars

func purgeCalendarFolders(ctx context.Context, gc *connector.GraphConnector, boundary time.Time) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		cfs, err := exchange.GetAllCalendars(gs, uid, prefix)
		if err != nil {
			return nil, err
		}

		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid, fid string) error {
		return exchange.DeleteCalendar(gs, uid, fid)
	}

	return purgeFolders(ctx, gc, boundary, "calendar", getter, deleter)
}

// ----- contacts

func purgeContactFolders(ctx context.Context, gc *connector.GraphConnector, boundary time.Time) error {
	getter := func(gs graph.Service, uid, prefix string) ([]purgable, error) {
		cfs, err := exchange.GetAllContactFolders(gs, uid, prefix)
		if err != nil {
			return nil, err
		}

		purgables := make([]purgable, len(cfs))

		for i, v := range cfs {
			purgables[i] = v
		}

		return purgables, nil
	}

	deleter := func(gs graph.Service, uid, fid string) error {
		return exchange.DeleteContactFolder(gs, uid, fid)
	}

	return purgeFolders(ctx, gc, boundary, "contact", getter, deleter)
}

// ----- controller

func purgeFolders(
	ctx context.Context,
	gc *connector.GraphConnector,
	boundary time.Time,
	data string,
	getter func(graph.Service, string, string) ([]purgable, error),
	deleter func(graph.Service, string, string) error,
) error {
	// get them folders
	fs, err := getter(gc.Service(), user, prefix)
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "retrieving %s folders", data))
	}

	stLen := len(common.SimpleDateTimeFormat)

	// delete any that don't meet the boundary
	for _, fld := range fs {
		// compare the folder time to the deletion boundary time first
		var (
			del         bool
			displayName = *fld.GetDisplayName()
			dnLen       = len(displayName)
		)

		if dnLen > stLen {
			dnSuff := displayName[dnLen-stLen:]

			dnTime, err := common.ParseTime(dnSuff)
			if err != nil {
				Info(ctx, errors.Wrapf(err, "Error: deleting %s folder [%s]", data, displayName))
				continue
			}

			del = dnTime.Before(boundary)
		}

		if !del {
			continue
		}

		Infof(ctx, "Deleting %s folder: %s", data, displayName)

		err = deleter(gc.Service(), user, *fld.GetId())
		if err != nil {
			Info(ctx, errors.Wrapf(err, "Error: deleting %s folder [%s]", data, displayName))
		}
	}

	return nil
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
	gc, err := connector.NewGraphConnector(acct)
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
