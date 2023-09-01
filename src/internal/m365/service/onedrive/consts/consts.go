package onedrive

import "github.com/alcionai/corso/src/pkg/path"

const (
	SitesPathDir = "sites"
	// const used as the root dir for the drive portion of a path prefix.
	// eg: tid/onedrive/ro/files/drives/driveid/...
	DrivesPathDir = "drives"
	// const used as the root-of-drive dir for the drive portion of a path prefix.
	// eg: tid/onedrive/ro/files/drives/driveid/root:/...
	RootPathDir = "root:"
	// root id for drive items
	RootID = "root"
)

func DriveFolderPrefixBuilder(driveID string) *path.Builder {
	return path.Builder{}.Append(DrivesPathDir, driveID, RootPathDir)
}
