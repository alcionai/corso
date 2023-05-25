package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"runtime/trace"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	// copyBufferSize is used for chunked upload
	// Microsoft recommends 5-10MB buffers
	// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession?view=graph-rest-1.0#best-practices
	copyBufferSize = 5 * 1024 * 1024

	// Maximum number of retries for upload failures
	maxUploadRetries = 3

	// Time interval to check for upload failures
	stallCheckInterval = 5 * time.Minute
)

type restoreCaches struct {
	Folders               *folderCache
	ParentDirToMeta       map[string]metadata.Metadata
	OldPermIDToNewID      map[string]string
	DriveIDToRootFolderID map[string]string
}

func NewRestoreCaches() *restoreCaches {
	return &restoreCaches{
		Folders:               NewFolderCache(),
		ParentDirToMeta:       map[string]metadata.Metadata{},
		OldPermIDToNewID:      map[string]string{},
		DriveIDToRootFolderID: map[string]string{},
	}
}

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	creds account.M365Config,
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
		caches         = NewRestoreCaches()
		el             = errs.Local()
	)

	ctx = clues.Add(
		ctx,
		"backup_version", backupVersion,
		"destination", dest.ContainerName)

	// Reorder collections so that the parents directories are created
	// before the child directories; a requirement for permissions.
	data.SortRestoreCollections(dcs)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err     error
			metrics support.CollectionMetrics
			ictx    = clues.Add(
				ctx,
				"category", dc.FullPath().Category(),
				"destination", clues.Hide(dest.ContainerName),
				"resource_owner", clues.Hide(dc.FullPath().ResourceOwner()),
				"full_path", dc.FullPath())
		)

		metrics, err = RestoreCollection(
			ictx,
			creds,
			backupVersion,
			service,
			dc,
			caches,
			OneDriveSource,
			dest.ContainerName,
			deets,
			opts.RestorePermissions,
			errs)
		if err != nil {
			el.AddRecoverable(err)
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
// - the updated metadata map that include metadata for folders in this collection
// - error, if any besides recoverable
func RestoreCollection(
	ctx context.Context,
	creds account.M365Config,
	backupVersion int,
	service graph.Servicer,
	dc data.RestoreCollection,
	caches *restoreCaches,
	source driveSource,
	restoreContainerName string,
	deets *details.Builder,
	restorePerms bool,
	errs *fault.Bus,
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

	if _, ok := caches.DriveIDToRootFolderID[drivePath.DriveID]; !ok {
		root, err := api.GetDriveRoot(ctx, service, drivePath.DriveID)
		if err != nil {
			return metrics, clues.Wrap(err, "getting drive root id")
		}

		caches.DriveIDToRootFolderID[drivePath.DriveID] = ptr.Val(root.GetId())
	}

	// Assemble folder hierarchy we're going to restore into (we recreate the folder hierarchy
	// from the backup under this the restore folder instead of root)
	// i.e. Restore into `<restoreContainerName>/<original folder path>`
	// the drive into which this folder gets restored is tracked separately in drivePath.
	restoreDir := path.Builder{}.Append(restoreContainerName).Append(drivePath.Folders...)

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
		backupVersion,
		restorePerms)
	if err != nil {
		return metrics, clues.Wrap(err, "getting permissions").WithClues(ctx)
	}

	// Create restore folders and get the folder ID of the folder the data stream will be restored in
	restoreFolderID, err := CreateRestoreFolders(
		ctx,
		creds,
		service,
		drivePath,
		restoreDir,
		dc.FullPath(),
		colMeta,
		caches,
		restorePerms)
	if err != nil {
		return metrics, clues.Wrap(err, "creating folders for restore")
	}

	caches.ParentDirToMeta[dc.FullPath().String()] = colMeta
	items := dc.Items(ctx, errs)

	semaphoreCh := make(chan struct{}, graph.Parallelism(path.OneDriveService).ItemUpload())
	defer close(semaphoreCh)

	deetsLock := sync.Mutex{}

	updateDeets := func(
		ctx context.Context,
		repoRef path.Path,
		locationRef *path.Builder,
		updated bool,
		info details.ItemInfo,
	) {
		deetsLock.Lock()
		defer deetsLock.Unlock()

		err = deets.Add(repoRef, locationRef, updated, info)
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
				complete = true
				break
			}

			wg.Add(1)
			semaphoreCh <- struct{}{}

			go func(ctx context.Context, itemData data.Stream) {
				defer wg.Done()
				defer func() { <-semaphoreCh }()

				// TODO(meain): Don't have to pass this in now that we create a
				// separate copyBuffer for each restore
				copyBuffer := make([]byte, copyBufferSize)

				ictx := clues.Add(ctx, "restore_item_id", itemData.UUID())

				itemPath, err := dc.FullPath().AppendItem(itemData.UUID())
				if err != nil {
					el.AddRecoverable(clues.Wrap(err, "appending item to full path").WithClues(ictx))
					return
				}

				itemInfo, skipped, err := restoreItem(
					ictx,
					creds,
					dc,
					backupVersion,
					source,
					service,
					drivePath,
					restoreFolderID,
					copyBuffer,
					caches,
					restorePerms,
					itemData,
					itemPath)

				// skipped items don't get counted, but they can error
				if !skipped {
					atomic.AddInt64(&metricsObjects, 1)
					atomic.AddInt64(&metricsBytes, int64(len(copyBuffer)))
				}

				if err != nil {
					el.AddRecoverable(clues.Wrap(err, "restoring item"))
					return
				}

				if skipped {
					logger.Ctx(ictx).With("item_path", itemPath).Debug("did not restore item")
					return
				}

				// TODO: implement locationRef
				updateDeets(ictx, itemPath, &path.Builder{}, true, itemInfo)

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
	creds account.M365Config,
	dc data.RestoreCollection,
	backupVersion int,
	source driveSource,
	service graph.Servicer,
	drivePath *path.DrivePath,
	restoreFolderID string,
	copyBuffer []byte,
	caches *restoreCaches,
	restorePerms bool,
	itemData data.Stream,
	itemPath path.Path,
) (details.ItemInfo, bool, error) {
	itemUUID := itemData.UUID()
	ctx = clues.Add(ctx, "item_id", itemUUID)

	if backupVersion < version.OneDrive1DataAndMetaFiles {
		itemInfo, err := restoreV0File(
			ctx,
			source,
			service,
			drivePath,
			dc,
			restoreFolderID,
			copyBuffer,
			itemData)
		if err != nil {
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
		if !restorePerms || backupVersion >= version.OneDrive4DirIncludesPermissions {
			return details.ItemInfo{}, true, nil
		}

		metaReader := itemData.ToReader()
		defer metaReader.Close()

		meta, err := getMetadata(metaReader)
		if err != nil {
			return details.ItemInfo{}, true, clues.Wrap(err, "getting directory metadata").WithClues(ctx)
		}

		trimmedPath := strings.TrimSuffix(itemPath.String(), metadata.DirMetaFileSuffix)
		caches.ParentDirToMeta[trimmedPath] = meta

		return details.ItemInfo{}, true, nil
	}

	// only items with DataFileSuffix from this point on

	if backupVersion < version.OneDrive6NameInMeta {
		itemInfo, err := restoreV1File(
			ctx,
			source,
			creds,
			service,
			drivePath,
			dc,
			restoreFolderID,
			copyBuffer,
			restorePerms,
			caches,
			itemPath,
			itemData)
		if err != nil {
			return details.ItemInfo{}, false, clues.Wrap(err, "v1 restore")
		}

		return itemInfo, false, nil
	}

	// only v6+ backups from this point on

	itemInfo, err := restoreV6File(
		ctx,
		source,
		creds,
		service,
		drivePath,
		dc,
		restoreFolderID,
		copyBuffer,
		restorePerms,
		caches,
		itemPath,
		itemData)
	if err != nil {
		return details.ItemInfo{}, false, clues.Wrap(err, "v6 restore")
	}

	return itemInfo, false, nil
}

func restoreV0File(
	ctx context.Context,
	source driveSource,
	service graph.Servicer,
	drivePath *path.DrivePath,
	fetcher fileFetcher,
	restoreFolderID string,
	copyBuffer []byte,
	itemData data.Stream,
) (details.ItemInfo, error) {
	_, itemInfo, err := restoreData(
		ctx,
		service,
		fetcher,
		itemData.UUID(),
		itemData,
		drivePath.DriveID,
		restoreFolderID,
		copyBuffer,
		source)
	if err != nil {
		return itemInfo, clues.Wrap(err, "restoring file")
	}

	return itemInfo, nil
}

type fileFetcher interface {
	Fetch(ctx context.Context, name string) (data.Stream, error)
}

func restoreV1File(
	ctx context.Context,
	source driveSource,
	creds account.M365Config,
	service graph.Servicer,
	drivePath *path.DrivePath,
	fetcher fileFetcher,
	restoreFolderID string,
	copyBuffer []byte,
	restorePerms bool,
	caches *restoreCaches,
	itemPath path.Path,
	itemData data.Stream,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.UUID(), metadata.DataFileSuffix)

	itemID, itemInfo, err := restoreData(
		ctx,
		service,
		fetcher,
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
	metaName := trimmedName + metadata.MetaFileSuffix

	meta, err := fetchAndReadMetadata(ctx, fetcher, metaName)
	if err != nil {
		return details.ItemInfo{}, clues.Wrap(err, "restoring file")
	}

	err = RestorePermissions(
		ctx,
		creds,
		service,
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
	source driveSource,
	creds account.M365Config,
	service graph.Servicer,
	drivePath *path.DrivePath,
	fetcher fileFetcher,
	restoreFolderID string,
	copyBuffer []byte,
	restorePerms bool,
	caches *restoreCaches,
	itemPath path.Path,
	itemData data.Stream,
) (details.ItemInfo, error) {
	trimmedName := strings.TrimSuffix(itemData.UUID(), metadata.DataFileSuffix)

	// Get metadata file so we can determine the file name.
	metaName := trimmedName + metadata.MetaFileSuffix

	meta, err := fetchAndReadMetadata(ctx, fetcher, metaName)
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

	itemID, itemInfo, err := restoreData(
		ctx,
		service,
		fetcher,
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

	err = RestorePermissions(
		ctx,
		creds,
		service,
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
	creds account.M365Config,
	service graph.Servicer,
	drivePath *path.DrivePath,
	restoreDir *path.Builder,
	folderPath path.Path,
	folderMetadata metadata.Metadata,
	caches *restoreCaches,
	restorePerms bool,
) (string, error) {
	id, err := createRestoreFolders(
		ctx,
		service,
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
		creds,
		service,
		drivePath.DriveID,
		id,
		folderPath,
		folderMetadata,
		caches)

	return id, err
}

// createRestoreFolders creates the restore folder hierarchy in the specified
// drive and returns the folder ID of the last folder entry in the hierarchy.
// folderCache is mutated, as a side effect of populating the items.
func createRestoreFolders(
	ctx context.Context,
	service graph.Servicer,
	drivePath *path.DrivePath,
	restoreDir *path.Builder,
	caches *restoreCaches,
) (string, error) {
	var (
		driveID        = drivePath.DriveID
		folders        = restoreDir.Elements()
		location       = path.Builder{}.Append(driveID)
		parentFolderID = caches.DriveIDToRootFolderID[drivePath.DriveID]
	)

	ctx = clues.Add(
		ctx,
		"drive_id", drivePath.DriveID,
		"root_folder_id", parentFolderID)

	for _, folder := range folders {
		location = location.Append(folder)
		ictx := clues.Add(
			ctx,
			"creating_restore_folder", folder,
			"restore_folder_location", location,
			"parent_of_restore_folder", parentFolderID)

		if fl, ok := caches.Folders.get(location); ok {
			parentFolderID = ptr.Val(fl.GetId())
			// folder was already created, move on to the child
			continue
		}

		folderItem, err := api.GetFolderByName(ictx, service, driveID, parentFolderID, folder)
		if err != nil && !errors.Is(err, api.ErrFolderNotFound) {
			return "", clues.Wrap(err, "getting folder by display name")
		}

		// folder found, moving to next child
		if err == nil {
			parentFolderID = ptr.Val(folderItem.GetId())
			caches.Folders.set(location, folderItem)

			continue
		}

		// create the folder if not found
		folderItem, err = CreateItem(ictx, service, driveID, parentFolderID, newItem(folder, true))
		if err != nil {
			return "", clues.Wrap(err, "creating folder")
		}

		parentFolderID = ptr.Val(folderItem.GetId())
		caches.Folders.set(location, folderItem)

		logger.Ctx(ictx).Debug("resolved restore destination")
	}

	return parentFolderID, nil
}

func copyBufferWithStallCheck(dst io.Writer, src io.Reader, buffer []byte, stallTimeout time.Duration) (int64, error) {
	timer := time.NewTimer(stallTimeout)
	defer timer.Stop()

	type ce struct {
		complete bool
		err      error
	}

	totalCopied := int64(0)
	cont := make(chan ce)

	for {
		go func() {
			n, rerr := src.Read(buffer)
			if rerr != nil && rerr != io.EOF {
				cont <- ce{err: clues.Wrap(rerr, "reading data")}
			}

			_, werr := dst.Write(buffer[:n])
			if werr != nil && werr != io.EOF {
				cont <- ce{err: clues.Wrap(werr, "writing data")}
			}

			totalCopied += int64(n)

			if rerr == io.EOF {
				// Copy completed successfully
				cont <- ce{complete: true}
			}

			cont <- ce{complete: false}
		}()

		timer.Reset(stallTimeout)

		select {
		case <-timer.C:
			return 0, clues.New("copy stalled")
		case in := <-cont:
			if in.err != nil {
				return 0, in.err
			}

			if in.complete {
				return totalCopied, nil
			}

			// Continue copying
			continue
		}
	}
}

// restoreData will create a new item in the specified `parentFolderID` and upload the data.Stream
func restoreData(
	ctx context.Context,
	service graph.Servicer,
	fetcher fileFetcher,
	name string,
	itemData data.Stream,
	driveID, parentFolderID string,
	copyBuffer []byte,
	source driveSource,
) (string, details.ItemInfo, error) {
	ctx, end := diagnostics.Span(ctx, "gc:oneDrive:restoreItem", diagnostics.Label("item_uuid", itemData.UUID()))
	defer end()

	trace.Log(ctx, "gc:oneDrive:restoreItem", itemData.UUID())

	// Get the stream size (needed to create the upload session)
	ss, ok := itemData.(data.StreamSize)
	if !ok {
		return "", details.ItemInfo{}, clues.New("item does not implement DataStreamInfo").WithClues(ctx)
	}

	// Create Item
	newItem, err := CreateItem(ctx, service, driveID, parentFolderID, newItem(name, false))
	if err != nil {
		return "", details.ItemInfo{}, err
	}

	var (
		w       io.Writer
		written int64
	)

	// This is just to retry file upload, the uploadSession creation is not retried here
	// We need extra logic to retry file upload as we have to pull the file again from kopia
	for i := 0; i <= maxUploadRetries; i++ {
		// Get a drive item writer
		w, err = driveItemWriter(ctx, service, driveID, ptr.Val(newItem.GetId()), ss.Size())
		if err != nil {
			return "", details.ItemInfo{}, err
		}

		pname := name
		iReader := itemData.ToReader()

		if i > 0 {
			pname = fmt.Sprintf("%s (retry %d)", name, i)

			// If it is not the first try, we have to pull the file
			// again from kopia. Ideally we could just seek the stream
			// but we don't have a Seeker available here.
			itemData, err := fetcher.Fetch(ctx, itemData.UUID())
			if err != nil {
				return "", details.ItemInfo{}, clues.Wrap(err, "getting data file")
			}

			iReader = itemData.ToReader()
		}

		progReader, closer := observe.ItemProgress(
			ctx,
			iReader,
			observe.ItemRestoreMsg,
			clues.Hide(pname),
			ss.Size())

		go closer()

		// Upload the stream data
		written, err = copyBufferWithStallCheck(w, progReader, copyBuffer, stallCheckInterval)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", details.ItemInfo{}, clues.Wrap(err, "uploading file")
	}

	dii := details.ItemInfo{}

	switch source {
	case SharePointSource:
		dii.SharePoint = sharePointItemInfo(newItem, written)
	default:
		dii.OneDrive = oneDriveItemInfo(newItem, written)
	}

	return ptr.Val(newItem.GetId()), dii, nil
}

func fetchAndReadMetadata(
	ctx context.Context,
	fetcher fileFetcher,
	metaName string,
) (metadata.Metadata, error) {
	ctx = clues.Add(ctx, "meta_file_name", metaName)

	metaFile, err := fetcher.Fetch(ctx, metaName)
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

// Augment restore path to add extra files(meta) needed for restore as
// well as do any other ordering operations on the paths
//
// Only accepts StoragePath/RestorePath pairs where the RestorePath is
// at least as long as the StoragePath. If the RestorePath is longer than the
// StoragePath then the first few (closest to the root) directories will use
// default permissions during restore.
func AugmentRestorePaths(
	backupVersion int,
	paths []path.RestorePaths,
) ([]path.RestorePaths, error) {
	// Keyed by each value's StoragePath.String() which corresponds to the RepoRef
	// of the directory.
	colPaths := map[string]path.RestorePaths{}

	for _, p := range paths {
		first := true

		for {
			sp, err := p.StoragePath.Dir()
			if err != nil {
				return nil, err
			}

			drivePath, err := path.ToDrivePath(sp)
			if err != nil {
				return nil, err
			}

			if len(drivePath.Folders) == 0 {
				break
			}

			if len(p.RestorePath.Elements()) < len(sp.Elements()) {
				return nil, clues.New("restorePath shorter than storagePath").
					With("restore_path", p.RestorePath, "storage_path", sp)
			}

			rp := p.RestorePath

			// Make sure the RestorePath always points to the level of the current
			// collection. We need to track if it's the first iteration because the
			// RestorePath starts out at the collection level to begin with.
			if !first {
				rp, err = p.RestorePath.Dir()
				if err != nil {
					return nil, err
				}
			}

			paths := path.RestorePaths{
				StoragePath: sp,
				RestorePath: rp,
			}

			colPaths[sp.String()] = paths
			p = paths
			first = false
		}
	}

	// Adds dirmeta files as we need to make sure collections for all
	// directories involved are created and not just the final one. No
	// need to add `.meta` files (metadata for files) as they will
	// anyways be looked up automatically.
	// TODO: Stop populating .dirmeta for newer versions once we can
	// get files from parent directory via `Fetch` in a collection.
	// As of now look up metadata for parent directories from a
	// collection.
	for _, p := range colPaths {
		el := p.StoragePath.Elements()

		if backupVersion >= version.OneDrive6NameInMeta {
			mPath, err := p.StoragePath.AppendItem(".dirmeta")
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: p.RestorePath})
		} else if backupVersion >= version.OneDrive4DirIncludesPermissions {
			mPath, err := p.StoragePath.AppendItem(el.Last() + ".dirmeta")
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: p.RestorePath})
		} else if backupVersion >= version.OneDrive1DataAndMetaFiles {
			pp, err := p.StoragePath.Dir()
			if err != nil {
				return nil, err
			}

			mPath, err := pp.AppendItem(el.Last() + ".dirmeta")
			if err != nil {
				return nil, err
			}

			prp, err := p.RestorePath.Dir()
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: prp})
		}
	}

	// This sort is done primarily to order `.meta` files after `.data`
	// files. This is only a necessity for OneDrive as we are storing
	// metadata for files/folders in separate meta files and we the
	// data to be restored before we can restore the metadata.
	//
	// This sorting assumes stuff in the same StoragePath directory end up in the
	// same RestorePath collection.
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].StoragePath.String() < paths[j].StoragePath.String()
	})

	return paths, nil
}
