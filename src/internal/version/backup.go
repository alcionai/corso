package version

import "math"

const Backup = 5

// Various labels to refer to important version changes.
// Labels don't need 1:1 service:version representation.  Add a new
// label when it's important to mark a delta in behavior that's handled
// somewhere in the logic.
// Labels should state their application, the backup version number,
// and the colloquial purpose of the label.
const (
	// OneDrive1DataAndMetaFiles is the corso backup format version
	// in which we split from storing just the data to storing both
	// the data and metadata in two files.
	OneDrive1DataAndMetaFiles = 1

	// OneDrive3IsMetaMarker is a small improvement on
	// VersionWithDataAndMetaFiles, but has a marker IsMeta which
	// specifies if the file is a meta file or a data file.
	OneDrive3IsMetaMarker = 3

	// OneDrive4IncludesPermissions includes permissions for folders in the same
	// collection as the folder itself.
	OneDrive4DirIncludesPermissions = 4

	// OneDrive5DirMetaNoName changed the directory metadata file name from
	// <dirname>.dirmeta to just .dirmeta to avoid issues with folder renames
	// during incremental backups.
	OneDrive5DirMetaNoName = 5

	// OneDriveXNameInMeta points to the backup format version where we begin
	// storing files in kopia with their item ID instead of their OneDrive file
	// name.
	// TODO(ashmrtn): Update this to a real value when we merge the file name
	// change. Set to MAXINT for now to keep the if-check using it working.
	OneDriveXNameInMeta = math.MaxInt
)
