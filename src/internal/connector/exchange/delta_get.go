package exchange

import (
	"context"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
)

//nolint:lll
const (
	mailURLTemplate     = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}"
	contactsURLTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/contacts/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}"
)

// The following functions are based off the code in v0.41.0 of msgraph-sdk-go
// for sending delta requests with query parameters.

func createGetRequestInformationWithRequestConfiguration(
	baseRequestInfoFunc func() (*abs.RequestInformation, error),
	requestConfig *DeltaRequestBuilderGetRequestConfiguration,
	template string,
) (*abs.RequestInformation, error) {
	requestInfo, err := baseRequestInfoFunc()
	if err != nil {
		return nil, err
	}

	requestInfo.UrlTemplate = template

	if requestConfig != nil {
		if requestConfig.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfig.QueryParameters))
		}

		requestInfo.AddRequestHeaders(requestConfig.Headers)
		requestInfo.AddRequestOptions(requestConfig.Options)
	}

	return requestInfo, nil
}

// TODO: ash --> verify
// From source: *msuser.users_item_mail_folders_item_messages_delta_request_builder.go
func sendMessagesDeltaGet(
	ctx context.Context,
	m *msuser.UsersItemMailFoldersItemMessagesDeltaRequestBuilder,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
	adapter abs.RequestAdapter,
) (models.UsersItemMailFoldersItemMessagesDeltaResponseable, error) {
	requestInfo, err := createGetRequestInformationWithRequestConfiguration(
		func() (*abs.RequestInformation, error) {
			return m.CreateGetRequestInformation(ctx, nil)
		},
		requestConfiguration,
		mailURLTemplate,
	)
	if err != nil {
		return nil, err
	}

	errorMapping := abs.ErrorMappings{
		"4XX": odataerrors.CreateODataErrorFromDiscriminatorValue,
		"5XX": odataerrors.CreateODataErrorFromDiscriminatorValue,
	}

	res, err := adapter.SendAsync(
		ctx,
		requestInfo,
		models.CreateUsersItemMailFoldersItemMessagesDeltaResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(models.UsersItemMailFoldersItemMessagesDeltaResponseable), nil
}

// From source: *msuser.users_item_contact_folders_item_contacts_delta_request_builder.go
func sendContactsDeltaGet(
	ctx context.Context,
	m *msuser.UsersItemContactFoldersItemContactsDeltaRequestBuilder,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
	adapter abs.RequestAdapter,
) (models.UsersItemContactFoldersItemContactsDeltaResponseable, error) {
	requestInfo, err := createGetRequestInformationWithRequestConfiguration(
		func() (*abs.RequestInformation, error) {
			return m.CreateGetRequestInformation(ctx, nil)
		},
		requestConfiguration,
		contactsURLTemplate,
	)
	if err != nil {
		return nil, err
	}

	errorMapping := abs.ErrorMappings{
		"4XX": odataerrors.CreateODataErrorFromDiscriminatorValue,
		"5XX": odataerrors.CreateODataErrorFromDiscriminatorValue,
	}

	res, err := adapter.SendAsync(
		ctx,
		requestInfo,
		models.CreateUsersItemContactFoldersItemContactsDeltaResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(models.UsersItemContactFoldersItemContactsDeltaResponseable), nil
}
