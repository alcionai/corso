package onedrive

import (
	"context"
	"encoding/json"
	"io"
	"runtime/trace"
	"sort"
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// copyBufferSize is used for chunked upload
// Microsoft recommends 5-10MB buffers
// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession?view=graph-rest-1.0#best-practices
const copyBufferSize = 5 * 1024 * 1024

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	backupVersion int,
	service graph.Servicer,
	dest control.RestoreDestination,
	opts control.Options,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		metrics        support.CollectionMetrics
		folderMetas    map[string]Metadata

		// permissionIDMappings is used to map between old and new id
		// of permissions as we restore them
		permissionIDMappings = map[string]string{}
	)

	ctx = clues.Add(
		ctx,
		"backup_version", backupVersion,
		"destination", dest.ContainerName)

	// Reorder collections so that the parents directories are created
	// before the child directories
	sort.Slice(dcs, func(i, j int) bool {
		return dcs[i].FullPath().String() < dcs[j].FullPath().String()
	})

	var (
		el          = errs.Local()
		parentMetas = map[string]Metadata{}
	)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err  error
			ictx = clues.Add(
				ctx,
				"resource_owner", dc.FullPath().ResourceOwner(), // TODO: pii
				"category", dc.FullPath().Category(),
				"path", dc.FullPath()) // TODO: pii
		)

		metrics, folderMetas, permissionIDMappings, err = RestoreCollection(
			ictx,
			backupVersion,
			service,
			dc,
			parentMetas,
			OneDriveSource,
			dest.ContainerName,
			deets,
			permissionIDMappings,
			opts.RestorePermissions,
			errs)
		if err != nil {
			el.AddRecoverable(err)
		}

		for k, v := range folderMetas {
			parentMetas[k] = v
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if errors.Is(err, context.Canceled) {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		dest.ContainerName)

	return status, el.Failure()
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
	parentMetas map[string]Metadata,
	source driveSource,
	restoreContainerName string,
	deets *details.Builder,
	permissionIDMappings map[string]string,
	restorePerms bool,
	errs *fault.Bus,
) (support.CollectionMetrics, map[string]Metadata, map[string]string, error) {
	var (
		metrics     = support.CollectionMetrics{}
		copyBuffer  = make([]byte, copyBufferSize)
		directory   = dc.FullPath()
		itemInfo    details.ItemInfo
		folderMetas = map[string]Metadata{}
	)

	ctx, end := D.Span(ctx, "gc:oneDrive:restoreCollection", D.Label("path", directory))
	defer end()

	drivePath, err := path.ToOneDrivePath(directory)
	if err != nil {
		return metrics, folderMetas, permissionIDMappings, clues.Wrap(err, "creating drive path").WithClues(ctx)
	}

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<drive>/root:/<restoreContainerName>/<original folder path>`

	restoreFolderElements := []string{restoreContainerName}
	restoreFolderElements = append(restoreFolderElements, drivePath.Folders...)

	ctx = clues.Add(
		ctx,
		"directory", dc.FullPath().Folder(false),
		"destination_elements", restoreFolderElements,
		"drive_id", drivePath.DriveID)

	trace.Log(ctx, "gc:oneDrive:restoreCollection", directory.String())
	logger.Ctx(ctx).Info("restoring onedrive collection")

	colMeta, err := getCollectionMetadata(
		ctx,
		drivePath,
		dc,
		parentMetas,
		backupVersion,
		restorePerms)
	if err != nil {
		return metrics, folderMetas, permissionIDMappings, clues.Wrap(err, "getting permissions").WithClues(ctx)
	}

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := createRestoreFoldersWithPermissions(
		ctx,
		service,
		drivePath.DriveID,
		restoreFolderElements,
		colMeta,
		permissionIDMappings,
	)
	if err != nil {
		return metrics, folderMetas, permissionIDMappings, clues.Wrap(err, "creating folders for restore")
	}

	var (
		el    = errs.Local()
		items = dc.Items(ctx, errs)
	)

	for {
		if el.Failure() != nil {
			break
		}

		select {
		case <-ctx.Done():
			return metrics, folderMetas, permissionIDMappings, err

		case itemData, ok := <-items:
			if !ok {
				return metrics, folderMetas, permissionIDMappings, nil
			}

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				el.AddRecoverable(clues.Wrap(err, "appending item to full path").WithClues(ctx))
				continue
			}

			if source == OneDriveSource && backupVersion >= version.OneDrive1DataAndMetaFiles {
				name := itemData.UUID()

				if strings.HasSuffix(name, DataFileSuffix) {
					metrics.Objects++
					metrics.Bytes += int64(len(copyBuffer))

					var (
						itemInfo details.ItemInfo
						err      error
					)

					if backupVersion < version.OneDriveXNameInMeta {
						itemInfo, err = restoreV1File(
							ctx,
							source,
							service,
							drivePath,
							dc,
							restoreFolderID,
							copyBuffer,
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
							permissionIDMappings,
							restorePerms,
							itemData,
						)
					}

					if err != nil {
						el.AddRecoverable(err)
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
					// Only the version.OneDrive1DataAndMetaFiles needed to deserialize the
					// permission for child folders here. Later versions can request
					// permissions inline when processing the collection.
					if !restorePerms || backupVersion >= version.OneDrive4DirIncludesPermissions {
						continue
					}

					metaReader := itemData.ToReader()
					defer metaReader.Close()

					meta, err := getMetadata(metaReader)
					if err != nil {
						el.AddRecoverable(clues.Wrap(err, "getting directory metadata").WithClues(ctx))
						continue
					}

					trimmedPath := strings.TrimSuffix(itemPath.String(), DirMetaFileSuffix)
					folderMetas[trimmedPath] = meta

				}
			} else {
				metrics.Objects++
				metrics.Bytes += int64(len(copyBuffer))

				// No permissions stored at the moment for SharePoint
				_, itemInfo, err = restoreData(
					ctx,
					service,
					itemData.UUID(),
					itemData,
					drivePath.DriveID,
					restoreFolderID,
					copyBuffer,
					source)
				if err != nil {
					el.AddRecoverable(err)
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

	return metrics, folderMetas, permissionIDMappings, el.Failure()
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
		return details.ItemInfo{}, clues.Wrap(err, "restoring file")
	}

	err = restorePermissions(
		ctx,
		service,
		drivePath.DriveID,
		itemID,
		meta,
		permissionIDMappings,
	)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring item permissions")
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
	permissionIDMappings map[string]string,
	restorePerms bool,
	itemData data.Stream,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.UUID(), DataFileSuffix)

	// Get metadata file so we can determine the file name.
	metaName := trimmedName + MetaFileSuffix

	meta, err := fetchAndReadMetadata(ctx, fetcher, metaName)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring file")
	}

	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "deserializing item metadata")
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
		meta,
		permissionIDMappings,
	)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring item permissions")
	}

	return itemInfo, nil
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
		return "", graph.Wrap(ctx, err, "getting drive root")
	}

	parentFolderID := ptr.Val(driveRoot.GetId())
	ctx = clues.Add(ctx, "drive_root_id", parentFolderID)

	logger.Ctx(ctx).Debug("found drive root")

	for _, folder := range restoreFolders {
		folderItem, err := getFolder(ctx, service, driveID, parentFolderID, folder)
		if err == nil {
			parentFolderID = ptr.Val(folderItem.GetId())
			continue
		}

		if !errors.Is(err, errFolderNotFound) {
			return "", clues.Wrap(err, "folder not found").With("folder_id", folder).WithClues(ctx)
		}

		folderItem, err = createItem(ctx, service, driveID, parentFolderID, newItem(folder, true))
		if err != nil {
			return "", clues.Wrap(err, "creating folder")
		}

		parentFolderID = ptr.Val(folderItem.GetId())

		logger.Ctx(ctx).Debugw("resolved restore destination", "dest_id", parentFolderID)
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

	ctx = clues.Add(ctx, "item_name", itemData.UUID())

	itemName := itemData.UUID()
	trace.Log(ctx, "gc:oneDrive:restoreItem", itemName)

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.StreamSize)
	if !ok {
		return "", details.ItemInfo{}, clues.New("item does not implement DataStreamInfo").WithClues(ctx)
	}

	// Create Item
	newItem, err := createItem(ctx, service, driveID, parentFolderID, newItem(name, false))
	if err != nil {
		return "", details.ItemInfo{}, clues.Wrap(err, "creating item")
	}

	// Get a drive item writer
	w, err := driveItemWriter(ctx, service, driveID, *newItem.GetId(), ss.Size())
	if err != nil {
		return "", details.ItemInfo{}, clues.Wrap(err, "creating item writer")
	}

	iReader := itemData.ToReader()
	progReader, closer := observe.ItemProgress(ctx, iReader, observe.ItemRestoreMsg, observe.PII(itemName), ss.Size())

	go closer()

	// Upload the stream data
	written, err := io.CopyBuffer(w, progReader, copyBuffer)
	if err != nil {
		return "", details.ItemInfo{}, graph.Wrap(ctx, err, "writing item bytes")
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
