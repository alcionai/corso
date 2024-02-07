package restore

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	. "github.com/alcionai/canario/src/cli/print"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/selectors"
)

var restoreCommands = []func(cmd *cobra.Command) *cobra.Command{
	addExchangeCommands,
	addOneDriveCommands,
	addSharePointCommands,
	addGroupsCommands,
}

// AddCommands attaches all `corso restore * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	subCommand := restoreCmd()
	cmd.AddCommand(subCommand)

	for _, addRestoreTo := range restoreCommands {
		sc := addRestoreTo(subCommand)
		flags.AddAllProviderFlags(sc)
		flags.AddAllStorageFlags(sc)
	}
}

const restoreCommand = "restore"

// FIXME: this command doesn't build docs, and so these examples
// are not visible within the website.
//
//nolint:lll
const restoreCommandExamples = `# Restore email inbox messages to their original location
corso restore exchange \
	--backup 1234abcd-12ab-cd34-56de-1234abcd \
	--email-folder '/inbox' \
	--destination '/'

# Restore a specific OneDrive folder to the top-level destination named "recovered_june_releases"
corso restore onedrive \
	--backup 1234abcd-12ab-cd34-56de-1234abcd \
	--folder '/work/corso_june_releases' \
	--destination /recovered_june_releases

# Restore a calendar event, making a copy if the event already exists.
corso restore exchange \
	--backup 1234abcd-12ab-cd34-56de-1234abcd \
	--event-calendar 'Company Events' \
	--event abdef0101 \
	--collisions copy

# Restore a SharePoint library in-place, replacing any conflicting files.
corso restore sharepoint \
	--backup 1234abcd-12ab-cd34-56de-1234abcd \
	--library documents \
	--destination '/' \
	--collisions replace`

// The restore category of commands.
// `corso restore [<subcommand>] [<flag>...]`
func restoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     restoreCommand,
		Short:   "Restore your service data",
		Long:    `Restore the data stored in one of your M365 services.`,
		RunE:    handleRestoreCmd,
		Args:    cobra.NoArgs,
		Example: restoreCommandExamples,
	}
}

// Handler for flat calls to `corso restore`.
// Produces the same output as `corso restore --help`.
func handleRestoreCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ---------------------------------------------------------------------------
// common handlers
// ---------------------------------------------------------------------------

func runRestore(
	ctx context.Context,
	cmd *cobra.Command,
	urco utils.RestoreCfgOpts,
	sel selectors.Selector,
	backupID, serviceName string,
) error {
	if err := utils.ValidateRestoreConfigFlags(urco); err != nil {
		return Only(ctx, err)
	}

	r, _, err := utils.GetAccountAndConnect(ctx, cmd, sel.PathService())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ro, err := r.NewRestore(ctx, backupID, sel, utils.MakeRestoreConfig(ctx, urco))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize "+serviceName+" restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+flags.BackupIDFV))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run "+serviceName+" restore"))
	}

	Info(ctx, "Restore Complete")

	skipped := ro.Counter.Get(count.CollisionSkip)
	if skipped > 0 {
		Infof(ctx, "Skipped %d items due to collision", skipped)
	}

	dis := ds.Items()

	Outf(ctx, "Restored %d items", len(dis))
	dis.MaybePrintEntries(ctx)

	return nil
}
