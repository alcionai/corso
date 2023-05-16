package restore

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
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
		options.AddFailFastFlag(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	exchangeServiceCommandRestoreExamples = `# Restore emails with ID 98765abcdef and 12345abcdef from Alice's last backup (1234abcd...)
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --email 98765abcdef,12345abcdef

# Restore emails with subject containing "Hello world" in the "Inbox"
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --email-subject "Hello world" --email-folder Inbox

# Restore an entire calendar
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-calendar Calendar

# Restore the contact with ID abdef0101
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

	if utils.RunModeFV == utils.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateExchangeRestoreFlags(utils.BackupIDFV, opts); err != nil {
		return err
	}

	r, _, err := utils.GetAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(dttm.HumanReadable)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, utils.BackupIDFV, sel.Selector, dest)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize Exchange restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+utils.BackupIDFV))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run Exchange restore"))
	}

	ds.Items().PrintEntries(ctx)

	return nil
}
