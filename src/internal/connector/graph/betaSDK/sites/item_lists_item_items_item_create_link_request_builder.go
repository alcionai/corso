package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemListsItemItemsItemCreateLinkRequestBuilder provides operations to call the createLink method.
type ItemListsItemItemsItemCreateLinkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListsItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemCreateLinkRequestBuilderInternal instantiates a new CreateLinkRequestBuilder and sets the default values.
func NewItemListsItemItemsItemCreateLinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemCreateLinkRequestBuilder) {
    m := &ItemListsItemItemsItemCreateLinkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/microsoft.graph.createLink";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemListsItemItemsItemCreateLinkRequestBuilder instantiates a new CreateLinkRequestBuilder and sets the default values.
func NewItemListsItemItemsItemCreateLinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemCreateLinkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemCreateLinkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation create a sharing link for a listItem. The **createLink** action creates a new sharing link if the specified link type doesn't already exist for the calling application.If a sharing link of the specified type already exists for the app, this action will return the existing sharing link. **listItem** resources inherit sharing permissions from the list the item resides in.
func (m *ItemListsItemItemsItemCreateLinkRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemListsItemItemsItemCreateLinkPostRequestBodyable, requestConfiguration *ItemListsItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post create a sharing link for a listItem. The **createLink** action creates a new sharing link if the specified link type doesn't already exist for the calling application.If a sharing link of the specified type already exists for the app, this action will return the existing sharing link. **listItem** resources inherit sharing permissions from the list the item resides in.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/listitem-createlink?view=graph-rest-1.0
func (m *ItemListsItemItemsItemCreateLinkRequestBuilder) Post(ctx context.Context, body ItemListsItemItemsItemCreateLinkPostRequestBodyable, requestConfiguration *ItemListsItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Permissionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreatePermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Permissionable), nil
}
