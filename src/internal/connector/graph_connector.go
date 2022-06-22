// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"bytes"
	"context"
	"fmt"
	"io"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/corso/internal/connector/support"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/logger"
)

const (
	numberOfRetries = 3
	mailCategory    = "mail"
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	tenant  string
	adapter msgraphsdk.GraphRequestAdapter
	client  msgraphsdk.GraphServiceClient
	Users   map[string]string                 //key<email> value<id>
	Streams string                            //Not implemented for ease of code check-in
	status  *support.ConnectorOperationStatus // contains the status of the last run status
}

func NewGraphConnector(tenantId, clientId, secret string) (*GraphConnector, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := az.NewClientSecretCredential(tenantId, clientId, secret, nil)
	if err != nil {
		return nil, err
	}
	auth, err := ka.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, err
	}
	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}
	gc := GraphConnector{
		tenant:  tenantId,
		adapter: *adapter,
		client:  *msgraphsdk.NewGraphServiceClient(adapter),
		Users:   make(map[string]string, 0),
		status:  nil,
	}
	// TODO: Revisit Query all users.
	err = gc.setTenantUsers()
	if err != nil {
		return nil, err
	}
	return &gc, nil
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
	response, err := gc.client.Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return err
	}
	if response == nil {
		err = support.WrapAndAppend("general access", errors.New("connector failed: No access"), err)
		return err
	}
	userIterator, err := msgraphgocore.NewPageIterator(response, &gc.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return err
	}
	var iterateError error
	callbackFunc := func(userItem interface{}) bool {
		user, ok := userItem.(models.Userable)
		if !ok {
			err = support.WrapAndAppend(gc.adapter.GetBaseUrl(), errors.New("user iteration failure"), err)
			return true
		}
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	iterateError = userIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = support.WrapAndAppend(gc.adapter.GetBaseUrl(), iterateError, err)
	}
	return err
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	return buildFromMap(true, gc.Users)
}

func (gc *GraphConnector) GetUsersIds() []string {
	return buildFromMap(false, gc.Users)
}
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
// TODO: https://github.com/alcionai/corso/issues/135
//  Add iota to this call -> mail, contacts, calendar,  etc.
func (gc *GraphConnector) ExchangeDataCollection(ctx context.Context, user string) ([]DataCollection, error) {
	// TODO replace with completion of Issue 124:

	//TODO: Retry handler to convert return: (DataCollection, error)
	return gc.serializeMessages(ctx, user)
}

// optionsForMailFolders creates transforms the 'select' into a more dynamic call for MailFolders.
// var moreOps is a comma separated string of options(e.g. "displayName, isHidden")
// return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(moreOps []string) *msfolder.MailFoldersRequestBuilderGetRequestConfiguration {
	selecting := append(moreOps, "id")
	requestParameters := &msfolder.MailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msfolder.MailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options
}

// restoreMessages: Utility function to connect to M365 backstore
// and upload messages from DataCollection.
// FullPath: tenantId, userId, <mailCategory>, FolderId
func (gc *GraphConnector) restoreMessages(ctx context.Context, dc DataCollection) error {
	var errs error
	// must be user.GetId(), PrimaryName no longer works 6-15-2022
	user := dc.FullPath()[1]
	for {
		data, err := dc.NextItem()
		if err == io.EOF {
			break
		}

		buf := &bytes.Buffer{}
		_, err = buf.ReadFrom(data.ToReader())
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
		sentMessage, err := gc.client.UsersById(user).MailFoldersById(address).Messages().Post(clone)
		if err != nil {
			errs = support.WrapAndAppend(data.UUID()+": "+
				support.ConnectorStackErrorTrace(ctx, err), err, errs)
			continue
			// TODO: Add to retry Handler for the for failure
		}

		if sentMessage == nil && err == nil {
			errs = support.WrapAndAppend(data.UUID(), errors.New("Message not Sent: Blocked by server"), errs)

		}
		// This completes the restore loop for a message..
	}
	return errs
}

// serializeMessages: Temp Function as place Holder until Collections have been added
// to the GraphConnector struct.
func (gc *GraphConnector) serializeMessages(ctx context.Context, user string) ([]DataCollection, error) {
	options := optionsForMailFolders([]string{})
	response, err := gc.client.UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("unable to access folders for %s", user)
	}
	folderList := make([]string, 0)
	for _, folderable := range response.GetValue() {
		folderList = append(folderList, *folderable.GetId())
	}
	// Time to create Exchange data Holder
	var byteArray []byte
	collections := make([]DataCollection, 0)
	var errs error
	var totalItems, success int
	for _, aFolder := range folderList {

		result, err := gc.client.UsersById(user).MailFoldersById(aFolder).Messages().Get()
		if err != nil {
			errs = support.WrapAndAppend(user, err, errs)
		}
		if result == nil {
			errs = support.WrapAndAppend(user, fmt.Errorf("nil response on message query, folder: %s", aFolder), errs)
			continue
		}

		pageIterator, err := msgraphgocore.NewPageIterator(result, &gc.adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
		if err != nil {
			errs = support.WrapAndAppend(user, fmt.Errorf("iterator failed initialization: %v", err), errs)
			continue
		}
		objectWriter := kw.NewJsonSerializationWriter()
		edc := NewExchangeDataCollection(user, []string{gc.tenant, user, mailCategory, aFolder})

		callbackFunc := func(messageItem interface{}) bool {
			totalItems++
			message, ok := messageItem.(models.Messageable)
			if !ok {
				errs = support.WrapAndAppend(user, fmt.Errorf("non-message return for user: %s", user), errs)
				return true
			}
			if *message.GetHasAttachments() {
				// Retry Loop
				for count := 0; count < numberOfRetries; count++ {
					attached, err := gc.client.UsersById(user).MessagesById(*message.GetId()).Attachments().Get()
					if err == nil && attached != nil {
						message.SetAttachments(attached.GetValue())
						break
					}
				}
				if err != nil {
					logger.Ctx(ctx).Debug("exceeded maximum retries")
					errs = support.WrapAndAppend(*message.GetId(), fmt.Errorf("attachment failed: %v ", err), errs)
				}
			}
			err = objectWriter.WriteObjectValue("", message)
			if err != nil {
				errs = support.WrapAndAppend(*message.GetId(), err, errs)
				return true
			}
			byteArray, err = objectWriter.GetSerializedContent()
			objectWriter.Close()
			if err != nil {
				errs = support.WrapAndAppend(*message.GetId(), err, errs)
				return true
			}
			if byteArray != nil {
				edc.PopulateCollection(ExchangeData{id: *message.GetId(), message: byteArray})
			}
			return true
		}
		err = pageIterator.Iterate(callbackFunc)

		if err != nil {
			errs = support.WrapAndAppend(user, err, errs)
		}

		// Todo Retry Handler to be implemented
		edc.FinishPopulation()
		success += edc.Length()

		collections = append(collections, &edc)
	}

	status, err := support.CreateStatus(support.Backup, totalItems, success, len(folderList), errs)
	if err == nil {
		gc.UpdateStatus(*status)
		logger.Ctx(ctx).Debugw(gc.Status())
	}
	return collections, errs
}

// UpdateStatus helper function to status that has e
func (gc *GraphConnector) UpdateStatus(cos support.ConnectorOperationStatus) {
	gc.status = &cos
}

func (gc *GraphConnector) Status() string {
	if gc.status == nil {
		return ""
	}
	return gc.status.String()
}
