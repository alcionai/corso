package graph

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

// AllMetadataFileNames produces the standard set of filenames used to store graph
// metadata such as delta tokens and folderID->path references.
func AllMetadataFileNames() []string {
	return []string{DeltaURLsFileName, PreviousPathFileName}
}

type QueryParams struct {
	Category      path.CategoryType
	ResourceOwner string
	Credentials   account.M365Config
}

var _ Servicer = &Service{}

type Service struct {
	adapter *msgraphsdk.GraphRequestAdapter
	client  *msgraphsdk.GraphServiceClient
}

func NewService(adapter *msgraphsdk.GraphRequestAdapter) *Service {
	return &Service{
		adapter: adapter,
		client:  msgraphsdk.NewGraphServiceClient(adapter),
	}
}

func (s Service) Adapter() *msgraphsdk.GraphRequestAdapter {
	return s.adapter
}

func (s Service) Client() *msgraphsdk.GraphServiceClient {
	return s.client
}

type Servicer interface {
	// Client() returns msgraph Service client that can be used to process and execute
	// the majority of the queries to the M365 Backstore
	Client() *msgraphsdk.GraphServiceClient
	// Adapter() returns GraphRequest adapter used to process large requests, create batches
	// and page iterators
	Adapter() *msgraphsdk.GraphRequestAdapter
}

// Idable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of an ID.
type Idable interface {
	GetId() *string
}

// Descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type Descendable interface {
	Idable
	GetParentFolderId() *string
}

// Displayable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a display name.
type Displayable interface {
	Idable
	GetDisplayName() *string
}

// Additionalable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of an additional data map.
type Additionalable interface {
	Idable
	GetAdditionalData() map[string]any
}

type Container interface {
	Descendable
	Displayable
	Additionalable
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
	// conclude its search. Default input is "".
	Populate(ctx context.Context, baseFolderID string, baseContainerPather ...string) error

	// PathInCache performs a look up of a path reprensentation
	// and returns the m365ID of directory iff the pathString
	// matches the path of a container within the cache.
	// @returns bool represents if m365ID was found.
	PathInCache(pathString string) (string, bool)

	AddToCache(ctx context.Context, m365Container Container) error

	// Items returns the containers in the cache.
	Items() []CachedContainer
}
