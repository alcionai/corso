package testdata

import "github.com/alcionai/corso/src/pkg/selectors"

// BackupFolderScope is the standard folder scope that should be used
// in integration backups with sharepoint.
func BackupFolderScope(sel *selectors.SharePointBackup) []selectors.SharePointScope {
	return sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch())
}
