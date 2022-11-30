package exchange

import (
	"context"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	mscontactdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/delta"
	msmaildelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/delta"
)

//nolint:lll
const (
	mailURLTemplate     = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}"
	contactsURLTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/contacts/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}"
)

// The following functions are based off the code in v0.41.0 of msgraph-sdk-go
// for sending delta requests with query parameters.

//nolint:unused
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

//nolint:unused
func sendMessagesDeltaGet(
	ctx context.Context,
	m *msmaildelta.DeltaRequestBuilder,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
	adapter abs.RequestAdapter,
) (msmaildelta.DeltaResponseable, error) {
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
		msmaildelta.CreateDeltaResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(msmaildelta.DeltaResponseable), nil
}

//nolint:unused
func sendContactsDeltaGet(
	ctx context.Context,
	m *mscontactdelta.DeltaRequestBuilder,
	requestConfiguration *DeltaRequestBuilderGetRequestConfiguration,
	adapter abs.RequestAdapter,
) (mscontactdelta.DeltaResponseable, error) {
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
		mscontactdelta.CreateDeltaResponseFromDiscriminatorValue,
		errorMapping,
	)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(mscontactdelta.DeltaResponseable), nil
}
