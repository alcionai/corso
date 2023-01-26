package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder provides operations to count the resources in the collection.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters get the number of the resource
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
//
//nolint:lll
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/$count{?%24search,%24filter}"
	urlTplParams := make(map[string]string)

	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}

	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter

	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
//
//nolint:lll, revive
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl

	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateGetRequestInformation get the number of the resource
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration) (*int32, error) {
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
