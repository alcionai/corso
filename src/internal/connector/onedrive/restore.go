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

// dirEntry is used to keep track of the data that we will need
// for subsequent folder creations.
type dirEntry struct {
	ID   string
	meta ItemMeta
}

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

	pids := map[string]dirEntry{}

	// Create initial root container to hold the restore
	containerId, err := CreateRestoreFolder(ctx, service, driveID, dest.ContainerName, []UserPermission{}, []UserPermission{}, *driveRoot.GetId())
	if err != nil {
		return nil, errors.New("unable to create restore container")
	}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		elms := dc.FullPath().Elements()
		path := strings.Join(elms, "/")
		parentPath := strings.Join(elms[:len(elms)-1], "/")

		parent, ok := pids[parentPath]
		if !ok {
			if len(dc.FullPath().Elements()) == 8 { // base folders
				parent = dirEntry{ID: containerId, meta: ItemMeta{}}
			} else {
				errUpdater(dc.FullPath().String(), errors.New("unable to find parent path"))
				continue
			}
		}

		id, pmeta, metrics, cancelled := RestoreCollection(
			ctx, service, dc,
			OneDriveSource, deets,
			errUpdater,
			parent)
		restoreMetrics.Combine(metrics)
		if cancelled {
			break
		}
		pids[path] = dirEntry{ID: id, meta: pmeta}
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
	parent dirEntry,
) (string, ItemMeta, support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics    = support.CollectionMetrics{}
		copyBuffer = make([]byte, copyBufferSize)
		directory  = dc.FullPath()
	)

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		return "", ItemMeta{}, metrics, false
	}

	elms := directory.Elements()
	folder := elms[len(elms)-1]
	meta, err := getItemMeta(dc.Meta())
	if err != nil {
		errUpdater(folder, errors.Wrapf(err, "failed to parse metadata %v", folder))
		return "", ItemMeta{}, metrics, false
	}

	permAdded, permRemoved := getChildPermissions(meta.Permissions, parent.meta.Permissions)
	parentFolderID, err := CreateRestoreFolder(ctx, service, drivePath.DriveID, folder, permAdded, permRemoved, parent.ID)
	if err != nil {
		errUpdater(folder, errors.Wrapf(err, "failed to create folder %v", folder))
		return "", ItemMeta{}, metrics, false
	}

	// Restore items from the collection
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return parentFolderID, meta, metrics, true

		case itemData, ok := <-items:
			if !ok {
				return parentFolderID, meta, metrics, false
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

func getItemMeta(metar io.ReadCloser) (ItemMeta, error) {
	var meta ItemMeta
	// `metar` will be nil for the top level container folder
	if metar != nil {
		metaraw, err := ioutil.ReadAll(metar)
		if err != nil {
			return ItemMeta{}, err
		}

		err = json.Unmarshal(metaraw[4:], &meta) // TODO(meain): First 4 bytes are somehow 00 00 00 01
		if err != nil {
			return ItemMeta{}, err
		}
	}

	return meta, nil
}

// getChildPermissions is to filter out permissions present in the
// parent from the ones that are available for child. This is
// necessary as we store the nested permissions in the child. We
// cannot avoid storing the nested permissions as it is possible that
// a file in a folder can remove the nested permission that is present
// on itself.
func getChildPermissions(childPermissions, parentPermissions []UserPermission) ([]UserPermission, []UserPermission) {
	addedPermissions := []UserPermission{}
	removedPermissions := []UserPermission{}

	for _, cp := range childPermissions {
		found := false
		for _, pp := range parentPermissions {
			if cp.ID == pp.ID {
				found = true
				break
			}
		}
		if !found {
			addedPermissions = append(addedPermissions, cp)
		}
	}

	for _, pp := range parentPermissions {
		found := false
		for _, cp := range childPermissions {
			if pp.ID == cp.ID {
				found = true
				break
			}
		}
		if !found {
			removedPermissions = append(removedPermissions, pp)
		}
	}

	return addedPermissions, removedPermissions
}

func CreateRestoreFolder(ctx context.Context, service graph.Servicer, driveID, folder string, permAdded, permRemoved []UserPermission, parentFolderID string) (string, error) {
	folderItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(folder, true), permAdded, permRemoved)
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
	// TODO(meain): handle file permissions
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(itemData.UUID(), false), []UserPermission{}, []UserPermission{})
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
