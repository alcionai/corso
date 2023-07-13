package testdata

import (
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// SharePointBackupFolderScope is the standard folder scope that should be used
// in integration backups with sharepoint.
func SharePointBackupFolderScope(sel *selectors.SharePointBackup) []selectors.SharePointScope {
	return sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch())
}

// FilterSharePointRestoreLibraryScope adds filters to include only the two
// standard document libraries: Documents and 'More Documents'.
func SharePointRestoreStandardLibraryFilter(sel *selectors.SharePointRestore) {
	sel.Filter(
		sel.Library(tconfig.LibraryDocuments),
		sel.Library(tconfig.LibraryMoreDocuments))
}
