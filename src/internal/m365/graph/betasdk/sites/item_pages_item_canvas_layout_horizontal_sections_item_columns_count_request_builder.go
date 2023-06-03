package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder
// provides operations to count the resources in the collection.
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetQueryParameters get the number of the resource
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetQueryParameters
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
//
//nolint:lll
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/$count{?%24search,%24filter}"
	urlTplParams := make(map[string]string)

	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}

	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter

	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
//
//nolint:lll
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder(rawURL string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL

	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateGetRequestInformation get the number of the resource
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderGetRequestConfiguration) (*int32, error) {
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
