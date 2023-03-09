package backup

import (
	"context"

	"github.com/hashicorp/go-multierror"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ---------------------------------------------------------------------------
// adding commands to cobra
// ---------------------------------------------------------------------------

var subCommandFuncs = []func() *cobra.Command{
	createCmd,
	listCmd,
	detailsCmd,
	deleteCmd,
}

var serviceCommands = []func(cmd *cobra.Command) *cobra.Command{
	addExchangeCommands,
	addOneDriveCommands,
	addSharePointCommands,
}

// AddCommands attaches all `corso backup * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	backupC := backupCmd()
	cmd.AddCommand(backupC)

	for _, sc := range subCommandFuncs {
		subCommand := sc()
		backupC.AddCommand(subCommand)

		for _, addBackupTo := range serviceCommands {
			addBackupTo(subCommand)
		}
	}
}

// ---------------------------------------------------------------------------
// adding flags to cobra commands
// ---------------------------------------------------------------------------

// ==============================================
// Folder Object flags
// These options are flags for indicating
// that a time-based filter should be used for
// within returning objects for details.
// Used by: OneDrive, SharePoint
// ================================================
var (
	fileCreatedAfter   string
	fileCreatedBefore  string
	fileModifiedAfter  string
	fileModifiedBefore string
)

// list output filter flags
var (
	failedItemsFN    = "failed-items"
	listFailedItems  string
	skippedItemsFN   = "skipped-items"
	listSkippedItems string
)

func addFailedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&listFailedItems, failedItemsFN, "show",
		"Toggles showing or hiding the list of items that failed.")
}

func addSkippedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&listSkippedItems, skippedItemsFN, "show",
		"Toggles showing or hiding the list of items that were skipped.")
}

// ---------------------------------------------------------------------------
// commands
// ---------------------------------------------------------------------------

// The backup category of commands.
// `corso backup [<subcommand>] [<flag>...]`
func backupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "backup",
		Short: "Backup your service data",
		Long:  `Backup the data stored in one of your M365 services.`,
		RunE:  handleBackupCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso backup`.
// Produces the same output as `corso backup --help`.
func handleBackupCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup create subcommand.
// `corso backup create <service> [<flag>...]`
var createCommand = "create"

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   createCommand,
		Short: "Backup an M365 Service",
		RunE:  handleCreateCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso backup create`.
// Produces the same output as `corso backup create --help`.
func handleCreateCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup list subcommand.
// `corso backup list <service> [<flag>...]`
var listCommand = "list"

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   listCommand,
		Short: "List the history of backups",
		RunE:  handleListCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso backup list`.
// Produces the same output as `corso backup list --help`.
func handleListCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup details subcommand.
// `corso backup details <service> [<flag>...]`
var detailsCommand = "details"

func detailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   detailsCommand,
		Short: "Shows the details of a backup",
		RunE:  handleDetailsCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso backup details`.
// Produces the same output as `corso backup details --help`.
func handleDetailsCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The backup delete subcommand.
// `corso backup delete <service> [<flag>...]`
var deleteCommand = "delete"

func deleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   deleteCommand,
		Short: "Deletes a backup",
		RunE:  handleDeleteCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso backup delete`.
// Produces the same output as `corso backup delete --help`.
func handleDeleteCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ---------------------------------------------------------------------------
// common handlers
// ---------------------------------------------------------------------------

func runBackups(
	ctx context.Context,
	r repository.Repository,
	serviceName, resourceOwnerType string,
	selectorSet []selectors.Selector,
) error {
	var (
		merrs *multierror.Error
		bIDs  []model.StableID
	)

	for _, discSel := range selectorSet {
		bo, err := r.NewBackup(ctx, discSel)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to initialize %s backup for %s %s",
				serviceName,
				resourceOwnerType,
				discSel.DiscreteOwner))

			continue
		}

		err = bo.Run(ctx)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to run %s backup for %s %s",
				serviceName,
				resourceOwnerType,
				discSel.DiscreteOwner))

			continue
		}

		bIDs = append(bIDs, bo.Results.BackupID)
	}

	bups, ferrs := r.Backups(ctx, bIDs)
	// TODO: print/log recoverable errors
	if ferrs.Failure() != nil {
		return Only(ctx, errors.Wrap(ferrs.Failure(), "Unable to retrieve backup results from storage"))
	}

	backup.PrintAll(ctx, bups)

	if e := merrs.ErrorOrNil(); e != nil {
		return Only(ctx, e)
	}

	return nil
}

// genericDeleteCommand is a helper function that all services can use
// for the removal of an entry from the repository
func genericDeleteCommand(cmd *cobra.Command, bID, designation string, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	r, _, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackup(ctx, model.StableID(bID)); err != nil {
		return Only(ctx, errors.Wrapf(err, "Deleting backup %s", bID))
	}

	Infof(ctx, "Deleted %s backup %s", designation, bID)

	return nil
}

// genericListCommand is a helper function that all services can use
// to display the backup IDs saved within the repository
func genericListCommand(cmd *cobra.Command, bID string, service path.ServiceType, args []string) error {
	ctx := cmd.Context()

	r, _, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	if len(backupID) > 0 {
		b, err := r.Backup(ctx, model.StableID(bID))
		if err != nil {
			if errors.Is(err, data.ErrNotFound) {
				return Only(ctx, errors.Errorf("No backup exists with the id %s", bID))
			}

			return Only(ctx, errors.Wrap(err, "Failed to find backup "+bID))
		}

		b.Print(ctx)

		fe, _, errs := r.GetBackupErrors(ctx, string(b.ID))
		if errs.Failure() != nil {
			return Only(ctx, errors.Wrap(err, "Failed to find errors in backup"))
		}

		fe.PrintItems(ctx, listFailedItems != "show", listSkippedItems != "show")

		return nil
	}

	bs, err := r.BackupsByTag(ctx, store.Service(service))
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to list backups in the repository"))
	}

	backup.PrintAll(ctx, bs)

	return nil
}

func getAccountAndConnect(ctx context.Context) (repository.Repository, *account.Account, error) {
	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return nil, nil, err
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider)
	}

	return r, &cfg.Account, nil
}
