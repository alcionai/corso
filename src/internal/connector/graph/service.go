package graph

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/account"
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

// ContainerResolver houses functions for getting information about containers
// from remote APIs (i.e. resolve folder paths with Graph API). Resolvers may
// cache information about containers.
type ContainerResolver interface {
	// IDToPath takes an m365 container ID and converts it to a hierarchical path
	// to that container. The path has a similar format to paths on the local
	// file system.
	IDToPath(ctx context.Context, m365ID string) (*path.Builder, error)
	// Populate performs any setup logic the resolver may need.
	Populate(context.Context) error
}
