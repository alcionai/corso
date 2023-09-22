package backup

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
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
	addGroupsCommands,
}

// AddCommands attaches all `corso backup * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	backupC := backupCmd()
	cmd.AddCommand(backupC)

	for _, sc := range subCommandFuncs {
		subCommand := sc()
		flags.AddAllProviderFlags(subCommand)
		flags.AddAllStorageFlags(subCommand)
		backupC.AddCommand(subCommand)

		for _, addBackupTo := range serviceCommands {
			addBackupTo(subCommand)
		}
	}
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

func genericCreateCommand(
	ctx context.Context,
	r repository.Repositoryer,
	serviceName string,
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
				logger.Ctx(ctx).Infow("service not enabled",
					"resource_owner_id", bo.ResourceOwner.ID(),
					"service", serviceName)

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

	if len(bups) > 0 {
		Info(ctx, "Completed Backups:")
	}

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
	designation string,
	bID, args []string,
) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	ctx := clues.Add(cmd.Context(), "delete_backup_id", bID)

	r, _, err := utils.GetAccountAndConnect(ctx, cmd, pst)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackups(ctx, true, bID...); err != nil {
		return Only(ctx, clues.Wrap(err, fmt.Sprintf("Deleting backup %v", bID)))
	}

	Infof(ctx, "Deleted %s backup %v", designation, bID)

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

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	r, _, err := utils.GetAccountAndConnect(ctx, cmd, service)
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

func genericDetailsCommand(
	cmd *cobra.Command,
	backupID string,
	sel selectors.Selector,
) (*details.Details, error) {
	ctx := cmd.Context()

	r, rdao, err := utils.GetAccountAndConnect(ctx, cmd, path.OneDriveService)
	if err != nil {
		return nil, clues.Stack(err)
	}

	defer utils.CloseRepo(ctx, r)

	return genericDetailsCore(
		ctx,
		r,
		backupID,
		sel,
		rdao.Opts)
}

func genericDetailsCore(
	ctx context.Context,
	bg repository.BackupGetter,
	backupID string,
	sel selectors.Selector,
	opts control.Options,
) (*details.Details, error) {
	ctx = clues.Add(ctx, "backup_id", backupID)

	sel.Configure(selectors.Config{OnlyMatchItemNames: true})

	d, _, errs := bg.GetBackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Failure() != nil {
		if errors.Is(errs.Failure(), data.ErrNotFound) {
			return nil, clues.New("no backup exists with the id " + backupID)
		}

		return nil, clues.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	if opts.SkipReduce {
		return d, nil
	}

	d, err := sel.Reduce(ctx, d, errs)
	if err != nil {
		return nil, clues.Wrap(err, "filtering backup details to selection")
	}

	return d, nil
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

func ifShow(flag string) bool {
	return strings.ToLower(strings.TrimSpace(flag)) == "show"
}

func printBackupStats(ctx context.Context, r repository.Repositoryer, bid string) {
	b, err := r.Backup(ctx, bid)
	if err != nil {
		logger.CtxErr(ctx, err).Error("finding backup immediately after backup operation completion")
	}

	b.ToPrintable().Stats.Print(ctx)
	Info(ctx, " ")
}
