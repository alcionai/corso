// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"context"
	"net/http"
	"runtime/trace"
	"sync"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
)

// ---------------------------------------------------------------------------
// Graph Connector
// ---------------------------------------------------------------------------

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
	ResourceOwnerIDToName map[string]string
	ResourceOwnerNameToID map[string]string

	// wg is used to track completion of GC tasks
	wg     *sync.WaitGroup
	region *trace.Region

	// mutex used to synchronize updates to `status`
	mu     sync.Mutex
	status support.ConnectorOperationStatus // contains the status of the last run status
}

func NewGraphConnector(
	ctx context.Context,
	itemClient *http.Client,
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
		itemClient:  itemClient,
		Discovery:   discovery,
		tenant:      acct.ID(),
		wg:          &sync.WaitGroup{},
		credentials: creds,
		ownerLookup: rc,
		Service:     service,
	}

	return &gc, nil
}

// ---------------------------------------------------------------------------
// Owner Lookup
// ---------------------------------------------------------------------------

// PopulateOwnerIDAndNamesFrom takes the provided owner identifier and produces
// the owner's name and ID from that value.  Returns an error if the owner is
// not recognized by the current tenant.
//
// The id-name maps are optional.  Some processes will look up all owners in
// the tenant before reaching this step.  In that case, the data gets handed
// down for this func to consume instead of performing further queries.  The
// maps get stored inside the gc instance for later re-use.
//
// TODO: If the maps are nil or empty, this func will perform a lookup on the given
// owner, and populate each map with that owner's id and name for downstream
// guarantees about that data being present.  Optional performance enhancement
// idea: downstream from here, we should _only_ need the given user's id and name,
// and could store minimal map copies with that info instead of the whole tenant.
func (gc *GraphConnector) PopulateOwnerIDAndNamesFrom(
	ctx context.Context,
	owner string, // input value, can be either id or name
	idToName, nameToID map[string]string, // optionally pre-populated lookups
) (string, string, error) {
	// ensure the maps exist, even if they aren't populated so that
	// getOwnerIDAndNameFrom can populate any values it looks up.
	if len(idToName) == 0 {
		idToName = map[string]string{}
	}

	if len(nameToID) == 0 {
		nameToID = map[string]string{}
	}

	id, name, err := gc.ownerLookup.getOwnerIDAndNameFrom(ctx, gc.Discovery, owner, idToName, nameToID)
	if err != nil {
		return "", "", errors.Wrap(err, "resolving resource owner details")
	}

	gc.ResourceOwnerIDToName = idToName
	gc.ResourceOwnerNameToID = nameToID

	return id, name, nil
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

// Status returns the current status of the graphConnector operaion.
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
// Resource Handling
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
		idToName, nameToID map[string]string,
	) (
		ownerID string,
		ownerName string,
		err error,
	)
}

var ErrResourceOwnerNotFound = clues.New("resource owner not found in tenant")

// getOwnerIDAndNameFrom looks up the owner's canonical id and display name.
// if idToName and nameToID are populated, and the owner is a key of one of
// those maps, then those values are returned.
//
// As a fallback, the resource calls the discovery api to fetch the user or
// site using the owner value. This fallback assumes that the owner is a well
// formed ID or display name of appropriate design (PrincipalName for users,
// WebURL for sites). If the fallback lookup is used, the maps are populated
// to contain the id and name references.
//
// Consumers are allowed to pass in a path suffix (eg: /sites/foo) as a site
// owner, but only if they also pass in a nameToID map.  A nil map will cascade
// to the fallback, which will fail for having a malformed id value.
func (r resourceClient) getOwnerIDAndNameFrom(
	ctx context.Context,
	discovery api.Client,
	owner string,
	idToName, nameToID map[string]string,
) (string, string, error) {
	if n, ok := idToName[owner]; ok {
		return owner, n, nil
	} else if id, ok := nameToID[owner]; ok {
		return id, owner, nil
	}

	ctx = clues.Add(ctx, "owner_identifier", owner)

	var (
		id, name string
		err      error
	)

	// check if the provided owner is a suffix of a weburl in the lookup map
	if r.enum == Sites {
		url, _, ok := filters.PathSuffix([]string{owner}).CompareAny(maps.Keys(nameToID)...)
		if ok {
			return nameToID[url], url, nil
		}
	}

	id, name, err = r.getter.GetIDAndName(ctx, owner)
	if err != nil {
		return "", "", err
	}

	if len(id) == 0 || len(name) == 0 {
		return "", "", clues.Stack(ErrResourceOwnerNotFound)
	}

	idToName[id] = name
	nameToID[name] = id

	return id, name, nil
}
