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
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, sharePointRestoreCmd())

		c.Use = c.Use + " " + sharePointServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --site) and more frequently used flags take precedence.
		fs.SortFlags = false

		utils.AddBackupIDFlag(c, true)
		utils.AddAWSCredsFlags(c)
		utils.AddAzureCredsFlags(c)
		utils.AddSharePointDetailsAndRestoreFlags(c)

		options.AddRestorePermissionsFlag(c)
		options.AddFailFastFlag(c)
	}

	return c
}

const (
	sharePointServiceCommand          = "sharepoint"
	sharePointServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	sharePointServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef in Bob's latest backup (1234abcd...)
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore the file with ID 98765abcdef along with its associated permissions
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file 98765abcdef --restore-permissions

# Restore files named "ServerRenderTemplate.xsl" in the folder "Display Templates/Style Sheets".
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file "ServerRenderTemplate.xsl" --folder "Display Templates/Style Sheets"

# Restore all files in the folder "Display Templates/Style Sheets" that were created before 2020.
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd 
    --file-created-before 2020-01-01T00:00:00 --folder "Display Templates/Style Sheets"

# Restore all files in the "Documents" library.
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd 
    --library Documents --folder "Display Templates/Style Sheets" `
)

// `corso restore sharepoint [<flag>...]`
func sharePointRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     sharePointServiceCommand,
		Short:   "Restore M365 SharePoint service data",
		RunE:    restoreSharePointCmd,
		Args:    cobra.NoArgs,
		Example: sharePointServiceCommandRestoreExamples,
	}
}

// processes an sharepoint service restore.
func restoreSharePointCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeSharePointOpts(cmd)

	if utils.RunModeFV == utils.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateSharePointRestoreFlags(utils.BackupIDFV, opts); err != nil {
		return err
	}

	r, _, err := utils.GetAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(dttm.HumanReadableDriveItem)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeSharePointRestoreDataSelectors(ctx, opts)
	utils.FilterSharePointRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, utils.BackupIDFV, sel.Selector, dest)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize SharePoint restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+utils.BackupIDFV))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run SharePoint restore"))
	}

	ds.Items().MaybePrintEntries(ctx)

	return nil
}
