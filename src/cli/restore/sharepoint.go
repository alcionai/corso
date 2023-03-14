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
	listItems   []string
	listPaths   []string
	pageFolders []string
	pages       []string
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

		fs.StringVar(
			&backupID,
			utils.BackupFN, "",
			"ID of the backup to restore. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))

		fs.StringSliceVar(
			&utils.Site,
			utils.SiteFN, nil,
			"Restore data by site ID; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(
			&utils.WebURL,
			utils.WebURLFN, nil,
			"Restore data by site webURL; accepts '"+utils.Wildcard+"' to select all sites.")

		// sharepoint hierarchy (path/name) flags

		fs.StringVar(
			&utils.Library,
			utils.LibraryFN, "",
			"Restore files within a library. Default includes all libraries.")

		fs.StringSliceVar(
			&utils.FolderPaths,
			utils.FolderFN, nil,
			"Restore files by folder; defaults to root.")

		fs.StringSliceVar(
			&utils.FileNames,
			utils.FileFN, nil,
			"Restore files by name.")

		fs.StringSliceVar(
			&listPaths,
			utils.ListFN, nil,
			"Restore list items by SharePoint list ID")

		fs.StringSliceVar(
			&listItems,
			utils.ListItemFN, nil,
			"Restore list items by ID")

		fs.StringSliceVar(
			&pageFolders,
			utils.PageFolderFN, nil,
			"Restore Site pages by page folder name")

		fs.StringSliceVar(
			&pages,
			utils.PagesFN, nil,
			"Restore site pages by file name(s)")

		// sharepoint info flags

		fs.StringVar(
			&utils.FileCreatedAfter,
			utils.FileCreatedAfterFN, "",
			"Restore files created after this datetime.")

		fs.StringVar(
			&utils.FileCreatedBefore,
			utils.FileCreatedBeforeFN, "",
			"Restore files created before this datetime.")

		fs.StringVar(
			&utils.FileModifiedAfter,
			utils.FileModifiedAfterFN, "",
			"Restore files modified after this datetime.")
		fs.StringVar(
			&utils.FileModifiedBefore,
			utils.FileModifiedBeforeFN, "",
			"Restore files modified before this datetime.")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const (
	sharePointServiceCommand          = "sharepoint"
	sharePointServiceCommandUseSuffix = "--backup <backupId>"

	//nolint:lll
	sharePointServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore a Site's file named "ServerRenderTemplate.xsl in "Display Templates/Style Sheets" from a specific backup
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --web-url https://example.com --file "ServerRenderTemplate.xsl" --folder "Display Templates/Style Sheets"

# Restore all files from a Site that were created before 2020 when captured in a specific backup
corso restore sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd 
      --web-url https://example.com --folder "Display Templates/Style Sheets" --file-created-before 2020-01-01T00:00:00`
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

	opts := utils.SharePointOpts{
		FileNames:          utils.FileNames,
		FolderPaths:        utils.FolderPaths,
		Library:            utils.Library,
		ListItems:          listItems,
		ListPaths:          listPaths,
		PageFolders:        pageFolders,
		Pages:              pages,
		Sites:              utils.Site,
		WebURLs:            utils.WebURL,
		FileCreatedAfter:   utils.FileCreatedAfter,
		FileCreatedBefore:  utils.FileCreatedBefore,
		FileModifiedAfter:  utils.FileModifiedAfter,
		FileModifiedBefore: utils.FileModifiedBefore,
		Populated:          utils.GetPopulatedFlags(cmd),
	}

	if err := utils.ValidateSharePointRestoreFlags(backupID, opts); err != nil {
		return err
	}

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(common.SimpleDateTimeOneDrive)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeSharePointRestoreDataSelectors(opts)
	utils.FilterSharePointRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, dest)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize SharePoint restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, errors.Errorf("Backup or backup details missing for id %s", backupID))
		}

		return Only(ctx, errors.Wrap(err, "Failed to run SharePoint restore"))
	}

	ds.PrintEntries(ctx)

	return nil
}
