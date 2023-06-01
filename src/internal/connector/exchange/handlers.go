package exchange

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// restore
// ---------------------------------------------------------------------------

type restoreHandler interface {
	itemRestorer
	containerAPI
	newContainerCache(userID string) graph.ContainerResolver
	formatRestoreDestination(
		destinationContainerName string,
		collectionFullPath path.Path,
	) *path.Builder
}

// runs the item restoration (ie: item creation) process
// for a single item, whose summary contents are held in
// the body property.
type itemRestorer interface {
	restore(
		ctx context.Context,
		body []byte,
		userID, destinationID string,
		errs *fault.Bus,
	) (*details.ExchangeInfo, error)
}

// runs the actual graph API post request.
type itemPoster[T any] interface {
	PostItem(
		ctx context.Context,
		userID, dirID string,
		body T,
	) (T, error)
}

// produces structs that interface with the graph/cache_container
// CachedContainer interface.
type containerAPI interface {
	// POSTs the creation of a new container
	CreateContainer(
		ctx context.Context,
		userID, containerName, parentContainerID string,
	) (graph.Container, error)

	// GETs a container by name.
	// if containerByNamer is nil, this functionality is not supported
	// and should be skipped by the caller.
	// normally, we'd alias the func directly.  The indirection here
	// is because not all types comply with GetContainerByName.
	containerSearcher() containerByNamer

	// returns either the provided value (assumed to be the root
	// folder for that cache tree), or the default root container
	// (if the category uses a root folder that exists above the
	// restore location path).
	orRootContainer(string) string
}

type containerByNamer interface {
	// searches for a container by name.
	GetContainerByName(
		ctx context.Context,
		userID, containerName string,
	) (graph.Container, error)
}

// primary interface controller for all per-cateogry restoration behavior.
func restoreHandlers(
	ac api.Client,
) map[path.CategoryType]restoreHandler {
	return map[path.CategoryType]restoreHandler{
		path.ContactsCategory: newContactRestoreHandler(ac),
		path.EmailCategory:    newMailRestoreHandler(ac),
		path.EventsCategory:   newEventRestoreHandler(ac),
	}
}
