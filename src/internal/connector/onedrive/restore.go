package onedrive

import (
	"context"
	"io"
	"runtime/trace"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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
func RestoreCollections(
	ctx context.Context,
	service graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.Collection,
	deets *details.Details,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		restoreErrors  error
	)

	errUpdater := func(id string, err error) {
		restoreErrors = support.WrapAndAppend(id, err, restoreErrors)
	}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		temp, canceled := RestoreCollection(ctx, service, dc, OneDriveSource, dest.ContainerName, deets, errUpdater)

		restoreMetrics.Combine(temp)

		if canceled {
			break
		}
	}

	return support.CreateStatus(
			ctx,
			support.Restore,
			len(dcs),
			restoreMetrics,
			restoreErrors,
			dest.ContainerName),
		nil
}

// RestoreCollection handles restoration of an individual collection.
// returns:
// - the collection's item and byte count metrics
// - the context cancellation state (true if the context is cancelled)
func RestoreCollection(
	ctx context.Context,
	service graph.Servicer,
	dc data.Collection,
	source driveSource,
	restoreContainerName string,
	deets *details.Details,
	errUpdater func(string, error),
) (support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics    = support.CollectionMetrics{}
		copyBuffer = make([]byte, copyBufferSize)
		directory  = dc.FullPath()
	)

	drivePath, err := toOneDrivePath(directory)
	if err != nil {
		errUpdater(directory.String(), err)
		return metrics, false
	}

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

	restoreFolderElements := []string{restoreContainerName}
	restoreFolderElements = append(restoreFolderElements, drivePath.folders...)

	trace.Log(ctx, "gc:oneDrive:restoreCollection", directory.String())
	logger.Ctx(ctx).Debugf("Restore target for %s is %v", dc.FullPath(), restoreFolderElements)

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := createRestoreFolders(ctx, service, drivePath.driveID, restoreFolderElements)
	if err != nil {
		errUpdater(directory.String(), errors.Wrapf(err, "failed to create folders %v", restoreFolderElements))
		return metrics, false
	}

	// Restore items from the collection
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return metrics, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, false
			}
			metrics.Objects++

			metrics.TotalBytes += int64(len(copyBuffer))

			itemInfo, err := restoreItem(ctx,
				service,
				itemData,
				drivePath.driveID,
				restoreFolderID,
				copyBuffer,
				source)
			if err != nil {
				errUpdater(itemData.UUID(), err)
				continue
			}

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				logger.Ctx(ctx).DPanicw("transforming item to full path", "error", err)
				errUpdater(itemData.UUID(), err)

				continue
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				itemInfo)

			metrics.Successes++
		}
	}
}

// createRestoreFolders creates the restore folder hieararchy in the specified drive and returns the folder ID
// of the last folder entry in the hiearchy
func createRestoreFolders(ctx context.Context, service graph.Servicer, driveID string, restoreFolders []string,
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

		if !errors.Is(err, errFolderNotFound) {
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
func restoreItem(
	ctx context.Context,
	service graph.Servicer,
	itemData data.Stream,
	driveID, parentFolderID string,
	copyBuffer []byte,
	source driveSource,
) (details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreItem", D.Label("item_uuid", itemData.UUID()))
	defer end()

	itemName := itemData.UUID()
	trace.Log(ctx, "gc:oneDrive:restoreItem", itemName)

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.StreamSize)
	if !ok {
		return details.ItemInfo{}, errors.Errorf("item %q does not implement DataStreamInfo", itemName)
	}

	// Create Item
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(itemData.UUID(), false))
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to create item %s", itemName)
	}

	// Get a drive item writer
	w, err := driveItemWriter(ctx, service, driveID, *newItem.GetId(), ss.Size())
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to create item upload session %s", itemName)
	}

	iReader := itemData.ToReader()
	progReader, closer := observe.ItemProgress(iReader, observe.ItemRestoreMsg, itemName, ss.Size())

	go closer()

	// Upload the stream data
	written, err := io.CopyBuffer(w, progReader, copyBuffer)
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to upload data: item %s", itemName)
	}

	dii := details.ItemInfo{}

	switch source {
	case SharePointSource:
		dii.SharePoint = sharePointItemInfo(newItem, written)
	default:
		dii.OneDrive = oneDriveItemInfo(newItem, written)
	}

	return dii, nil
}
