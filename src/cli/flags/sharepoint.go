package flags

import (
	"github.com/spf13/cobra"
)

const (
	DataLibraries = "libraries"
	DataPages     = "pages"
)

const (
	LibraryFN    = "library"
	ListFolderFN = "list"
	ListItemFN   = "list-item"
	PageFolderFN = "page-folder"
	PageFN       = "page"
	SiteFN       = "site"    // site only accepts WebURL values
	SiteIDFN     = "site-id" // site-id accepts actual site ids
)

var (
	LibraryFV    string
	ListFolderFV []string
	ListItemFV   []string
	PageFolderFV []string
	PageFV       []string
	SiteIDFV     []string
	WebURLFV     []string
)

// AddSharePointDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddSharePointDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	// libraries

	fs.StringVar(
		&LibraryFV,
		LibraryFN, "",
		"Select only this library; defaults to all libraries.")
	fs.StringSliceVar(
		&FolderPathFV,
		FolderFN, nil,
		"Select by folder; defaults to root.")
	fs.StringSliceVar(
		&FileNameFV,
		FileFN, nil,
		"Select by file name.")
	fs.StringVar(
		&FileCreatedAfterFV,
		FileCreatedAfterFN, "",
		"Select files created after this datetime.")
	fs.StringVar(
		&FileCreatedBeforeFV,
		FileCreatedBeforeFN, "",
		"Select files created before this datetime.")
	fs.StringVar(
		&FileModifiedAfterFV,
		FileModifiedAfterFN, "",
		"Select files modified after this datetime.")
	fs.StringVar(
		&FileModifiedBeforeFV,
		FileModifiedBeforeFN, "",
		"Select files modified before this datetime.")

	// lists

	fs.StringSliceVar(
		&ListFolderFV,
		ListFolderFN, nil,
		"Select lists by name; accepts '"+Wildcard+"' to select all lists.")
	cobra.CheckErr(fs.MarkHidden(ListFolderFN))
	fs.StringSliceVar(
		&ListItemFV,
		ListItemFN, nil,
		"Select lists by item name; accepts '"+Wildcard+"' to select all lists.")
	cobra.CheckErr(fs.MarkHidden(ListItemFN))

	// pages

	fs.StringSliceVar(
		&PageFolderFV,
		PageFolderFN, nil,
		"Select pages by folder name; accepts '"+Wildcard+"' to select all pages.")
	cobra.CheckErr(fs.MarkHidden(PageFolderFN))
	fs.StringSliceVar(
		&PageFV,
		PageFN, nil,
		"Select pages by item name; accepts '"+Wildcard+"' to select all pages.")
	cobra.CheckErr(fs.MarkHidden(PageFN))
}

// AddSiteIDFlag adds the --site-id flag, which accepts site ID values.
// This flag is hidden, since we expect users to prefer the --site url
// and do not want to encourage confusion.
// TODO(meain): --site is the primary one, but it would be useful to have comepletion for this as well
func AddSiteIDFlag(cmd *cobra.Command, multiple bool) {
	fs := cmd.Flags()

	message := "ID of the site to operate on"
	if multiple {
		//nolint:lll
		message += "; accepts '" + Wildcard + "' to select all sites.  Args cannot be comma-delimited and must use multiple flags."
	}

	// note string ARRAY var.  IDs naturally contain commas, so we cannot accept
	// duplicate values within a flag declaration.  ie: --site-id a,b,c does not
	// work.  Users must call --site-id a --site-id b --site-id c.
	fs.StringArrayVar(&SiteIDFV, SiteIDFN, nil, message)
	cobra.CheckErr(fs.MarkHidden(SiteIDFN))
}

// AddSiteFlag adds the --site flag, which accepts webURL values.
func AddSiteFlag(
	cmd *cobra.Command,
	multiple bool,
	completionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective),
) {
	message := "Web URL of the site to operate on"
	if multiple {
		message += "; accepts '" + Wildcard + "' to select all sites."
	}

	cmd.Flags().StringSliceVar(&WebURLFV, SiteFN, nil, message)

	cobra.CheckErr(cmd.RegisterFlagCompletionFunc(SiteFN, completionFunc))
}
