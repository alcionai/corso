package backup

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
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
	// awaiting release
	// addGroupsCommands,
	// addTeamsCommands,
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

		// delete after release
		if len(os.Getenv("CORSO_ENABLE_GROUPS")) > 0 {
			addGroupsCommands(subCommand)
			addTeamsCommands(subCommand)
		}
	}
}

// ---------------------------------------------------------------------------
// common flags and flag attachers for commands
// ---------------------------------------------------------------------------

func addFailedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&flags.ListFailedItemsFV, flags.FailedItemsFN, "show",
		"Toggles showing or hiding the list of items that failed.")
}

func addSkippedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&flags.ListSkippedItemsFV, flags.SkippedItemsFN, "show",
		"Toggles showing or hiding the list of items that were skipped.")
}

func addRecoveredErrorsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&flags.ListRecoveredErrorsFV, flags.RecoveredErrorsFN, "show",
		"Toggles showing or hiding the list of errors which corso recovered from.")
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

// standard set of selector behavior that we want used in the cli
var defaultSelectorConfig = selectors.Config{OnlyMatchItemNames: true}

func runBackups(
	ctx context.Context,
	r repository.Repository,
	serviceName, resourceOwnerType string,
	selectorSet []selectors.Selector,
	ins idname.Cacher,
) error {
	var (
		bIDs []string
		errs = []error{}
	)

	for _, discSel := range selectorSet {
		discSel.Configure(defaultSelectorConfig)

		var (
			owner = discSel.DiscreteOwner
			ictx  = clues.Add(ctx, "resource_owner_selected", owner)
		)

		bo, err := r.NewBackupWithLookup(ictx, discSel, ins)
		if err != nil {
			errs = append(errs, clues.Wrap(err, owner).WithClues(ictx))
			Errf(ictx, "%v\n", err)

			continue
		}

		ictx = clues.Add(
			ctx,
			"resource_owner_id", bo.ResourceOwner.ID(),
			"resource_owner_name", bo.ResourceOwner.Name())

		err = bo.Run(ictx)
		if err != nil {
			if errors.Is(err, graph.ErrServiceNotEnabled) {
				logger.Ctx(ctx).Infow("service not enabled", "resource_owner_name", bo.ResourceOwner.Name())

				continue
			}

			errs = append(errs, clues.Wrap(err, owner).WithClues(ictx))
			Errf(ictx, "%v\n", err)

			continue
		}

		bIDs = append(bIDs, string(bo.Results.BackupID))

		if !DisplayJSONFormat() {
			Infof(ctx, "Done\n")
			printBackupStats(ctx, r, string(bo.Results.BackupID))
		} else {
			Infof(ctx, "Done - ID: %v\n", bo.Results.BackupID)
		}
	}

	bups, berrs := r.Backups(ctx, bIDs)
	if berrs.Failure() != nil {
		return Only(ctx, clues.Wrap(berrs.Failure(), "Unable to retrieve backup results from storage"))
	}

	Info(ctx, "Completed Backups:")
	backup.PrintAll(ctx, bups)

	if len(errs) > 0 {
		sb := fmt.Sprintf("%d of %d backups failed:\n", len(errs), len(selectorSet))

		for i, e := range errs {
			logger.CtxErr(ctx, e).Errorf("Backup %d of %d failed", i+1, len(selectorSet))
			sb += "∙ " + e.Error() + "\n"
		}

		return Only(ctx, clues.New(sb))
	}

	return nil
}

// genericDeleteCommand is a helper function that all services can use
// for the removal of an entry from the repository
func genericDeleteCommand(
	cmd *cobra.Command,
	pst path.ServiceType,
	bID, designation string,
	args []string,
) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctx := clues.Add(cmd.Context(), "delete_backup_id", bID)

	r, _, _, _, err := utils.GetAccountAndConnect(ctx, pst, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackups(ctx, true, bID); err != nil {
		return Only(ctx, clues.Wrap(err, "Deleting backup "+bID))
	}

	Infof(ctx, "Deleted %s backup %s", designation, bID)

	return nil
}

// genericListCommand is a helper function that all services can use
// to display the backup IDs saved within the repository
func genericListCommand(
	cmd *cobra.Command,
	bID string,
	service path.ServiceType,
	args []string,
) error {
	ctx := cmd.Context()

	r, _, _, _, err := utils.GetAccountAndConnect(ctx, service, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	if len(bID) > 0 {
		fe, b, errs := r.GetBackupErrors(ctx, bID)
		if errs.Failure() != nil {
			if errors.Is(errs.Failure(), data.ErrNotFound) {
				return Only(ctx, clues.New("No backup exists with the id "+bID))
			}

			return Only(ctx, clues.Wrap(errs.Failure(), "Failed to list backup id "+bID))
		}

		b.Print(ctx)
		fe.PrintItems(
			ctx,
			!ifShow(flags.ListFailedItemsFV),
			!ifShow(flags.ListSkippedItemsFV),
			!ifShow(flags.ListRecoveredErrorsFV))

		return nil
	}

	bs, err := r.BackupsByTag(ctx, store.Service(service))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to list backups in the repository"))
	}

	backup.PrintAll(ctx, bs)

	return nil
}

func ifShow(flag string) bool {
	return strings.ToLower(strings.TrimSpace(flag)) == "show"
}

func printBackupStats(ctx context.Context, r repository.Repository, bid string) {
	b, err := r.Backup(ctx, bid)
	if err != nil {
		logger.CtxErr(ctx, err).Error("finding backup immediately after backup operation completion")
	}

	b.ToPrintable().Stats.Print(ctx)
	Info(ctx, " ")
}
