package version

import "math"

const (
	// Restore1DataAndMetaFiles is the corso backup format version
	// in which we split from storing just the data to storing both
	// the data and metadata in two files.
	Restore1DataAndMetaFiles = 1
	// Restore2IsMetaMarker is a small improvement on
	// VersionWithDataAndMetaFiles, but has a marker IsMeta which
	// specifies if the file is a meta file or a data file.
	Restore2IsMetaMarker = 3
	// RestoreCurrNameInMeta points to the backup format version where we begin
	// storing files in kopia with their item ID instead of their OneDrive file
	// name.
	// TODO(ashmrtn): Update this to a real value when we merge the file name
	// change. Set to MAXINT for now to keep the if-check using it working.
	RestoreCurrNameInMeta = math.MaxInt
)
