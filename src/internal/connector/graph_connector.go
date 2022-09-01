// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/selectors"
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	graphService
	tenant           string
	Users            map[string]string                 // key<email> value<id>
	status           *support.ConnectorOperationStatus // contains the status of the last run status
	statusCh         chan *support.ConnectorOperationStatus
	awaitingMessages int32
	credentials      account.M365Config
}

// Service returns the GC's embedded graph.Service
func (gc GraphConnector) Service() graph.Service {
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

func NewGraphConnector(acct account.Account) (*GraphConnector, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return nil, errors.Wrap(err, "retrieving m356 account configuration")
	}

	gc := GraphConnector{
		tenant:      m365.TenantID,
		Users:       make(map[string]string, 0),
		status:      nil,
		statusCh:    make(chan *support.ConnectorOperationStatus),
		credentials: m365,
	}

	aService, err := gc.createService(false)
	if err != nil {
		return nil, err
	}

	gc.graphService = *aService

	err = gc.setTenantUsers()
	if err != nil {
		return nil, err
	}

	return &gc, nil
}

// createService constructor for graphService component
func (gc *GraphConnector) createService(shouldFailFast bool) (*graphService, error) {
	adapter, err := graph.CreateAdapter(
		gc.credentials.TenantID,
		gc.credentials.ClientID,
		gc.credentials.ClientSecret,
	)
	if err != nil {
		return nil, err
	}

	connector := graphService{
		adapter:  *adapter,
		client:   *msgraphsdk.NewGraphServiceClient(adapter),
		failFast: shouldFailFast,
	}

	return &connector, err
}

func (gs *graphService) EnableFailFast() {
	gs.failFast = true
}

// setTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated during this method
// iff the return value is true
func (gc *GraphConnector) setTenantUsers() error {
	response, err := exchange.GetAllUsersForTenant(gc.graphService, "")
	if err != nil {
		return errors.Wrapf(
			err,
			"tenant %s M365 query: %s",
			gc.tenant,
			support.ConnectorStackErrorTrace(err),
		)
	}

	if response == nil {
		err = support.WrapAndAppend("general access", errors.New("connector failed: No access"), err)
		return err
	}

	userIterator, err := msgraphgocore.NewPageIterator(
		response,
		&gc.graphService.adapter,
		models.CreateUserCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	callbackFunc := func(userItem interface{}) bool {
		user, ok := userItem.(models.Userable)
		if !ok {
			err = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), errors.New("user iteration failure"), err)
			return true
		}

		if user.GetUserPrincipalName() == nil {
			err = support.WrapAndAppend(
				gc.graphService.adapter.GetBaseUrl(),
				fmt.Errorf("no email address for User: %s", *user.GetId()),
				err,
			)

			return true
		}

		// *user.GetId() is populated for every M365 entityable object by M365 backstore
		gc.Users[*user.GetUserPrincipalName()] = *user.GetId()

		return true
	}

	iterateError := userIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), iterateError, err)
	}

	return err
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	return buildFromMap(true, gc.Users)
}

// GetUsersIds returns the M365 id for the user
func (gc *GraphConnector) GetUsersIds() []string {
	return buildFromMap(false, gc.Users)
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

// ExchangeDataStream returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
// Assumption: User exists
//  Add iota to this call -> mail, contacts, calendar,  etc.
func (gc *GraphConnector) ExchangeDataCollection(
	ctx context.Context,
	selector selectors.Selector,
) ([]data.Collection, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, errors.Wrap(err, "exchangeDataCollection: unable to parse selector")
	}

	var (
		scopes      = eb.DiscreteScopes(gc.GetUsers())
		collections = []data.Collection{}
		errs        error
	)

	for _, scope := range scopes {
		// Creates a map of collections based on scope
		dcs, err := gc.createCollections(ctx, scope)
		if err != nil {
			user := scope.Get(selectors.ExchangeUser)
			return nil, support.WrapAndAppend(user[0], err, errs)
		}

		for _, collection := range dcs {
			collections = append(collections, collection)
		}
	}

	return collections, errs
}

// RestoreExchangeDataCollection: Utility function to connect to M365 backstore
// and upload messages from DataCollection.
// FullPath: tenantId, userId, <collectionCategory>, FolderId
func (gc *GraphConnector) RestoreExchangeDataCollection(
	ctx context.Context,
	dcs []data.Collection,
) error {
	var (
		pathCounter         = map[string]bool{}
		attempts, successes int
		errs                error
		folderID            string
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	for _, dc := range dcs {
		var (
			directory = strings.Join(dc.FullPath(), "")
			user      = dc.FullPath()[1]
			items     = dc.Items()
			category  = dc.FullPath()[2]
			exit      bool
		)

		if _, ok := pathCounter[directory]; !ok {
			pathCounter[directory] = true
			folderID, errs = exchange.GetRestoreContainer(&gc.graphService, user, category)

			if errs != nil {
				return errs
			}
		}

		for !exit {
			select {
			case <-ctx.Done():
				return support.WrapAndAppend("context cancelled", ctx.Err(), errs)
			case itemData, ok := <-items:
				if !ok {
					exit = true
					break
				}
				attempts++

				buf := &bytes.Buffer{}

				_, err := buf.ReadFrom(itemData.ToReader())
				if err != nil {
					errs = support.WrapAndAppend(itemData.UUID(), err, errs)
					continue
				}

				err = exchange.RestoreExchangeObject(ctx, buf.Bytes(), category, policy, &gc.graphService, folderID, user)

				if err != nil {
					errs = support.WrapAndAppend(itemData.UUID(), err, errs)
					continue
				}
				successes++
			}
		}
	}

	gc.incrementAwaitingMessages()

	status := support.CreateStatus(ctx, support.Restore, attempts, successes, len(pathCounter), errs)
	// set the channel asynchronously so that this func doesn't block.
	go func(cos *support.ConnectorOperationStatus) {
		gc.statusCh <- cos
	}(status)

	return errs
}

// createCollection - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are stored.
// to the GraphConnector struct.
func (gc *GraphConnector) createCollections(
	ctx context.Context,
	scope selectors.ExchangeScope,
) ([]*exchange.Collection, error) {
	var (
		errs                           error
		transformer, query, gIter, err = exchange.SetupExchangeCollectionVars(scope)
	)

	if err != nil {
		return nil, support.WrapAndAppend(gc.Service().Adapter().GetBaseUrl(), err, nil)
	}

	users := scope.Get(selectors.ExchangeUser)
	allCollections := make([]*exchange.Collection, 0)
	// Create collection of ExchangeDataCollection
	for _, user := range users {
		collections := make(map[string]*exchange.Collection)

		response, err := query(&gc.graphService, user)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"user %s M365 query: %s",
				user, support.ConnectorStackErrorTrace(err))
		}

		pageIterator, err := msgraphgocore.NewPageIterator(response, &gc.graphService.adapter, transformer)
		if err != nil {
			return nil, err
		}

		// callbackFunc iterates through all M365 object target and fills exchange.Collection.jobs[]
		// with corresponding item M365IDs. New collections are created for each directory.
		// Each directory used the M365 Identifier. The use of ID stops collisions betweens users
		callbackFunc := gIter(ctx, user, scope, errs, gc.failFast, gc.credentials, collections, gc.statusCh)
		iterateError := pageIterator.Iterate(callbackFunc)

		if iterateError != nil {
			errs = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), iterateError, errs)
		}

		if errs != nil {
			return nil, errs // return error if snapshot is incomplete
		}

		for _, collection := range collections {
			gc.incrementAwaitingMessages()

			allCollections = append(allCollections, collection)
		}
	}

	return allCollections, errs
}

// AwaitStatus updates status field based on item within statusChannel.
func (gc *GraphConnector) AwaitStatus() *support.ConnectorOperationStatus {
	if gc.awaitingMessages > 0 {
		atomic.AddInt32(&gc.awaitingMessages, -1)
		gc.status = <-gc.statusCh
	}

	return gc.status
}

// Status returns the current status of the graphConnector operaion.
func (gc *GraphConnector) Status() *support.ConnectorOperationStatus {
	return gc.status
}

// PrintableStatus returns a string formatted version of the GC status.
func (gc *GraphConnector) PrintableStatus() string {
	if gc.status == nil {
		return ""
	}

	return gc.status.String()
}

func (gc *GraphConnector) incrementAwaitingMessages() {
	atomic.AddInt32(&gc.awaitingMessages, 1)
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
