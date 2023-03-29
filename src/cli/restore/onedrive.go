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

		utils.AddBackupIDFlag(c, true)
		utils.AddOneDriveDetailsAndRestoreFlags(c)

		// restore permissions
		options.AddRestorePermissionsFlag(c)

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

	opts := utils.OneDriveOpts{
		Users:              utils.User,
		FileNames:          utils.FileName,
		FolderPaths:        utils.FolderPath,
		FileCreatedAfter:   utils.FileCreatedAfter,
		FileCreatedBefore:  utils.FileCreatedBefore,
		FileModifiedAfter:  utils.FileModifiedAfter,
		FileModifiedBefore: utils.FileModifiedBefore,

		Populated: utils.GetPopulatedFlags(cmd),
	}

	if err := utils.ValidateOneDriveRestoreFlags(utils.BackupID, opts); err != nil {
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

	dest := control.DefaultRestoreDestination(common.SimpleDateTimeOneDrive)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeOneDriveRestoreDataSelectors(opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, utils.BackupID, sel.Selector, dest)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize OneDrive restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+utils.BackupID))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run OneDrive restore"))
	}

	ds.PrintEntries(ctx)

	return nil
}
