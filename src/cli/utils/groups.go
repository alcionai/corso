package utils

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GroupsOpts struct {
	Groups        []string
	Channels      []string
	Messages      []string
	Conversations []string
	Posts         []string

	MessageCreatedAfter    string
	MessageCreatedBefore   string
	MessageLastReplyAfter  string
	MessageLastReplyBefore string

	SiteID             []string
	WebURL             []string
	Library            string
	FileName           []string // for libraries, to duplicate onedrive interface
	FolderPath         []string // for libraries, to duplicate onedrive interface
	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Lists []string

	PageFolder []string
	Page       []string

	RestoreCfg RestoreCfgOpts
	ExportCfg  ExportCfgOpts

	Populated flags.PopulatedFlags
}

func GroupsAllowedCategories() map[string]struct{} {
	return map[string]struct{}{
		flags.DataLibraries:     {},
		flags.DataMessages:      {},
		flags.DataConversations: {},
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
		case flags.DataConversations:
			sel.Include(sel.ConversationPosts(selectors.Any(), selectors.Any()))
		}
	}

	return sel
}

func MakeGroupsOpts(cmd *cobra.Command) GroupsOpts {
	return GroupsOpts{
		Groups:        flags.GroupFV,
		Channels:      flags.ChannelFV,
		Messages:      flags.MessageFV,
		Conversations: flags.ConversationFV,
		Posts:         flags.PostFV,
		WebURL:        flags.WebURLFV,
		SiteID:        flags.SiteIDFV,

		Library:                flags.LibraryFV,
		FileName:               flags.FileNameFV,
		FolderPath:             flags.FolderPathFV,
		FileCreatedAfter:       flags.FileCreatedAfterFV,
		FileCreatedBefore:      flags.FileCreatedBeforeFV,
		FileModifiedAfter:      flags.FileModifiedAfterFV,
		FileModifiedBefore:     flags.FileModifiedBeforeFV,
		MessageCreatedAfter:    flags.MessageCreatedAfterFV,
		MessageCreatedBefore:   flags.MessageCreatedBeforeFV,
		MessageLastReplyAfter:  flags.MessageLastReplyAfterFV,
		MessageLastReplyBefore: flags.MessageLastReplyBeforeFV,

		Lists: flags.ListFV,

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
func ValidateGroupsRestoreFlags(backupID string, opts GroupsOpts, isRestore bool) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	// The user has to explicitly specify which resource to restore. In
	// this case, since we can only restore sites, the user is supposed
	// to specify which site to restore.
	if isRestore {
		if len(opts.WebURL)+len(opts.SiteID) == 0 {
			return clues.New("web URL of the site to restore is required. Use --" + flags.SiteFN + " to provide one.")
		} else if len(opts.WebURL)+len(opts.SiteID) > 1 {
			return clues.New("only a single site can be selected for restore")
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

	if _, ok := opts.Populated[flags.MessageCreatedAfterFN]; ok && !IsValidTimeFormat(opts.MessageCreatedAfter) {
		return clues.New("invalid time format for " + flags.MessageCreatedAfterFN)
	}

	if _, ok := opts.Populated[flags.MessageCreatedBeforeFN]; ok && !IsValidTimeFormat(opts.MessageCreatedBefore) {
		return clues.New("invalid time format for " + flags.MessageCreatedBeforeFN)
	}

	if _, ok := opts.Populated[flags.MessageLastReplyAfterFN]; ok && !IsValidTimeFormat(opts.MessageLastReplyAfter) {
		return clues.New("invalid time format for " + flags.MessageLastReplyAfterFN)
	}

	if _, ok := opts.Populated[flags.MessageLastReplyBeforeFN]; ok && !IsValidTimeFormat(opts.MessageLastReplyBefore) {
		return clues.New("invalid time format for " + flags.MessageLastReplyBeforeFN)
	}

	return nil
}

// AddGroupsFilter adds the scope of the provided values to the selector's
// filter set
func AddGroupsFilter(
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
	var (
		groups                 = opts.Groups
		folderPaths, fileNames = len(opts.FolderPath), len(opts.FileName)
		lists                  = len(opts.Lists)
		pageFolders, pageItems = len(opts.PageFolder), len(opts.Page)
		chans, chanMsgs        = len(opts.Channels), len(opts.Messages)
		convs, convPosts       = len(opts.Conversations), len(opts.Posts)
	)

	if len(opts.Groups) == 0 {
		groups = selectors.Any()
	}

	sel := selectors.NewGroupsRestore(groups)

	if folderPaths+fileNames+
		lists+
		pageFolders+pageItems+
		chans+chanMsgs+
		convs+convPosts == 0 {
		sel.Include(sel.AllData())
		return sel
	}

	// sharepoint site selectors

	if folderPaths+fileNames+
		lists+
		pageFolders+pageItems > 0 {
		if folderPaths+fileNames > 0 {
			if fileNames == 0 {
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

		configureSharepointListsSelector(sel, opts.Lists)

		if pageFolders+pageItems > 0 {
			if pageItems == 0 {
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
	}

	// channel and message selectors

	if chans+chanMsgs > 0 {
		// if no channel is specified, include all channels
		if chans == 0 {
			opts.Channels = selectors.Any()
		}

		// if no message is specified, only select channels
		// otherwise, look for channel/message pairs
		if chanMsgs == 0 {
			sel.Include(sel.Channels(opts.Channels))
		} else {
			sel.Include(sel.ChannelMessages(opts.Channels, opts.Messages))
		}
	}

	// conversation and post selectors

	if convs+convPosts > 0 {
		// if no conversation is specified, include all conversations
		if convs == 0 {
			opts.Conversations = selectors.Any()
		}

		// if no post is specified, only select conversations;
		// otherwise, look for channel/message pairs
		if chanMsgs == 0 {
			sel.Include(sel.Conversation(opts.Conversations))
		} else {
			sel.Include(sel.ConversationPosts(opts.Conversations, opts.Posts))
		}
	}

	return sel
}

// FilterGroupsRestoreInfoSelectors builds the common info-selector filters.
func FilterGroupsRestoreInfoSelectors(
	sel *selectors.GroupsRestore,
	opts GroupsOpts,
) {
	var site string

	// It is possible that neither exists as this filter is only
	// mandatory for restore and not for export.
	if len(opts.SiteID) > 0 {
		site = opts.SiteID[0]
	} else if len(opts.WebURL) > 0 {
		site = opts.WebURL[0]
	}

	// sel.Site can accept both ID and URL and so irrespective of
	// which flag the user uses, we can process both weburl and siteid
	// Also since the url will start with `https://` in the data that
	// we store and the id is a uuid, we can grantee that there will be
	// no collisions.
	AddGroupsFilter(sel, site, sel.Site)

	AddGroupsFilter(sel, opts.Library, sel.Library)
	AddGroupsFilter(sel, opts.FileCreatedAfter, sel.CreatedAfter)
	AddGroupsFilter(sel, opts.FileCreatedBefore, sel.CreatedBefore)
	AddGroupsFilter(sel, opts.FileModifiedAfter, sel.ModifiedAfter)
	AddGroupsFilter(sel, opts.FileModifiedBefore, sel.ModifiedBefore)
	AddGroupsFilter(sel, opts.MessageCreatedAfter, sel.MessageCreatedAfter)
	AddGroupsFilter(sel, opts.MessageCreatedBefore, sel.MessageCreatedBefore)
	AddGroupsFilter(sel, opts.MessageLastReplyAfter, sel.MessageLastReplyAfter)
	AddGroupsFilter(sel, opts.MessageLastReplyBefore, sel.MessageLastReplyBefore)
}
