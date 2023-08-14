package exchange

import (
	"bytes"
	"context"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// RestoreCollection handles restoration of an individual collection.
func RestoreCollection(
	ctx context.Context,
	ir itemRestorer,
	dc data.RestoreCollection,
	resourceID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	deets *details.Builder,
	errs *fault.Bus,
	ctr *count.Bus,
) (support.CollectionMetrics, error) {
	ctx, end := diagnostics.Span(ctx, "m365:exchange:restoreCollection", diagnostics.Label("path", dc.FullPath()))
	defer end()

	var (
		el       = errs.Local()
		metrics  support.CollectionMetrics
		items    = dc.Items(ctx, errs)
		fullPath = dc.FullPath()
		category = fullPath.Category()
	)

	colProgress := observe.CollectionProgress(
		ctx,
		category.String(),
		fullPath.Folder(false))
	defer close(colProgress)

	for {
		select {
		case <-ctx.Done():
			return metrics, clues.Wrap(ctx.Err(), "context cancelled").WithClues(ctx)

		case itemData, ok := <-items:
			if !ok || el.Failure() != nil {
				return metrics, el.Failure()
			}

			ictx := clues.Add(ctx, "item_id", itemData.ID())
			trace.Log(ictx, "m365:exchange:restoreCollection:item", itemData.ID())
			metrics.Objects++

			buf := &bytes.Buffer{}

			_, err := buf.ReadFrom(itemData.ToReader())
			if err != nil {
				el.AddRecoverable(ctx, clues.Wrap(err, "reading item bytes").WithClues(ictx))
				continue
			}

			body := buf.Bytes()

			info, err := ir.restore(
				ictx,
				body,
				resourceID,
				destinationID,
				collisionKeyToItemID,
				collisionPolicy,
				errs,
				ctr)
			if err != nil {
				if !graph.IsErrItemAlreadyExistsConflict(err) {
					el.AddRecoverable(ictx, err)
				}

				continue
			}

			metrics.Bytes += int64(len(body))
			metrics.Successes++

			// FIXME: this may be the incorrect path.  If we restored within a top-level
			// destination folder, then the restore path no longer matches the fullPath.
			itemPath, err := fullPath.AppendItem(itemData.ID())
			if err != nil {
				el.AddRecoverable(ctx, clues.Wrap(err, "adding item to collection path").WithClues(ctx))
				continue
			}

			locationRef := path.Builder{}.Append(itemPath.Folders()...)

			err = deets.Add(
				itemPath,
				locationRef,
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

// CreateDestination creates folders in sequence
// [root leaf1 leaf2] similar to a linked list.
// @param directory is the desired path from the root to the container
// that the items will be restored into.
func CreateDestination(
	ctx context.Context,
	ca containerAPI,
	destination *path.Builder,
	resourceID string,
	gcr graph.ContainerResolver,
	errs *fault.Bus,
) (string, graph.ContainerResolver, error) {
	var (
		cache             = gcr
		restoreLoc        = &path.Builder{}
		containerParentID string
	)

	for _, container := range destination.Elements() {
		restoreLoc = restoreLoc.Append(container)

		ictx := clues.Add(
			ctx,
			"container_parent_id", containerParentID,
			"container_name", container,
			"restore_location", restoreLoc)

		containerID, err := getOrPopulateContainer(
			ictx,
			ca,
			cache,
			restoreLoc,
			resourceID,
			containerParentID,
			container,
			errs)
		if err != nil {
			return "", cache, clues.Stack(err)
		}

		containerParentID = containerID
	}

	// containerParentID now identifies the last created container,
	// not its parent.
	return containerParentID, cache, nil
}

func getOrPopulateContainer(
	ctx context.Context,
	ca containerAPI,
	gcr graph.ContainerResolver,
	restoreLoc *path.Builder,
	resourceID, containerParentID, containerName string,
	errs *fault.Bus,
) (string, error) {
	cached, ok := gcr.LocationInCache(restoreLoc.String())
	if ok {
		return cached, nil
	}

	c, err := ca.CreateContainer(ctx, resourceID, containerParentID, containerName)

	// 409 handling case:
	// attempt to fetch the container by name and add that result to the cache.
	// This is rare, but may happen if CreateContainer() POST fails with 5xx:
	// sometimes the backend will create the folder despite the 5xx response,
	// leaving our local containerResolver with inconsistent state.
	if graph.IsErrFolderExists(err) {
		cc, e := ca.GetContainerByName(ctx, resourceID, containerParentID, containerName)
		if e != nil {
			err = clues.Stack(err, e)
		} else {
			c = cc
			err = nil
		}
	}

	if err != nil {
		return "", clues.Wrap(err, "creating restore container")
	}

	folderID := ptr.Val(c.GetId())

	if err = gcr.AddToCache(ctx, c); err != nil {
		return "", clues.Wrap(err, "adding container to cache")
	}

	return folderID, nil
}

func uploadAttachments(
	ctx context.Context,
	ap attachmentPoster,
	as []models.Attachmentable,
	resourceID, destinationID, itemID string,
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
			resourceID,
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

			el.AddRecoverable(ctx, clues.Wrap(err, "uploading mail attachment").WithClues(ctx))
		}
	}

	return el.Failure()
}
