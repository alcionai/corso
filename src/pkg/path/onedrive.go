package path

import "github.com/alcionai/clues"

// drivePath is used to represent path components
// of an item within the drive i.e.
// Given `drives/b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA/root:/Folder1/Folder2/file`
//
// driveID is `b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA` and
// folders[] is []{"Folder1", "Folder2"}
type DrivePath struct {
	DriveID string
	Folders []string
}

func ToOneDrivePath(p Path) (*DrivePath, error) {
	folders := p.Folders()

	// Must be at least `drives/<driveID>/root:`
	if len(folders) < 3 {
		return nil, clues.
			New("folder path doesn't match expected format for OneDrive items").
			With("folders", p.Folder())
	}

	return &DrivePath{DriveID: folders[1], Folders: folders[3:]}, nil
}

// Returns the path to the folder within the drive (i.e. under `root:`)
func GetDriveFolderPath(p Path) (string, error) {
	drivePath, err := ToOneDrivePath(p)
	if err != nil {
		return "", err
	}

	return Builder{}.Append(drivePath.Folders...).String(), nil
}
