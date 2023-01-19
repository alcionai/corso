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

	parentPermissions := map[string][]UserPermission{}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		parentPerms, ok := parentPermissions[dc.FullPath().String()]
		if !ok {
			// TODO(meain): validate this change with actual backup and restore
			if len(dc.FullPath().Elements()) == 7 {
				// root directory will not have permissions
				parentPerms = []UserPermission{}
			} else {
				errUpdater(dc.FullPath().String(), fmt.Errorf("unable to find parent permissions"))
			}
		}

		metrics, folderPerms, canceled := RestoreCollection(ctx, service, dc, parentPerms, OneDriveSource, dest.ContainerName, deets, errUpdater)

		for k, v := range folderPerms {
			parentPermissions[k] = v
		}

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
// - the context cancellation state (true if the context is canceled)
func RestoreCollection(
	ctx context.Context,
	service graph.Servicer,
	dc data.Collection,
	parentPerms []UserPermission,
	source driveSource,
	restoreContainerName string,
	deets *details.Builder,
	errUpdater func(string, error),
) (support.CollectionMetrics, map[string][]UserPermission, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics     = support.CollectionMetrics{}
		copyBuffer  = make([]byte, copyBufferSize)
		directory   = dc.FullPath()
		restoredIDs = map[string]string{}
		itemInfo    details.ItemInfo
		itemID      string
		folderPerms = map[string][]UserPermission{}
	)

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		errUpdater(directory.String(), err)
		return metrics, folderPerms, false
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
		return metrics, folderPerms, false
	}

	// Restore items from the collection
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return metrics, folderPerms, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, folderPerms, false
			}

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				logger.Ctx(ctx).DPanicw("transforming item to full path", "error", err)
				errUpdater(itemData.UUID(), err)
				continue
			}

			if source == OneDriveSource {
				name := itemData.UUID()
				if strings.HasSuffix(name, DataFileSuffix) {
					metrics.Objects++
					metrics.TotalBytes += int64(len(copyBuffer))
					trimmedName := strings.TrimSuffix(name, DataFileSuffix)

					itemID, itemInfo, err = restoreData(ctx, service, trimmedName, itemData, drivePath.DriveID, restoreFolderID, copyBuffer, source)
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					restoredIDs[trimmedName] = itemID

					deets.Add(itemPath.String(), itemPath.ShortRef(), "", true, itemInfo)
				} else if strings.HasSuffix(name, MetaFileSuffix) {
					meta, err := getMetadata(itemData.ToReader())
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					trimmedName := strings.TrimSuffix(name, MetaFileSuffix)
					restoreID, ok := restoredIDs[trimmedName]
					if !ok {
						errUpdater(itemData.UUID(), fmt.Errorf("item not available to restore permissions"))
						continue
					}

					err = restorePermissions(ctx, service, drivePath.DriveID, restoreID, parentPerms, meta.Permissions)
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					// Objects count is incremented when we restore a
					// data file and success count is incremented when
					// we restore a meta file as every data file
					// should have an associated meta file
					metrics.Successes++
				} else if strings.HasSuffix(name, DirMetaFileSuffix) {
					meta, err := getMetadata(itemData.ToReader())
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					trimmedName := strings.TrimSuffix(name, DirMetaFileSuffix)
					_, err = createRestoreFolder(ctx, service, drivePath.DriveID, trimmedName, restoreFolderID, parentPerms, meta.Permissions)
					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					trimmedPath := strings.TrimSuffix(itemPath.String(), DirMetaFileSuffix)
					folderPerms[trimmedPath] = meta.Permissions
				} else {
					if !ok {
						// TODO(meain): support older backups?
						errUpdater(itemData.UUID(), fmt.Errorf("invalid backup format, you might be using an old backup"))
						continue
					}
				}
			} else {
				metrics.Objects++
				metrics.TotalBytes += int64(len(copyBuffer))

				// No permissions stored at the moment for SharePoint
				_, itemInfo, err = restoreData(ctx,
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

				deets.Add(itemPath.String(), itemPath.ShortRef(), "", true, itemInfo)
				metrics.Successes++ // Counted as success when we have also restored metadata
			}

		}
	}
}

// Creates a folder with its permissions
func createRestoreFolder(ctx context.Context,
	service graph.Servicer,
	driveID, folder, parentFolderID string,
	parentPerms, childPerms []UserPermission,
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

	err = restorePermissions(ctx, service, driveID, *folderItem.GetId(), parentPerms, childPerms)
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

// restoreData will create a new item in the specified `parentFolderID` and upload the data.Stream
func restoreData(
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

		err = json.Unmarshal(metaraw, &meta)
		if err != nil {
			panic(err)
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
	parentPerms []UserPermission,
	childPerms []UserPermission,
) error {
	permAdded, permRemoved := getChildPermissions(childPerms, parentPerms)

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

	return nil
}
