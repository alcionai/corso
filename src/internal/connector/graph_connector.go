// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"context"
	"net/http"
	"runtime/trace"
	"sync"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// Graph Connector
// ---------------------------------------------------------------------------

// must comply with BackupProducer and RestoreConsumer
var (
	_ inject.BackupProducer  = &GraphConnector{}
	_ inject.RestoreConsumer = &GraphConnector{}
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	Service    graph.Servicer
	Discovery  api.Client
	itemClient *http.Client // configured to handle large item downloads

	tenant      string
	credentials account.M365Config

	ownerLookup getOwnerIDAndNamer
	// maps of resource owner ids to names, and names to ids.
	// not guaranteed to be populated, only here as a post-population
	// reference for processes that choose to populate the values.
	IDNameLookup common.IDNameSwapper

	// wg is used to track completion of GC tasks
	wg     *sync.WaitGroup
	region *trace.Region

	// mutex used to synchronize updates to `status`
	mu     sync.Mutex
	status support.ConnectorOperationStatus // contains the status of the last run status
}

func NewGraphConnector(
	ctx context.Context,
	acct account.Account,
	r resource,
	errs *fault.Bus,
) (*GraphConnector, error) {
	creds, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "retrieving m365 account configuration").WithClues(ctx)
	}

	service, err := createService(creds)
	if err != nil {
		return nil, clues.Wrap(err, "creating service connection").WithClues(ctx)
	}

	discovery, err := api.NewClient(creds)
	if err != nil {
		return nil, clues.Wrap(err, "creating api client").WithClues(ctx)
	}

	rc, err := r.resourceClient(discovery)
	if err != nil {
		return nil, clues.Wrap(err, "creating resource client").WithClues(ctx)
	}

	gc := GraphConnector{
		Discovery:    discovery,
		IDNameLookup: common.IDsNames{},
		Service:      service,

		credentials: creds,
		itemClient:  graph.HTTPClient(graph.NoTimeout()),
		ownerLookup: rc,
		tenant:      acct.ID(),
		wg:          &sync.WaitGroup{},
	}

	return &gc, nil
}

// ---------------------------------------------------------------------------
// Service Client
// ---------------------------------------------------------------------------

// createService constructor for graphService component
func createService(creds account.M365Config) (*graph.Service, error) {
	adapter, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret)
	if err != nil {
		return &graph.Service{}, err
	}

	return graph.NewService(adapter), nil
}

// ---------------------------------------------------------------------------
// Processing Status
// ---------------------------------------------------------------------------

// AwaitStatus waits for all gc tasks to complete and then returns status
func (gc *GraphConnector) Wait() *data.CollectionStats {
	defer func() {
		if gc.region != nil {
			gc.region.End()
			gc.region = nil
		}
	}()
	gc.wg.Wait()

	// clean up and reset statefulness
	dcs := data.CollectionStats{
		Folders:   gc.status.Folders,
		Objects:   gc.status.Metrics.Objects,
		Successes: gc.status.Metrics.Successes,
		Bytes:     gc.status.Metrics.Bytes,
		Details:   gc.status.String(),
	}

	gc.wg = &sync.WaitGroup{}
	gc.status = support.ConnectorOperationStatus{}

	return &dcs
}

// UpdateStatus is used by gc initiated tasks to indicate completion
func (gc *GraphConnector) UpdateStatus(status *support.ConnectorOperationStatus) {
	defer gc.wg.Done()

	if status == nil {
		return
	}

	gc.mu.Lock()
	defer gc.mu.Unlock()
	gc.status = support.MergeStatus(gc.status, *status)
}

// Status returns the current status of the graphConnector operation.
func (gc *GraphConnector) Status() support.ConnectorOperationStatus {
	return gc.status
}

// PrintableStatus returns a string formatted version of the GC status.
func (gc *GraphConnector) PrintableStatus() string {
	return gc.status.String()
}

func (gc *GraphConnector) incrementAwaitingMessages() {
	gc.wg.Add(1)
}

func (gc *GraphConnector) incrementMessagesBy(num int) {
	gc.wg.Add(num)
}

// ---------------------------------------------------------------------------
// Resource Lookup Handling
// ---------------------------------------------------------------------------

type resource int

const (
	UnknownResource resource = iota
	AllResources             // unused
	Users
	Sites
)

func (r resource) resourceClient(discovery api.Client) (*resourceClient, error) {
	switch r {
	case Users:
		return &resourceClient{enum: r, getter: discovery.Users()}, nil
	case Sites:
		return &resourceClient{enum: r, getter: discovery.Sites()}, nil
	default:
		return nil, clues.New("unrecognized owner resource enum").With("resource_enum", r)
	}
}

type resourceClient struct {
	enum   resource
	getter getIDAndNamer
}

type getIDAndNamer interface {
	GetIDAndName(ctx context.Context, owner string) (
		ownerID string,
		ownerName string,
		err error,
	)
}

var _ getOwnerIDAndNamer = &resourceClient{}

type getOwnerIDAndNamer interface {
	getOwnerIDAndNameFrom(
		ctx context.Context,
		discovery api.Client,
		owner string,
		ins common.IDNameSwapper,
	) (
		ownerID string,
		ownerName string,
		err error,
	)
}

// getOwnerIDAndNameFrom looks up the owner's canonical id and display name.
// If the owner is present in the idNameSwapper, then that interface's id and
// name values are returned.  As a fallback, the resource calls the discovery
// api to fetch the user or site using the owner value. This fallback assumes
// that the owner is a well formed ID or display name of appropriate design
// (PrincipalName for users, WebURL for sites).
func (r resourceClient) getOwnerIDAndNameFrom(
	ctx context.Context,
	discovery api.Client,
	owner string,
	ins common.IDNameSwapper,
) (string, string, error) {
	if ins != nil {
		if n, ok := ins.NameOf(owner); ok {
			return owner, n, nil
		} else if i, ok := ins.IDOf(owner); ok {
			return i, owner, nil
		}
	}

	ctx = clues.Add(ctx, "owner_identifier", owner)

	var (
		id, name string
		err      error
	)

	id, name, err = r.getter.GetIDAndName(ctx, owner)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return "", "", clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		return "", "", err
	}

	if len(id) == 0 || len(name) == 0 {
		return "", "", clues.Stack(graph.ErrResourceOwnerNotFound)
	}

	return id, name, nil
}

// PopulateOwnerIDAndNamesFrom takes the provided owner identifier and produces
// the owner's name and ID from that value.  Returns an error if the owner is
// not recognized by the current tenant.
//
// The id-name swapper is optional.  Some processes will look up all owners in
// the tenant before reaching this step.  In that case, the data gets handed
// down for this func to consume instead of performing further queries.  The
// data gets stored inside the gc instance for later re-use.
func (gc *GraphConnector) PopulateOwnerIDAndNamesFrom(
	ctx context.Context,
	owner string, // input value, can be either id or name
	ins common.IDNameSwapper,
) (string, string, error) {
	// move this to GC method
	id, name, err := gc.ownerLookup.getOwnerIDAndNameFrom(ctx, gc.Discovery, owner, ins)
	if err != nil {
		return "", "", clues.Wrap(err, "identifying resource owner")
	}

	gc.IDNameLookup = common.IDsNames{
		IDToName: map[string]string{id: name},
		NameToID: map[string]string{name: id},
	}

	return id, name, nil
}
