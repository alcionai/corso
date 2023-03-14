// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"context"
	"fmt"
	"net/http"
	"runtime/trace"
	"strings"
	"sync"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	D "github.com/alcionai/corso/src/internal/diagnostics"
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
	Owners     api.Client
	itemClient *http.Client // configured to handle large item downloads

	tenant      string
	Sites       map[string]string // key<???> value<???>
	credentials account.M365Config

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

	if r == AllResources || r == Sites {
		if err = gc.setTenantSites(ctx, errs); err != nil {
			return nil, errors.Wrap(err, "retrieveing tenant site list")
		}
	}

	return &gc, nil
}

// createService constructor for graphService component
func (gc *GraphConnector) createService() (*graph.Service, error) {
	adapter, err := graph.CreateAdapter(
		gc.credentials.AzureTenantID,
		gc.credentials.AzureClientID,
		gc.credentials.AzureClientSecret,
	)
	if err != nil {
		return &graph.Service{}, err
	}

	return graph.NewService(adapter), nil
}

// setTenantSites queries the M365 to identify the sites in the
// workspace. The sites field is updated during this method
// iff the returned error is nil.
func (gc *GraphConnector) setTenantSites(ctx context.Context, errs *fault.Bus) error {
	gc.Sites = map[string]string{}

	ctx, end := D.Span(ctx, "gc:setTenantSites")
	defer end()

	sites, err := getResources(
		ctx,
		gc.Service,
		gc.tenant,
		sharepoint.GetAllSitesForTenant,
		models.CreateSiteCollectionResponseFromDiscriminatorValue,
		identifySite,
		errs)
	if err != nil {
		return err
	}

	gc.Sites = sites

	return nil
}

var errKnownSkippableCase = errors.New("case is known and skippable")

const personalSitePath = "sharepoint.com/personal/"

// Transforms an interface{} into a key,value pair representing
// siteName:siteID.
func identifySite(item any) (string, string, error) {
	m, ok := item.(models.Siteable)
	if !ok {
		return "", "", clues.New("non-Siteable item").With("item_type", fmt.Sprintf("%T", item))
	}

	id := ptr.Val(m.GetId())
	url, ok := ptr.ValOK(m.GetWebUrl())

	if m.GetName() == nil {
		// the built-in site at "https://{tenant-domain}/search" never has a name.
		if ok && strings.HasSuffix(url, "/search") {
			// TODO: pii siteID, on this and all following cases
			return "", "", clues.Stack(errKnownSkippableCase).With("site_id", id)
		}

		return "", "", clues.New("site has no name").With("site_id", id)
	}

	// personal (ie: oneDrive) sites have to be filtered out server-side.
	if ok && strings.Contains(url, personalSitePath) {
		return "", "", clues.Stack(errKnownSkippableCase).With("site_id", id)
	}

	return url, id, nil
}

// GetSiteWebURLs returns the WebURLs of sharepoint sites within the tenant.
func (gc *GraphConnector) GetSiteWebURLs() []string {
	return maps.Keys(gc.Sites)
}

// GetSiteIds returns the canonical site IDs in the tenant
func (gc *GraphConnector) GetSiteIDs() []string {
	return maps.Values(gc.Sites)
}

// UnionSiteIDsAndWebURLs reduces the id and url slices into a single slice of site IDs.
// WebURLs will run as a path-suffix style matcher.  Callers may provide partial urls, though
// each element in the url must fully match.  Ex: the webURL value "foo" will match "www.ex.com/foo",
// but not match "www.ex.com/foobar".
// The returned IDs are reduced to a set of unique values.
func (gc *GraphConnector) UnionSiteIDsAndWebURLs(
	ctx context.Context,
	ids, urls []string,
	errs *fault.Bus,
) ([]string, error) {
	if len(gc.Sites) == 0 {
		if err := gc.setTenantSites(ctx, errs); err != nil {
			return nil, err
		}
	}

	idm := map[string]struct{}{}

	for _, id := range ids {
		idm[id] = struct{}{}
	}

	match := filters.PathSuffix(urls)

	for url, id := range gc.Sites {
		if !match.Compare(url) {
			continue
		}

		idm[id] = struct{}{}
	}

	idsl := make([]string, 0, len(idm))
	for id := range idm {
		idsl = append(idsl, id)
	}

	return idsl, nil
}

// AwaitStatus waits for all gc tasks to complete and then returns status
func (gc *GraphConnector) AwaitStatus() *support.ConnectorOperationStatus {
	defer func() {
		if gc.region != nil {
			gc.region.End()
		}
	}()
	gc.wg.Wait()

	return &gc.status
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
// Helper Funcs
// ---------------------------------------------------------------------------

func getResources(
	ctx context.Context,
	gs graph.Servicer,
	tenantID string,
	query func(context.Context, graph.Servicer) (serialization.Parsable, error),
	parser func(parseNode serialization.ParseNode) (serialization.Parsable, error),
	identify func(any) (string, string, error),
	errs *fault.Bus,
) (map[string]string, error) {
	resources := map[string]string{}

	response, err := query(ctx, gs)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "retrieving tenant's resources")
	}

	iter, err := msgraphgocore.NewPageIterator(response, gs.Adapter(), parser)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	el := errs.Local()

	callbackFunc := func(item any) bool {
		if el.Failure() != nil {
			return false
		}

		k, v, err := identify(item)
		if err != nil {
			if !errors.Is(err, errKnownSkippableCase) {
				el.AddRecoverable(clues.Stack(err).
					WithClues(ctx).
					With("query_url", gs.Adapter().GetBaseUrl()))
			}

			return true
		}

		resources[k] = v

		return true
	}

	if err := iter.Iterate(ctx, callbackFunc); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resources, el.Failure()
}
