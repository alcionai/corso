package export

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
)

// called by export.go to map subcommands to provider-specific handling.
func addOneDriveCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case exportCommand:
		c, fs = utils.AddCommand(cmd, oneDriveExportCmd())

		c.Use = c.Use + " " + oneDriveServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddOneDriveDetailsAndRestoreFlags(c)
		flags.AddExportConfigFlags(c)
		flags.AddFailFastFlag(c)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
	}

	return c
}

const (
	oneDriveServiceCommand          = "onedrive"
	oneDriveServiceCommandUseSuffix = "<destination> --backup <backupId>"

	//nolint:lll
	oneDriveServiceCommandExportExamples = `# Export file with ID 98765abcdef in Bob's last backup (1234abcd...) to my-exports directory
corso export onedrive my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Export files named "FY2021 Planning.xlsx" in "Documents/Finance Reports" to current directory
corso export onedrive . --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Export all files and folders in folder "Documents/Finance Reports" that were created before 2020 to my-exports
corso export onedrive my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso export onedrive [<flag>...] <destination>`
func oneDriveExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   oneDriveServiceCommand,
		Short: "Export M365 OneDrive service data",
		RunE:  exportOneDriveCmd,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing export destination")
			}

			return nil
		},
		Example: oneDriveServiceCommandExportExamples,
	}
}

// processes an onedrive service export.
func exportOneDriveCmd(cmd *cobra.Command, args []string) error {
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

	sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	return runExport(ctx, cmd, args, opts.ExportCfg, sel.Selector, flags.BackupIDFV, "OneDrive")
}
