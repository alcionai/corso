package utils

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	ListItemFN   = "list-item"
	ListFN       = "list"
	PageFolderFN = "page-folder"
	PagesFN      = "page"
)

// flag population variables
var (
	PageFolder []string
	Page       []string
)

type SharePointOpts struct {
	Library    string
	FileName   []string // for libraries, to duplicate onedrive interface
	FolderPath []string // for libraries, to duplicate onedrive interface

	ListItem []string
	ListPath []string

	PageFolder []string
	Page       []string

	SiteID []string
	WebURL []string

	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Populated PopulatedFlags
}

// AddSharePointDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddSharePointDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringVar(
		&Library,
		LibraryFN, "",
		"Select only this library; defaults to all libraries.")

	fs.StringSliceVar(
		&FolderPath,
		FolderFN, nil,
		"Select by folder; defaults to root.")

	fs.StringSliceVar(
		&FileName,
		FileFN, nil,
		"Select by file name.")

	fs.StringSliceVar(
		&PageFolder,
		PageFolderFN, nil,
		"Select pages by folder name; accepts '"+Wildcard+"' to select all folders.")
	cobra.CheckErr(fs.MarkHidden(PageFolderFN))

	fs.StringSliceVar(
		&Page,
		PagesFN, nil,
		"Select pages by item name; accepts '"+Wildcard+"' to select all pages within the site.")
	cobra.CheckErr(fs.MarkHidden(PagesFN))

	fs.StringVar(
		&FileCreatedAfter,
		FileCreatedAfterFN, "",
		"Select files created after this datetime.")

	fs.StringVar(
		&FileCreatedBefore,
		FileCreatedBeforeFN, "",
		"Select files created before this datetime.")

	fs.StringVar(
		&FileModifiedAfter,
		FileModifiedAfterFN, "",
		"Select files modified after this datetime.")
	fs.StringVar(
		&FileModifiedBefore,
		FileModifiedBeforeFN, "",
		"Select files modified before this datetime.")
}

// ValidateSharePointRestoreFlags checks common flags for correctness and interdependencies
func ValidateSharePointRestoreFlags(backupID string, opts SharePointOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
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
func IncludeSharePointRestoreDataSelectors(opts SharePointOpts) *selectors.SharePointRestore {
	sites := opts.SiteID

	lfp, lfn := len(opts.FolderPath), len(opts.FileName)
	ls, lwu := len(opts.SiteID), len(opts.WebURL)
	slp, sli := len(opts.ListPath), len(opts.ListItem)
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

		opts.ListPath = trimFolderSlash(opts.ListPath)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.ListPath)

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
		opts.WebURL = trimFolderSlash(opts.WebURL)
		containsURLs, suffixURLs := splitFoldersIntoContainsAndPrefix(opts.WebURL)

		if len(containsURLs) > 0 {
			sel.Include(sel.WebURL(containsURLs))
		}

		if len(suffixURLs) > 0 {
			sel.Include(sel.WebURL(suffixURLs, selectors.SuffixMatch()))
		}
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
