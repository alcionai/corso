package restore

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
)

// called by restore.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, exchangeRestoreCmd())

		c.Use = c.Use + " " + exchangeServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		// general flags
		fs.SortFlags = false

		utils.AddBackupIDFlag(c, true)
		utils.AddExchangeDetailsAndRestoreFlags(c)

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = "--backup <backupId>"

	exchangeServiceCommandRestoreExamples = `# Restore emails with ID 98765abcdef and 12345abcdef from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --email 98765abcdef,12345abcdef

# Restore Alice's emails with subject containing "Hello world" in "Inbox" from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --user alice@example.com --email-subject "Hello world" --email-folder Inbox

# Restore Bobs's entire calendar from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --user bob@example.com --event-calendar Calendar

# Restore contact with ID abdef0101 from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --contact abdef0101`
)

// `corso restore exchange [<flag>...]`
func exchangeRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     exchangeServiceCommand,
		Short:   "Restore M365 Exchange service data",
		RunE:    restoreExchangeCmd,
		Args:    cobra.NoArgs,
		Example: exchangeServiceCommandRestoreExamples,
	}
}

// processes an exchange service restore.
func restoreExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeExchangeOpts(cmd)

	if utils.RunMode == utils.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateExchangeRestoreFlags(utils.BackupID, opts); err != nil {
		return err
	}

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to connect to the "+cfg.Storage.Provider.String()+" repository"))
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(common.SimpleDateTime)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, utils.BackupID, sel.Selector, dest)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize Exchange restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+utils.BackupID))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run Exchange restore"))
	}

	ds.PrintEntries(ctx)

	return nil
}
