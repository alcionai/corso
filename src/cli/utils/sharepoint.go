package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	LibraryItemFN = "library-item"
	LibraryFN     = "library"
	ListItemFN    = "list-item"
	ListFN        = "list"
	WebURLFN      = "web-url"
)

type SharePointOpts struct {
	LibraryItems []string
	LibraryPaths []string
	ListItems    []string
	ListPaths    []string
	Sites        []string
	WebURLs      []string

	Populated PopulatedFlags
}

// ValidateSharePointRestoreFlags checks common flags for correctness and interdependencies
func ValidateSharePointRestoreFlags(backupID string, opts SharePointOpts) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	// if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
	// 	return errors.New("invalid time format for created-after")
	// }

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

	lp, li := len(opts.LibraryPaths), len(opts.LibraryItems)
	ls, lwu := len(opts.Sites), len(opts.WebURLs)
	slp, sli := len(opts.ListPaths), len(opts.ListItems)

	if ls == 0 {
		sites = selectors.Any()
	}

	sel := selectors.NewSharePointRestore(sites)

	if lp+li+lwu+slp+sli == 0 {
		sel.Include(sel.Sites(sites))
		return sel
	}

	if lp+li > 0 {
		if li == 0 {
			opts.LibraryItems = selectors.Any()
		}

		opts.LibraryPaths = trimFolderSlash(opts.LibraryPaths)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.LibraryPaths)

		if len(containsFolders) > 0 {
			sel.Include(sel.LibraryItems(sites, containsFolders, opts.LibraryItems))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.LibraryItems(sites, prefixFolders, opts.LibraryItems, selectors.PrefixMatch()))
		}
	}

	if slp+sli > 0 {
		if sli == 0 {
			opts.ListItems = selectors.Any()
		}

		opts.ListPaths = trimFolderSlash(opts.ListPaths)
		containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.ListPaths)

		if len(containsFolders) > 0 {
			sel.Include(sel.ListItems(opts.Sites, containsFolders, opts.ListItems))
		}

		if len(prefixFolders) > 0 {
			sel.Include(sel.ListItems(opts.Sites, prefixFolders, opts.ListItems, selectors.PrefixMatch()))
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
	// AddSharePointFilter(sel, opts.FileCreatedAfter, sel.CreatedAfter)
}
