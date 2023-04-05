package utils

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/selectors"
)

type OneDriveOpts struct {
	Users []string

	FileName           []string
	FolderPath         []string
	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Populated PopulatedFlags
}

func MakeOneDriveOpts(cmd *cobra.Command) OneDriveOpts {
	return OneDriveOpts{
		Users: UserFV,

		FileName:           FileNameFV,
		FolderPath:         FolderPathFV,
		FileCreatedAfter:   FileCreatedAfterFV,
		FileCreatedBefore:  FileCreatedBeforeFV,
		FileModifiedAfter:  FileModifiedAfterFV,
		FileModifiedBefore: FileModifiedBeforeFV,

		Populated: GetPopulatedFlags(cmd),
	}
}

// AddOneDriveDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddOneDriveDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringSliceVar(
		&FolderPathFV,
		FolderFN, nil,
		"Select files by OneDrive folder; defaults to root.")

	fs.StringSliceVar(
		&FileNameFV,
		FileFN, nil,
		"Select files by name.")

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
}

// ValidateOneDriveRestoreFlags checks common flags for correctness and interdependencies
func ValidateOneDriveRestoreFlags(backupID string, opts OneDriveOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	if _, ok := opts.Populated[FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		return clues.New("invalid time format for created-after")
	}

	if _, ok := opts.Populated[FileCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.FileCreatedBefore) {
		return clues.New("invalid time format for created-before")
	}

	if _, ok := opts.Populated[FileModifiedAfterFN]; ok && !IsValidTimeFormat(opts.FileModifiedAfter) {
		return clues.New("invalid time format for modified-after")
	}

	if _, ok := opts.Populated[FileModifiedBeforeFN]; ok && !IsValidTimeFormat(opts.FileModifiedBefore) {
		return clues.New("invalid time format for modified-before")
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

	lp, ln := len(opts.FolderPath), len(opts.FileName)

	// only use the inclusion if either a path or item name
	// is specified
	if lp+ln == 0 {
		sel.Include(sel.AllData())
		return sel
	}

	opts.FolderPath = trimFolderSlash(opts.FolderPath)

	if ln == 0 {
		opts.FileName = selectors.Any()
	}

	containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(opts.FolderPath)

	if len(containsFolders) > 0 {
		sel.Include(sel.Items(containsFolders, opts.FileName))
	}

	if len(prefixFolders) > 0 {
		sel.Include(sel.Items(prefixFolders, opts.FileName, selectors.PrefixMatch()))
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
