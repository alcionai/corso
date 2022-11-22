// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"context"
	"runtime/trace"
	"strings"
	"sync"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

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
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Graph Connector
// ---------------------------------------------------------------------------

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	graphService
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

// Service returns the GC's embedded graph.Service
func (gc *GraphConnector) Service() graph.Service {
	return gc.graphService
}

var _ graph.Service = &graphService{}

type graphService struct {
	client   msgraphsdk.GraphServiceClient
	adapter  msgraphsdk.GraphRequestAdapter
	failFast bool // if true service will exit sequence upon encountering an error
}

func (gs graphService) Client() *msgraphsdk.GraphServiceClient {
	return &gs.client
}

func (gs graphService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &gs.adapter
}

func (gs graphService) ErrPolicy() bool {
	return gs.failFast
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

	aService, err := gc.createService(false)
	if err != nil {
		return nil, errors.Wrap(err, "creating service connection")
	}

	gc.graphService = *aService

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
func (gc *GraphConnector) createService(shouldFailFast bool) (*graphService, error) {
	adapter, err := graph.CreateAdapter(
		gc.credentials.AzureTenantID,
		gc.credentials.AzureClientID,
		gc.credentials.AzureClientSecret,
	)
	if err != nil {
		return nil, err
	}

	connector := graphService{
		adapter:  *adapter,
		client:   *msgraphsdk.NewGraphServiceClient(adapter),
		failFast: shouldFailFast,
	}

	return &connector, nil
}

func (gs *graphService) EnableFailFast() {
	gs.failFast = true
}

// setTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated during this method
// iff the returned error is nil
func (gc *GraphConnector) setTenantUsers(ctx context.Context) error {
	ctx, end := D.Span(ctx, "gc:setTenantUsers")
	defer end()

	users, err := getResources(
		ctx,
		gc.graphService,
		gc.tenant,
		exchange.GetAllUsersForTenant,
		models.CreateUserCollectionResponseFromDiscriminatorValue,
		identifyUser,
	)
	if err != nil {
		return err
	}

	gc.Users = users

	return nil
}

// Transforms an interface{} into a key,value pair representing
// userPrincipalName:userID.
func identifyUser(item any) (string, string, error) {
	m, ok := item.(models.Userable)
	if !ok {
		return "", "", errors.New("iteration retrieved non-User item")
	}

	if m.GetUserPrincipalName() == nil {
		return "", "", errors.Errorf("no principal name for User: %s", *m.GetId())
	}

	return *m.GetUserPrincipalName(), *m.GetId(), nil
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	return buildFromMap(true, gc.Users)
}

// GetUsersIds returns the M365 id for the user
func (gc *GraphConnector) GetUsersIds() []string {
	return buildFromMap(false, gc.Users)
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
		gc.graphService,
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
		// the built-in site at "htps://{tenant-domain}/search" never has a name.
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

	return *m.GetName(), *m.GetId(), nil
}

// GetSites returns the siteIDs of sharepoint sites within tenant.
func (gc *GraphConnector) GetSites() []string {
	return buildFromMap(true, gc.Sites)
}

// GetSiteIds returns the M365 id for the user
func (gc *GraphConnector) GetSiteIds() []string {
	return buildFromMap(false, gc.Sites)
}

// buildFromMap helper function for returning []string from map.
// Returns list of keys iff true; otherwise returns a list of values
func buildFromMap(isKey bool, mapping map[string]string) []string {
	returnString := make([]string, 0)

	if isKey {
		for k := range mapping {
			returnString = append(returnString, k)
		}
	} else {
		for _, v := range mapping {
			returnString = append(returnString, v)
		}
	}

	return returnString
}

// RestoreDataCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: gc.status is updated at the completion of operation
func (gc *GraphConnector) RestoreDataCollections(
	ctx context.Context,
	selector selectors.Selector,
	dest control.RestoreDestination,
	dcs []data.Collection,
) (*details.Details, error) {
	ctx, end := D.Span(ctx, "connector:restore")
	defer end()

	var (
		status *support.ConnectorOperationStatus
		err    error
		deets  = &details.Details{}
	)

	switch selector.Service {
	case selectors.ServiceExchange:
		status, err = exchange.RestoreExchangeDataCollections(ctx, gc.graphService, dest, dcs, deets)
	case selectors.ServiceOneDrive:
		status, err = onedrive.RestoreCollections(ctx, gc, dest, dcs, deets)
	default:
		err = errors.Errorf("restore data from service %s not supported", selector.Service.String())
	}

	gc.incrementAwaitingMessages()
	gc.UpdateStatus(status)

	return deets, err
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
	gs graph.Service,
	tenantID string,
	query func(context.Context, graph.Service) (serialization.Parsable, error),
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
