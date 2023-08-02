package utils

import (
	"context"
	"net/url"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointOpts struct {
	SiteID []string
	WebURL []string

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

func MakeSharePointOpts(cmd *cobra.Command) SharePointOpts {
	return SharePointOpts{
		SiteID: flags.SiteIDFV,
		WebURL: flags.WebURLFV,

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

// ValidateSharePointRestoreFlags checks common flags for correctness and interdependencies
func ValidateSharePointRestoreFlags(backupID string, opts SharePointOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	// ensure url can parse all weburls provided by --site.
	if _, ok := opts.Populated[flags.SiteFN]; ok {
		for _, wu := range opts.WebURL {
			if _, err := url.Parse(wu); err != nil {
				return clues.New("invalid site url: " + wu)
			}
		}
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
func IncludeSharePointRestoreDataSelectors(ctx context.Context, opts SharePointOpts) *selectors.SharePointRestore {
	sites := opts.SiteID

	lfp, lfn := len(opts.FolderPath), len(opts.FileName)
	ls, lwu := len(opts.SiteID), len(opts.WebURL)
	slp, sli := len(opts.ListFolder), len(opts.ListItem)
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

	if lwu > 0 {
		urls := make([]string, 0, len(opts.WebURL))

		for _, wu := range opts.WebURL {
			// for normalization, ensure the site has a https:// prefix.
			wu = strings.TrimPrefix(wu, "https://")
			wu = strings.TrimPrefix(wu, "http://")

			// don't add a prefix to path-only values
			if len(wu) > 0 && wu != "*" && !strings.HasPrefix(wu, "/") {
				wu = "https://" + wu
			}

			u, err := url.Parse(wu)
			if err != nil {
				// shouldn't be possible to err, if we called validation first.
				logger.Ctx(ctx).With("web_url", wu).Error("malformed web url")
				continue
			}

			urls = append(urls, u.String())
		}

		sel.Include(sel.WebURL(urls))
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
