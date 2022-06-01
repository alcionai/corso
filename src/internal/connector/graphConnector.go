// Package connector uploads and retrieves data from M365 through
// the msgraph-go-sdk.
package connector

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/internal/connector/datautil"
)

// GraphConnector is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type GraphConnector struct {
	tenant  string
	adapter msgraphsdk.GraphRequestAdapter
	client  msgraphsdk.GraphServiceClient
	Users   map[string]string //key<email> value<id>
	errors  datautil.ErrorList
	Streams string //Not implemented for ease of code check-in
}

func NewGraphConnector(tenantId string, clientId string, secret string) GraphConnector {
	gc := GraphConnector{
		tenant: tenantId,
		users:  make(map[string]string, 0),
		errors: datautil.NewErrorList(),
	}
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := GetClientCredential(tenantId, clientId, secret)
	if err != nil {
		gc.errors.AddError(&err)
	}

	permissions := []string{"https://graph.microsoft.com/.default"}
	adapter, err := GetAdapterWithPermissions(cred, permissions)
	if err != nil {
		gc.errors.AddError(&err)
	}

	gc.SetAdapter(adapter)
	gc.SetClient(msgraphsdk.NewGraphServiceClient(adapter))
	gc.SetTenantUsers()
	return gc
}

// SetTenantUsers queries the M365 to identify the users in the
// workspace. The users field is updated from this return.
func (gc *GraphConnector) SetTenantUsers() {
	selecting := []string{"id, mail"}
	requestParams := &msuser.UsersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	response, err := gc.client.Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
	userIterator, err2 := msgraphgocore.NewPageIterator(response, &gc.adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil || err2 != nil {
		if err != nil {
			gc.errors.AddError(&err)
		}
		if err2 != nil {
			gc.errors.AddError(&err2)
		}
		fmt.Printf("Users not Updated\n%s\n", gc.errors.GetDetailedErrors())
		return
	}
	var hasFailed error
	callbackFunc := func(userItem interface{}) bool {
		if hasFailed != nil {
			fmt.Printf("Experienced err: %v\nOperation terminated", hasFailed)
			gc.errors.AddError(&hasFailed)
			return true
		}
		user := userItem.(models.Userable)
		gc.users[*user.GetMail()] = *user.GetId()
		return true
	}
	hasFailed = userIterator.Iterate(callbackFunc)
}

// DisplayErrorLogs prints the errors experienced during the session.
func (gc *GraphConnector) DisplayErrorLogs() {
	errorLog := gc.errors.GetDetailedErrors()
	fmt.Println(errorLog)
}

// GetAdapterWithPermissions is a utility method for creating an
// GraphRequestAdapter. The input is an Azure credential and a string defining
// the scope of application access. This is inline with the Azure application.
// The return is an GraphRequestAdapter and an error encountered during the
// authentication process.
func GetAdapterWithPermissions(cred azcore.TokenCredential, permission []string) (*msgraphsdk.GraphRequestAdapter, error) {
	auth, _ := ka.NewAzureIdentityAuthenticationProviderWithScopes(cred, permission)
	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	return adapter, err

}

// GetClientCredential is a credentialing method through Azure. Inputs are
// strings associated with application created in Azure Portal. It returns
// an Azure Credential and any validation error experienced.
func GetClientCredential(tenant string, clientId string, secret string) (*az.ClientSecretCredential, error) {
	cred, err := az.NewClientSecretCredential(
		tenant,
		clientId,
		secret,
		nil,
	)

	return cred, err
}

// GetUsers returns the email address of users within tenant.
func (gc *GraphConnector) GetUsers() []string {
	keys := make([]string, len(gc.users))
	for k := range gc.users {
		keys = append(keys, k)
	}
	return keys
}

func (gc *GraphConnector) GetUsersIds() []string {
	values := make([]string, len(gc.users))
	for _, v := range gc.users {
		values = append(values, v)
	}
	return values
}

// HasConnectionErrors is a helper method that returns true iff an error was encountered.
func (gc *GraphConnector) HasConnectorErrors() bool {
	if gc.errors.GetLength() != 0 {
		return true
	} else {
		return false
	}
}

func (gc *GraphConnector) SetAdapter(adapt *msgraphsdk.GraphRequestAdapter) {
	gc.adapter = *adapt
}

func (gc *GraphConnector) SetClient(client *msgraphsdk.GraphServiceClient) {
	gc.client = *client
}
