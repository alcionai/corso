package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesCountRequestBuilder provides operations to count the resources in the collection.
type ItemPagesCountRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesCountRequestBuilderGetQueryParameters get the number of the resource
type ItemPagesCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
}

// ItemPagesCountRequestBuilderGetRequestConfiguration configuration for the request
// such as headers, query parameters, and middleware options.
type ItemPagesCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesCountRequestBuilderGetQueryParameters
}

// NewItemPagesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesCountRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemPagesCountRequestBuilder {
	m := &ItemPagesCountRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/$count{?%24search,%24filter}"
	urlTplParams := make(map[string]string)

	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}

	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter

	return m
}

// NewItemPagesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesCountRequestBuilder(
	rawURL string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemPagesCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL

	return NewItemPagesCountRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateGetRequestInformation get the number of the resource
//
//nolint:lll
func (m *ItemPagesCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "text/plain")

	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}

		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}

	return requestInfo, nil
}

// Get get the number of the resource
//
//nolint:lll
func (m *ItemPagesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesCountRequestBuilderGetRequestConfiguration) (*int32, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}

	res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*int32), nil
}
