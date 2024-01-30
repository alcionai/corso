package backup

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	sharePointServiceCommand                 = "sharepoint"
	sharePointServiceCommandCreateUseSuffix  = "--site <siteURL> | '" + flags.Wildcard + "'"
	sharePointServiceCommandDeleteUseSuffix  = "--backups <backupId>"
	sharePointServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	sharePointServiceCommandCreateExamples = `# Backup SharePoint data in the HR Site
corso backup create sharepoint --site https://example.com/hr

# Backup SharePoint for the HR and Team sites
corso backup create sharepoint --site https://example.com/hr,https://example.com/team

# Backup all SharePoint data for all Sites
corso backup create sharepoint --site '*'

# Backup all SharePoint list data for a Site
corso backup create sharepoint --site https://example.com/hr --data lists
`

	sharePointServiceCommandDeleteExamples = `# Delete SharePoint backup with ID 1234abcd-12ab-cd34-56de-1234abcd \
and 1234abcd-12ab-cd34-56de-1234abce
corso backup delete sharepoint --backups 1234abcd-12ab-cd34-56de-1234abcd,1234abcd-12ab-cd34-56de-1234abce`

	sharePointServiceCommandDetailsExamples = `# Explore items in the HR site's latest backup (1234abcd...)
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore files in the folder "Reports" named "Fiscal 22"
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file-name "Fiscal 22" --folder "Reports"

# Explore files in the folder ""Display Templates/Style Sheets"" created before the end of 2015.
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file-created-before 2015-01-01T00:00:00 --folder "Display Templates/Style Sheets"

# Explore all files within the document library "Work Documents"
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --library "Work Documents"

# Explore lists by their name(s)
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --list "list-name-1,list-name-2"

# Explore lists created after a given time
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --list-created-after 2024-01-01T12:23:34

# Explore lists created before a given time
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --list-created-before 2024-01-01T12:23:34

# Explore lists modified before a given time
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --list-modified-before 2024-01-01T12:23:34

# Explore lists modified after a given time
corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --list-modified-after 2024-01-01T12:23:34`
)

// called by backup.go to map subcommands to provider-specific handling.
func addSharePointCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case createCommand:
		c, _ = utils.AddCommand(cmd, sharePointCreateCmd())

		c.Use = c.Use + " " + sharePointServiceCommandCreateUseSuffix
		c.Example = sharePointServiceCommandCreateExamples

		flags.AddSiteFlag(c, true)
		flags.AddSiteIDFlag(c, true)
		// [TODO](hitesh) to add lists flag to invoke backup for lists
		// when explicit invoke is not required anymore
		flags.AddDataFlag(c, []string{flags.DataLibraries}, true)
		flags.AddGenericBackupFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(cmd, sharePointListCmd())

		flags.AddBackupIDFlag(c, false)
		flags.AddAllBackupListFlags(c)

	case detailsCommand:
		c, _ = utils.AddCommand(cmd, sharePointDetailsCmd())

		c.Use = c.Use + " " + sharePointServiceCommandDetailsUseSuffix
		c.Example = sharePointServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)
		flags.AddBackupIDFlag(c, true)
		flags.AddSharePointDetailsAndRestoreFlags(c)

	case deleteCommand:
		c, _ = utils.AddCommand(cmd, sharePointDeleteCmd())

		c.Use = c.Use + " " + sharePointServiceCommandDeleteUseSuffix
		c.Example = sharePointServiceCommandDeleteExamples

		flags.AddMultipleBackupIDsFlag(c, false)
		flags.AddBackupIDFlag(c, false)
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

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := validateSharePointBackupCreateFlags(flags.SiteIDFV, flags.WebURLFV, flags.CategoryDataFV); err != nil {
		return err
	}

	r, acct, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		path.SharePointService)
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	svcCli, err := m365.NewM365Client(ctx, *acct)
	if err != nil {
		return Only(ctx, clues.Stack(err))
	}

	ins, err := svcCli.SitesMap(ctx, errs)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to retrieve M365 sites"))
	}

	sel, err := sharePointBackupCreateSelectors(ctx, ins, flags.SiteIDFV, flags.WebURLFV, flags.CategoryDataFV)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving up sharepoint sites by ID and URL"))
	}

	selectorSet := []selectors.Selector{}

	for _, discSel := range sel.SplitByResourceOwner(ins.IDs()) {
		selectorSet = append(selectorSet, discSel.Selector)
	}

	return genericCreateCommand(
		ctx,
		r,
		"SharePoint",
		selectorSet,
		ins)
}

func validateSharePointBackupCreateFlags(sites, weburls, cats []string) error {
	if len(sites) == 0 && len(weburls) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.SiteFN + " urls, or the wildcard --" +
				flags.SiteFN + " *")
	}

	allowedCats := utils.SharePointAllowedCategories()

	for _, d := range cats {
		if _, ok := allowedCats[d]; !ok {
			return clues.New(
				d + " is an unrecognized data type; only  " + flags.DataLibraries + " supported")
		}
	}

	return nil
}

// TODO: users might specify a data type, this only supports AllData().
func sharePointBackupCreateSelectors(
	ctx context.Context,
	ins idname.Cacher,
	sites, weburls, cats []string,
) (*selectors.SharePointBackup, error) {
	if len(sites) == 0 && len(weburls) == 0 {
		return selectors.NewSharePointBackup(selectors.None()), nil
	}

	if filters.PathContains(sites).Compare(flags.Wildcard) {
		return includeAllSitesWithCategories(ins, cats), nil
	}

	if filters.PathContains(weburls).Compare(flags.Wildcard) {
		return includeAllSitesWithCategories(ins, cats), nil
	}

	sel := selectors.NewSharePointBackup(append(slices.Clone(sites), weburls...))

	return utils.AddCategories(sel, cats), nil
}

func includeAllSitesWithCategories(ins idname.Cacher, categories []string) *selectors.SharePointBackup {
	return utils.AddCategories(selectors.NewSharePointBackup(ins.IDs()), categories)
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
	return genericListCommand(cmd, flags.BackupIDFV, path.SharePointService, args)
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
	backupIDValue := []string{}

	if len(flags.BackupIDsFV) > 0 {
		backupIDValue = flags.BackupIDsFV
	} else if len(flags.BackupIDFV) > 0 {
		backupIDValue = append(backupIDValue, flags.BackupIDFV)
	} else {
		return clues.New("either --backup or --backups flag is required")
	}

	return genericDeleteCommand(cmd, path.SharePointService, "SharePoint", backupIDValue, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details SharePoint [<flag>...]`
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

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	return runDetailsSharePointCmd(cmd)
}

func runDetailsSharePointCmd(cmd *cobra.Command) error {
	ctx := cmd.Context()
	opts := utils.MakeSharePointOpts(cmd)

	sel := utils.IncludeSharePointRestoreDataSelectors(ctx, opts)
	sel.Configure(selectors.Config{OnlyMatchItemNames: true})
	utils.FilterSharePointRestoreInfoSelectors(sel, opts)

	ds, err := genericDetailsCommand(cmd, flags.BackupIDFV, sel.Selector)
	if err != nil {
		return Only(ctx, err)
	}

	if len(ds.Entries) > 0 {
		ds.PrintEntries(ctx)
	} else {
		Info(ctx, selectors.ErrorNoMatchingItems)
	}

	return nil
}
