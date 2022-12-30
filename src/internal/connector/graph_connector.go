// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"context"
	"runtime/trace"
	"strings"
	"sync"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Graph Connector
// ---------------------------------------------------------------------------

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	Service     graph.Servicer
	tenant      string
	Users       map[string]string // key<email> value<id>
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

func NewGraphConnector(ctx context.Context, acct account.Account, r resource) (*GraphConnector, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return nil, errors.Wrap(err, "retrieving m365 account configuration")
	}

	gc := GraphConnector{
		tenant:      m365.AzureTenantID,
		Users:       make(map[string]string, 0),
		wg:          &sync.WaitGroup{},
		credentials: m365,
	}

	gService, err := gc.createService()
	if err != nil {
		return nil, errors.Wrap(err, "creating service connection")
	}

	gc.Service = gService

	// TODO(ashmrtn): When selectors only encapsulate a single resource owner that
	// is not a wildcard don't populate users or sites when making the connector.
	// For now this keeps things functioning if callers do pass in a selector like
	// "*" instead of.
	if r == AllResources || r == Users {
		if err = gc.setTenantUsers(ctx); err != nil {
			return nil, errors.Wrap(err, "retrieving tenant user list")
		}
	}

	if r == AllResources || r == Sites {
		if err = gc.setTenantSites(ctx); err != nil {
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

// setTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated during this method
// iff the returned error is nil
func (gc *GraphConnector) setTenantUsers(ctx context.Context) error {
	ctx, end := D.Span(ctx, "gc:setTenantUsers")
	defer end()

	users, err := discovery.Users(ctx, gc.Service, gc.tenant)
	if err != nil {
		return err
	}

	gc.Users = make(map[string]string, len(users))

	for _, u := range users {
		gc.Users[*u.GetUserPrincipalName()] = *u.GetId()
	}

	return nil
}

// GetUsers returns the email address of users within the tenant.
func (gc *GraphConnector) GetUsers() []string {
	return maps.Keys(gc.Users)
}

// GetUsersIds returns the M365 id for the user
func (gc *GraphConnector) GetUsersIds() []string {
	return maps.Values(gc.Users)
}

// setTenantSites queries the M365 to identify the sites in the
// workspace. The sites field is updated during this method
// iff the returned error is nil.
func (gc *GraphConnector) setTenantSites(ctx context.Context) error {
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
	)
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
		return "", "", errors.New("iteration retrieved non-Site item")
	}

	if m.GetName() == nil {
		// the built-in site at "https://{tenant-domain}/search" never has a name.
		if m.GetWebUrl() != nil && strings.HasSuffix(*m.GetWebUrl(), "/search") {
			return "", "", errKnownSkippableCase
		}

		return "", "", errors.Errorf("no name for Site: %s", *m.GetId())
	}

	// personal (ie: oneDrive) sites have to be filtered out server-side.
	url := m.GetWebUrl()
	if url != nil && strings.Contains(*url, personalSitePath) {
		return "", "", errKnownSkippableCase
	}

	return *m.GetWebUrl(), *m.GetId(), nil
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
func (gc *GraphConnector) UnionSiteIDsAndWebURLs(ctx context.Context, ids, urls []string) ([]string, error) {
	if len(gc.Sites) == 0 {
		if err := gc.setTenantSites(ctx); err != nil {
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

// RestoreDataCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: gc.status is updated at the completion of operation
func (gc *GraphConnector) RestoreDataCollections(
	ctx context.Context,
	acct account.Account,
	selector selectors.Selector,
	dest control.RestoreDestination,
	dcs []data.Collection,
) (*details.Details, error) {
	ctx, end := D.Span(ctx, "connector:restore")
	defer end()

	var (
		status *support.ConnectorOperationStatus
		err    error
		deets  = &details.Builder{}
	)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, errors.Wrap(err, "malformed azure credentials")
	}

	switch selector.Service {
	case selectors.ServiceExchange:
		status, err = exchange.RestoreExchangeDataCollections(ctx, creds, gc.Service, dest, dcs, deets)
	case selectors.ServiceOneDrive:
		status, err = onedrive.RestoreCollections(ctx, gc.Service, dest, dcs, deets)
	case selectors.ServiceSharePoint:
		status, err = sharepoint.RestoreCollections(ctx, gc.Service, dest, dcs, deets)
	default:
		err = errors.Errorf("restore data from service %s not supported", selector.Service.String())
	}

	gc.incrementAwaitingMessages()
	gc.UpdateStatus(status)

	return deets.Details(), err
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
) (map[string]string, error) {
	resources := map[string]string{}

	response, err := query(ctx, gs)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"retrieving resources for tenant %s: %s",
			tenantID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	iter, err := msgraphgocore.NewPageIterator(response, gs.Adapter(), parser)
	if err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	var iterErrs error

	callbackFunc := func(item any) bool {
		k, v, err := identify(item)
		if err != nil {
			if errors.Is(err, errKnownSkippableCase) {
				return true
			}

			iterErrs = support.WrapAndAppend(gs.Adapter().GetBaseUrl(), err, iterErrs)

			return true
		}

		resources[k] = v

		return true
	}

	if err := iter.Iterate(ctx, callbackFunc); err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	return resources, iterErrs
}

// IsRecoverableError returns true iff error is a RecoverableGCEerror
func IsRecoverableError(e error) bool {
	var recoverable support.RecoverableGCError
	return errors.As(e, &recoverable)
}

// IsNonRecoverableError returns true iff error is a NonRecoverableGCEerror
func IsNonRecoverableError(e error) bool {
	var nonRecoverable support.NonRecoverableGCError
	return errors.As(e, &nonRecoverable)
}
