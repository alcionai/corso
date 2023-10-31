package groups

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// itemer standardizes common behavior that can be expected from all
// items within a groups collection backup.
type groupsItemer interface {
	serialization.Parsable
	graph.GetIDer
	graph.GetLastModifiedDateTimer
}

type backupHandler[C graph.GetIDer, I groupsItemer] interface {
	getItemer[I]

	// gets all containers for the resource
	getContainers(
		ctx context.Context,
		cc api.CallConfig,
	) ([]container[C], error)

	// gets all item IDs (by delta, if possible) in the container
	getContainerItemIDs(
		ctx context.Context,
		containerPath path.Elements,
		prevDelta string,
		cc api.CallConfig,
	) (pagers.AddedAndRemoved, error)

	// includeContainer evaluates whether the container is included
	// in the provided scope.
	includeContainer(
		ctx context.Context,
		qp graph.QueryParams,
		c C,
		scope selectors.GroupsScope,
	) bool

	// canonicalPath constructs the service and category specific path for
	// the given builder.
	canonicalPath(
		folders path.Elements,
		tenantID string,
	) (path.Path, error)

	// canMakeDeltaQueries evaluates whether the data type can support a
	// delta query enumeration.
	canMakeDeltaQueries() bool
}

type getItemer[I groupsItemer] interface {
	GetItem(
		ctx context.Context,
		protectedResource string,
		containerIDs path.Elements,
		itemID string,
	) (I, *details.GroupsInfo, error)
}

// ---------------------------------------------------------------------------
// Container management
// ---------------------------------------------------------------------------

type container[C graph.GetIDer] struct {
	folders             path.Elements
	location            path.Elements
	canMakeDeltaQueries bool
	container           C
}
