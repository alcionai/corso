// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"bytes"
	"context"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

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
	tenant        string
	queryService  *subConnector
	Users         map[string]string                 //key<email> value<id>
	status        *support.ConnectorOperationStatus // contains the status of the last run status
	statusChannel chan *support.ConnectorOperationStatus
	credentials   account.M365Config
}

type subConnector struct {
	client  msgraphsdk.GraphServiceClient
	adapter msgraphsdk.GraphRequestAdapter
}

func NewGraphConnector(acct account.Account) (*GraphConnector, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return nil, errors.Wrap(err, "retrieving m356 account configuration")
	}
	adapter, err := createAdapter(m365.TenantID, m365.ClientID, m365.ClientSecret)
	if err != nil {
		return nil, err
	}
	gc := GraphConnector{
		tenant: m365.TenantID,
		queryService: &subConnector{
			adapter: *adapter,
			client:  *msgraphsdk.NewGraphServiceClient(adapter),
		},
		Users:         make(map[string]string, 0),
		status:        nil,
		statusChannel: make(chan *support.ConnectorOperationStatus),
		credentials:   m365,
	}
	// TODO: Revisit Query all users.
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

// createSubConnector private constructor method for subConnector
func (gc *GraphConnector) createSubConnector() (*subConnector, error) {
	adapter, err := createAdapter(gc.credentials.TenantID, gc.credentials.ClientID, gc.credentials.ClientSecret)
	if err != nil {
		return nil, err
	}
	connector := subConnector{
		adapter: *adapter,
		client:  *msgraphsdk.NewGraphServiceClient(adapter),
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
	response, err := gc.queryService.client.Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return err
	}
	if response == nil {
		err = support.WrapAndAppend("general access", errors.New("connector failed: No access"), err)
		return err
	}
	userIterator, err := msgraphgocore.NewPageIterator(response, &gc.queryService.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return err
	}
	var iterateError error
	callbackFunc := func(userItem interface{}) bool {
		user, ok := userItem.(models.Userable)
		if !ok {
			err = support.WrapAndAppend(gc.queryService.adapter.GetBaseUrl(), errors.New("user iteration failure"), err)
			return true
		}
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	iterateError = userIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = support.WrapAndAppend(gc.queryService.adapter.GetBaseUrl(), iterateError, err)
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
			if user == selectors.All {
				errs = support.WrapAndAppend(
					"all-users",
					errors.New("all users selector currently not handled"),
					errs)
				continue
			}
			dcs, tasklist, err := gc.serializeMessages(ctx, user)
			if err != nil {
				return nil, support.WrapAndAppend(user, err, errs)
			}
			sub, err := gc.createSubConnector()
			if err != nil {
				return nil, support.WrapAndAppend(user, err, errs)
			}
			// async call to populate
			go populateFromTaskList(ctx, tasklist, *sub, dcs, gc.statusChannel)
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
func (gc *GraphConnector) RestoreMessages(ctx context.Context, dc DataCollection) error {
	var errs error
	// must be user.GetId(), PrimaryName no longer works 6-15-2022
	user := dc.FullPath()[1]
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			return support.WrapAndAppend("context cancelled", ctx.Err(), errs)
		case data, ok := <-items:
			if !ok {
				return errs
			}

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
			// details on valueId settings: https://docs.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxprops/77844470-22ca-43fb-993d-c53e96cf9cd6
			valueId := "Integer 0x0E07"
			enableValue := "4"
			sv := models.NewSingleValueLegacyExtendedProperty()
			sv.SetId(&valueId)
			sv.SetValue(&enableValue)
			svlep := []models.SingleValueLegacyExtendedPropertyable{sv}
			clone.SetSingleValueExtendedProperties(svlep)
			draft := false
			clone.SetIsDraft(&draft)
			sentMessage, err := gc.queryService.client.UsersById(user).MailFoldersById(address).Messages().Post(clone)
			if err != nil {
				errs = support.WrapAndAppend(data.UUID()+": "+
					support.ConnectorStackErrorTrace(err), err, errs)
				continue
				// TODO: Add to retry Handler for the for failure
			}

			if sentMessage == nil && err == nil {
				errs = support.WrapAndAppend(data.UUID(), errors.New("Message not Sent: Blocked by server"), errs)

			}
			// This completes the restore loop for a message..
		}
	}
}

// serializeMessages: Temp Function as place Holder until Collections have been added
// to the GraphConnector struct.
func (gc *GraphConnector) serializeMessages(ctx context.Context, user string) ([]ExchangeDataCollection, TaskList, error) {
	options := optionsForMessageSnapshot()
	response, err := gc.queryService.client.UsersById(user).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return nil, nil, err
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, &gc.queryService.adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, nil, err
	}
	tasklist := NewTaskList() // map[folder][] messageIds
	callbackFunc := func(messageItem any) bool {
		message, ok := messageItem.(models.Messageable)
		if !ok {
			err = support.WrapAndAppendf(gc.queryService.adapter.GetBaseUrl(), errors.New("message iteration failure"), err)
			return true
		}
		// Saving to messages to list. Indexed by folder
		tasklist.AddTask(*message.GetParentFolderId(), *message.GetId())
		return true
	}
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = support.WrapAndAppend(gc.queryService.adapter.GetBaseUrl(), iterateError, err)
	}
	if err != nil {
		return nil, nil, err // return error if snapshot is incomplete
	}
	// Create collection of ExchangeDataCollection and create  data Holder
	collections := make([]ExchangeDataCollection, 0)

	for aFolder := range tasklist {
		// prep the items for handoff to the backup consumer
		edc := NewExchangeDataCollection(user, []string{gc.tenant, user, mailCategory, aFolder})
		collections = append(collections, edc)
	}

	return collections, tasklist, err
}

// populateFromTaskList async call to fill DataCollection via a channel
func populateFromTaskList(
	context context.Context,
	tasklist TaskList,
	sc subConnector,
	collections []ExchangeDataCollection,
	statusChannel chan<- *support.ConnectorOperationStatus, // All with vairable must be made to channel functions
) {
	var errs error
	var attemptedItems, success int
	objectWriter := kw.NewJsonSerializationWriter()
	//Todo this has to return all the errors in the status
	for aFolder, tasks := range tasklist {
		// Get the same folder
		edc := SelectCollectionByLastIndex(aFolder, collections)
		if edc == nil {
			for _, task := range tasks {
				errs = support.WrapAndAppend(task, errors.New("unable to query: collection not found during populateFromTaskList"), errs)
			}
			continue
		}

		for _, task := range tasks {
			response, err := sc.client.UsersById(edc.user).MessagesById(task).Get()
			if err != nil {
				details := support.ConnectorStackErrorTrace(err)
				errs = support.WrapAndAppend(edc.user, errors.Wrapf(err, "unable to retrieve %s, %s", task, details), errs)
				continue
			}
			err = messageToDataCollection(&sc.client, context, objectWriter, *edc, response, edc.user)

			if err != nil {
				errs = support.WrapAndAppendf(edc.user, err, errs)
			}
		}
		edc.FinishPopulation()
		attemptedItems += len(tasks)
		success += edc.Length()
	}
	status, err := support.CreateStatus(support.Backup, attemptedItems, success, len(tasklist), errs)
	if err == nil {
		logger.Ctx(context).Debugw(status.String())
		statusChannel <- status
	}
}

func messageToDataCollection(
	client *msgraphsdk.GraphServiceClient,
	ctx context.Context,
	objectWriter *kw.JsonSerializationWriter,
	edc ExchangeDataCollection,
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
		edc.PopulateCollection(&ExchangeData{id: *aMessage.GetId(), message: byteArray})
	}
	return nil
}

// SetStatus helper function
func (gc *GraphConnector) SetStatus(cos support.ConnectorOperationStatus) {
	gc.status = &cos
}

func (gc *GraphConnector) RetrieveStatusFromChannel() *support.ConnectorOperationStatus {
	gc.status = <-gc.statusChannel
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

// IsRecoverableError returns true iff error is a RecoverableGCEerror
func IsRecoverableError(e error) bool {
	var recoverable *support.RecoverableGCError
	return errors.As(e, &recoverable)
}

// IsNonRecoverableError returns true iff error is a NonRecoverableGCEerror
func IsNonRecoverableError(e error) bool {
	var nonRecoverable *support.NonRecoverableGCError
	return errors.As(e, &nonRecoverable)
}
