package backup

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

// sharePoint bucket info from flags
var sharepointData []string

const (
	dataLibraries = "libraries"
	dataPages     = "pages"
)

const (
	sharePointServiceCommand                 = "sharepoint"
	sharePointServiceCommandCreateUseSuffix  = "--site <siteURL> | '" + utils.Wildcard + "'"
	sharePointServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	sharePointServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	sharePointServiceCommandCreateExamples = `# Backup SharePoint data for a Site
corso backup create sharepoint --site <siteURL>

# Backup SharePoint for two sites: HR and Team
corso backup create sharepoint --site https://example.com/hr,https://example.com/team

# Backup all SharePoint data for all Sites
corso backup create sharepoint --site '*'`

	sharePointServiceCommandDeleteExamples = `# Delete SharePoint backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd`

	sharePointServiceCommandDetailsExamples = `# Explore a site's files from backup 1234abcd-12ab-cd34-56de-1234abcd
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd

# Find all files that were created before a certain date.
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file-created-before 2015-01-01T00:00:00 --folder "Display Templates/Style Sheets"

# Find all files within a specific library.
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --library documents --folder "Display Templates/Style Sheets"
`
)

// called by backup.go to map subcommands to provider-specific handling.
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, sharePointCreateCmd())

		c.Use = c.Use + " " + sharePointServiceCommandCreateUseSuffix
		c.Example = sharePointServiceCommandCreateExamples

		utils.AddSiteFlag(c)
		utils.AddSiteIDFlag(c)

		fs.StringSliceVar(
			&sharepointData,
			utils.DataFN, nil,
			"Select one or more types of data to backup: "+dataLibraries+" or "+dataPages+".")
		cobra.CheckErr(fs.MarkHidden(utils.DataFN))

		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, sharePointListCmd())
		fs.SortFlags = false

		utils.AddBackupIDFlag(c, false)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, _ = utils.AddCommand(cmd, sharePointDetailsCmd())

		c.Use = c.Use + " " + sharePointServiceCommandDetailsUseSuffix
		c.Example = sharePointServiceCommandDetailsExamples

		options.AddSkipReduceFlag(c)
		utils.AddBackupIDFlag(c, true)
		utils.AddSharePointDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, _ = utils.AddCommand(cmd, sharePointDeleteCmd())

		c.Use = c.Use + " " + sharePointServiceCommandDeleteUseSuffix
		c.Example = sharePointServiceCommandDeleteExamples

		utils.AddBackupIDFlag(c, true)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create sharepoint [<flag>...]`
func sharePointCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     sharePointServiceCommand,
		Short:   "Backup M365 SharePoint service data",
		RunE:    createSharePointCmd,
		Args:    cobra.NoArgs,
		Example: sharePointServiceCommandCreateExamples,
	}
}

// processes an sharepoint service backup.
func createSharePointCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateSharePointBackupCreateFlags(utils.SiteID, utils.WebURL, sharepointData); err != nil {
		return err
	}

	r, acct, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), *acct, connector.Sites, errs)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to connect to Microsoft APIs"))
	}

	sel, err := sharePointBackupCreateSelectors(ctx, utils.SiteID, utils.WebURL, sharepointData, gc)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Retrieving up sharepoint sites by ID and URL"))
	}

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(gc.GetSiteIDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return runBackups(
		ctx,
		r,
		"SharePoint", "site",
		selectorSet,
	)
}

func validateSharePointBackupCreateFlags(sites, weburls, cats []string) error {
	if len(sites) == 0 && len(weburls) == 0 {
		return errors.New(
			"requires one or more --" +
				utils.SiteFN + " urls, or the wildcard --" +
				utils.SiteFN + " *",
		)
	}

	for _, d := range cats {
		if d != dataLibraries && d != dataPages {
			return errors.New(
				d + " is an unrecognized data type; either  " + dataLibraries + "or " + dataPages,
			)
		}
	}

	return nil
}

// TODO: users might specify a data type, this only supports AllData().
func sharePointBackupCreateSelectors(
	ctx context.Context,
	sites, weburls, cats []string,
	gc *connector.GraphConnector,
) (*selectors.SharePointBackup, error) {
	if len(sites) == 0 && len(weburls) == 0 {
		return selectors.NewSharePointBackup(selectors.None()), nil
	}

	for _, site := range sites {
		if site == utils.Wildcard {
			return includeAllSitesWithCategories(cats), nil
		}
	}

	for _, wURL := range weburls {
		if wURL == utils.Wildcard {
			return includeAllSitesWithCategories(cats), nil
		}
	}

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	union, err := gc.UnionSiteIDsAndWebURLs(ctx, sites, weburls, errs)
	if err != nil {
		return nil, err
	}

	sel := selectors.NewSharePointBackup(union)

	return addCategories(sel, cats), nil
}

func includeAllSitesWithCategories(categories []string) *selectors.SharePointBackup {
	sel := addCategories(
		selectors.NewSharePointBackup(selectors.Any()),
		categories)

	return sel
}

func addCategories(sel *selectors.SharePointBackup, cats []string) *selectors.SharePointBackup {
	// Issue #2631: Libraries are the only supported feature for SharePoint at this time.
	if len(cats) == 0 {
		sel.Include(sel.LibraryFolders(selectors.Any()))
	}

	for _, d := range cats {
		switch d {
		case dataLibraries:
			sel.Include(sel.LibraryFolders(selectors.Any()))
		case dataPages:
			sel.Include(sel.Pages(selectors.Any()))
		}
	}

	return sel
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list sharepoint [<flag>...]`
func sharePointListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   sharePointServiceCommand,
		Short: "List the history of M365 SharePoint service backups",
		RunE:  listSharePointCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listSharePointCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, utils.BackupID, path.SharePointService, args)
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete sharepoint [<flag>...]`
func sharePointDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     sharePointServiceCommand,
		Short:   "Delete backed-up M365 SharePoint service data",
		RunE:    deleteSharePointCmd,
		Args:    cobra.NoArgs,
		Example: sharePointServiceCommandDeleteExamples,
	}
}

// deletes a sharePoint service backup.
func deleteSharePointCmd(cmd *cobra.Command, args []string) error {
	return genericDeleteCommand(cmd, utils.BackupID, "SharePoint", args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details onedrive [<flag>...]`
func sharePointDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:     sharePointServiceCommand,
		Short:   "Shows the details of a M365 SharePoint service backup",
		RunE:    detailsSharePointCmd,
		Args:    cobra.NoArgs,
		Example: sharePointServiceCommandDetailsExamples,
	}
}

// lists the history of backup operations
func detailsSharePointCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctx := cmd.Context()
	opts := utils.SharePointOpts{
		FolderPath:         utils.FolderPath,
		FileName:           utils.FileName,
		Library:            utils.Library,
		SiteID:             utils.SiteID,
		WebURL:             utils.WebURL,
		FileCreatedAfter:   fileCreatedAfter,
		FileCreatedBefore:  fileCreatedBefore,
		FileModifiedAfter:  fileModifiedAfter,
		FileModifiedBefore: fileModifiedBefore,

		Populated: utils.GetPopulatedFlags(cmd),
	}

	r, _, err := getAccountAndConnect(ctx)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	ctrlOpts := options.Control()

	ds, err := runDetailsSharePointCmd(ctx, r, utils.BackupID, opts, ctrlOpts.SkipReduce)
	if err != nil {
		return Only(ctx, err)
	}

	if len(ds.Entries) == 0 {
		Info(ctx, selectors.ErrorNoMatchingItems)
		return nil
	}

	ds.PrintEntries(ctx)

	return nil
}

// runDetailsSharePointCmd actually performs the lookup in backup details.
// the fault.Errors return is always non-nil.  Callers should check if
// errs.Failure() == nil.
func runDetailsSharePointCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.SharePointOpts,
	skipReduce bool,
) (*details.Details, error) {
	if err := utils.ValidateSharePointRestoreFlags(backupID, opts); err != nil {
		return nil, err
	}

	ctx = clues.Add(ctx, "backup_id", backupID)

	d, _, errs := r.GetBackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Failure() != nil {
		if errors.Is(errs.Failure(), data.ErrNotFound) {
			return nil, errors.Errorf("no backup exists with the id %s", backupID)
		}

		return nil, errors.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	ctx = clues.Add(ctx, "details_entries", len(d.Entries))

	if !skipReduce {
		sel := utils.IncludeSharePointRestoreDataSelectors(opts)
		utils.FilterSharePointRestoreInfoSelectors(sel, opts)
		d = sel.Reduce(ctx, d, errs)
	}

	return d, nil
}
