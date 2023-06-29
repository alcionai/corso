package restore

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
)

// called by restore.go to map subcommands to provider-specific handling.
func addOneDriveCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, oneDriveRestoreCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddOneDriveDetailsAndRestoreFlags(c)
		flags.AddRestorePermissionsFlag(c)
		flags.AddFailFastFlag(c)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
	}

	return c
}

const (
	oneDriveServiceCommand          = "onedrive"
	oneDriveServiceCommandUseSuffix = "--backup <backupId>"

	oneDriveServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef in Bob's last backup (1234abcd...)
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore the file with ID 98765abcdef along with its associated permissions
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef --restore-permissions

# Restore files named "FY2021 Planning.xlsx" in "Documents/Finance Reports"
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Restore all files and folders in folder "Documents/Finance Reports" that were created before 2020
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd 
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso restore onedrive [<flag>...]`
func oneDriveRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Restore M365 OneDrive service data",
		RunE:    restoreOneDriveCmd,
		Args:    cobra.NoArgs,
		Example: oneDriveServiceCommandRestoreExamples,
	}
}

// processes an onedrive service restore.
func restoreOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeOneDriveOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateOneDriveRestoreFlags(flags.BackupIDFV, opts); err != nil {
		return err
	}

	r, _, _, err := utils.GetAccountAndConnect(ctx, repo.S3Overrides())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	restoreCfg := control.DefaultRestoreConfig(dttm.HumanReadableDriveItem)
	Infof(ctx, "Restoring to folder %s", restoreCfg.Location)

	sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, flags.BackupIDFV, sel.Selector, restoreCfg)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize OneDrive restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+flags.BackupIDFV))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run OneDrive restore"))
	}

	ds.Items().MaybePrintEntries(ctx)

	return nil
}
