package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"runtime/trace"
	"sort"
	"strings"

	"github.com/alcionai/clues"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	// VersionWithDataAndMetaFiles is the corso backup format version
	// in which we split from storing just the data to storing both
	// the data and metadata in two files.
	VersionWithDataAndMetaFiles = 1
	// VersionWithIsMetaMarker is a small improvement on
	// VersionWithDataAndMetaFiles, but has a marker IsMeta which
	// specifies if the file is a meta file or a data file.
	VersionWithIsMetaMarker = 3
	// VersionWithNameInMeta points to the backup format version where we begin
	// storing files in kopia with their item ID instead of their OneDrive file
	// name.
	// TODO(ashmrtn): Update this to a real value when we merge the file name
	// change. Set to MAXINT for now to keep the if-check using it working.
	VersionWithNameInMeta = math.MaxInt
)

func getParentPermissions(
	parentPath path.Path,
	parentPermissions map[string][]UserPermission,
) ([]UserPermission, error) {
	parentPerms, ok := parentPermissions[parentPath.String()]
	if !ok {
		onedrivePath, err := path.ToOneDrivePath(parentPath)
		if err != nil {
			return nil, errors.Wrap(err, "invalid restore path")
		}

		if len(onedrivePath.Folders) != 0 {
			return nil, errors.Wrap(err, "unable to compute item permissions")
		}

		parentPerms = []UserPermission{}
	}

	return parentPerms, nil
}

func getParentAndCollectionPermissions(
	drivePath *path.DrivePath,
	collectionPath path.Path,
	permissions map[string][]UserPermission,
	restorePerms bool,
) ([]UserPermission, []UserPermission, error) {
	if !restorePerms {
		return nil, nil, nil
	}

	var (
		err         error
		parentPerms []UserPermission
		colPerms    []UserPermission
	)

	// Only get parent permissions if we're not restoring the root.
	if len(drivePath.Folders) > 0 {
		parentPath, err := collectionPath.Dir()
		if err != nil {
			return nil, nil, clues.Wrap(err, "getting parent path")
		}

		parentPerms, err = getParentPermissions(parentPath, permissions)
		if err != nil {
			return nil, nil, clues.Wrap(err, "getting parent permissions")
		}
	}

	// TODO(ashmrtn): For versions after this pull the permissions from the
	// current collection with Fetch().
	colPerms, err = getParentPermissions(collectionPath, permissions)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting collection permissions")
	}

	return parentPerms, colPerms, nil
}

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	backupVersion int,
	service graph.Servicer,
	dest control.RestoreDestination,
	opts control.Options,
	dcs []data.RestoreCollection,
	deets *details.Builder,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		restoreErrors  error
		metrics        support.CollectionMetrics
		folderPerms    map[string][]UserPermission
		canceled       bool

		// permissionIDMappings is used to map between old and new id
		// of permissions as we restore them
		permissionIDMappings = map[string]string{}
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
		metrics, folderPerms, permissionIDMappings, canceled = RestoreCollection(
			ctx,
			backupVersion,
			service,
			dc,
			parentPermissions,
			OneDriveSource,
			dest.ContainerName,
			deets,
			errUpdater,
			permissionIDMappings,
			opts.RestorePermissions,
		)

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
	backupVersion int,
	service graph.Servicer,
	dc data.RestoreCollection,
	parentPermissions map[string][]UserPermission,
	source driveSource,
	restoreContainerName string,
	deets *details.Builder,
	errUpdater func(string, error),
	permissionIDMappings map[string]string,
	restorePerms bool,
) (support.CollectionMetrics, map[string][]UserPermission, map[string]string, bool) {
	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics     = support.CollectionMetrics{}
		copyBuffer  = make([]byte, copyBufferSize)
		directory   = dc.FullPath()
		itemInfo    details.ItemInfo
		folderPerms = map[string][]UserPermission{}
	)

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		errUpdater(directory.String(), err)
		return metrics, folderPerms, permissionIDMappings, false
	}

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

	restoreFolderElements := []string{restoreContainerName}
	restoreFolderElements = append(restoreFolderElements, drivePath.Folders...)

	trace.Log(ctx, "gc:oneDrive:restoreCollection", directory.String())
	logger.Ctx(ctx).Infow(
		"restoring to destination",
		"origin", dc.FullPath().Folder(false),
		"destination", restoreFolderElements)

	parentPerms, colPerms, err := getParentAndCollectionPermissions(
		drivePath,
		dc.FullPath(),
		parentPermissions,
		restorePerms)
	if err != nil {
		errUpdater(directory.String(), err)
		return metrics, folderPerms, permissionIDMappings, false
	}

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := createRestoreFoldersWithPermissions(
		ctx,
		service,
		drivePath.DriveID,
		restoreFolderElements,
		parentPerms,
		colPerms,
		permissionIDMappings,
	)
	if err != nil {
		errUpdater(directory.String(), errors.Wrapf(err, "failed to create folders %v", restoreFolderElements))
		return metrics, folderPerms, permissionIDMappings, false
	}

	// Restore items from the collection
	items := dc.Items(ctx, nil) // TODO: fault.Errors instead of nil

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return metrics, folderPerms, permissionIDMappings, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, folderPerms, permissionIDMappings, false
			}

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				logger.Ctx(ctx).DPanicw("transforming item to full path", "error", err)

				errUpdater(itemData.UUID(), err)

				continue
			}

			if source == OneDriveSource && backupVersion >= VersionWithDataAndMetaFiles {
				name := itemData.UUID()
				if strings.HasSuffix(name, DataFileSuffix) {
					metrics.Objects++
					metrics.TotalBytes += int64(len(copyBuffer))

					var (
						itemInfo details.ItemInfo
						err      error
					)

					if backupVersion < VersionWithNameInMeta {
						itemInfo, err = restoreV1File(
							ctx,
							source,
							service,
							drivePath,
							dc,
							restoreFolderID,
							copyBuffer,
							colPerms,
							permissionIDMappings,
							restorePerms,
							itemData,
						)
					} else {
						itemInfo, err = restoreV2File(
							ctx,
							source,
							service,
							drivePath,
							dc,
							restoreFolderID,
							copyBuffer,
							colPerms,
							permissionIDMappings,
							restorePerms,
							itemData,
						)
					}

					if err != nil {
						errUpdater(itemData.UUID(), err)
						continue
					}

					deets.Add(
						itemPath.String(),
						itemPath.ShortRef(),
						"",
						"", // TODO: implement locationRef
						true,
						itemInfo)
					metrics.Successes++
				} else if strings.HasSuffix(name, MetaFileSuffix) {
					// Just skip this for the moment since we moved the code to the above
					// item restore path. We haven't yet stopped fetching these items in
					// RestoreOp, so we still need to handle them in some way.
					continue
				} else if strings.HasSuffix(name, DirMetaFileSuffix) {
					if !restorePerms {
						continue
					}

					metaReader := itemData.ToReader()
					meta, err := getMetadata(metaReader)
					metaReader.Close()

					if err != nil {
						errUpdater(itemData.UUID(), clues.Wrap(err, "folder metadata"))
						continue
					}

					trimmedPath := strings.TrimSuffix(itemPath.String(), DirMetaFileSuffix)
					folderPerms[trimmedPath] = meta.Permissions
				} else {
					if !ok {
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

				deets.Add(
					itemPath.String(),
					itemPath.ShortRef(),
					"",
					"", // TODO: implement locationRef
					true,
					itemInfo)
				metrics.Successes++
			}
		}
	}
}

type fileFetcher interface {
	Fetch(ctx context.Context, name string) (data.Stream, error)
}

func restoreV1File(
	ctx context.Context,
	source driveSource,
	service graph.Servicer,
	drivePath *path.DrivePath,
	fetcher fileFetcher,
	restoreFolderID string,
	copyBuffer []byte,
	parentPerms []UserPermission,
	permissionIDMappings map[string]string,
	restorePerms bool,
	itemData data.Stream,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.UUID(), DataFileSuffix)

	itemID, itemInfo, err := restoreData(
		ctx,
		service,
		trimmedName,
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		copyBuffer,
		source)
	if err != nil {
		return details.ItemInfo{}, err
	}

	// Mark it as success without processing .meta
	// file if we are not restoring permissions
	if !restorePerms {
		return itemInfo, nil
	}

	// Fetch item permissions from the collection and restore them.
	metaName := trimmedName + MetaFileSuffix

	meta, err := fetchAndReadMetadata(ctx, fetcher, metaName)
	if err != nil {
		err = clues.Wrap(err, "restoring file")
		return details.ItemInfo{}, err
	}

	err = restorePermissions(
		ctx,
		service,
		drivePath.DriveID,
		itemID,
		parentPerms,
		meta.Permissions,
		permissionIDMappings,
	)
	if err != nil {
		err = clues.Wrap(err, "restoring item permissions")
		return details.ItemInfo{}, err
	}

	return itemInfo, nil
}

func restoreV2File(
	ctx context.Context,
	source driveSource,
	service graph.Servicer,
	drivePath *path.DrivePath,
	fetcher fileFetcher,
	restoreFolderID string,
	copyBuffer []byte,
	parentPerms []UserPermission,
	permissionIDMappings map[string]string,
	restorePerms bool,
	itemData data.Stream,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.UUID(), DataFileSuffix)

	// Get metadata file so we can determine the file name.
	metaName := trimmedName + MetaFileSuffix

	meta, err := fetchAndReadMetadata(ctx, fetcher, metaName)
	if err != nil {
		err = clues.Wrap(err, "restoring file")
		return details.ItemInfo{}, err
	}

	if err != nil {
		err = clues.Wrap(err, "deserializing item metadata")
		return details.ItemInfo{}, err
	}

	// TODO(ashmrtn): Future versions could attempt to do the restore in a
	// different location like "lost+found" and use the item ID if we want to do
	// as much as possible to restore the data.
	if len(meta.FileName) == 0 {
		return details.ItemInfo{}, clues.New("item with empty name")
	}

	itemID, itemInfo, err := restoreData(
		ctx,
		service,
		meta.FileName,
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		copyBuffer,
		source)
	if err != nil {
		return details.ItemInfo{}, err
	}

	// Mark it as success without processing .meta
	// file if we are not restoring permissions
	if !restorePerms {
		return itemInfo, nil
	}

	err = restorePermissions(
		ctx,
		service,
		drivePath.DriveID,
		itemID,
		parentPerms,
		meta.Permissions,
		permissionIDMappings,
	)
	if err != nil {
		err = clues.Wrap(err, "restoring item permissions")
		return details.ItemInfo{}, err
	}

	return itemInfo, nil
}

// createRestoreFoldersWithPermissions creates the restore folder hierarchy in
// the specified drive and returns the folder ID of the last folder entry in the
// hierarchy. Permissions are only applied to the last folder in the hierarchy.
// Passing nil for the permissions results in just creating the folder(s).
func createRestoreFoldersWithPermissions(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	restoreFolders []string,
	parentPermissions []UserPermission,
	folderPermissions []UserPermission,
	permissionIDMappings map[string]string,
) (string, error) {
	id, err := CreateRestoreFolders(ctx, service, driveID, restoreFolders)
	if err != nil {
		return "", err
	}

	err = restorePermissions(
		ctx,
		service,
		driveID,
		id,
		parentPermissions,
		folderPermissions,
		permissionIDMappings)

	return id, err
}

// CreateRestoreFolders creates the restore folder hierarchy in the specified
// drive and returns the folder ID of the last folder entry in the hierarchy.
func CreateRestoreFolders(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	restoreFolders []string,
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
	progReader, closer := observe.ItemProgress(ctx, iReader, observe.ItemRestoreMsg, observe.PII(itemName), ss.Size())

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

func fetchAndReadMetadata(
	ctx context.Context,
	fetcher fileFetcher,
	metaName string,
) (Metadata, error) {
	metaFile, err := fetcher.Fetch(ctx, metaName)
	if err != nil {
		err = clues.Wrap(err, "getting item metadata").With("meta_file_name", metaName)
		return Metadata{}, err
	}

	metaReader := metaFile.ToReader()
	defer metaReader.Close()

	meta, err := getMetadata(metaReader)
	if err != nil {
		err = clues.Wrap(err, "deserializing item metadata").With("meta_file_name", metaName)
		return Metadata{}, err
	}

	return meta, nil
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
	permissionIDMappings map[string]string,
) error {
	permAdded, permRemoved := getChildPermissions(childPerms, parentPerms)

	for _, p := range permRemoved {
		err := service.Client().DrivesById(driveID).ItemsById(itemID).
			PermissionsById(permissionIDMappings[p.ID]).Delete(ctx, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to remove permission for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}
	}

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

		np, err := service.Client().DrivesById(driveID).ItemsById(itemID).Invite().Post(ctx, pbody, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to set permission for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}

		permissionIDMappings[p.ID] = *np.GetValue()[0].GetId()
	}

	return nil
}
