package testdata

import "github.com/alcionai/corso/src/pkg/selectors"

// SharePointBackupFolderScope is the standard folder scope that should be used
// in integration backups with sharepoint.
func SharePointBackupFolderScope(sel *selectors.SharePointBackup) []selectors.SharePointScope {
	return sel.LibraryFolders([]string{TestFolderName}, selectors.PrefixMatch())
}
