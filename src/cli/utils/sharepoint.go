package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	LibraryItemFN = "library-item"
	LibraryFN     = "library"
	WebURLFN      = "web-url"
)

type SharePointOpts struct {
	LibraryItems []string
	LibraryPaths []string
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
func IncludeSharePointRestoreDataSelectors(
	sel *selectors.SharePointRestore,
	opts SharePointOpts,
) {
	lp, ln, lwu := len(opts.LibraryPaths), len(opts.LibraryItems), len(opts.WebURLs)

	// only use the inclusion if either a path or item name
	// is specified
	if lp+ln == 0 {
		return
	}

	if len(opts.Sites) == 0 {
		opts.Sites = selectors.Any()
	}

	if lp+ln+lwu == 0 {
		sel.Include(sel.Sites(opts.Sites))

		return
	}

	opts.LibraryPaths = trimFolderSlash(opts.LibraryPaths)

	if ln == 0 {
		opts.LibraryItems = selectors.Any()
	}

	containsURLs, suffixURLs := splitFoldersIntoContainsAndPrefix(opts.WebURLs)

	if len(containsURLs) > 0 {
		sel.Include(sel.WebURL(containsURLs))
	}

	if len(suffixURLs) > 0 {
		sel.Include(sel.WebURL(suffixURLs, selectors.SuffixMatch()))
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
