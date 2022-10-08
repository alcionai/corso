package graph

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type QueryParams struct {
	User        string
	Scope       selectors.ExchangeScope
	Credentials account.M365Config
	FailFast    bool
}

type Service interface {
	// Client() returns msgraph Service client that can be used to process and execute
	// the majority of the queries to the M365 Backstore
	Client() *msgraphsdk.GraphServiceClient
	// Adapter() returns GraphRequest adapter used to process large requests, create batches
	// and page iterators
	Adapter() *msgraphsdk.GraphRequestAdapter
	// ErrPolicy returns if the service is implementing a Fast-Fail policy or not
	ErrPolicy() bool
}

// descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type Descendable interface {
	GetId() *string
	GetParentFolderId() *string
}

// displayable represents objects that implement msgraph-sdk-fo/models.entityable
// and have the concept of a display name.
type Displayable interface {
	GetId() *string
	GetDisplayName() *string
}

// container is an interface that implements both the descendable and displayble interface.
type Container interface {
	Descendable
	Displayable
}

// ContainerResolver houses functions for getting information about containers
// from remote APIs (i.e. resolve folder paths with Graph API). Resolvers may
// cache information about containers.
type ContainerResolver interface {
	// IDToPath takes an m365 container ID and converts it to a hierarchical path
	// to that container. The path has a similar format to paths on the local
	// file system.
	IDToPath(ctx context.Context, m365ID string) (*path.Builder, error)
	// Populate performs initialization steps for the resolver
	// @param ctx is necessary param for Graph API tracing
	// @param baseFolderID represents the M365ID base that the resolver will
	// conclude its search.
	// @param baseFolderPath is the set of path elements of the baseFolder.
	Populate(ctx context.Context, baseFolderID string, baseFolderPath ...string) error

	// PathInCache verifies if M365 container exists within the cache based
	// by comparing the pathString representation to the paths of cachedContainers saved
	PathInCache(pathString string) (string, bool)

	AddToCache(m365Container Container) error
}
