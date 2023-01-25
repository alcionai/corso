package sites

import (
	"context"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemColumnsItemSourceColumnRequestBuilder provides operations to manage the sourceColumn property of the microsoft.graph.columnDefinition entity.
type ItemColumnsItemSourceColumnRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemColumnsItemSourceColumnRequestBuilderGetQueryParameters the source column for content type column.
type ItemColumnsItemSourceColumnRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemColumnsItemSourceColumnRequestBuilderGetQueryParameters
}

// NewItemColumnsItemSourceColumnRequestBuilderInternal instantiates a new SourceColumnRequestBuilder and sets the default values.
func NewItemColumnsItemSourceColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemColumnsItemSourceColumnRequestBuilder {
	m := &ItemColumnsItemSourceColumnRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/columns/{columnDefinition%2Did}/sourceColumn{?%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemColumnsItemSourceColumnRequestBuilder instantiates a new SourceColumnRequestBuilder and sets the default values.
func NewItemColumnsItemSourceColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemColumnsItemSourceColumnRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemColumnsItemSourceColumnRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateGetRequestInformation the source column for content type column.
func (m *ItemColumnsItemSourceColumnRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// Get the source column for content type column.
func (m *ItemColumnsItemSourceColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ColumnDefinitionable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
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
