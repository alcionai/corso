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
	getContainerser[C]
	getContainerItemIDser
	getItemAndAugmentInfoer[C, I]
	includeContainerer[C]
	canonicalPather
	canMakeDeltaQuerieser
}

type getItemAndAugmentInfoer[C graph.GetIDer, I groupsItemer] interface {
	getItemer[I]
	augmentItemInfoer[C]
}

type augmentItemInfoer[C graph.GetIDer] interface {
	// augmentItemInfo completes the groupInfo population with any data
	// owned by the container and not accessible to the item.
	augmentItemInfo(*details.GroupsInfo, C)
}

type getItemer[I groupsItemer] interface {
	getItem(
		ctx context.Context,
		protectedResource string,
		containerIDs path.Elements,
		itemID string,
	) (I, *details.GroupsInfo, error)
}

// gets all containers for the resource
type getContainerser[C graph.GetIDer] interface {
	getContainers(
		ctx context.Context,
		cc api.CallConfig,
	) ([]container[C], error)
}

// gets all item IDs (by delta, if possible) in the container
type getContainerItemIDser interface {
	getContainerItemIDs(
		ctx context.Context,
		containerPath path.Elements,
		prevDelta string,
		cc api.CallConfig,
	) (pagers.AddedAndRemoved, error)
}

// includeContainer evaluates whether the container is included
// in the provided scope.
type includeContainerer[C graph.GetIDer] interface {
	includeContainer(
		c C,
		scope selectors.GroupsScope,
	) bool
}

// canonicalPath constructs the service and category specific path for
// the given builder.
type canonicalPather interface {
	canonicalPath(storageDir path.Elements) (path.Path, error)
}

// canMakeDeltaQueries evaluates whether the handler can support a
// delta query when enumerating its items.
type canMakeDeltaQuerieser interface {
	canMakeDeltaQueries() bool
}

// ---------------------------------------------------------------------------
// Container management
// ---------------------------------------------------------------------------

type container[C graph.GetIDer] struct {
	storageDirFolders   path.Elements
	humanLocation       path.Elements
	canMakeDeltaQueries bool
	container           C
}
