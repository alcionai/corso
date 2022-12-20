package onedrive

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"runtime/trace"
	"strings"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

const (
	// copyBufferSize is used for chunked upload
	// Microsoft recommends 5-10MB buffers
	// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession?view=graph-rest-1.0#best-practices
	copyBufferSize = 5 * 1024 * 1024
)

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	service graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.Collection,
	deets *details.Builder,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		restoreErrors  error
	)

	errUpdater := func(id string, err error) {
		restoreErrors = support.WrapAndAppend(id, err, restoreErrors)
	}

	if len(dcs) == 0 {
		return nil, errors.New("no data available to restore")
	}

	directory := dcs[0].FullPath() // drive will be same for all the paths
	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		errUpdater(directory.String(), err)
	}

	driveID := drivePath.DriveID
	driveRoot, err := service.Client().DrivesById(driveID).Root().Get(ctx, nil)
	if err != nil {
		errUpdater(directory.String(), errors.Wrapf(err,
			"failed to get drive root. details: %s",
			support.ConnectorStackErrorTrace(err),
		))
		return nil, errors.New("unable to connect to drive")
	}

	logger.Ctx(ctx).Debugf("Found Root for Drive %s with ID %s", driveID, *driveRoot.GetId())

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

	pids := map[string]string{}

	// Create initial root container to hold the restore
	containerId, err := CreateRestoreFolder(ctx, service, driveID, dest.ContainerName, nil, *driveRoot.GetId())
	if err != nil {
		return nil, errors.New("unable to create restore container")
	}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		elms := dc.FullPath().Elements()
		path := strings.Join(elms, "/")
		parentPath := strings.Join(elms[:len(elms)-1], "/")

		pid, ok := pids[parentPath]
		if !ok {
			if len(dc.FullPath().Elements()) == 8 { // base folders
				pid = containerId
			} else {
				errUpdater(dc.FullPath().String(), errors.New("unable to find parent path"))
				continue 
			}
		}

		id, metrics, cancelled := RestoreCollection(ctx, service, dc, OneDriveSource, deets, errUpdater, pid)
		restoreMetrics.Combine(metrics)
		if cancelled {
			break
		}
		pids[path] = id
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
	deets *details.Builder,
	errUpdater func(string, error),
	parentFolderID string,
) (string, support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics    = support.CollectionMetrics{}
		copyBuffer = make([]byte, copyBufferSize)
		directory  = dc.FullPath()
	)

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		return "", metrics, false
	}

	elms := directory.Elements()
	folder := elms[len(elms)-1]
	parentFolderID, err = CreateRestoreFolder(ctx, service, drivePath.DriveID, folder, dc.Meta(), parentFolderID)
	if err != nil {
		errUpdater(folder, errors.Wrapf(err, "failed to create folder %v", folder))
		return "", metrics, false
	}

	// Restore items from the collection
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return parentFolderID, metrics, true

		case itemData, ok := <-items:
			if !ok {
				return parentFolderID, metrics, false
			}
			metrics.Objects++

			metrics.TotalBytes += int64(len(copyBuffer))

			itemInfo, err := restoreItem(ctx,
				service,
				itemData,
				drivePath.DriveID,
				parentFolderID,
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
				true,
				itemInfo)

			metrics.Successes++
		}
	}
}

func CreateRestoreFolder(ctx context.Context, service graph.Servicer, driveID, folder string, metar io.ReadCloser, parentFolderID string) (string, error) {
	var meta ItemMeta
	// `metar` will be nil for the top level container folder
	// TODO(meain): Should we make `metar` into ItemMeta instead of io.ReadCloser?
	if metar != nil {
		metaraw, err := ioutil.ReadAll(metar)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal(metaraw[4:], &meta) // TODO(meain): First 4 bytes are somehow 00 00 00 01
		if err != nil {
			return "", err
		}
	}

	folderItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(folder, true), meta)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to create folder %s/%s. details: %s", parentFolderID, folder,
			support.ConnectorStackErrorTrace(err),
		)
	}

	logger.Ctx(ctx).Debugf("Resolved %s in %s to %s", folder, parentFolderID, *folderItem.GetId())

	return *folderItem.GetId(), nil
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

	imReader, err := itemData.ToMetaReader()
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to fetch metadata: item %s", itemName)
	}

	rmeta, err := ioutil.ReadAll(imReader)
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to read metadata: item %s", itemName)
	}

	var meta ItemMeta
	err = json.Unmarshal(rmeta, &meta)
	if err != nil {
		return details.ItemInfo{}, errors.Wrapf(err, "failed to parse metadata: item %s", itemName)
	}

	// Create Item
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(itemData.UUID(), false), meta)
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
