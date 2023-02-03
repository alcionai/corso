package path

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"
)

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
		return nil, errors.Errorf(
			"folder path doesn't match expected format for OneDrive items: %s",
			p.Folder(),
		)
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

func OneDriveResourcePath(itemName, itemPath string, isItem bool) (Path, error) {
	p, err := FromDataLayerPath(itemPath, isItem)
	if err != nil {
		return nil, clues.Wrap(err, "building OneDrive path")
	}

	return &oneDriveResourcePath{
		Path:     p,
		itemName: itemName,
	}, nil
}

type oneDriveResourcePath struct {
	Path
	itemName string
}

// ShortRef returns a hash of the contents of the path plus the itemName. This
// ensures the returned string is updated even if the path doesn't change but
// the item has been renamed in an external service.
func (p oneDriveResourcePath) ShortRef() string {
	return p.ToBuilder().Append(p.itemName).ShortRef()
}
