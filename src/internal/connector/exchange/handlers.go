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
	containerCacheHandler
}

type itemRestorer interface {
	restore(
		ctx context.Context,
		body []byte,
		userID, destinationID string,
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

type containerCacheHandler interface {
	newContainerCache(userID string) graph.ContainerResolver
	containerFactory() containerCreator
	containerSearcher() (containerByNamer, bool)
}

type containerCreator interface {
	CreateContainer(
		ctx context.Context,
		userID, containerName, parentContainerID string,
	) (graph.Container, error)
}

type containerByNamer interface {
	GetContainerByName(
		ctx context.Context,
		userID, containerName string,
	) (graph.Container, error)
}

func restoreHandlers(
	ac api.Client,
) map[path.CategoryType]restoreHandler {
	return map[path.CategoryType]restoreHandler{
		path.ContactsCategory: newContactRestoreHandler(ac),
		path.EmailCategory:    newMailRestoreHandler(ac),
		path.EventsCategory:   newEventRestoreHandler(ac),
	}
}
