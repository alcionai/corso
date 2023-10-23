package groups

import (
	"context"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type backupHandler interface {
	getItemByIDer

	// gets all containers for the resource
	getContainers(
		ctx context.Context,
	) ([]models.Channelable, error)

	// gets all item IDs (by delta, if possible) in the container
	getContainerItemIDs(
		ctx context.Context,
		containerID, prevDelta string,
		canMakeDeltaQueries bool,
	) (map[string]time.Time, bool, []string, api.DeltaUpdate, error)

	// includeContainer evaluates whether the container is included
	// in the provided scope.
	includeContainer(
		ctx context.Context,
		qp graph.QueryParams,
		ch models.Channelable,
		scope selectors.GroupsScope,
	) bool

	// canonicalPath constructs the service and category specific path for
	// the given builder.
	canonicalPath(
		folders *path.Builder,
		tenantID string,
	) (path.Path, error)
}

type getItemByIDer interface {
	GetItemByID(
		ctx context.Context,
		resourceID, containerID, itemID string,
	) (models.ChatMessageable, *details.GroupsInfo, error)
}
