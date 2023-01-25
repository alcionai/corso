package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef "betasdk/models/termstore"
)

// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.term entity.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetQueryParameters children of current term.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetQueryParameters struct {
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
// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetQueryParameters
}
// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderInternal instantiates a new ChildrenRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) {
    m := &ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/children{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder instantiates a new ChildrenRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) Count()(*ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) {
    return NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation children of current term.
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to children for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get children of current term.
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderGetRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.TermCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateTermCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.TermCollectionResponseable), nil
}
// Post create new navigation property to children for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) Post(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderPostRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable), nil
}
