package utils

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GroupsOpts struct {
	Groups []string

	SiteID             []string
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

	RestoreCfg RestoreCfgOpts
	ExportCfg  ExportCfgOpts

	Populated flags.PopulatedFlags
}

func GroupsAllowedCategories() map[string]struct{} {
	return map[string]struct{}{
		flags.DataLibraries: {},
		flags.DataMessages:  {},
	}
}

func AddGroupsCategories(sel *selectors.GroupsBackup, cats []string) *selectors.GroupsBackup {
	if len(cats) == 0 {
		sel.Include(sel.AllData())
	}

	for _, d := range cats {
		switch d {
		case flags.DataLibraries:
			sel.Include(sel.LibraryFolders(selectors.Any()))
		case flags.DataMessages:
			sel.Include(sel.ChannelMessages(selectors.Any(), selectors.Any()))
		}
	}

	return sel
}

func MakeGroupsOpts(cmd *cobra.Command) GroupsOpts {
	return GroupsOpts{
		Groups: flags.GroupFV,

		SiteID: flags.SiteIDFV,

		Library:            flags.LibraryFV,
		FileName:           flags.FileNameFV,
		FolderPath:         flags.FolderPathFV,
		FileCreatedAfter:   flags.FileCreatedAfterFV,
		FileCreatedBefore:  flags.FileCreatedBeforeFV,
		FileModifiedAfter:  flags.FileModifiedAfterFV,
		FileModifiedBefore: flags.FileModifiedBeforeFV,

		ListFolder: flags.ListFolderFV,
		ListItem:   flags.ListItemFV,

		Page:       flags.PageFV,
		PageFolder: flags.PageFolderFV,

		RestoreCfg: makeRestoreCfgOpts(cmd),
		ExportCfg:  makeExportCfgOpts(cmd),

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// ValidateGroupsRestoreFlags checks common flags for correctness and interdependencies
func ValidateGroupsRestoreFlags(backupID string, opts GroupsOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	if _, ok := opts.Populated[flags.FileCreatedAfterFN]; ok && !IsValidTimeFormat(opts.FileCreatedAfter) {
		return clues.New("invalid time format for " + flags.FileCreatedAfterFN)
	}

	if _, ok := opts.Populated[flags.FileCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.FileCreatedBefore) {
		return clues.New("invalid time format for " + flags.FileCreatedBeforeFN)
	}

	if _, ok := opts.Populated[flags.FileModifiedAfterFN]; ok && !IsValidTimeFormat(opts.FileModifiedAfter) {
		return clues.New("invalid time format for " + flags.FileModifiedAfterFN)
	}

	if _, ok := opts.Populated[flags.FileModifiedBeforeFN]; ok && !IsValidTimeFormat(opts.FileModifiedBefore) {
		return clues.New("invalid time format for " + flags.FileModifiedBeforeFN)
	}

	return validateRestoreConfigFlags(flags.CollisionsFV, opts.RestoreCfg)
}

// AddGroupInfo adds the scope of the provided values to the selector's
// filter set
func AddGroupInfo(
	sel *selectors.GroupsRestore,
	v string,
	f func(string) []selectors.GroupsScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// IncludeGroupsRestoreDataSelectors builds the common data-selector
// inclusions for Group commands.
func IncludeGroupsRestoreDataSelectors(ctx context.Context, opts GroupsOpts) *selectors.GroupsRestore {
	groups := opts.Groups

	lg := len(opts.Groups)

	// TODO(meain): handle sites once we add non-root site backup
	// ls := len(opts.SiteID)

	lfp, lfn := len(opts.FolderPath), len(opts.FileName)
	slp, sli := len(opts.ListFolder), len(opts.ListItem)
	pf, pi := len(opts.PageFolder), len(opts.Page)

	if lg == 0 {
		groups = selectors.Any()
	}

	sel := selectors.NewGroupsRestore(groups)

	if lfp+lfn+slp+sli+pf+pi == 0 {
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

	return sel
}

// FilterGroupsRestoreInfoSelectors builds the common info-selector filters.
func FilterGroupsRestoreInfoSelectors(
	sel *selectors.GroupsRestore,
	opts GroupsOpts,
) {
	AddGroupInfo(sel, opts.Library, sel.Library)
	AddGroupInfo(sel, opts.FileCreatedAfter, sel.CreatedAfter)
	AddGroupInfo(sel, opts.FileCreatedBefore, sel.CreatedBefore)
	AddGroupInfo(sel, opts.FileModifiedAfter, sel.ModifiedAfter)
	AddGroupInfo(sel, opts.FileModifiedBefore, sel.ModifiedBefore)
}
