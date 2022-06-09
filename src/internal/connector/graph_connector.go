// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"fmt"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	"github.com/pkg/errors"
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	tenant  string
	adapter msgraphsdk.GraphRequestAdapter
	client  msgraphsdk.GraphServiceClient
	Users   map[string]string //key<email> value<id>
	Streams string            //Not implemented for ease of code check-in
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
		err = WrapAndAppend("general access", errors.New("connector failed: No access"), err)
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
			err = WrapAndAppend(gc.adapter.GetBaseUrl(), errors.New("user iteration failure"), err)
			return true
		}
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	iterateError = userIterator.Iterate(callbackFunc)
	if iterateError != nil {
		err = WrapAndAppend(gc.adapter.GetBaseUrl(), iterateError, err)
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
func (gc *GraphConnector) ExchangeDataCollection(user string) (DataCollection, error) {
	// TODO replace with completion of Issue 124:
	collection := NewExchangeDataCollection(user, []string{gc.tenant, user})
	//TODO: Retry handler to convert return: (DataCollection, error)
	return gc.serializeMessages(user, collection)

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

// serializeMessages: Temp Function as place Holder until Collections have been added
// to the GraphConnector struct.
func (gc *GraphConnector) serializeMessages(user string, dc ExchangeDataCollection) (DataCollection, error) {
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
	var errs error
	for _, aFolder := range folderList {
		result, err := gc.client.UsersById(user).MailFoldersById(aFolder).Messages().Get()
		if err != nil {
			errs = WrapAndAppend(user, err, errs)
		}
		if result == nil {
			errs = WrapAndAppend(user, fmt.Errorf("nil response on message query, folder: %s", aFolder), errs)
			continue
		}

		pageIterator, err := msgraphgocore.NewPageIterator(result, &gc.adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
		if err != nil {
			errs = WrapAndAppend(user, fmt.Errorf("iterator failed initialization: %v", err), errs)
			continue
		}
		objectWriter := kw.NewJsonSerializationWriter()

		callbackFunc := func(messageItem interface{}) bool {
			message, ok := messageItem.(models.Messageable)
			if !ok {
				errs = WrapAndAppend(user, fmt.Errorf("non-message return for user: %s", user), errs)
				return true
			}
			if *message.GetHasAttachments() {
				attached, err := gc.client.UsersById(user).MessagesById(*message.GetId()).Attachments().Get()
				if err == nil && attached != nil {
					message.SetAttachments(attached.GetValue())
				}
				if err != nil {
					errs = WrapAndAppend(*message.GetId(), fmt.Errorf("attachment failed: %v ", err), errs)
				}
			}
			err = objectWriter.WriteObjectValue("", message)
			if err != nil {
				errs = WrapAndAppend(*message.GetId(), err, errs)
				return true
			}
			byteArray, err = objectWriter.GetSerializedContent()
			objectWriter.Close()
			if err != nil {
				errs = WrapAndAppend(*message.GetId(), err, errs)
				return true
			}
			if byteArray != nil {
				dc.PopulateCollection(ExchangeData{id: *message.GetId(), message: byteArray})
			}
			return true
		}
		err = pageIterator.Iterate(callbackFunc)

		if err != nil {
			errs = WrapAndAppend(user, err, errs)
		}
	}
	fmt.Printf("Returning ExchangeDataColection with %d items\n", dc.Length())
	fmt.Printf("Errors: \n%s\n", errs.Error())
	dc.FinishPopulation()
	return &dc, errs
}
