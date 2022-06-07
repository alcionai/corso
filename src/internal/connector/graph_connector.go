// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"errors"
	"fmt"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
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
	errorList := make([]error, 0)
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
		return errors.New("connector unable to complete queries. Verify credentials")
	}
	userIterator, err := msgraphgocore.NewPageIterator(response, &gc.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return err
	}
	var hasFailed error
	callbackFunc := func(userItem interface{}) bool {
		user, ok := userItem.(models.Userable)
		if !ok {
			errorList = append(errorList, errors.New("unable to iterable to user"))
			return true
		}
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	hasFailed = userIterator.Iterate(callbackFunc)
	if len(errorList) > 0 {
		return errors.New(ConvertErrorList(errorList))
	}
	return hasFailed
}

// ConvertsErrorList takes a list of errors and converts returns
// a string
// TODO: Place in error package after merged
func ConvertErrorList(errorList []error) string {
	errorLog := ""
	for idx, err := range errorList {
		errorLog = errorLog + fmt.Sprintf("Error# %d\t%v\n", idx, err)
	}
	return errorLog
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
	errorList := make([]error, 0)
	for _, folderable := range response.GetValue() {
		folderList = append(folderList, *folderable.GetId())
	}
	fmt.Printf("Folder List: %v\n", folderList)
	// Time to create Exchange data Holder
	var byteArray []byte
	var iterateError error
	for _, aFolder := range folderList {
		result, err := gc.client.UsersById(user).MailFoldersById(aFolder).Messages().Get()
		if err != nil {
			errorList = append(errorList, err)
		}
		if result == nil {
			fmt.Println("Cannot Get result")
		}

		pageIterator, err := msgraphgocore.NewPageIterator(result, &gc.adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
		if err != nil {
			errorList = append(errorList, err)
		}
		objectWriter := kw.NewJsonSerializationWriter()

		callbackFunc := func(messageItem interface{}) bool {
			message, ok := messageItem.(models.Messageable)
			if !ok {
				errorList = append(errorList, fmt.Errorf("unable to iterate on message for user: %s", user))
				return true
			}
			if *message.GetHasAttachments() {
				attached, err := gc.client.UsersById(user).MessagesById(*message.GetId()).Attachments().Get()
				if err == nil && attached != nil {
					message.SetAttachments(attached.GetValue())
				}
				if err != nil {
					err = fmt.Errorf("Attachment Error: " + err.Error())
					errorList = append(errorList, err)
				}
			}

			err = objectWriter.WriteObjectValue("", message)
			if err != nil {
				errorList = append(errorList, err)
				return true
			}
			byteArray, err = objectWriter.GetSerializedContent()
			objectWriter.Close()
			if err != nil {
				errorList = append(errorList, err)
				return true
			}
			if byteArray != nil {
				dc.PopulateCollection(ExchangeData{id: *message.GetId(), message: byteArray})
			}
			return true
		}
		iterateError = pageIterator.Iterate(callbackFunc)

		if iterateError != nil {
			errorList = append(errorList, err)
		}
	}
	fmt.Printf("Returning ExchangeDataColection with %d items\n", dc.Length())
	fmt.Printf("Errors: \n%s\n", ConvertErrorList(errorList))
	var errs error
	if len(errorList) > 0 {
		errs = errors.New(ConvertErrorList(errorList))
	}
	return &dc, errs
}
