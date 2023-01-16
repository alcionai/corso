package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"runtime/trace"
	"sort"
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
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	// Reorder collections so that the parents directories are created
	// before the child directories
	sort.Slice(dcs, func(i, j int) bool {
		return dcs[i].FullPath().String() < dcs[j].FullPath().String()
	})

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		metrics, canceled := RestoreCollection(ctx, service, dc, OneDriveSource, dest.ContainerName, deets, errUpdater)

		restoreMetrics.Combine(metrics)

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
	deets *details.Builder,
	errUpdater func(string, error),
) (support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics    = support.CollectionMetrics{}
		copyBuffer = make([]byte, copyBufferSize)
		directory  = dc.FullPath()

		restoredIDs = map[string]string{}
		itemInfo    details.ItemInfo
		itemID      string
	)

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		errUpdater(directory.String(), err)
		return metrics, false
	}

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

	restoreFolderElements := []string{restoreContainerName}
	restoreFolderElements = append(restoreFolderElements, drivePath.Folders...)

	trace.Log(ctx, "gc:oneDrive:restoreCollection", directory.String())
	logger.Ctx(ctx).Infow(
		"restoring to destination",
		"origin", dc.FullPath().Folder(),
		"destination", restoreFolderElements)

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := CreateRestoreFolders(ctx, service, drivePath.DriveID, restoreFolderElements)
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
			fmt.Println("Restoring", itemData.UUID())

			// TODO(meain): refactor this if else into a separate func
			if source == OneDriveSource {
				name := itemData.UUID() // TODO(meain): Is this bound to change?
				if strings.HasSuffix(name, ".data") {
					metrics.Objects++
					metrics.TotalBytes += int64(len(copyBuffer))
					trimmedName := name[:len(name)-5]

					itemID, itemInfo, err = restoreItem(ctx,
						service,
						trimmedName,
						itemData,
						drivePath.DriveID,
						restoreFolderID,
						copyBuffer,
						source)
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					// fmt.Println("Adding", trimmedName, "as", itemID)
					restoredIDs[trimmedName] = itemID

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

				} else if strings.HasSuffix(name, ".meta") {
					parentPerm := []UserPermission{} // TODO(meain): get it from calling func

					meta, err := getMetadata(itemData.ToReader())
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}
					permAdded, permRemoved := getChildPermissions(meta.Permissions, parentPerm)

					trimmedName := name[:len(name)-5]
					restoreID, ok := restoredIDs[trimmedName]
					// fmt.Println("Got id of", trimmedName, "as", restoreID, "with", ok)
					if !ok {
						errUpdater(itemData.UUID(), fmt.Errorf("item not available to restore permissions"))
						continue
					}
					// fmt.Println("Restoring permissions for", trimmedName, permAdded)
					restorePermissions(
						ctx, service,
						drivePath.DriveID, restoreID,
						permAdded, permRemoved,
					)
					metrics.Successes++
				} else if strings.HasSuffix(name, ".dirmeta") {
					trimmedName := name[:len(name)-8]
					parentPerm := []UserPermission{} // TODO(meain): get it from calling func

					meta, err := getMetadata(itemData.ToReader())
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}
					permAdded, permRemoved := getChildPermissions(meta.Permissions, parentPerm)
					createRestoreFolder(ctx, service, drivePath.DriveID, trimmedName, restoreFolderID, permAdded, permRemoved)
				}
			} else {
				metrics.Objects++
				metrics.TotalBytes += int64(len(copyBuffer))

				// No permissions stored at the moment for SharePoint
				_, itemInfo, err = restoreItem(ctx,
					service,
					itemData.UUID(),
					itemData,
					drivePath.DriveID,
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
					true,
					itemInfo)
			}

		}
	}
}

// Creates a folder with its permissions
func createRestoreFolder(ctx context.Context,
	service graph.Servicer,
	driveID, folder, parentFolderID string,
	permAdded, permRemoved []UserPermission,
) (string, error) {
	folderItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(folder, true))
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to create folder %s/%s. details: %s", parentFolderID, folder,
			support.ConnectorStackErrorTrace(err),
		)
	}

	logger.Ctx(ctx).Debugf("Resolved %s in %s to %s", folder, parentFolderID, *folderItem.GetId())

	err = restorePermissions(ctx, service, driveID, *folderItem.GetId(), permAdded, permRemoved)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to set folder permissions %s/%s. details: %s", parentFolderID, folder,
			support.ConnectorStackErrorTrace(err),
		)
	}

	return *folderItem.GetId(), nil
}

// createRestoreFolders creates the restore folder hierarchy in the specified drive and returns the folder ID
// of the last folder entry in the hierarchy
func CreateRestoreFolders(ctx context.Context, service graph.Servicer, driveID string, restoreFolders []string,
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

		logger.Ctx(ctx).Debugw("resolved restore destination",
			"dest_name", folder,
			"parent", parentFolderID,
			"dest_id", *folderItem.GetId())

		parentFolderID = *folderItem.GetId()
	}

	return parentFolderID, nil
}

// restoreItem will create a new item in the specified `parentFolderID` and upload the data.Stream
func restoreItem(
	ctx context.Context,
	service graph.Servicer,
	name string,
	itemData data.Stream,
	driveID, parentFolderID string,
	copyBuffer []byte,
	source driveSource,
) (string, details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreItem", D.Label("item_uuid", itemData.UUID()))
	defer end()

	itemName := itemData.UUID()
	trace.Log(ctx, "gc:oneDrive:restoreItem", itemName)

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.StreamSize)
	if !ok {
		return "", details.ItemInfo{}, errors.Errorf("item %q does not implement DataStreamInfo", itemName)
	}

	// Create Item
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(name, false))
	if err != nil {
		return "", details.ItemInfo{}, errors.Wrapf(err, "failed to create item %s", itemName)
	}

	// Get a drive item writer
	w, err := driveItemWriter(ctx, service, driveID, *newItem.GetId(), ss.Size())
	if err != nil {
		return "", details.ItemInfo{}, errors.Wrapf(err, "failed to create item upload session %s", itemName)
	}

	iReader := itemData.ToReader()
	progReader, closer := observe.ItemProgress(ctx, iReader, observe.ItemRestoreMsg, itemName, ss.Size())

	go closer()

	// Upload the stream data
	written, err := io.CopyBuffer(w, progReader, copyBuffer)
	if err != nil {
		return "", details.ItemInfo{}, errors.Wrapf(err, "failed to upload data: item %s", itemName)
	}

	dii := details.ItemInfo{}

	switch source {
	case SharePointSource:
		dii.SharePoint = sharePointItemInfo(newItem, written)
	default:
		dii.OneDrive = oneDriveItemInfo(newItem, written)
	}

	return *newItem.GetId(), dii, nil
}

// getMetadata read and parses the metadata info for an item
func getMetadata(metar io.ReadCloser) (Metadata, error) {
	var meta Metadata
	// `metar` will be nil for the top level container folder
	if metar != nil {
		metaraw, err := io.ReadAll(metar)
		if err != nil {
			return Metadata{}, err
		}

		// TODO(meain): First 4 bytes are somehow 00 00 00 01 for
		// folder metadata
		start := 0
		if metaraw[0] != '{' {
			start = 4
		}

		err = json.Unmarshal(metaraw[start:], &meta)
		if err != nil {
			return Metadata{}, err
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

// restorePermissions takes in the permissions that were added and the
// removed(ones present in parent but not in child) and adds/removes
// the necessary permissions on onedrive objects.
func restorePermissions(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	itemID string,
	permAdded []UserPermission,
	permRemoved []UserPermission,
) error {
	for _, p := range permAdded {
		pbody := msdrive.NewItemsItemInvitePostRequestBody()
		pbody.SetRoles(p.Roles)

		if p.Expiration != nil {
			expiry := p.Expiration.String()
			pbody.SetExpirationDateTime(&expiry)
		}

		si := false
		pbody.SetSendInvitation(&si)

		rs := true
		pbody.SetRequireSignIn(&rs)

		rec := models.NewDriveRecipient()
		rec.SetEmail(&p.Email)
		pbody.SetRecipients([]models.DriveRecipientable{rec})

		_, err := service.Client().DrivesById(driveID).ItemsById(itemID).Invite().Post(ctx, pbody, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to set permissions for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}
	}

	for _, p := range permRemoved {
		err := service.Client().DrivesById(driveID).ItemsById(itemID).PermissionsById(p.ID).Delete(ctx, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to set permissions for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}
	}

	// TODO(meain): should it return effect permissions after
	// accepting child and parent permissions?
	return nil
}
