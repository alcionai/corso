package testdata

import (
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	TestChannelName = "Test"
	TestChatTopic   = "Test"
)

// GroupsBackupFolderScope is the standard folder scope that should be used
// in integration backups with groups when interacting with libraries.
func GroupsBackupLibraryFolderScope(sel *selectors.GroupsBackup) []selectors.GroupsScope {
	return sel.LibraryFolders([]string{TestFolderName}, selectors.PrefixMatch())
}

// GroupsBackupChannelScope is the standard folder scope that should be used
// in integration backups with groups when interacting with channels.
func GroupsBackupChannelScope(sel *selectors.GroupsBackup) []selectors.GroupsScope {
	return sel.Channels([]string{TestChannelName})
}

// GroupsBackupConversationScope is the standard folder scope that should be used
// in integration backups with groups when interacting with conversations.
func GroupsBackupConversationScope(sel *selectors.GroupsBackup) []selectors.GroupsScope {
	// there's no way to easily specify a test conversation by name.
	return sel.Conversation(selectors.Any())
}

// TeamsChatsBackupChatScope is the standard folder scope that should be used
// in integration backups with teams chats when interacting with chats.
func TeamsChatsBackupChatScope(sel *selectors.TeamsChatsBackup) []selectors.TeamsChatsScope {
	return sel.Chats([]string{TestChatTopic})
}
