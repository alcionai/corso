package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal instantiates a new GetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) {
    m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/{webPart%2Did}/microsoft.graph.getPositionOfWebPart";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder instantiates a new GetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getPositionOfWebPart
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getPositionOfWebPart
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/webpart-getposition?view=graph-rest-1.0
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.WebPartPositionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateWebPartPositionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.WebPartPositionable), nil
}
