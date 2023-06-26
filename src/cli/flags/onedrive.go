package flags

import (
	"github.com/spf13/cobra"
)

const (
	FileFN   = "file"
	FolderFN = "folder"

	FileCreatedAfterFN   = "file-created-after"
	FileCreatedBeforeFN  = "file-created-before"
	FileModifiedAfterFN  = "file-modified-after"
	FileModifiedBeforeFN = "file-modified-before"
)

var (
	FolderPathFV []string
	FileNameFV   []string

	FileCreatedAfterFV   string
	FileCreatedBeforeFV  string
	FileModifiedAfterFV  string
	FileModifiedBeforeFV string
)

// AddOneDriveDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddOneDriveDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringSliceVar(
		&FolderPathFV,
		FolderFN, nil,
		"Select files by OneDrive folder; defaults to root.")

	fs.StringSliceVar(
		&FileNameFV,
		FileFN, nil,
		"Select files by name.")

	fs.StringVar(
		&FileCreatedAfterFV,
		FileCreatedAfterFN, "",
		"Select files created after this datetime.")
	fs.StringVar(
		&FileCreatedBeforeFV,
		FileCreatedBeforeFN, "",
		"Select files created before this datetime.")

	fs.StringVar(
		&FileModifiedAfterFV,
		FileModifiedAfterFN, "",
		"Select files modified after this datetime.")

	fs.StringVar(
		&FileModifiedBeforeFV,
		FileModifiedBeforeFN, "",
		"Select files modified before this datetime.")
}
