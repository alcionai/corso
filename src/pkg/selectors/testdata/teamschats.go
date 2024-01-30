package testdata

import "github.com/alcionai/corso/src/pkg/selectors"

// TeamsChatsBackupChatScope is the standard folder scope that should be used
// in integration backups with teams chats when interacting with chats.
func TeamsChatsBackupChatScope(sel *selectors.TeamsChatsBackup) []selectors.TeamsChatsScope {
	return sel.Chats([]string{TestChatTopic})
}
