package utils

import (
	"context"
	"net/url"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	ListFolderFN = "list"
	ListItemFN   = "list-item"
	PageFolderFN = "page-folder"
	PageFN       = "page"
)

// flag population variables
var (
	ListFolder []string
	ListItem   []string
	PageFolder []string
	Page       []string
)

type SharePointOpts struct {
	SiteID []string
	WebURL []string

	Library            string
	FileName           []string // for libraries, to duplicate onedrive interface
	FolderPath         []string // for libraries, to duplicate onedrive interface
	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	ListFolder []string
	ListItem   []string

	PageFolder []string
	Page       []string

	Populated PopulatedFlags
}

func MakeSharePointOpts(cmd *cobra.Command) SharePointOpts {
	return SharePointOpts{
		SiteID: SiteIDFV,
		WebURL: WebURLFV,

		Library:            LibraryFV,
		FileName:           FileNameFV,
		FolderPath:         FolderPathFV,
		FileCreatedAfter:   FileCreatedAfterFV,
		FileCreatedBefore:  FileCreatedBeforeFV,
		FileModifiedAfter:  FileModifiedAfterFV,
		FileModifiedBefore: FileModifiedBeforeFV,

		ListFolder: ListFolder,
		ListItem:   ListItem,

		Page:       Page,
		PageFolder: PageFolder,

		Populated: GetPopulatedFlags(cmd),
	}
}

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
		&ListFolder,
		ListFolderFN, nil,
		"Select lists by name; accepts '"+Wildcard+"' to select all lists.")
	cobra.CheckErr(fs.MarkHidden(ListFolderFN))
	fs.StringSliceVar(
		&ListItem,
		ListItemFN, nil,
		"Select lists by item name; accepts '"+Wildcard+"' to select all lists.")
	cobra.CheckErr(fs.MarkHidden(ListItemFN))

	// pages

	fs.StringSliceVar(
		&PageFolder,
		PageFolderFN, nil,
		"Select pages by folder name; accepts '"+Wildcard+"' to select all pages.")
	cobra.CheckErr(fs.MarkHidden(PageFolderFN))
	fs.StringSliceVar(
		&Page,
		PageFN, nil,
		"Select pages by item name; accepts '"+Wildcard+"' to select all pages.")
	cobra.CheckErr(fs.MarkHidden(PageFN))
}

// ValidateSharePointRestoreFlags checks common flags for correctness and interdependencies
func ValidateSharePointRestoreFlags(backupID string, opts SharePointOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	// ensure url can parse all weburls provided by --site.
	if _, ok := opts.Populated[SiteFN]; ok {
		for _, wu := range opts.WebURL {
			if _, err := url.Parse(wu); err != nil {
				return clues.New("invalid site url: " + wu)
			}
		}
	}

	if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		return clues.New("invalid time format for " + FileCreatedAfterFN)
	}

	if _, ok := opts.Populated[FileCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.FileCreatedBefore) {
		return clues.New("invalid time format for " + FileCreatedBeforeFN)
	}

	if _, ok := opts.Populated[FileModifiedAfterFN]; ok && !IsValidTimeFormat(opts.FileModifiedAfter) {
		return clues.New("invalid time format for " + FileModifiedAfterFN)
	}

	if _, ok := opts.Populated[FileModifiedBeforeFN]; ok && !IsValidTimeFormat(opts.FileModifiedBefore) {
		return clues.New("invalid time format for " + FileModifiedBeforeFN)
	}

	return nil
}

// AddSharePointInfo adds the scope of the provided values to the selector's
// filter set
func AddSharePointInfo(
	sel *selectors.SharePointRestore,
	v string,
	f func(string) []selectors.SharePointScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// IncludeSharePointRestoreDataSelectors builds the common data-selector
// inclusions for SharePoint commands.
func IncludeSharePointRestoreDataSelectors(ctx context.Context, opts SharePointOpts) *selectors.SharePointRestore {
	sites := opts.SiteID

	lfp, lfn := len(opts.FolderPath), len(opts.FileName)
	ls, lwu := len(opts.SiteID), len(opts.WebURL)
	slp, sli := len(opts.ListFolder), len(opts.ListItem)
	pf, pi := len(opts.PageFolder), len(opts.Page)

	if ls == 0 {
		sites = selectors.Any()
	}

	sel := selectors.NewSharePointRestore(sites)

	if lfp+lfn+lwu+slp+sli+pf+pi == 0 {
		sel.Include(sel.AllData())
		return sel
	}

	if lfp+lfn > 0 {
		if lfn == 0 {
			opts.FileName = selectors.Any()
		}

		opts.FolderPath = trimFolderSlash(opts.FolderPath)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.FolderPath)

		if len(containsFolders) > 0 {
			sel.Include(sel.LibraryItems(containsFolders, opts.FileName))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.LibraryItems(prefixFolders, opts.FileName, selectors.PrefixMatch()))
		}
	}

	if slp+sli > 0 {
		if sli == 0 {
			opts.ListItem = selectors.Any()
		}

		opts.ListFolder = trimFolderSlash(opts.ListFolder)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.ListFolder)

		if len(containsFolders) > 0 {
			sel.Include(sel.ListItems(containsFolders, opts.ListItem))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.ListItems(prefixFolders, opts.ListItem, selectors.PrefixMatch()))
		}
	}

	if pf+pi > 0 {
		if pi == 0 {
			opts.Page = selectors.Any()
		}

		opts.PageFolder = trimFolderSlash(opts.PageFolder)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.PageFolder)

		if len(containsFolders) > 0 {
			sel.Include(sel.PageItems(containsFolders, opts.Page))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.PageItems(prefixFolders, opts.Page, selectors.PrefixMatch()))
		}
	}

	if lwu > 0 {
		urls := make([]string, len(opts.WebURL))

		for _, wu := range opts.WebURL {
			u, err := url.Parse(wu)
			if err != nil {
				// shouldn't be possible to err, if we called validation first.
				logger.Ctx(ctx).With("web_url", wu).Error("malformed web url")
				continue
			}

			// for normalization, remove the scheme
			urls = append(urls, strings.TrimPrefix(u.Scheme, u.String()))
		}

		sel.Include(sel.WebURL(urls))
	}

	return sel
}

// FilterSharePointRestoreInfoSelectors builds the common info-selector filters.
func FilterSharePointRestoreInfoSelectors(
	sel *selectors.SharePointRestore,
	opts SharePointOpts,
) {
	AddSharePointInfo(sel, opts.Library, sel.Library)
	AddSharePointInfo(sel, opts.FileCreatedAfter, sel.CreatedAfter)
	AddSharePointInfo(sel, opts.FileCreatedBefore, sel.CreatedBefore)
	AddSharePointInfo(sel, opts.FileModifiedAfter, sel.ModifiedAfter)
	AddSharePointInfo(sel, opts.FileModifiedBefore, sel.ModifiedBefore)
}
