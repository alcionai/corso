package backup

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

// sharePoint bucket info from flags
var (
	libraryItems []string
	libraryPaths []string
	pageFolders  []string
	page         []string
	site         []string
	weburl       []string

	sharepointData []string
)

const (
	dataLibraries = "libraries"
	dataPages     = "pages"
)

const (
	sharePointServiceCommand                 = "sharepoint"
	sharePointServiceCommandCreateUseSuffix  = "--site <siteId> | '" + utils.Wildcard + "'"
	sharePointServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	sharePointServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	sharePointServiceCommandCreateExamples = `# Backup SharePoint data for <site>
corso backup create sharepoint --site <site_id>

# Backup SharePoint for Alice and Bob
corso backup create sharepoint --site <site_id_1>,<site_id_2>

# TODO: Site IDs may contain commas.  We'll need to warn the site about escaping them.

# Backup all SharePoint data for all sites
corso backup create sharepoint --site '*'`

	sharePointServiceCommandDeleteExamples = `# Delete SharePoint backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd`

	sharePointServiceCommandDetailsExamples = `# Explore <site>'s files from backup 1234abcd-12ab-cd34-56de-1234abcd

corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd --site <site_id>`
)

// called by backup.go to map subcommands to provider-specific handling.
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, sharePointCreateCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + sharePointServiceCommandCreateUseSuffix
		c.Example = sharePointServiceCommandCreateExamples

		fs.StringArrayVar(&site,
			utils.SiteFN, nil,
			"Backup SharePoint data by site ID; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(&weburl,
			utils.WebURLFN, nil,
			"Restore data by site webURL; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(
			&sharepointData,
			utils.DataFN, nil,
			"Select one or more types of data to backup: "+dataLibraries+" or "+dataPages+".")
		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, sharePointListCmd(), utils.MarkPreReleaseCommand())

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to retrieve.")

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, sharePointDetailsCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + sharePointServiceCommandDetailsUseSuffix
		c.Example = sharePointServiceCommandDetailsExamples

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to retrieve.")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))

		// sharepoint hierarchy flags

		fs.StringSliceVar(
			&libraryPaths,
			utils.LibraryFN, nil,
			"Select backup details by Library name.")

		fs.StringSliceVar(
			&libraryItems,
			utils.LibraryItemFN, nil,
			"Select backup details by library item name or ID.")

		fs.StringArrayVar(&site,
			utils.SiteFN, nil,
			"Select backup details by site ID; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(&weburl,
			utils.WebURLFN, nil,
			"Select backup data by site webURL; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(
			&pageFolders,
			utils.PageFN, nil,
			"Select backup data by site ID; accepts '"+utils.Wildcard+"' to select all sites.")

		fs.StringSliceVar(
			&page,
			utils.PageItemFN, nil,
			"Select backup data by file name; accepts '"+utils.Wildcard+"' to select all pages within the site.",
		)

		// info flags

		// fs.StringVar(
		// 	&fileCreatedAfter,
		// 	utils.FileCreatedAfterFN, "",
		// 	"Select backup details for items created after this datetime.")

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, sharePointDeleteCmd(), utils.MarkPreReleaseCommand())

		c.Use = c.Use + " " + sharePointServiceCommandDeleteUseSuffix
		c.Example = sharePointServiceCommandDeleteExamples

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to delete. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))
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

	if err := validateSharePointBackupCreateFlags(site, weburl, sharepointData); err != nil {
		return err
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to connect to Microsoft APIs"))
	}

	sel, err := sharePointBackupCreateSelectors(ctx, site, weburl, sharepointData, gc)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Retrieving up sharepoint sites by ID and WebURL"))
	}

	var (
		errs *multierror.Error
		bIDs []model.StableID
	)

	for _, discSel := range sel.SplitByResourceOwner(gc.GetSiteIDs()) {
		bo, err := r.NewBackup(ctx, discSel.Selector)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(
				err,
				"Failed to initialize SharePoint backup for site %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		err = bo.Run(ctx)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(
				err,
				"Failed to run SharePoint backup for site %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		bIDs = append(bIDs, bo.Results.BackupID)
	}

	bups, ferrs := r.Backups(ctx, bIDs)
	// TODO: print/log recoverable errors
	if ferrs.Err() != nil {
		return Only(ctx, errors.Wrap(ferrs.Err(), "Unable to retrieve backup results from storage"))
	}

	backup.PrintAll(ctx, bups)

	if e := errs.ErrorOrNil(); e != nil {
		return Only(ctx, e)
	}

	return nil
}

func validateSharePointBackupCreateFlags(sites, weburls, data []string) error {
	if len(sites) == 0 && len(weburls) == 0 {
		return errors.New(
			"requires one or more --" +
				utils.SiteFN + " ids, --" +
				utils.WebURLFN + " urls, or the wildcard --" +
				utils.SiteFN + " *",
		)
	}

	for _, d := range data {
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
	sites, weburls, data []string,
	gc *connector.GraphConnector,
) (*selectors.SharePointBackup, error) {
	if len(sites) == 0 && len(weburls) == 0 {
		return selectors.NewSharePointBackup(selectors.None()), nil
	}

	for _, site := range sites {
		if site == utils.Wildcard {
			sel := selectors.NewSharePointBackup(selectors.Any())
			sel.Include(sel.AllData())

			return sel, nil
		}
	}

	for _, wURL := range weburls {
		if wURL == utils.Wildcard {
			sel := selectors.NewSharePointBackup(selectors.Any())
			sel.Include(sel.AllData())

			return sel, nil
		}
	}

	union, err := gc.UnionSiteIDsAndWebURLs(ctx, sites, weburls)
	if err != nil {
		return nil, err
	}

	sel := selectors.NewSharePointBackup(union)
	if len(data) == 0 {
		sel.Include(sel.AllData())

		return sel, nil
	}

	for _, d := range data {
		switch d {
		case dataLibraries:
			sel.Include(sel.Libraries(selectors.Any()))
		case dataPages:
			sel.Include(sel.Pages(selectors.Any()))
		}
	}

	return sel, nil
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
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if len(backupID) > 0 {
		b, err := r.Backup(ctx, model.StableID(backupID))
		if err != nil {
			if errors.Is(err, data.ErrNotFound) {
				return Only(ctx, errors.Errorf("No backup exists with the id %s", backupID))
			}

			return Only(ctx, errors.Wrap(err, "Failed to find backup "+backupID))
		}

		b.Print(ctx)

		return nil
	}

	bs, err := r.BackupsByTag(ctx, store.Service(path.SharePointService))
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to list backups in the repository"))
	}

	backup.PrintAll(ctx, bs)

	return nil
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
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if err := r.DeleteBackup(ctx, model.StableID(backupID)); err != nil {
		return Only(ctx, errors.Wrapf(err, "Deleting backup %s", backupID))
	}

	Info(ctx, "Deleted SharePoint backup ", backupID)

	return nil
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
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, acct, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	opts := utils.SharePointOpts{
		LibraryItems: libraryItems,
		LibraryPaths: libraryPaths,
		Sites:        site,
		WebURLs:      weburl,

		Populated: utils.GetPopulatedFlags(cmd),
	}

	ds, err := runDetailsSharePointCmd(ctx, r, backupID, opts)
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
// errs.Err() == nil.
func runDetailsSharePointCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.SharePointOpts,
) (*details.Details, error) {
	if err := utils.ValidateSharePointRestoreFlags(backupID, opts); err != nil {
		return nil, err
	}

	d, _, errs := r.BackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Err() != nil {
		if errors.Is(errs.Err(), data.ErrNotFound) {
			return nil, errors.Errorf("no backup exists with the id %s", backupID)
		}

		return nil, errors.Wrap(errs.Err(), "Failed to get backup details in the repository")
	}

	sel := utils.IncludeSharePointRestoreDataSelectors(opts)
	utils.FilterSharePointRestoreInfoSelectors(sel, opts)

	return sel.Reduce(ctx, d, errs), nil
}
