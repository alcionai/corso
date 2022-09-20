package onedrive

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// copyBufferSize is used for chunked upload
	// Microsoft recommends 5-10MB buffers
	// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession?view=graph-rest-1.0#best-practices
	copyBufferSize = 5 * 1024 * 1024
)

// drivePath is used to represent path components
// of an item within the drive i.e.
// Given `drives/b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA/root:/Folder1/Folder2/file`
//
// driveID is `b!X_8Z2zuXpkKkXZsr7gThk9oJpuj0yXVGnK5_VjRRPK-q725SX_8ZQJgFDK8PlFxA` and
// folders[] is []{"Folder1", "Folder2"}
type drivePath struct {
	driveID string
	folders []string
}

func toOneDrivePath(p path.Path) (*drivePath, error) {
	folders := p.Folders()

	// Must be at least `drives/<driveID>/root:`
	if len(folders) < 3 {
		return nil, errors.Errorf("folder path doesn't match expected format for OneDrive items: %s", p.Folder())
	}

	return &drivePath{driveID: folders[1], folders: folders[3:]}, nil
}

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(ctx context.Context, service graph.Service, dcs []data.Collection,
) (*support.ConnectorOperationStatus, error) {
	var (
		total, restored      int
		restoreErrors        error
		copyBuffer           = make([]byte, copyBufferSize)
		restoreContainerName = fmt.Sprintf("Corso_Restore_%s", common.FormatNow(common.SimpleDateTimeFormatOneDrive))
	)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {

		directory := dc.FullPath()

		drivePath, err := toOneDrivePath(directory)
		if err != nil {
			restoreErrors = support.WrapAndAppend(directory.String(), err, restoreErrors)
			continue
		}

		// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
		// from the backup under this the restore folder instead of root)
		// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

		restoreFolderElements := []string{restoreContainerName}

		restoreFolderElements = append(restoreFolderElements, drivePath.folders...)

		logger.Ctx(ctx).Debugf("Restore target for %s is %v", dc.FullPath(), restoreFolderElements)

		// Create restore folders and get the folder ID of the folder the data stream will be restored in
		restoreFolderID, err := createRestoreFolders(ctx, service, drivePath.driveID, restoreFolderElements)
		if err != nil {
			restoreErrors = support.WrapAndAppend(directory.String(), errors.Wrapf(err, "failed to create folders %v",
				restoreFolderElements), restoreErrors)
			continue
		}

		// Restore items from the collection
		exit := false
		items := dc.Items()
		for !exit {
			select {
			case <-ctx.Done():
				return nil, support.WrapAndAppend("context cancelled", ctx.Err(), restoreErrors)
			case itemData, ok := <-items:
				if !ok {
					exit = true
					break
				}
				total++

				err := restoreItem(ctx, service, itemData, drivePath.driveID, restoreFolderID, copyBuffer)
				if err != nil {
					restoreErrors = support.WrapAndAppend(itemData.UUID(), err, restoreErrors)
					continue
				}

				restored++
			}
		}
	}

	return support.CreateStatus(ctx, support.Restore, total, restored, 0, restoreErrors), nil
}

// createRestoreFolders creates the restore folder hieararchy in the specified drive and returns the folder ID
// of the last folder entry in the hiearchy
func createRestoreFolders(ctx context.Context, service graph.Service, driveID string, restoreFolders []string,
) (string, error) {
	driveRoot, err := service.Client().DrivesById(driveID).Root().Get(ctx, nil)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to get drive root. details: %s",
			support.ConnectorStackErrorTrace(err),
		)
	}
	logger.Ctx(ctx).Debugf("Found Root for Drive %s with ID %s", driveID, *driveRoot.GetId())

	parentFolderID := *driveRoot.GetId()
	for _, folder := range restoreFolders {
		folderItem, err := getFolder(ctx, service, driveID, parentFolderID, folder)
		if err == nil {
			parentFolderID = *folderItem.GetId()
			logger.Ctx(ctx).Debugf("Found %s with ID %s", folder, parentFolderID)
			continue
		}
		if err != errFolderNotFound {
			return "", errors.Wrapf(err, "folder %s not found in drive(%s) parentFolder(%s)", folder, driveID, parentFolderID)
		}

		folderItem, err = createItem(ctx, service, driveID, parentFolderID, newItem(folder, true))
		if err != nil {
			return "", errors.Wrapf(
				err,
				"failed to create folder %s/%s. details: %s", parentFolderID, folder,
				support.ConnectorStackErrorTrace(err),
			)
		}

		logger.Ctx(ctx).Debugf("Resolved %s in %s to %s", folder, parentFolderID, *folderItem.GetId())
		parentFolderID = *folderItem.GetId()
	}

	return parentFolderID, nil
}

// restoreItem will create a new item in the specified `parentFolderID` and upload the data.Stream
func restoreItem(ctx context.Context, service graph.Service, itemData data.Stream, driveID, parentFolderID string,
	copyBuffer []byte,
) error {
	itemName := itemData.UUID()

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.StreamSize)
	if !ok {
		return errors.Errorf("item %q does not implement DataStreamInfo", itemName)
	}

	// Create Item
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(itemData.UUID(), false))
	if err != nil {
		return errors.Wrapf(err, "failed to create item %s", itemName)
	}

	// Get a drive item writer
	w, err := driveItemWriter(ctx, service, driveID, *newItem.GetId(), ss.Size())
	if err != nil {
		return errors.Wrapf(err, "failed to create item upload session %s", itemName)
	}

	// Upload the stream data
	_, err = io.CopyBuffer(w, itemData.ToReader(), copyBuffer)
	if err != nil {
		return errors.Wrapf(err, "failed to upload data: item %s", itemName)
	}

	return nil
}
