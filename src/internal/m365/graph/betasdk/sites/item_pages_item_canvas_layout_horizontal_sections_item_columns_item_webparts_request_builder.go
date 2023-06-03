package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/m365/graph/betasdk/models"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters get the webPart resources from a sitePage. Sort by the order in which they appear on the page.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters struct {
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

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal instantiates a new WebpartsRequestBuilder and sets the default values.
//
//nolint:lll,wsl
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder instantiates a new WebpartsRequestBuilder and sets the default values.
//
//nolint:lll,revive,wsl
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Count() *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder {
	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CreateGetRequestInformation get the webPart resources from a sitePage. Sort by the order in which they appear on the page.
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// CreatePostRequestInformation create new navigation property to webparts for sites
//
//nolint:lll,wsl,errcheck
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// Get get the webPart resources from a sitePage. Sort by the order in which they appear on the page.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/webpart-list?view=graph-rest-1.0
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartCollectionResponseable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateWebPartCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartCollectionResponseable), nil
}

// Post create new navigation property to webparts for sites
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Post(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable, error) {
	requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateWebPartFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable), nil
}
