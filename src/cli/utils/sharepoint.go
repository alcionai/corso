package utils

import (
	"errors"
	"fmt"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	ListItemFN   = "list-item"
	ListFN       = "list"
	PageFolderFN = "page-folders"
	PagesFN      = "pages"
)

type SharePointOpts struct {
	FileNames          []string // for libraries, to duplicate onedrive interface
	FolderPaths        []string // for libraries, to duplicate onedrive interface
	Library            string
	ListItems          []string
	ListPaths          []string
	PageFolders        []string
	Pages              []string
	Sites              []string
	WebURLs            []string
	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Populated PopulatedFlags
}

// ValidateSharePointRestoreFlags checks common flags for correctness and interdependencies
func ValidateSharePointRestoreFlags(backupID string, opts SharePointOpts) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		fmt.Printf("What was I sent: %v\n", opts.FileCreatedAfter)
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
