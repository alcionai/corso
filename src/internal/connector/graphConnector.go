// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"errors"
	"fmt"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
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

	err = gc.SetTenantUsers()
	if err != nil {
		return nil, err
	}
	return &gc, nil
}

// SetTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated during this method
// iff the return value is true
func (gc *GraphConnector) SetTenantUsers() error {
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
		if hasFailed != nil {
			fmt.Printf("Iteration err: %v\n", hasFailed)
			errorList = append(errorList, hasFailed)
			return true
		}
		user := userItem.(models.Userable)
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	hasFailed = userIterator.Iterate(callbackFunc)
	if len(errorList) > 0 {
		return errors.New(ConvertErrorList(errorList))
	}
	return nil
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
