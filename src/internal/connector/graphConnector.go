// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
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
	errors  []error
	Streams string //Not implemented for ease of code check-in
}

func NewGraphConnector(tenantId string, clientId string, secret string) (*GraphConnector, error) {
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
		errors:  make([]error, 0),
	}
	gc.SetTenantUsers()
	return &gc, nil
}

// SetTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated during this method
// iff the return value is true
func (gc *GraphConnector) SetTenantUsers() bool {
	selecting := []string{"id, mail"}
	requestParams := &msuser.UsersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	response, err := gc.client.Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		gc.errors = append(gc.errors, err)
		return false
	}
	userIterator, err := msgraphgocore.NewPageIterator(response, &gc.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		gc.errors = append(gc.errors, err)
		return false
	}
	var hasFailed error
	callbackFunc := func(userItem interface{}) bool {
		if hasFailed != nil {
			fmt.Printf("Experienced err: %v\nOperation terminated", hasFailed)
			gc.errors = append(gc.errors, hasFailed)
			return true
		}
		user := userItem.(models.Userable)
		gc.Users[*user.GetMail()] = *user.GetId()
		return true
	}
	hasFailed = userIterator.Iterate(callbackFunc)
	return true
}

// DisplayErrorLogs prints the errors experienced during the session.
func (gc *GraphConnector) DisplayErrorLogs() string {
	errorLog := ""
	for idx, err := range gc.errors {
		errorLog = errorLog + fmt.Sprintf("Error# %d\t%v\n", idx, err)
	}
	return errorLog
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	keys := make([]string, len(gc.Users))
	for k := range gc.Users {
		keys = append(keys, k)
	}
	return keys
}

func (gc *GraphConnector) GetUsersIds() []string {
	values := make([]string, len(gc.Users))
	for _, v := range gc.Users {
		values = append(values, v)
	}
	return values
}

// HasConnectionErrors is a helper method that returns true iff an error was encountered.
func (gc *GraphConnector) HasConnectorErrors() bool {
	return len(gc.errors) > 0
}
