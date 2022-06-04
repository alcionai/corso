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
func ConvertErrorList(errorList []error) string {
	errorLog := ""
	for idx, err := range errorList {
		errorLog = errorLog + fmt.Sprintf("Error# %d\t%v\n", idx, err)
	}
	return errorLog
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	keys := make([]string, 0)
	for k := range gc.Users {
		keys = append(keys, k)
	}
	return keys
}

func (gc *GraphConnector) GetUsersIds() []string {
	values := make([]string, 0)
	for _, v := range gc.Users {
		values = append(values, v)
	}
	return values
}

// ExchangeDataStream returns a DataCollection that the caller can
// use to read mailbox data out for the specified user
// Assumption: The user asked for exists
// TODO: https://github.com/alcionai/corso/issues/135
//  Add iota to this call -> mail, contacts, calendar,  etc.
func (gc *GraphConnector) ExchangeDataCollection(user string) (DataCollection, error) {
	// Call M365. Find out how many emails the individual has Step I
	// Step 2: get the number..
	// Step 3: Query again get the
	total, err := gc.getCountForMail(user)
	if err != nil {
		return nil, err
	}

	// TODO replace with completion of Issue 124:
	//return &ExchangeDataCollection{user: user, ExpectedItems: total}, nil
	return gc.SerializeMessages(user, total)

}

// Internal Helper that is invoked when the data collection is created to populate it
// TODO: https://github.com/alcionai/corso/issues/124
func populateCollection(ExchangeDataCollection) error {
	// TODO: Read data for `ed.user` and add to collection
	return nil
}

func (gc *GraphConnector) getCountForMail(user string) (int, error) {
	options := optionsForMailFolders("totalItemCount")

	response, err := gc.client.UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil || response == nil {
		return 0, err
	}
	values := response.GetValue()
	sum := 0
	for _, folderable := range values {
		sum += int(*folderable.GetTotalItemCount())
	}
	fmt.Printf("Total messages is %d\n", sum)
	return sum, nil

}

func optionsForMailFolders(moreOps string) *msfolder.MailFoldersRequestBuilderGetRequestConfiguration {
	selecting := []string{fmt.Sprintf("id, %s", moreOps)}
	requestParameters := &msfolder.MailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msfolder.MailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options
}

func (gc *GraphConnector) SerializeMessages(user string, aInt int) (DataCollection, error) {
	options := optionsForMailFolders("")
	response, err := gc.client.UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("unable to access folders for %s", user)
	}
	folderList := make([]string, 0)
	errorList := make([]error, 0)
	values := response.GetValue()
	for _, folderable := range values {
		folderList = append(folderList, *folderable.GetId())
	}
	fmt.Printf("Folder List: %v\n", folderList)
	// Time to create Exchange data Holder
	collection := NewExchangeDataCollection(user, aInt, []string{gc.tenant, user})
	var byteArray []byte
	for idx, aFolder := range folderList {
		result, err := gc.client.UsersById(user).MailFoldersById(aFolder).Messages().Get()
		if err != nil {
			errorList = append(errorList, err)
		}
		pageIterator, err := msgraphgocore.NewPageIterator(result, &gc.adapter, models.CreateMessageFromDiscriminatorValue)
		if err != nil {
			errorList = append(errorList, err)
		}
		objectWriter := kw.NewJsonSerializationWriter()

		callbackFunc := func(messageItem interface{}) bool {
			if messageItem == nil {
				return true
			}
			message, ok := messageItem.(models.Messageable)
			if !ok || message == nil {
				errorList = append(errorList, fmt.Errorf("unable to iterable to user"))
				return true
			}
			if *message.GetHasAttachments() {
				attached, err := gc.client.UsersById(user).MessagesById(*message.GetId()).Attachments().Get()
				if err == nil && attached != nil {
					//message.SetAttachments(attached.GetValue())
					fmt.Println("Attachment fails some how")
				}

			}
			err = objectWriter.WriteObjectValue("", message)
			if err != nil {
				errorList = append(errorList, err)
				return true
			}
			byteArray, err = objectWriter.GetSerializedContent()
			if err != nil {
				errorList = append(errorList, err)
				return true
			}
			fmt.Printf("Fail on round: %d\n", idx)
			collection.PopulateCollection(ExchangeData{id: *message.GetId(), message: byteArray})
			return true
		}
		iterateError := pageIterator.Iterate(callbackFunc)
		if err != nil {
			errorList = append(errorList, err)
		}
		if iterateError != nil {
			errorList = append(errorList, err)
		}
	}
	fmt.Printf("Returning ExchangeDataColection with %d items\n", collection.GetLength())
	if len(errorList) > 0 {
		return &collection, errors.New(ConvertErrorList(errorList))
	}
	return &collection, nil
}
