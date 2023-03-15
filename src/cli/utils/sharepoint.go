package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/spf13/cobra"
)

const (
	ListItemFN   = "list-item"
	ListFN       = "list"
	PageFolderFN = "page-folders"
	PagesFN      = "pages"
)

// flag population variables
var (
	PageFolders []string
	Pages       []string
)

type SharePointOpts struct {
	Library     string
	FileNames   []string // for libraries, to duplicate onedrive interface
	FolderPaths []string // for libraries, to duplicate onedrive interface

	ListItems []string
	ListPaths []string

	PageFolders []string
	Pages       []string

	Sites   []string
	WebURLs []string

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
		"Select only this library. Default includes all libraries.")

	fs.StringSliceVar(
		&FolderPaths,
		FolderFN, nil,
		"Select by folder; defaults to root.")

	fs.StringSliceVar(
		&FileNames,
		FileFN, nil,
		"Select by file name.")

	fs.StringSliceVar(
		&PageFolders,
		PageFolderFN, nil,
		"Select pages by folder name; accepts '"+Wildcard+"' to select all folders.")
	cobra.CheckErr(fs.MarkHidden(PageFolderFN))

	fs.StringSliceVar(
		&Pages,
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
		return errors.New("a backup ID is required")
	}

	if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		return errors.New("invalid time format for " + FileCreatedAfterFN)
	}

	if _, ok := opts.Populated[FileCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.FileCreatedBefore) {
		return errors.New("invalid time format for " + FileCreatedBeforeFN)
	}

	if _, ok := opts.Populated[FileModifiedAfterFN]; ok && !IsValidTimeFormat(opts.FileModifiedAfter) {
		return errors.New("invalid time format for " + FileModifiedAfterFN)
	}

	if _, ok := opts.Populated[FileModifiedBeforeFN]; ok && !IsValidTimeFormat(opts.FileModifiedBefore) {
		return errors.New("invalid time format for " + FileModifiedBeforeFN)
	}

	return nil
}

// AddSharePointFilter adds the scope of the provided values to the selector's
// filter set
func AddSharePointFilter(
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
	sites := opts.Sites

	lfp, lfn := len(opts.FolderPaths), len(opts.FileNames)
	ls, lwu := len(opts.Sites), len(opts.WebURLs)
	slp, sli := len(opts.ListPaths), len(opts.ListItems)
	pf, pi := len(opts.PageFolders), len(opts.Pages)

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
			opts.FileNames = selectors.Any()
		}

		opts.FolderPaths = trimFolderSlash(opts.FolderPaths)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.FolderPaths)

		if len(containsFolders) > 0 {
			sel.Include(sel.LibraryItems(containsFolders, opts.FileNames))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.LibraryItems(prefixFolders, opts.FileNames, selectors.PrefixMatch()))
		}
	}

	if slp+sli > 0 {
		if sli == 0 {
			opts.ListItems = selectors.Any()
		}

		opts.ListPaths = trimFolderSlash(opts.ListPaths)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.ListPaths)

		if len(containsFolders) > 0 {
			sel.Include(sel.ListItems(containsFolders, opts.ListItems))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.ListItems(prefixFolders, opts.ListItems, selectors.PrefixMatch()))
		}
	}

	if pf+pi > 0 {
		if pi == 0 {
			opts.Pages = selectors.Any()
		}

		opts.PageFolders = trimFolderSlash(opts.PageFolders)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.PageFolders)

		if len(containsFolders) > 0 {
			sel.Include(sel.PageItems(containsFolders, opts.Pages))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.PageItems(prefixFolders, opts.Pages, selectors.PrefixMatch()))
		}
	}

	if lwu > 0 {
		opts.WebURLs = trimFolderSlash(opts.WebURLs)
		containsURLs, suffixURLs := splitFoldersIntoContainsAndPrefix(opts.WebURLs)

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
	AddSharePointFilter(sel, opts.Library, sel.Library)
	AddSharePointFilter(sel, opts.FileCreatedAfter, sel.CreatedAfter)
	AddSharePointFilter(sel, opts.FileCreatedBefore, sel.CreatedBefore)
	AddSharePointFilter(sel, opts.FileModifiedAfter, sel.ModifiedAfter)
	AddSharePointFilter(sel, opts.FileModifiedBefore, sel.ModifiedBefore)
}
