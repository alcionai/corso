package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef "betasdk/models/termstore"
)

// ItemTermStoreGroupsItemSetsRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type ItemTermStoreGroupsItemSetsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreGroupsItemSetsRequestBuilderGetQueryParameters get a list of the set objects and their properties.
type ItemTermStoreGroupsItemSetsRequestBuilderGetQueryParameters struct {
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
// ItemTermStoreGroupsItemSetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsRequestBuilderGetQueryParameters
}
// ItemTermStoreGroupsItemSetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTermStoreGroupsItemSetsRequestBuilderInternal instantiates a new SetsRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreGroupsItemSetsRequestBuilder instantiates a new SetsRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemTermStoreGroupsItemSetsRequestBuilder) Count()(*ItemTermStoreGroupsItemSetsCountRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the set objects and their properties.
func (m *ItemTermStoreGroupsItemSetsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to sets for sites
func (m *ItemTermStoreGroupsItemSetsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, requestConfiguration *ItemTermStoreGroupsItemSetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get a list of the set objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/termstore-group-list-sets?view=graph-rest-1.0
func (m *ItemTermStoreGroupsItemSetsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsRequestBuilderGetRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.SetCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.SetCollectionResponseable), nil
}
// Post create new navigation property to sets for sites
func (m *ItemTermStoreGroupsItemSetsRequestBuilder) Post(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, requestConfiguration *ItemTermStoreGroupsItemSetsRequestBuilderPostRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable), nil
}
