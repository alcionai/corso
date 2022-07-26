// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/selectors"
)

const (
	numberOfRetries = 4
	mailCategory    = "mail"
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

// PopulateFunc collection of functions can be run serially or asynchronously to populate DataCollections
type PopulateFunc func(context.Context, *graphService, ExchangeDataCollection, chan<- *support.ConnectorOperationStatus)

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

func createAdapter(tenant, client, secret string) (*msgraphsdk.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := az.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, err
	}
	auth, err := ka.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, err
	}
	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	return adapter, err
}

// createService private constructor method for creating graphService
func (gc *GraphConnector) createService(shouldFailFast bool) (*graphService, error) {
	adapter, err := createAdapter(gc.credentials.TenantID, gc.credentials.ClientID, gc.credentials.ClientSecret)
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
func (gc *GraphConnector) ExchangeDataCollection(ctx context.Context, selector selectors.Selector) ([]DataCollection, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, errors.Wrap(err, "collecting exchange data")
	}

	collections := []DataCollection{}
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
			dcs, err := gc.createCollections(ctx, scope)
			if err != nil {
				return nil, support.WrapAndAppend(user, err, errs)
			}

			if len(dcs) > 0 {
				for _, collection := range dcs {
					collections = append(collections, &collection)
				}
			}
		}
	}
	return collections, errs
}

// RestoreMessages: Utility function to connect to M365 backstore
// and upload messages from DataCollection.
// FullPath: tenantId, userId, <mailCategory>, FolderId
func (gc *GraphConnector) RestoreMessages(ctx context.Context, dcs []DataCollection) error {
	var (
		pathCounter         = map[string]bool{}
		attempts, successes int
		errs                error
	)

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
	gc.SetStatus(*status)
	logger.Ctx(ctx).Debug(gc.PrintableStatus())
	return errs
}

// createCollections utility function to build set of DataCollections based on jobs in selector.
func (gc *GraphConnector) createCollections(
	ctx context.Context,
	scope selectors.ExchangeScope,

) ([]ExchangeDataCollection, error) {
	// Checked prior to being called for mail... will add selector
	var (
		response absser.Parsable
		gq GraphQuery
		transformer absser.ParsableFactory
		populationFunction PopulateFunc
		gIter GraphIterateFunc
		err error
	)
	if !scope.IncludesCategory(selectors.ExchangeMail) {
		return nil, errors.New("selector only supports mails")
	}
	// MailSpecific Setters --> TODO create function to return them
	fmt.Printf("%v\n", scope)
	description := scope.Get(selectors.ExchangeMailFolder)
	aUser := scope.Get(selectors.ExchangeUser)
	fmt.Printf("%v\n%s\n", description, aUser[0])
	populationFunction, transformer, gq, gIter = SetupExchangeCollectionVars(scope)
	user := aUser[0]
	response, err = gq(gc.graphService, []string{user})
	if err != nil {
		return nil, err
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, &gc.graphService.adapter, transformer)
	if err != nil {
		return nil, err
	}
	tasklist := NewTaskList() // map[folder][] messageIds
	// 3 Items in this
	callbackFunc := gIter(err, &tasklist) //Mail-Specific
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = support.WrapAndAppend(gc.graphService.adapter.GetBaseUrl(), iterateError, err)
	}
	if err != nil {
		return nil, err // return error if snapshot is incomplete
	}
	// Create collection of ExchangeDataCollection and create  data Holder
	collections := make([]ExchangeDataCollection, 0)

	for aFolder, tasks := range tasklist {
		// prep the items for handoff to the backup consumer
		service, err := gc.createService(gc.failFast)
		if err != nil {
			return nil, support.WrapAndAppend(user, err, err)
		}
		edc := NewExchangeDataCollection(
			ctx,
			user,
			tasks,
			[]string{gc.tenant, user, mailCategory, aFolder},
			*service,
			populationFunction,
			gc.statusCh)

		collections = append(collections, edc)
		gc.incrementAwaitingMessages()
	}
	return collections, err
}

func SetupExchangeCollectionVars(
	scope selectors.ExchangeScope,
) (
	PopulateFunc,
	absser.ParsableFactory,
	GraphQuery,
	GraphIterateFunc) {
	if scope.IncludesCategory(selectors.ExchangeMail) {

		return populateDataCollectionFromMail,
			models.CreateMessageCollectionResponseFromDiscriminatorValue,
			GetAllMessagesForUser,
			IterateMessagesCollection
	}
	return nil, nil, nil, nil
}

// type PopulateFunc func(context.Context, graphService, ExchangeDataCollection, chan *support.ConnectorOperationStatus)
func populateDataCollectionFromMail(
	ctx context.Context,
	gs *graphService,
	edc ExchangeDataCollection,
	statusChannel chan<- *support.ConnectorOperationStatus,
) {
	var errs error
	var success int
	objectWriter := kw.NewJsonSerializationWriter()

	for _, task := range edc.tasks {
		response, err := gs.client.UsersById(edc.user).MessagesById(task).Get()
		if err != nil {
			details := support.ConnectorStackErrorTrace(err)
			errs = support.WrapAndAppend(edc.user, errors.Wrapf(err, "unable to retrieve %s, %s", task, details), errs)
			continue
		}
		err = messageToDataCollection(&gs.client, ctx, objectWriter, edc.data, response, edc.user)
		if err != nil {
			errs = support.WrapAndAppendf(edc.user, err, errs)
		} else {
			success++
		}
		if errs != nil && gs.failFast {
			break
		}
	}
	edc.FinishPopulation()
	status := support.CreateStatus(ctx, support.Backup, len(edc.tasks), success, 1, errs)
	logger.Ctx(ctx).Debug(status.String())
	statusChannel <- status
}

func messageToDataCollection(
	client *msgraphsdk.GraphServiceClient,
	ctx context.Context,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- DataStream,
	message models.Messageable,
	user string,
) error {
	var err error
	aMessage := message
	adtl := message.GetAdditionalData()
	if len(adtl) > 2 {
		aMessage, err = support.ConvertFromMessageable(adtl, message)
		if err != nil {
			return err
		}
	}
	if *aMessage.GetHasAttachments() {
		// getting all the attachments might take a couple attempts due to filesize
		var retriesErr error
		for count := 0; count < numberOfRetries; count++ {
			attached, err := client.
				UsersById(user).
				MessagesById(*aMessage.GetId()).
				Attachments().
				Get()
			retriesErr = err
			if err == nil && attached != nil {
				aMessage.SetAttachments(attached.GetValue())
				break
			}
		}
		if retriesErr != nil {
			logger.Ctx(ctx).Debug("exceeded maximum retries")
			return support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(retriesErr, "attachment failed"), nil)
		}
	}
	err = objectWriter.WriteObjectValue("", aMessage)
	if err != nil {
		return support.SetNonRecoverableError(errors.Wrapf(err, "%s", *aMessage.GetId()))
	}

	byteArray, err := objectWriter.GetSerializedContent()
	objectWriter.Close()
	if err != nil {
		return support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(err, "serializing mail content"), nil)
	}
	if byteArray != nil {
		dataChannel <- &ExchangeData{id: *aMessage.GetId(), message: byteArray, info: exchange.MessageInfo(aMessage)}
	}
	return nil
}

// SetStatus helper function
func (gc *GraphConnector) SetStatus(cos support.ConnectorOperationStatus) {
	gc.status = &cos
}

// AwaitStatus updates status field based on item within statusChannel.
func (gc *GraphConnector) AwaitStatus() *support.ConnectorOperationStatus {
	if gc.awaitingMessages > 0 {
		gc.status = <-gc.statusCh
		atomic.AddInt32(&gc.awaitingMessages, -1)
		return gc.status
	}
	return nil
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
