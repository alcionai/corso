package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	// Maximum number of retries for upload failures
	maxUploadRetries = 3
)

// RestoreCollection handles restoration of an individual collection.
// returns:
// - the collection's item and byte count metrics
// - the updated metadata map that include metadata for folders in this collection
// - error, if any besides recoverable
func RestoreCollection(
	ctx context.Context,
	rh RestoreHandler,
	rcc inject.RestoreConsumerConfig,
	dc data.RestoreCollection,
	caches *restoreCaches,
	deets *details.Builder,
	fallbackDriveName string,
	errs *fault.Bus,
	ctr *count.Bus,
) (support.CollectionMetrics, error) {
	var (
		metrics        = support.CollectionMetrics{}
		directory      = dc.FullPath()
		el             = errs.Local()
		metricsObjects int64
		metricsBytes   int64
		metricsSuccess int64
		wg             sync.WaitGroup
		complete       bool
	)

	ctx, end := diagnostics.Span(ctx, "gc:drive:restoreCollection", diagnostics.Label("path", directory))
	defer end()

	drivePath, err := path.ToDrivePath(directory)
	if err != nil {
		return metrics, clues.Wrap(err, "creating drive path").WithClues(ctx)
	}

	di, err := ensureDriveExists(
		ctx,
		rh,
		caches,
		drivePath,
		rcc.ProtectedResource.ID(),
		fallbackDriveName)
	if err != nil {
		return metrics, clues.Wrap(err, "ensuring drive exists")
	}

	// clobber the drivePath details with the details retrieved
	// in the ensure func, as they might have changed to reflect
	// a different drive as a restore location.
	drivePath.DriveID = di.id
	drivePath.Root = di.rootFolderID

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<restoreContainerName>/<original folder path>`
	// the drive into which this folder gets restored is tracked separately in drivePath.
	restoreDir := &path.Builder{}

	if len(rcc.RestoreConfig.Location) > 0 {
		restoreDir = restoreDir.Append(rcc.RestoreConfig.Location)
	}

	restoreDir = restoreDir.Append(drivePath.Folders...)

	ctx = clues.Add(
		ctx,
		"directory", dc.FullPath().Folder(false),
		"restore_destination", restoreDir,
		"drive_id", drivePath.DriveID)

	trace.Log(ctx, "gc:oneDrive:restoreCollection", directory.String())
	logger.Ctx(ctx).Info("restoring onedrive collection")

	colMeta, err := getCollectionMetadata(
		ctx,
		drivePath,
		dc,
		caches,
		rcc.BackupVersion,
		rcc.RestoreConfig.IncludePermissions)
	if err != nil {
		return metrics, clues.Wrap(err, "getting permissions").WithClues(ctx)
	}

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := CreateRestoreFolders(
		ctx,
		rh,
		drivePath,
		restoreDir,
		dc.FullPath(),
		colMeta,
		caches,
		rcc.RestoreConfig.IncludePermissions)
	if err != nil {
		return metrics, clues.Wrap(err, "creating folders for restore")
	}

	collisionKeyToItemID, err := rh.GetItemsInContainerByCollisionKey(ctx, drivePath.DriveID, restoreFolderID)
	if err != nil {
		return metrics, clues.Wrap(err, "generating map of item collision keys")
	}

	caches.collisionKeyToItemID = collisionKeyToItemID
	caches.ParentDirToMeta.Store(dc.FullPath().String(), colMeta)
	items := dc.Items(ctx, errs)

	semaphoreCh := make(chan struct{}, graph.Parallelism(path.OneDriveService).ItemUpload())
	defer close(semaphoreCh)

	deetsLock := sync.Mutex{}

	updateDeets := func(
		ctx context.Context,
		repoRef path.Path,
		locationRef *path.Builder,
		info details.ItemInfo,
	) {
		deetsLock.Lock()
		defer deetsLock.Unlock()

		err = deets.Add(repoRef, locationRef, info)
		if err != nil {
			// Not critical enough to need to stop restore operation.
			logger.CtxErr(ctx, err).Infow("adding restored item to details")
		}
	}

	for {
		if el.Failure() != nil || complete {
			break
		}

		select {
		case <-ctx.Done():
			return metrics, err

		case itemData, ok := <-items:
			if !ok {
				// We've processed all items in this collection, exit the loop
				complete = true
				break
			}

			wg.Add(1)
			semaphoreCh <- struct{}{}

			go func(ctx context.Context, itemData data.Item) {
				defer wg.Done()
				defer func() { <-semaphoreCh }()

				copyBufferPtr := caches.pool.Get().(*[]byte)
				defer caches.pool.Put(copyBufferPtr)

				copyBuffer := *copyBufferPtr
				ictx := clues.Add(ctx, "restore_item_id", itemData.ID())

				itemPath, err := dc.FullPath().AppendItem(itemData.ID())
				if err != nil {
					el.AddRecoverable(ctx, clues.Wrap(err, "appending item to full path").WithClues(ictx))
					return
				}

				itemInfo, skipped, err := restoreItem(
					ictx,
					rh,
					rcc,
					dc,
					drivePath,
					restoreFolderID,
					copyBuffer,
					caches,
					itemData,
					itemPath,
					ctr)

				// skipped items don't get counted, but they can error
				if !skipped {
					atomic.AddInt64(&metricsObjects, 1)
					atomic.AddInt64(&metricsBytes, int64(len(copyBuffer)))
				}

				if err != nil {
					el.AddRecoverable(ctx, clues.Wrap(err, "restoring item"))
					return
				}

				if skipped {
					logger.Ctx(ictx).With("item_path", itemPath).Debug("did not restore item")
					return
				}

				// TODO: implement locationRef
				updateDeets(ictx, itemPath, &path.Builder{}, itemInfo)

				atomic.AddInt64(&metricsSuccess, 1)
			}(ctx, itemData)
		}
	}

	wg.Wait()

	metrics.Objects = int(metricsObjects)
	metrics.Bytes = metricsBytes
	metrics.Successes = int(metricsSuccess)

	return metrics, el.Failure()
}

// restores an item, according to correct backup version behavior.
// returns the item info, a bool (true = restore was skipped), and an error
func restoreItem(
	ctx context.Context,
	rh RestoreHandler,
	rcc inject.RestoreConsumerConfig,
	fibn data.FetchItemByNamer,
	drivePath *path.DrivePath,
	restoreFolderID string,
	copyBuffer []byte,
	caches *restoreCaches,
	itemData data.Item,
	itemPath path.Path,
	ctr *count.Bus,
) (details.ItemInfo, bool, error) {
	itemUUID := itemData.ID()
	ctx = clues.Add(ctx, "item_id", itemUUID)

	if rcc.BackupVersion < version.OneDrive1DataAndMetaFiles {
		itemInfo, err := restoreV0File(
			ctx,
			rh,
			rcc.RestoreConfig,
			drivePath,
			fibn,
			restoreFolderID,
			copyBuffer,
			caches.collisionKeyToItemID,
			itemData,
			ctr)
		if err != nil {
			if errors.Is(err, graph.ErrItemAlreadyExistsConflict) && rcc.RestoreConfig.OnCollision == control.Skip {
				return details.ItemInfo{}, true, nil
			}

			return details.ItemInfo{}, false, clues.Wrap(err, "v0 restore")
		}

		return itemInfo, false, nil
	}

	// only v1+ backups from this point on

	if strings.HasSuffix(itemUUID, metadata.MetaFileSuffix) {
		// Just skip this for the moment since we moved the code to the above
		// item restore path. We haven't yet stopped fetching these items in
		// RestoreOp, so we still need to handle them in some way.
		return details.ItemInfo{}, true, nil
	}

	if strings.HasSuffix(itemUUID, metadata.DirMetaFileSuffix) {
		// Only the version.OneDrive1DataAndMetaFiles needed to deserialize the
		// permission for child folders here. Later versions can request
		// permissions inline when processing the collection.
		if !rcc.RestoreConfig.IncludePermissions || rcc.BackupVersion >= version.OneDrive4DirIncludesPermissions {
			return details.ItemInfo{}, true, nil
		}

		metaReader := itemData.ToReader()
		defer metaReader.Close()

		meta, err := getMetadata(metaReader)
		if err != nil {
			return details.ItemInfo{}, true, clues.Wrap(err, "getting directory metadata").WithClues(ctx)
		}

		trimmedPath := strings.TrimSuffix(itemPath.String(), metadata.DirMetaFileSuffix)
		caches.ParentDirToMeta.Store(trimmedPath, meta)

		return details.ItemInfo{}, true, nil
	}

	// only items with DataFileSuffix from this point on

	if rcc.BackupVersion < version.OneDrive6NameInMeta {
		itemInfo, err := restoreV1File(
			ctx,
			rh,
			rcc,
			drivePath,
			fibn,
			restoreFolderID,
			copyBuffer,
			caches,
			itemPath,
			itemData,
			ctr)
		if err != nil {
			if errors.Is(err, graph.ErrItemAlreadyExistsConflict) && rcc.RestoreConfig.OnCollision == control.Skip {
				return details.ItemInfo{}, true, nil
			}

			return details.ItemInfo{}, false, clues.Wrap(err, "v1 restore")
		}

		return itemInfo, false, nil
	}

	// only v6+ backups from this point on

	itemInfo, err := restoreV6File(
		ctx,
		rh,
		rcc,
		drivePath,
		fibn,
		restoreFolderID,
		copyBuffer,
		caches,
		itemPath,
		itemData,
		ctr)
	if err != nil {
		if errors.Is(err, graph.ErrItemAlreadyExistsConflict) && rcc.RestoreConfig.OnCollision == control.Skip {
			return details.ItemInfo{}, true, nil
		}

		return details.ItemInfo{}, false, clues.Wrap(err, "v6 restore")
	}

	return itemInfo, false, nil
}

func restoreV0File(
	ctx context.Context,
	rh RestoreHandler,
	restoreCfg control.RestoreConfig,
	drivePath *path.DrivePath,
	fibn data.FetchItemByNamer,
	restoreFolderID string,
	copyBuffer []byte,
	collisionKeyToItemID map[string]api.DriveItemIDType,
	itemData data.Item,
	ctr *count.Bus,
) (details.ItemInfo, error) {
	_, itemInfo, err := restoreFile(
		ctx,
		restoreCfg,
		rh,
		fibn,
		itemData.ID(),
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		collisionKeyToItemID,
		copyBuffer,
		ctr)
	if err != nil {
		return itemInfo, clues.Wrap(err, "restoring file")
	}

	return itemInfo, nil
}

func restoreV1File(
	ctx context.Context,
	rh RestoreHandler,
	rcc inject.RestoreConsumerConfig,
	drivePath *path.DrivePath,
	fibn data.FetchItemByNamer,
	restoreFolderID string,
	copyBuffer []byte,
	caches *restoreCaches,
	itemPath path.Path,
	itemData data.Item,
	ctr *count.Bus,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.ID(), metadata.DataFileSuffix)

	itemID, itemInfo, err := restoreFile(
		ctx,
		rcc.RestoreConfig,
		rh,
		fibn,
		trimmedName,
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		caches.collisionKeyToItemID,
		copyBuffer,
		ctr)
	if err != nil {
		return details.ItemInfo{}, err
	}

	// Mark it as success without processing .meta
	// file if we are not restoring permissions
	if !rcc.RestoreConfig.IncludePermissions {
		return itemInfo, nil
	}

	// Fetch item permissions from the collection and restore them.
	metaName := trimmedName + metadata.MetaFileSuffix

	meta, err := FetchAndReadMetadata(ctx, fibn, metaName)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring file")
	}

	err = RestorePermissions(
		ctx,
		rh,
		drivePath.DriveID,
		itemID,
		itemPath,
		meta,
		caches)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring item permissions")
	}

	return itemInfo, nil
}

func restoreV6File(
	ctx context.Context,
	rh RestoreHandler,
	rcc inject.RestoreConsumerConfig,
	drivePath *path.DrivePath,
	fibn data.FetchItemByNamer,
	restoreFolderID string,
	copyBuffer []byte,
	caches *restoreCaches,
	itemPath path.Path,
	itemData data.Item,
	ctr *count.Bus,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.ID(), metadata.DataFileSuffix)

	// Get metadata file so we can determine the file name.
	metaName := trimmedName + metadata.MetaFileSuffix

	meta, err := FetchAndReadMetadata(ctx, fibn, metaName)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring file")
	}

	ctx = clues.Add(
		ctx,
		"count_perms", len(meta.Permissions),
		"restore_item_name", clues.Hide(meta.FileName))

	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "deserializing item metadata")
	}

	// TODO(ashmrtn): Future versions could attempt to do the restore in a
	// different location like "lost+found" and use the item ID if we want to do
	// as much as possible to restore the data.
	if len(meta.FileName) == 0 {
		return details.ItemInfo{}, clues.New("item with empty name")
	}

	itemID, itemInfo, err := restoreFile(
		ctx,
		rcc.RestoreConfig,
		rh,
		fibn,
		meta.FileName,
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		caches.collisionKeyToItemID,
		copyBuffer,
		ctr)
	if err != nil {
		return details.ItemInfo{}, err
	}

	// Mark it as success without processing .meta
	// file if we are not restoring permissions
	if !rcc.RestoreConfig.IncludePermissions {
		return itemInfo, nil
	}

	err = RestorePermissions(
		ctx,
		rh,
		drivePath.DriveID,
		itemID,
		itemPath,
		meta,
		caches)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring item permissions")
	}

	return itemInfo, nil
}

// CreateRestoreFolders creates the restore folder hierarchy in
// the specified drive and returns the folder ID of the last folder entry in the
// hierarchy. Permissions are only applied to the last folder in the hierarchy.
// Passing nil for the permissions results in just creating the folder(s).
// folderCache is mutated, as a side effect of populating the items.
func CreateRestoreFolders(
	ctx context.Context,
	rh RestoreHandler,
	drivePath *path.DrivePath,
	restoreDir *path.Builder,
	folderPath path.Path,
	folderMetadata metadata.Metadata,
	caches *restoreCaches,
	restorePerms bool,
) (string, error) {
	id, err := createRestoreFolders(
		ctx,
		rh,
		drivePath,
		restoreDir,
		caches)
	if err != nil {
		return "", err
	}

	if len(drivePath.Folders) == 0 {
		// No permissions for root folder
		return id, nil
	}

	if !restorePerms {
		return id, nil
	}

	err = RestorePermissions(
		ctx,
		rh,
		drivePath.DriveID,
		id,
		folderPath,
		folderMetadata,
		caches)

	return id, err
}

type folderRestorer interface {
	GetFolderByNamer
	PostItemInContainerer
}

// createRestoreFolders creates the restore folder hierarchy in the specified
// drive and returns the folder ID of the last folder entry in the hierarchy.
// folderCache is mutated, as a side effect of populating the items.
func createRestoreFolders(
	ctx context.Context,
	fr folderRestorer,
	drivePath *path.DrivePath,
	restoreDir *path.Builder,
	caches *restoreCaches,
) (string, error) {
	var (
		driveID             = drivePath.DriveID
		folders             = restoreDir.Elements()
		location            = path.Builder{}.Append(driveID)
		parentFolderMeta, _ = caches.DriveIDToDriveInfo.Load(drivePath.DriveID)
		parentFolderID      = parentFolderMeta.rootFolderID
	)

	ctx = clues.Add(
		ctx,
		"drive_id", drivePath.DriveID,
		"root_folder_id", parentFolderID)

	for _, folderName := range folders {
		location = location.Append(folderName)
		ictx := clues.Add(
			ctx,
			"creating_restore_folder", folderName,
			"restore_folder_location", location,
			"parent_of_restore_folder", parentFolderID)

		if fl, ok := caches.Folders.get(location); ok {
			parentFolderID = ptr.Val(fl.GetId())
			// folder was already created, move on to the child
			continue
		}

		// we assume this folder creation always uses the Replace
		// conflict policy, which means it will act as a GET if the
		// folder already exists.
		folderItem, err := createFolder(
			ictx,
			fr,
			driveID,
			parentFolderID,
			folderName)
		if err != nil {
			return "", clues.Wrap(err, "creating folder")
		}

		parentFolderID = ptr.Val(folderItem.GetId())
		caches.Folders.set(location, folderItem)

		logger.Ctx(ictx).Debug("resolved restore destination")
	}

	return parentFolderID, nil
}

func createFolder(
	ctx context.Context,
	piic PostItemInContainerer,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	// create the folder if not found
	// the Replace collision policy is used since collisions on that
	// policy will no-op and return the existing folder.  This has two
	// benefits: first, we get to treat the post as idempotent; and
	// second, we don't have to worry about race conditions.
	item, err := piic.PostItemInContainer(
		ctx,
		driveID,
		parentFolderID,
		newItem(folderName, true),
		control.Replace)

	// ErrItemAlreadyExistsConflict can only occur for folders if the
	// item being replaced is a file, not another folder.
	if err != nil && !errors.Is(err, graph.ErrItemAlreadyExistsConflict) {
		return nil, clues.Wrap(err, "creating folder")
	}

	if err == nil {
		return item, err
	}

	// if we made it here, then we tried to replace a file with a folder and
	// hit a conflict.  An unlikely occurrence, and we can try again with a copy
	// conflict behavior setting and probably succeed, though that will change
	// the location name of the restore.
	item, err = piic.PostItemInContainer(
		ctx,
		driveID,
		parentFolderID,
		newItem(folderName, true),
		control.Copy)
	if err != nil {
		return nil, clues.Wrap(err, "creating folder")
	}

	return item, err
}

type itemRestorer interface {
	DeleteItemer
	ItemInfoAugmenter
	NewItemContentUploader
	PostItemInContainerer
}

// restoreFile will create a new item in the specified `parentFolderID` and upload the data.Item
func restoreFile(
	ctx context.Context,
	restoreCfg control.RestoreConfig,
	ir itemRestorer,
	fibn data.FetchItemByNamer,
	name string,
	itemData data.Item,
	driveID, parentFolderID string,
	collisionKeyToItemID map[string]api.DriveItemIDType,
	copyBuffer []byte,
	ctr *count.Bus,
) (string, details.ItemInfo, error) {
	ctx, end := diagnostics.Span(ctx, "gc:oneDrive:restoreItem", diagnostics.Label("item_uuid", itemData.ID()))
	defer end()

	trace.Log(ctx, "gc:oneDrive:restoreItem", itemData.ID())

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.ItemSize)
	if !ok {
		return "", details.ItemInfo{}, clues.New("item does not implement DataStreamInfo").WithClues(ctx)
	}

	var (
		item                 = newItem(name, false)
		collisionKey         = api.DriveItemCollisionKey(item)
		collision            api.DriveItemIDType
		shouldDeleteOriginal bool
	)

	if dci, ok := collisionKeyToItemID[collisionKey]; ok {
		log := logger.Ctx(ctx).With("collision_key", clues.Hide(collisionKey))
		log.Debug("item collision")

		if restoreCfg.OnCollision == control.Skip {
			ctr.Inc(count.CollisionSkip)
			log.Debug("skipping item with collision")

			return "", details.ItemInfo{}, graph.ErrItemAlreadyExistsConflict
		}

		collision = dci
		shouldDeleteOriginal = restoreCfg.OnCollision == control.Replace && !dci.IsFolder
	}

	// drive items do not support PUT requests on the drive item data, so
	// when replacing a collision, first delete the existing item.  It would
	// be nice to be able to do a post-then-delete like we do with exchange,
	// but onedrive will conflict on the filename.  So until the api's built-in
	// conflict replace handling bug gets fixed, we either delete-post, and
	// risk failures in the middle, or we post w/ copy, then delete, then patch
	// the name, which could triple our graph calls in the worst case.
	if shouldDeleteOriginal {
		if err := ir.DeleteItem(ctx, driveID, collision.ItemID); err != nil && !graph.IsErrDeletedInFlight(err) {
			return "", details.ItemInfo{}, clues.New("deleting colliding item")
		}
	}

	// Create Item
	// the Copy collision policy is used since we've technically already handled
	// the collision behavior above.  At this point, copy is most likely to succeed.
	// We could go with control.Skip if we wanted to ensure no duplicate, but those
	// duplicates will only happen under very unlikely race conditions.
	newItem, err := ir.PostItemInContainer(
		ctx,
		driveID,
		parentFolderID,
		item,
		// notes on forced copy:
		// 1. happy path: any non-colliding item will restore as if no collision had occurred
		// 2. if a file-container collision is present, we assume the item being restored
		//    will get generated according to server-side copy rules.
		// 3. if restoreCfg specifies replace and a file-container collision is present, we
		//    make no changes to the original file, and do not delete it.
		control.Copy)
	if err != nil {
		return "", details.ItemInfo{}, err
	}

	w, uploadURL, err := driveItemWriter(ctx, ir, driveID, ptr.Val(newItem.GetId()), ss.Size())
	if err != nil {
		return "", details.ItemInfo{}, clues.Wrap(err, "get item upload session")
	}

	var (
		written          int64
		progReader       io.ReadCloser
		closeProgressBar func()
	)

	// This is just to retry file upload, the uploadSession creation is
	// not retried here We need extra logic to retry file upload as we
	// have to pull the file again from kopia If we fail a file upload,
	// we restart from scratch and try to upload again. Graph does not
	// show "register" any partial file uploads and so if we fail an
	// upload the file size will be 0.
	for i := 0; i <= maxUploadRetries; i++ {
		pname := name
		iReader := itemData.ToReader()

		if i > 0 {
			pname = fmt.Sprintf("%s (retry %d)", name, i)

			// If it is not the first try, we have to pull the file
			// again from kopia. Ideally we could just seek the stream
			// but we don't have a Seeker available here.
			itemData, err := fibn.FetchItemByName(ctx, itemData.ID())
			if err != nil {
				return "", details.ItemInfo{}, clues.Wrap(err, "get data file")
			}

			iReader = itemData.ToReader()
		}

		progReader, closeProgressBar = observe.ItemProgress(
			ctx,
			iReader,
			observe.ItemRestoreMsg,
			clues.Hide(pname),
			ss.Size())

		// Upload the stream data
		written, err = io.CopyBuffer(w, progReader, copyBuffer)
		if err == nil {
			break
		}

		// clear out the progress bar immediately on error
		closeProgressBar()

		// refresh the io.Writer to restart the upload
		// TODO: @vkamra verify if var session is the desired input
		w = graph.NewLargeItemWriter(ptr.Val(newItem.GetId()), uploadURL, ss.Size())
	}

	if err != nil {
		return "", details.ItemInfo{}, clues.Wrap(err, "uploading file")
	}

	defer closeProgressBar()

	dii := ir.AugmentItemInfo(details.ItemInfo{}, newItem, written, nil)

	if shouldDeleteOriginal {
		ctr.Inc(count.CollisionReplace)
	} else {
		ctr.Inc(count.NewItemCreated)
	}

	return ptr.Val(newItem.GetId()), dii, nil
}

func FetchAndReadMetadata(
	ctx context.Context,
	fibn data.FetchItemByNamer,
	metaName string,
) (metadata.Metadata, error) {
	ctx = clues.Add(ctx, "meta_file_name", metaName)

	metaFile, err := fibn.FetchItemByName(ctx, metaName)
	if err != nil {
		return metadata.Metadata{}, clues.Wrap(err, "getting item metadata")
	}

	metaReader := metaFile.ToReader()
	defer metaReader.Close()

	meta, err := getMetadata(metaReader)
	if err != nil {
		return metadata.Metadata{}, clues.Wrap(err, "deserializing item metadata")
	}

	return meta, nil
}

// getMetadata read and parses the metadata info for an item
func getMetadata(metar io.ReadCloser) (metadata.Metadata, error) {
	var meta metadata.Metadata
	// `metar` will be nil for the top level container folder
	if metar != nil {
		metaraw, err := io.ReadAll(metar)
		if err != nil {
			return metadata.Metadata{}, err
		}

		err = json.Unmarshal(metaraw, &meta)
		if err != nil {
			return metadata.Metadata{}, err
		}
	}

	return meta, nil
}

type PostDriveAndGetRootFolderer interface {
	PostDriver
	GetRootFolderer
}

// ensureDriveExists looks up the drive by its id.  If no drive is found with
// that ID, a new drive is generated with the same name.  If the name collides
// with an existing drive, a number is appended to the drive name.  Eg: foo ->
// foo 1.  This will repeat as many times as is needed.
// Returns the root folder of the drive
func ensureDriveExists(
	ctx context.Context,
	pdagrf PostDriveAndGetRootFolderer,
	caches *restoreCaches,
	drivePath *path.DrivePath,
	protectedResourceID, fallbackDriveName string,
) (driveInfo, error) {
	driveID := drivePath.DriveID

	// the drive might already be cached by ID.  it's okay
	// if the name has changed.  the ID is a better reference
	// anyway.
	if di, ok := caches.DriveIDToDriveInfo.Load(driveID); ok {
		return di, nil
	}

	var (
		newDriveName = fallbackDriveName
		newDrive     models.Driveable
		err          error
	)

	// if the drive wasn't found by ID, maybe we can find a
	// drive with the same name but different ID.
	// start by looking up the old drive's name
	oldName, ok := caches.BackupDriveIDName.NameOf(driveID)
	if ok {
		// check for drives that currently have the same name
		if di, ok := caches.DriveNameToDriveInfo.Load(oldName); ok {
			return di, nil
		}

		// if no current drives have the same name, we'll make
		// a new drive with that name.
		newDriveName = oldName
	}

	nextDriveName := newDriveName

	// For sharepoint, document libraries can collide by name with
	// item types beyond just drive.  Lists, for example, cannot share
	// names with document libraries (they're the same type, actually).
	// In those cases we need to rename the drive until we can create
	// one without a collision.
	for i := 1; ; i++ {
		ictx := clues.Add(ctx, "new_drive_name", clues.Hide(nextDriveName))

		newDrive, err = pdagrf.PostDrive(ictx, protectedResourceID, nextDriveName)
		if err != nil && !errors.Is(err, graph.ErrItemAlreadyExistsConflict) {
			return driveInfo{}, clues.Wrap(err, "creating new drive")
		}

		if err == nil {
			break
		}

		nextDriveName = fmt.Sprintf("%s %d", newDriveName, i)
	}

	if err := caches.AddDrive(ctx, newDrive, pdagrf); err != nil {
		return driveInfo{}, clues.Wrap(err, "adding drive to cache").OrNil()
	}

	di, _ := caches.DriveIDToDriveInfo.Load(ptr.Val(newDrive.GetId()))

	return di, nil
}
