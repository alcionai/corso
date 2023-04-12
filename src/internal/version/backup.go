package version

const Backup = 6

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

	// Version 2 switched Exchange calendars from using folder display names to
	// folder IDs in their RepoRef.

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

	// OneDrive6NameInMeta points to the backup format version where we begin
	// storing files in kopia with their item ID instead of their OneDrive file
	// name.
	OneDrive6NameInMeta = 6

	// OneDriveXLocationRef provides LocationRef information for Exchange,
	// OneDrive, and SharePoint libraries.
	OneDriveXLocationRef = Backup + 1
)
