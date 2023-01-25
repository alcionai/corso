package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
type ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetQueryParameters collection of horizontal sections on the SharePoint page.
type ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetQueryParameters
}
// ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Columns provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) Columns()(*ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) {
    return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) ColumnsById(id string)(*ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["horizontalSectionColumn%2Did"] = id
    }
    return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderInternal instantiates a new HorizontalSectionItemRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) {
    m := &ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder instantiates a new HorizontalSectionItemRequestBuilder and sets the default values.
func NewItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property horizontalSections for sites
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation collection of horizontal sections on the SharePoint page.
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property horizontalSections in sites
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property horizontalSections for sites
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of horizontal sections on the SharePoint page.
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable), nil
}
// Patch update the navigation property horizontalSections in sites
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) Patch(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.HorizontalSectionable), nil
}
