package testdata

import (
	"github.com/alcionai/canario/src/pkg/selectors"
)

const TestListName = "test-list"

// SharePointBackupFolderScope is the standard folder scope that should be used
// in integration backups with sharepoint.
func SharePointBackupFolderScope(sel *selectors.SharePointBackup) []selectors.SharePointScope {
	return sel.LibraryFolders([]string{TestFolderName}, selectors.PrefixMatch())
}

func SharePointBackupListsScope(sel *selectors.SharePointBackup) []selectors.SharePointScope {
	return sel.ListItems([]string{TestListName}, selectors.Any())
}
