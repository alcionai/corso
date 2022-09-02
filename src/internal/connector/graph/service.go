package graph

import (
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
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
