// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"bytes"
	"context"
	"strings"
	"sync/atomic"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/selectors"
)

const (
	mailCategory = "mail"
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	graphService
	tenant           string
	Users            map[string]string                 //key<email> value<id>
	status           *support.ConnectorOperationStatus // contains the status of the last run status
	statusCh         chan *support.ConnectorOperationStatus
	awaitingMessages int32
	credentials      account.M365Config
}

type graphService struct {
	client   msgraphsdk.GraphServiceClient
	adapter  msgraphsdk.GraphRequestAdapter
	failFast bool // if true service will exit sequence upon encountering an error
}

func (gs *graphService) Client() *msgraphsdk.GraphServiceClient {
	return &gs.client
}

func (gs *graphService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &gs.adapter
}

func (gs *graphService) ErrPolicy() bool {
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
	selecting := []string{"id, mail"}
	requestParams := &msuser.UsersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	response, err := gc.graphService.client.Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return err
	}
	if response == nil {
		err = support.WrapAndAppend("general access", errors.New("connector failed: No access"), err)
		return err
	}
	userIterator, err := msgraphgocore.NewPageIterator(response, &gc.graphService.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return err
	}
	var iterateError error
	callbackFunc := func(userItem interface{}) bool {
		user, ok := userItem.(models.Userable)
		if !ok {
			err = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), errors.New("user iteration failure"), err)
			return true
		}
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	iterateError = userIterator.Iterate(callbackFunc)
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
func (gc *GraphConnector) ExchangeDataCollection(ctx context.Context, selector selectors.Selector) ([]data.Collection, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, errors.Wrap(err, "collecting exchange data")
	}

	collections := []data.Collection{}
	scopes := eb.Scopes()
	var errs error

	// for each scope that includes mail messages, get all
	for _, scope := range scopes {
		if !scope.IncludesCategory(selectors.ExchangeMail) {
			continue
		}

		for _, user := range scope.Get(selectors.ExchangeUser) {
			// TODO: handle "get mail for all users"
			// this would probably no-op without this check,
			// but we want it made obvious that we're punting.
			if user == selectors.AnyTgt {
				errs = support.WrapAndAppend(
					"all-users",
					errors.New("all users selector currently not handled"),
					errs)
				continue
			}
			// Creates a map of collections based on scope
			dcs, err := gc.createCollections(ctx, scope)
			if err != nil {
				return nil, support.WrapAndAppend(user, err, errs)
			}

			if len(dcs) > 0 {
				for _, collection := range dcs {
					collections = append(collections, collection)
				}
			}
		}
	}
	return collections, errs
}

// RestoreMessages: Utility function to connect to M365 backstore
// and upload messages from DataCollection.
// FullPath: tenantId, userId, <mailCategory>, FolderId
func (gc *GraphConnector) RestoreMessages(ctx context.Context, dcs []data.Collection) error {
	var (
		pathCounter         = map[string]bool{}
		attempts, successes int
		errs                error
	)
	gc.incrementAwaitingMessages()

	for _, dc := range dcs {
		// must be user.GetId(), PrimaryName no longer works 6-15-2022
		user := dc.FullPath()[1]
		items := dc.Items()
		pathCounter[strings.Join(dc.FullPath(), "")] = true

		var exit bool
		for !exit {
			select {
			case <-ctx.Done():
				return support.WrapAndAppend("context cancelled", ctx.Err(), errs)
			case data, ok := <-items:
				if !ok {
					exit = true
					break
				}
				attempts++

				buf := &bytes.Buffer{}
				_, err := buf.ReadFrom(data.ToReader())
				if err != nil {
					errs = support.WrapAndAppend(data.UUID(), err, errs)
					continue
				}
				message, err := support.CreateMessageFromBytes(buf.Bytes())
				if err != nil {
					errs = support.WrapAndAppend(data.UUID(), err, errs)
					continue
				}
				clone := support.ToMessage(message)
				address := dc.FullPath()[3]
				valueId := "Integer 0x0E07"
				enableValue := "4"
				sv := models.NewSingleValueLegacyExtendedProperty()
				sv.SetId(&valueId)
				sv.SetValue(&enableValue)
				svlep := []models.SingleValueLegacyExtendedPropertyable{sv}
				clone.SetSingleValueExtendedProperties(svlep)
				draft := false
				clone.SetIsDraft(&draft)
				sentMessage, err := gc.graphService.client.UsersById(user).MailFoldersById(address).Messages().Post(clone)
				if err != nil {
					errs = support.WrapAndAppend(
						data.UUID()+": "+support.ConnectorStackErrorTrace(err),
						err, errs)
					continue
					// TODO: Add to retry Handler for the for failure
				}

				if sentMessage == nil {
					errs = support.WrapAndAppend(data.UUID(), errors.New("Message not Sent: Blocked by server"), errs)
					continue
				}

				successes++
				// This completes the restore loop for a message..
			}
		}
	}

	status := support.CreateStatus(ctx, support.Restore, attempts, successes, len(pathCounter), errs)
	// set the channel asynchronously so that this func doesn't block.
	go func(cos *support.ConnectorOperationStatus) {
		gc.statusCh <- cos
	}(status)
	logger.Ctx(ctx).Debug(gc.PrintableStatus())
	return errs
}

// createCollection - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are stored.
// to the GraphConnector struct.
func (gc *GraphConnector) createCollections(
	ctx context.Context,
	scope selectors.ExchangeScope,
) (map[string]*exchange.Collection, error) {
	var (
		transformer absser.ParsableFactory
		query       exchange.GraphQuery
		gIter       exchange.GraphIterateFunc
	)
	user := scope.Get(selectors.ExchangeUser)[0]
	transformer, query, gIter, err := exchange.SetupExchangeCollectionVars(scope)
	if err != nil {
		return nil, support.WrapAndAppend(user, err, nil)
	}
	response, err := query(&gc.graphService, []string{user})
	if err != nil {
		return nil, err
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, &gc.graphService.adapter, transformer)
	if err != nil {
		return nil, err
	}
	// Create collection of ExchangeDataCollection and create  data Holder
	collections := make(map[string]*exchange.Collection)
	var errs error
	// callbackFunc iterates through all models.Messageable and fills exchange.Collection.jobs[]
	// with corresponding item IDs. New collections are created for each directory
	callbackFunc := gIter(gc.tenant, scope, errs, gc.failFast, gc.credentials, collections, gc.statusCh)
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		errs = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), iterateError, errs)
	}
	if errs != nil {
		return nil, errs // return error if snapshot is incomplete
	}

	for range collections {
		gc.incrementAwaitingMessages()
	}

	return collections, errs
}

// AwaitStatus updates status field based on item within statusChannel.
func (gc *GraphConnector) AwaitStatus() *support.ConnectorOperationStatus {
	if gc.awaitingMessages > 0 {
		gc.status = <-gc.statusCh
		atomic.AddInt32(&gc.awaitingMessages, -1)
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
