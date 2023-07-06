package export

import (
	"io"
	"os"
	ospath "path"

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
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
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
	oneDriveServiceCommandUseSuffix = "--backup <backupId>"

	oneDriveServiceCommandExportExamples = `# Export file with ID 98765abcdef in Bob's last backup (1234abcd...)
corso export onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Export files named "FY2021 Planning.xlsx" in "Documents/Finance Reports"
corso export onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Export all files and folders in folder "Documents/Finance Reports" that were created before 2020
corso export onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd 
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso export onedrive [<flag>...]`
func oneDriveExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:     oneDriveServiceCommand,
		Short:   "Export M365 OneDrive service data",
		RunE:    exportOneDriveCmd,
		Args:    cobra.NoArgs,
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

	r, _, _, err := utils.GetAccountAndConnect(ctx, path.OneDriveService, repo.S3Overrides(cmd))
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	exportCfg := control.DefaultExportConfig()
	if flags.ArchiveFV {
		exportCfg.Archive = true
	}

	restoreLocation := flags.ExportDestFV
	if restoreLocation == "" {
		restoreLocation = control.DefaultRestoreLocation + dttm.FormatNow(dttm.HumanReadableDriveItem)
	}

	Infof(ctx, "Restoring to folder %s", restoreLocation)

	sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	eo, err := r.NewExport(ctx, flags.BackupIDFV, sel.Selector, exportCfg)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize OneDrive export"))
	}

	expColl, err := eo.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+flags.BackupIDFV))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run OneDrive export"))
	}

	// It would be better to give a progressbar than a spinner, but we
	// have know way of knowing how many files are available as of now.
	diskWriteComplete := observe.MessageWithCompletion(ctx, "Writing data to disk")
	defer close(diskWriteComplete)

	for _, col := range expColl {
		folder := ospath.Join(restoreLocation, col.GetBasePath())

		for item := range col.GetItems(ctx) {
			err := item.Error
			if err != nil {
				return Only(ctx, clues.Wrap(err, "getting item").With("dir_name", folder))
			}

			name := item.Data.Name
			fpath := ospath.Join(folder, name)

			err = os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return Only(ctx, clues.Wrap(err, "creating directory").With("dir_name", folder))
			}

			// In case the user tries to restore to a non-clean
			// directory, we might run into collisions an fail.
			f, err := os.Create(fpath)
			if err != nil {
				return Only(ctx, clues.Wrap(err, "creating file").With("file_name", name, "file_dir", folder))
			}

			_, err = io.Copy(f, item.Data.Body)
			if err != nil {
				return Only(ctx, clues.Wrap(err, "writing file").With("file_name", name, "file_dir", folder))
			}
		}
	}

	diskWriteComplete <- struct{}{}

	return nil
}
