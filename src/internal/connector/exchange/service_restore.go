package exchange

import (
	"bytes"
	"context"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type restoreHandler interface {
	itemRestorer
	containerCreator
}

type itemRestorer interface {
	restore(
		ctx context.Context,
		body []byte,
		destinationID string,
		errs *fault.Bus,
	) (*details.ExchangeInfo, error)
}

type itemPoster[T any] interface {
	PostItem(
		ctx context.Context,
		userID, dirID string,
		body T,
	) (T, error)
}

type containerCreator interface {
	newContainerCache() graph.ContainerResolver
	CreateContainer(
		ctx context.Context,
		userID, containerName, parentContainerID string,
	) (graph.Container, error)

	GetContainerByName(
		ctx context.Context,
		userID, containerName string,
	) (graph.Container, error)
	// TODO: remove when all handlers support GetContainerByName
	// as a create-collision fallback
	CanGetContainerByName() bool
}

// RestoreCollections restores M365 objects in data.RestoreCollection to MSFT
// store through GraphAPI.
func RestoreCollections(
	ctx context.Context,
	creds account.M365Config,
	ac api.Client,
	gs graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
) (*support.ConnectorOperationStatus, error) {
	if len(dcs) == 0 {
		return support.CreateStatus(ctx, support.Restore, 0, support.CollectionMetrics{}, ""), nil
	}

	var (
		userID         = dcs[0].FullPath().ResourceOwner()
		directoryCache = make(map[path.CategoryType]graph.ContainerResolver)
		handlers       = map[path.CategoryType]restoreHandler{
			path.ContactsCategory: newContactRestoreHandler(ac, userID),
			path.EmailCategory:    newMailRestoreHandler(ac, userID),
			path.EventsCategory:   newEventRestoreHandler(ac, userID),
		}
		metrics support.CollectionMetrics
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
		el     = errs.Local()
	)

	ctx = clues.Add(ctx, "resource_owner", clues.Hide(userID))

	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			category = dc.FullPath().Category()
			ictx     = clues.Add(
				ctx,
				"restore_category", category,
				"restore_full_path", dc.FullPath())
		)

		handler, ok := handlers[category]
		if !ok {
			el.AddRecoverable(clues.New("unsupported restore path category").WithClues(ictx))
		}

		containerID, dcc, err := createDestination(
			ictx,
			handler,
			dc.FullPath(),
			userID,
			dest.ContainerName,
			directoryCache[category],
			errs)
		if err != nil {
			el.AddRecoverable(err)
			continue
		}

		directoryCache[category] = dcc

		ictx = clues.Add(ictx, "restore_destination_id", containerID)

		temp, canceled := restoreCollection(
			ictx,
			handler,
			dc,
			userID,
			containerID,
			policy,
			deets,
			errs)

		metrics = support.CombineMetrics(metrics, temp)

		if canceled {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		metrics,
		dest.ContainerName)

	return status, el.Failure()
}

// restoreCollection handles restoration of an individual collection.
func restoreCollection(
	ctx context.Context,
	ir itemRestorer,
	dc data.RestoreCollection,
	userID, destinationID string,
	policy control.CollisionPolicy,
	deets *details.Builder,
	errs *fault.Bus,
) (support.CollectionMetrics, bool) {
	ctx, end := diagnostics.Span(ctx, "gc:exchange:restoreCollection", diagnostics.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics  support.CollectionMetrics
		items    = dc.Items(ctx, errs)
		fullPath = dc.FullPath()
		category = fullPath.Category()
	)

	colProgress, closer := observe.CollectionProgress(
		ctx,
		category.String(),
		clues.Hide(fullPath.Folder(false)))
	defer closer()
	defer close(colProgress)

	for {
		select {
		case <-ctx.Done():
			errs.AddRecoverable(clues.Wrap(ctx.Err(), "context cancelled").WithClues(ctx))
			return metrics, true

		case itemData, ok := <-items:
			if !ok || errs.Failure() != nil {
				return metrics, false
			}

			ictx := clues.Add(ctx, "item_id", itemData.UUID())
			trace.Log(ictx, "gc:exchange:restoreCollection:item", itemData.UUID())
			metrics.Objects++

			buf := &bytes.Buffer{}

			_, err := buf.ReadFrom(itemData.ToReader())
			if err != nil {
				errs.AddRecoverable(clues.Wrap(err, "reading item bytes").WithClues(ictx))
				continue
			}

			body := buf.Bytes()

			info, err := ir.restore(ictx, body, destinationID, errs)
			if err != nil {
				errs.AddRecoverable(err)
				continue
			}

			metrics.Bytes += int64(len(body))
			metrics.Successes++

			// FIXME: this may be the incorrect path.  If we restored within a top-level
			// destination folder, then the restore path no longer matches the fullPath.
			itemPath, err := fullPath.AppendItem(itemData.UUID())
			if err != nil {
				errs.AddRecoverable(clues.Wrap(err, "adding item to collection path").WithClues(ctx))
				continue
			}

			locationRef := path.Builder{}.Append(itemPath.Folders()...)

			err = deets.Add(
				itemPath,
				locationRef,
				true,
				details.ItemInfo{
					Exchange: info,
				})
			if err != nil {
				// These deets additions are for cli display purposes only.
				// no need to fail out on error.
				logger.Ctx(ctx).Infow("accounting for restored item", "error", err)
			}

			colProgress <- struct{}{}
		}
	}
}

// createDestination creates folders in sequence
// [root leaf1 leaf2] similar to a linked list.
// @param directory is the desired path from the root to the container
// that the items will be restored into.
// @param destinationName will be prepended onto the directory path
// for non-destructive restores.
func createDestination(
	ctx context.Context,
	cc containerCreator,
	directory path.Path,
	userID, destinationName string,
	gcr graph.ContainerResolver,
	errs *fault.Bus,
) (string, graph.ContainerResolver, error) {
	var (
		containers  = append([]string{destinationName}, directory.Folders()...)
		cache       = gcr
		containerID = rootFolderAlias
		pb          = &path.Builder{}
		isNewCache  bool
	)

	if gcr == nil {
		cache = cc.newContainerCache()
		isNewCache = true
	}

	ctx = clues.Add(ctx, "is_new_cache", isNewCache)

	for _, container := range containers {
		pb = pb.Append(container)

		fid, err := getOrPopulateContainer(
			ctx,
			cc,
			cache,
			pb,
			userID,
			container,
			containerID,
			isNewCache,
			errs)
		if err != nil {
			return "", cache, clues.Stack(err)
		}

		containerID = fid
	}

	return containerID, cache, nil
}

func getOrPopulateContainer(
	ctx context.Context,
	cc containerCreator,
	gcr graph.ContainerResolver,
	pb *path.Builder,
	userID, containerParentID, containerName string,
	isNewCache bool,
	errs *fault.Bus,
) (string, error) {
	cached, ok := gcr.LocationInCache(pb.String())
	if ok {
		return cached, nil
	}

	c, err := cc.CreateContainer(ctx, userID, containerName, containerParentID)

	// 409 handling case:
	// attempt to fetch the container by name and add that result to the cache.
	// This is rare, but may happen if CreateContainer() POST fails with 5xx:
	// sometimes the backend will create the folder despite the 5xx response,
	// leaving our local containerResolver with inconsistent state.
	if graph.IsErrFolderExists(err) && cc.CanGetContainerByName() {
		c, err = cc.GetContainerByName(ctx, userID, containerName)
	}

	if err != nil {
		return "", err
	}

	folderID := ptr.Val(c.GetId())

	// Only populate the cache if we actually had to create it. Since we set
	// newCache to false in this we'll only try to populate it once per function
	// call even if we make a new cache.
	if isNewCache {
		if err := gcr.Populate(ctx, errs, rootFolderAlias); err != nil {
			return "", clues.Wrap(err, "populating container cache")
		}
	}

	if err = gcr.AddToCache(ctx, c); err != nil {
		return "", clues.Wrap(err, "adding container to cache")
	}

	return folderID, nil
}

func uploadAttachments(
	ctx context.Context,
	ap attachmentPoster,
	as []models.Attachmentable,
	userID, destinationID, itemID string,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for _, a := range as {
		if el.Failure() != nil {
			return el.Failure()
		}

		err := uploadAttachment(
			ctx,
			ap,
			userID,
			destinationID,
			itemID,
			a)
		if err != nil {
			// FIXME: I don't know why we're swallowing this error case.
			// It needs investigation: https://github.com/alcionai/corso/issues/3498
			if ptr.Val(a.GetOdataType()) == "#microsoft.graph.itemAttachment" {
				name := ptr.Val(a.GetName())

				logger.CtxErr(ctx, err).
					With("attachment_name", name).
					Info("mail upload failed")

				continue
			}

			el.AddRecoverable(clues.Wrap(err, "uploading mail attachment").WithClues(ctx))
		}
	}

	return el.Failure()
}
