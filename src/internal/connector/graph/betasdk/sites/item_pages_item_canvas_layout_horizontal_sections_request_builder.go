package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters get a list of the horizontalSection objects and their properties. Sort by `id` in ascending order.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters struct {
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

// ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters
}

// ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderInternal instantiates a new HorizontalSectionsRequestBuilder and sets the default values.
//
//nolint:lll,wsl
func NewItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder instantiates a new HorizontalSectionsRequestBuilder and sets the default values.
//
//nolint:lll,wsl,revive
func NewItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder) Count() *ItemPagesItemCanvasLayoutHorizontalSectionsCountRequestBuilder {
	return NewItemPagesItemCanvasLayoutHorizontalSectionsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CreateGetRequestInformation get a list of the horizontalSection objects and their properties. Sort by `id` in ascending order.
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// CreatePostRequestInformation create new navigation property to horizontalSections for sites
//
//nolint:lll,wsl,errcheck
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// Get get a list of the horizontalSection objects and their properties. Sort by `id` in ascending order.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/horizontalsection-list?view=graph-rest-1.0
//
//nolint:lll, wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionCollectionResponseable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateHorizontalSectionCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionCollectionResponseable), nil
}

// nolint:lll,wsl
// Post create new navigation property to horizontalSections for sites
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilder) Post(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionable, error) {
	requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionable), nil
}
