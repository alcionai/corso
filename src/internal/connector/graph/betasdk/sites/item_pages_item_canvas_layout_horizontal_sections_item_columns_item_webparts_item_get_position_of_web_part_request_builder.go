package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal instantiates a new GetPositionOfWebPartRequestBuilder and sets the default values.
//
//nolint:lll,wsl
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/{webPart%2Did}/microsoft.graph.getPositionOfWebPart"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder instantiates a new GetPositionOfWebPartRequestBuilder and sets the default values.
//
//nolint:lll,revive,wsl
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}

// CreatePostRequestInformation invoke action getPositionOfWebPart
//
//nolint:lll,wsl
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// nolint:lll,wsl
// Post invoke action getPositionOfWebPart
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/webpart-getposition?view=graph-rest-1.0
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartPositionable, error) {
	requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateWebPartPositionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartPositionable), nil
}
