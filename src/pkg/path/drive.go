package path

import (
	"github.com/alcionai/clues"
)

// TODO: Move this into m365/collection/drive
// drivePath is used to represent path components
// of an item within the drive i.e.
// Given `drives/b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA/root:/Folder1/Folder2/file`
//
// driveID is `b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA` and
// folders[] is []{"Folder1", "Folder2"}
//
// Should be compatible with all drive-based services (ex: oneDrive, sharePoint Libraries, etc)
type DrivePath struct {
	DriveID string
	Root    string
	Folders Elements
}

func ToDrivePath(p Path) (*DrivePath, error) {
	folders := p.Folders()

	// Must be at least `drives/<driveID>/root:`
	if len(folders) < 3 {
		return nil, clues.
			New("folder path doesn't match expected format for Drive items").
			With("path_folders", p.Folder(false))
	}

	// FIXME(meain): Don't have any service specific code within this
	// function. Change this to either accept only the fragment of the
	// path that is the drive path or have a separate function for each
	// service.
	if p.Service() == GroupsService {
		// Groups have an extra /sites/<siteID> in the path
		return &DrivePath{DriveID: folders[3], Root: folders[4], Folders: folders[5:]}, nil
	}

	return &DrivePath{DriveID: folders[1], Root: folders[2], Folders: folders[3:]}, nil
}

// Returns the path to the folder within the drive (i.e. under `root:`)
func GetDriveFolderPath(p Path) (*Builder, error) {
	drivePath, err := ToDrivePath(p)
	if err != nil {
		return nil, err
	}

	return Builder{}.Append(drivePath.Folders...), nil
}

// BuildDriveLocation takes a driveID and a set of unescaped element names,
// including the root folder, and returns a *path.Builder containing the
// canonical path representation for the drive path.
func BuildDriveLocation(
	driveID string,
	unescapedElements ...string,
) *Builder {
	return Builder{}.Append("drives", driveID).Append(unescapedElements...)
}

// BuildGroupsDriveLocation is same as BuildDriveLocation, but for
// group drives and thus includes siteID.
func BuildGroupsDriveLocation(
	siteID string,
	driveID string,
	unescapedElements ...string,
) *Builder {
	return Builder{}.Append("sites", siteID, "drives", driveID).Append(unescapedElements...)
}
