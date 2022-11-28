package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	LibraryItemFN = "library-item"
	LibraryFN     = "library"
)

type SharePointOpts struct {
	Sites        []string
	LibraryItems []string
	LibraryPaths []string

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
func IncludeSharePointRestoreDataSelectors(
	sel *selectors.SharePointRestore,
	opts SharePointOpts,
) {
	lp, ln := len(opts.LibraryPaths), len(opts.LibraryItems)

	// only use the inclusion if either a path or item name
	// is specified
	if lp+ln == 0 {
		return
	}

	if len(opts.Sites) == 0 {
		opts.Sites = selectors.Any()
	}

	// either scope the request to a set of sites
	if lp+ln == 0 {
		sel.Include(sel.Sites(opts.Sites))

		return
	}

	opts.LibraryPaths = trimFolderSlash(opts.LibraryPaths)

	if ln == 0 {
		opts.LibraryItems = selectors.Any()
	}

	containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.LibraryPaths)

	if len(containsFolders) > 0 {
		sel.Include(sel.LibraryItems(opts.Sites, containsFolders, opts.LibraryItems))
	}

	if len(prefixFolders) > 0 {
		sel.Include(sel.LibraryItems(opts.Sites, prefixFolders, opts.LibraryItems, selectors.PrefixMatch()))
	}
}

// FilterSharePointRestoreInfoSelectors builds the common info-selector filters.
func FilterSharePointRestoreInfoSelectors(
	sel *selectors.SharePointRestore,
	opts SharePointOpts,
) {
	// AddSharePointFilter(sel, opts.FileCreatedAfter, sel.CreatedAfter)
}
