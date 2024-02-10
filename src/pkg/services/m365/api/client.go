package api

import (
	"context"
	"io"
	"net/http"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Client is used to fulfill queries that are traditionally
// backed by GraphAPI. A struct is used in this case, instead
// of deferring to pure function wrappers, so that the boundary
// separates the granular implementation of the graphAPI and
// kiota away from the other packages.
type Client struct {
	Credentials account.M365Config

	// The Stable service is re-usable for any request.
	// This allows us to maintain performance across async requests.
	Stable graph.Servicer

	// The LargeItem graph servicer is configured specifically for
	// downloading large items such as drive item content or outlook
	// mail and event attachments.
	LargeItem graph.Servicer

	// The Requester provides a client specifically for calling
	// arbitrary urls instead of constructing queries using the
	// graph api client.
	Requester graph.Requester

	counter *count.Bus

	options control.Options
}

// NewClient produces a new exchange api client.  Must be used in
// place of creating an ad-hoc client struct.
func NewClient(
	creds account.M365Config,
	co control.Options,
	counter *count.Bus,
	opts ...graph.Option,
) (Client, error) {
	s, err := NewService(creds, counter, opts...)
	if err != nil {
		return Client{}, err
	}

	li, err := newLargeItemService(creds, counter)
	if err != nil {
		return Client{}, err
	}

	azureAuth, err := graph.NewAzureAuth(creds)
	if err != nil {
		return Client{}, clues.Wrap(err, "generating azure authorizer")
	}

	rqr := graph.NewNoTimeoutHTTPWrapper(counter, graph.AuthorizeRequester(azureAuth))

	if co.DeltaPageSize < 1 || co.DeltaPageSize > maxDeltaPageSize {
		co.DeltaPageSize = maxDeltaPageSize
	}

	cli := Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
		Requester:   rqr,
		counter:     counter,
		options:     co,
	}

	return cli, nil
}

// initConcurrencyLimit ensures that the graph concurrency limiter is
// initialized, so that calls do not step over graph api's service limits.
// Limits are derived from the provided servie type.
// Callers will need to call this func before making api calls an api client.
func InitConcurrencyLimit(ctx context.Context, pst path.ServiceType) {
	graph.InitializeConcurrencyLimiter(ctx, pst == path.ExchangeService, 4)
}

// Service generates a new graph servicer.  New servicers are used for paged
// and other long-running requests instead of the client's stable service,
// so that in-flight state within the adapter doesn't get clobbered.
// Most calls should use the Client.Stable property instead of calling this
// func, unless it is explicitly necessary.
func (c Client) Service(
	counter *count.Bus,
	opts ...graph.Option,
) (graph.Servicer, error) {
	return NewService(c.Credentials, counter, opts...)
}

func NewService(
	creds account.M365Config,
	counter *count.Bus,
	opts ...graph.Option,
) (*graph.Service, error) {
	a, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		counter,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph api adapter")
	}

	return graph.NewService(a), nil
}

func newLargeItemService(
	creds account.M365Config,
	counter *count.Bus,
) (*graph.Service, error) {
	a, err := NewService(creds, counter, graph.NoTimeout())
	return a, clues.Wrap(err, "generating no-timeout graph adapter").OrNil()
}

type Getter interface {
	Get(
		ctx context.Context,
		url string,
		headers map[string]string,
	) (*http.Response, error)
}

// Get performs an ad-hoc get request using its graph.Requester
func (c Client) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return c.Requester.Request(ctx, http.MethodGet, url, nil, headers)
}

// Get performs an ad-hoc get request using its graph.Requester
func (c Client) Post(
	ctx context.Context,
	url string,
	headers map[string]string,
	body io.Reader,
) (*http.Response, error) {
	return c.Requester.Request(ctx, http.MethodGet, url, body, headers)
}

// ---------------------------------------------------------------------------
// per-call config
// ---------------------------------------------------------------------------

// CallConfig defines additional parameters to add to the query like field
// selection and limiting results.
//
// Not all calls may support all options. Check individual implementations to
// see what's supported and what's not.
type CallConfig struct {
	Expand              []string
	Select              []string
	CanMakeDeltaQueries bool
	// LimitResults limits the returned results to the given number. If 0, returns
	// all results.
	LimitResults int
}

// ---------------------------------------------------------------------------
// common interfaces
// ---------------------------------------------------------------------------

type GetByIDer[T any] interface {
	GetByID(
		ctx context.Context,
		identifier string,
		cc CallConfig,
	) (T, error)
}
