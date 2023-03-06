package restore

import (
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

var (
	folderPaths []string
	fileNames   []string

	fileCreatedAfter   string
	fileCreatedBefore  string
	fileModifiedAfter  string
	fileModifiedBefore string
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

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to restore. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))

		fs.StringSliceVar(&user,
			utils.UserFN, nil,
			"Restore data by user's email address; accepts '"+utils.Wildcard+"' to select all users.")

		// onedrive hierarchy (path/name) flags

		fs.StringSliceVar(
			&folderPaths,
			utils.FolderFN, nil,
			"Restore items by OneDrive folder; defaults to root")

		fs.StringSliceVar(
			&fileNames,
			utils.FileFN, nil,
			"Restore items by file name or ID")

		// permissions restore flag
		options.AddRestorePermissionsFlag(c)

		// onedrive info flags

		fs.StringVar(
			&fileCreatedAfter,
			utils.FileCreatedAfterFN, "",
			"Restore files created after this datetime")
		fs.StringVar(
			&fileCreatedBefore,
			utils.FileCreatedBeforeFN, "",
			"Restore files created before this datetime")

		fs.StringVar(
			&fileModifiedAfter,
			utils.FileModifiedAfterFN, "",
			"Restore files modified after this datetime")
		fs.StringVar(
			&fileModifiedBefore,
			utils.FileModifiedBeforeFN, "",
			"Restore files modified before this datetime")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const (
	oneDriveServiceCommand          = "onedrive"
	oneDriveServiceCommandUseSuffix = "--backup <backupId>"

	oneDriveServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore file with ID 98765abcdef along with its associated permissions
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef --restore-permissions

# Restore Alice's file named "FY2021 Planning.xlsx in "Documents/Finance Reports" from a specific backup
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Restore all files from Bob's folder that were created before 2020 when captured in a specific backup
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd 
      --user bob@example.com --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
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

	opts := utils.OneDriveOptions(user, folderPaths, fileNames, cmd)

	if err := utils.ValidateOneDriveRestoreFlags(backupID, opts); err != nil {
		return err
	}

	s, a, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, a, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(common.SimpleDateTimeOneDrive)

	sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, dest)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize OneDrive restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, errors.Errorf("Backup or backup details missing for id %s", backupID))
		}

		return Only(ctx, errors.Wrap(err, "Failed to run OneDrive restore"))
	}

	ds.PrintEntries(ctx)

	return nil
}
