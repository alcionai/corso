package sites

import (
	"context"

	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemContentTypesItemAssociateWithHubSitesRequestBuilder provides operations to call the associateWithHubSites method.
type ItemContentTypesItemAssociateWithHubSitesRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemContentTypesItemAssociateWithHubSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentTypesItemAssociateWithHubSitesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemContentTypesItemAssociateWithHubSitesRequestBuilderInternal instantiates a new AssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemContentTypesItemAssociateWithHubSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemAssociateWithHubSitesRequestBuilder {
	m := &ItemContentTypesItemAssociateWithHubSitesRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.associateWithHubSites"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemContentTypesItemAssociateWithHubSitesRequestBuilder instantiates a new AssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemContentTypesItemAssociateWithHubSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemAssociateWithHubSitesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemContentTypesItemAssociateWithHubSitesRequestBuilderInternal(urlParams, requestAdapter)
}

// CreatePostRequestInformation associate a [content type][contentType] with a list of hub sites.
func (m *ItemContentTypesItemAssociateWithHubSitesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemContentTypesItemAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemContentTypesItemAssociateWithHubSitesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// Post associate a [content type][contentType] with a list of hub sites.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-associatewithhubsites?view=graph-rest-1.0
func (m *ItemContentTypesItemAssociateWithHubSitesRequestBuilder) Post(ctx context.Context, body ItemContentTypesItemAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemContentTypesItemAssociateWithHubSitesRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}
