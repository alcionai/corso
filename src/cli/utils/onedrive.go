package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	FileFN               = "file"
	FolderFN             = "folder"
	NamesFN              = "name"
	PathsFN              = "path"
	FileCreatedAfterFN   = "file-created-after"
	FileCreatedBeforeFN  = "file-created-before"
	FileModifiedAfterFN  = "file-modified-after"
	FileModifiedBeforeFN = "file-modified-before"
)

type OneDriveOpts struct {
	Users              []string
	Names              []string
	Paths              []string
	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Populated PopulatedFlags
}

// ValidateOneDriveRestoreFlags checks common flags for correctness and interdependencies
func ValidateOneDriveRestoreFlags(backupID string, opts OneDriveOpts) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		return errors.New("invalid time format for created-after")
	}

	if _, ok := opts.Populated[FileCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.FileCreatedBefore) {
		return errors.New("invalid time format for created-before")
	}

	if _, ok := opts.Populated[FileModifiedAfterFN]; ok && !IsValidTimeFormat(opts.FileModifiedAfter) {
		return errors.New("invalid time format for modified-after")
	}

	if _, ok := opts.Populated[FileModifiedBeforeFN]; ok && !IsValidTimeFormat(opts.FileModifiedBefore) {
		return errors.New("invalid time format for modified-before")
	}

	return nil
}

// AddOneDriveFilter adds the scope of the provided values to the selector's
// filter set
func AddOneDriveFilter(
	sel *selectors.OneDriveRestore,
	v string,
	f func(string) []selectors.OneDriveScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// IncludeOneDriveRestoreDataSelectors builds the common data-selector
// inclusions for OneDrive commands.
func IncludeOneDriveRestoreDataSelectors(opts OneDriveOpts) *selectors.OneDriveRestore {
	users := opts.Users
	if len(users) == 0 {
		users = selectors.Any()
	}

	sel := selectors.NewOneDriveRestore(users)

	lp, ln := len(opts.Paths), len(opts.Names)

	// only use the inclusion if either a path or item name
	// is specified
	if lp+ln == 0 {
		sel.Include(sel.Users(opts.Users))

		return sel
	}

	opts.Paths = trimFolderSlash(opts.Paths)

	if ln == 0 {
		opts.Names = selectors.Any()
	}

	containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.Paths)

	if len(containsFolders) > 0 {
		sel.Include(sel.Items(users, containsFolders, opts.Names))
	}

	if len(prefixFolders) > 0 {
		sel.Include(sel.Items(users, prefixFolders, opts.Names, selectors.PrefixMatch()))
	}

	return sel
}

// FilterOneDriveRestoreInfoSelectors builds the common info-selector filters.
func FilterOneDriveRestoreInfoSelectors(
	sel *selectors.OneDriveRestore,
	opts OneDriveOpts,
) {
	AddOneDriveFilter(sel, opts.FileCreatedAfter, sel.CreatedAfter)
	AddOneDriveFilter(sel, opts.FileCreatedBefore, sel.CreatedBefore)
	AddOneDriveFilter(sel, opts.FileModifiedAfter, sel.ModifiedAfter)
	AddOneDriveFilter(sel, opts.FileModifiedBefore, sel.ModifiedBefore)
}
