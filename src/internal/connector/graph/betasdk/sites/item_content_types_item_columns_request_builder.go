package sites

import (
	"context"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemContentTypesItemColumnsRequestBuilder provides operations to manage the columns property of the microsoft.graph.contentType entity.
type ItemContentTypesItemColumnsRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemContentTypesItemColumnsRequestBuilderGetQueryParameters get the collection of columns, represented as [columnDefinition][columnDefinition] resources, in a [content type][contentType].
type ItemContentTypesItemColumnsRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemContentTypesItemColumnsRequestBuilderGetQueryParameters
}

// ItemContentTypesItemColumnsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentTypesItemColumnsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemContentTypesItemColumnsRequestBuilderInternal instantiates a new ColumnsRequestBuilder and sets the default values.
func NewItemContentTypesItemColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemColumnsRequestBuilder {
	m := &ItemContentTypesItemColumnsRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/columns{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemContentTypesItemColumnsRequestBuilder instantiates a new ColumnsRequestBuilder and sets the default values.
func NewItemContentTypesItemColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemColumnsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemContentTypesItemColumnsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemContentTypesItemColumnsRequestBuilder) Count() *ItemContentTypesItemColumnsCountRequestBuilder {
	return NewItemContentTypesItemColumnsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CreateGetRequestInformation get the collection of columns, represented as [columnDefinition][columnDefinition] resources, in a [content type][contentType].
func (m *ItemContentTypesItemColumnsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// CreatePostRequestInformation add a column to a [content type][contentType] in a site or list by specifying a [columnDefinition][columnDefinition].
func (m *ItemContentTypesItemColumnsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionable, requestConfiguration *ItemContentTypesItemColumnsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.Headers.Add("Accept", "application/json")
	requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// Get get the collection of columns, represented as [columnDefinition][columnDefinition] resources, in a [content type][contentType].
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-list-columns?view=graph-rest-1.0
func (m *ItemContentTypesItemColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionCollectionResponseable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateColumnDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionCollectionResponseable), nil
}

// Post add a column to a [content type][contentType] in a site or list by specifying a [columnDefinition][columnDefinition].
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-post-columns?view=graph-rest-1.0
func (m *ItemContentTypesItemColumnsRequestBuilder) Post(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionable, requestConfiguration *ItemContentTypesItemColumnsRequestBuilderPostRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionable, error) {
	requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateColumnDefinitionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionable), nil
}
