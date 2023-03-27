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
	Owners     api.Client
	itemClient *http.Client // configured to handle large item downloads

	tenant      string
	credentials account.M365Config

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

type resource int

const (
	UnknownResource resource = iota
	AllResources
	Users
	Sites
)

func NewGraphConnector(
	ctx context.Context,
	itemClient *http.Client,
	acct account.Account,
	r resource,
	errs *fault.Bus,
) (*GraphConnector, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "retrieving m365 account configuration").WithClues(ctx)
	}

	gc := GraphConnector{
		itemClient:  itemClient,
		tenant:      m365.AzureTenantID,
		wg:          &sync.WaitGroup{},
		credentials: m365,
	}

	gc.Service, err = gc.createService()
	if err != nil {
		return nil, clues.Wrap(err, "creating service connection").WithClues(ctx)
	}

	gc.Owners, err = api.NewClient(m365)
	if err != nil {
		return nil, clues.Wrap(err, "creating api client").WithClues(ctx)
	}

	return &gc, nil
}

// PopulateOwnerIDAndNamesFrom takes the provided owner identifier and produces
// the owner's name and ID from that value.  Returns an error if the owner is
// not recognized by the current tenant.
//
// The id-name swapper is optional.  Some processes will look up all owners in
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
	owner string, // input value, can be either id or name
	ins common.IDNameSwapper,
) (string, string, error) {
	// move this to GC method
	id, name, err := getOwnerIDAndNameFrom(owner, ins)
	if err != nil {
		return "", "", errors.Wrap(err, "resolving resource owner details")
	}

	gc.IDNameLookup = ins

	if ins == nil || (len(ins.IDs()) == 0 && len(ins.Names()) == 0) {
		gc.IDNameLookup = common.IDsNames{
			IDToName: map[string]string{id: name},
			NameToID: map[string]string{name: id},
		}
	}

	return id, name, nil
}

func getOwnerIDAndNameFrom(
	owner string,
	ins common.IDNameSwapper,
) (string, string, error) {
	if ins == nil {
		return owner, owner, nil
	}

	if n, ok := ins.NameOf(owner); ok {
		return owner, n, nil
	} else if i, ok := ins.IDOf(owner); ok {
		return i, owner, nil
	}

	// TODO: look-up user by owner, either id or name,
	// and populate with maps as a result.  Only
	// return owner, owner as a very last resort.

	return owner, owner, nil
}

// createService constructor for graphService component
func (gc *GraphConnector) createService() (*graph.Service, error) {
	adapter, err := graph.CreateAdapter(
		gc.credentials.AzureTenantID,
		gc.credentials.AzureClientID,
		gc.credentials.AzureClientSecret)
	if err != nil {
		return &graph.Service{}, err
	}

	return graph.NewService(adapter), nil
}

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
