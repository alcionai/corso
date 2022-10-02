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
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	folderPaths []string
	fileNames   []string

	fileCreatedAfter   string
	fileCreatedBefore  string
	fileModifiedAfter  string
	fileModifiedBefore string
)

// called by restore.go to map parent subcommands to provider-specific handling.
func addOneDriveCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(parent, oneDriveRestoreCmd())

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --all) and more frequently used flags take precedence.
		fs.SortFlags = false

		fs.StringVar(&backupID, "backup", "", "ID of the backup to restore")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

		fs.StringSliceVar(&user,
			"user", nil,
			"Restore all data by user ID; accepts "+utils.Wildcard+" to select all users")

		// onedrive hierarchy (path/name) flags

		fs.StringSliceVar(
			&folderPaths,
			"folder", nil,
			"Restore items by OneDrive folder; defaults to root")

		fs.StringSliceVar(
			&fileNames,
			"file-name", nil,
			"Restore items by OneDrive file name")

		// onedrive info flags

		fs.StringVar(
			&fileCreatedAfter,
			"file-created-after", "",
			"Restore files created after this datetime")
		fs.StringVar(
			&fileCreatedBefore,
			"file-created-before", "",
			"Restore files created before this datetime")

		fs.StringVar(
			&fileModifiedAfter,
			"file-modified-after", "",
			"Restore files modified after this datetime")
		fs.StringVar(
			&fileModifiedBefore,
			"file-modified-before", "",
			"Restore files modified before this datetime")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const oneDriveServiceCommand = "onedrive"

// `corso restore onedrive [<flag>...]`
func oneDriveRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:   oneDriveServiceCommand,
		Short: "Restore M365 OneDrive service data",
		RunE:  restoreOneDriveCmd,
		Args:  cobra.NoArgs,
	}
}

// processes an onedrive service restore.
func restoreOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := utils.ValidateOneDriveRestoreFlags(backupID); err != nil {
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

	opts := utils.OneDriveOpts{
		Users:          user,
		Paths:          folderPaths,
		Names:          fileNames,
		CreatedAfter:   fileCreatedAfter,
		CreatedBefore:  fileCreatedBefore,
		ModifiedAfter:  fileModifiedAfter,
		ModifiedBefore: fileModifiedBefore,
	}

	sel := selectors.NewOneDriveRestore()
	utils.IncludeOneDriveRestoreDataSelectors(sel, opts)
	utils.FilterOneDriveRestoreInfoSelectors(sel, opts)

	// if no selector flags were specified, get all data in the service.
	if len(sel.Scopes()) == 0 {
		sel.Include(sel.Users(selectors.Any()))
	}

	restoreDest := control.DefaultRestoreDestination(common.SimpleDateTimeFormatOneDrive)

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, restoreDest)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize OneDrive restore"))
	}

	if err := ro.Run(ctx); err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to run OneDrive restore"))
	}

	Infof(ctx, "Restored OneDrive in %s for user %s.\n", s.Provider, sel.ToPrintable().Resources())

	return nil
}
