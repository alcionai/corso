package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder provides operations to count the resources in the collection.
type ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetQueryParameters
}
// NewItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder) {
    m := &ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/verticalSection/webparts/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutVerticalSectionWebpartsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
