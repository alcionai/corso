package testdata

import (
	"github.com/alcionai/corso/src/pkg/selectors"
)

const TestChannelName = "Test"

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
