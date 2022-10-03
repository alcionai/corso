package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

type OneDriveOpts struct {
	Users          []string
	Names          []string
	Paths          []string
	CreatedAfter   string
	CreatedBefore  string
	ModifiedAfter  string
	ModifiedBefore string
}

// ValidateOneDriveRestoreFlags checks common flags for correctness and interdependencies
func ValidateOneDriveRestoreFlags(backupID string) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
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
func IncludeOneDriveRestoreDataSelectors(
	sel *selectors.OneDriveRestore,
	opts OneDriveOpts,
) {
	if len(opts.Users) == 0 {
		opts.Users = selectors.Any()
	}

	lp, ln := len(opts.Paths), len(opts.Names)

	// either scope the request to a set of users
	if lp+ln == 0 {
		sel.Include(sel.Users(opts.Users))

		return
	}

	if lp == 0 {
		opts.Paths = selectors.Any()
	}

	if ln == 0 {
		opts.Names = selectors.Any()
	}

	sel.Include(sel.Items(opts.Users, opts.Paths, opts.Names))
}

// FilterOneDriveRestoreInfoSelectors builds the common info-selector filters.
func FilterOneDriveRestoreInfoSelectors(
	sel *selectors.OneDriveRestore,
	opts OneDriveOpts,
) {
	AddOneDriveFilter(sel, opts.CreatedAfter, sel.CreatedAfter)
	AddOneDriveFilter(sel, opts.CreatedBefore, sel.CreatedBefore)
	AddOneDriveFilter(sel, opts.ModifiedAfter, sel.ModifiedAfter)
	AddOneDriveFilter(sel, opts.ModifiedBefore, sel.ModifiedBefore)
}
