package teamschats

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// itemer standardizes common behavior that can be expected from all
// items within a chats collection backup.
type chatsItemer interface {
	serialization.Parsable
	graph.GetIDer
	graph.GetLastUpdatedDateTimer
}

type backupHandler[I chatsItemer] interface {
	getContainerer[I]
	getItemAndAugmentInfoer[I]
	getItemer[I]
	getItemIDser[I]
	includeItemer[I]
	canonicalPather
}

// gets the container for the resource
// within this handler set, only one container (the root)
// is expected
type getContainerer[I chatsItemer] interface {
	getContainer(
		ctx context.Context,
		cc api.CallConfig,
	) (container[I], error)
}

type getItemAndAugmentInfoer[I chatsItemer] interface {
	getItemer[I]
	augmentItemInfoer[I]
}

type augmentItemInfoer[I chatsItemer] interface {
	// augmentItemInfo completes the teamChatsInfo population with any data
	// owned by the container and not accessible to the item.
	augmentItemInfo(*details.TeamsChatsInfo, I)
}

// gets all item IDs in the container
type getItemIDser[I chatsItemer] interface {
	getItemIDs(
		ctx context.Context,
		cc api.CallConfig,
	) ([]I, error)
}

type getItemer[I chatsItemer] interface {
	getItem(
		ctx context.Context,
		protectedResource string,
		itemID string,
	) (I, *details.TeamsChatsInfo, error)
}

// includeItemer evaluates whether the item is included
// in the provided scope.
type includeItemer[I chatsItemer] interface {
	includeItem(
		i I,
		scope selectors.TeamsChatsScope,
	) bool
}

// canonicalPath constructs the service and category specific path for
// the given builder.  The tenantID and protectedResourceID are assumed
// to be stored in the handler already.
type canonicalPather interface {
	canonicalPath() (path.Path, error)
}

// ---------------------------------------------------------------------------
// Container management
// ---------------------------------------------------------------------------

type container[I chatsItemer] struct {
	storageDirFolders path.Elements
	humanLocation     path.Elements
	container         I
}
