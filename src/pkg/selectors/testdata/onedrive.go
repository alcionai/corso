package testdata

import "github.com/alcionai/corso/src/pkg/selectors"

// OneDriveBackupFolderScope is the standard folder scope that should be used
// in integration backups with onedrive.
func OneDriveBackupFolderScope(sel *selectors.OneDriveBackup) []selectors.OneDriveScope {
	return sel.Folders([]string{"test"}, selectors.PrefixMatch())
}
