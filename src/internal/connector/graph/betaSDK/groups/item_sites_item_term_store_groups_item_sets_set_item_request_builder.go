package groups

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef "betasdk/models/termstore"
)

// ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Children()(*ItemSitesItemTermStoreGroupsItemSetsItemChildrenRequestBuilder) {
    return NewItemSitesItemTermStoreGroupsItemSetsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) ChildrenById(id string)(*ItemSitesItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return NewItemSitesItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) {
    m := &ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/groups/{group%2Did1}/sets/{set%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for groups
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all sets under the group in a term [store].
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sets in groups
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sets for groups
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get all sets under the group in a term [store].
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// ParentGroup provides operations to manage the parentGroup property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) ParentGroup()(*ItemSitesItemTermStoreGroupsItemSetsItemParentGroupRequestBuilder) {
    return NewItemSitesItemTermStoreGroupsItemSetsItemParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in groups
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Patch(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, requestConfiguration *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Setable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Relations()(*ItemSitesItemTermStoreGroupsItemSetsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermStoreGroupsItemSetsItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) RelationsById(id string)(*ItemSitesItemTermStoreGroupsItemSetsItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewItemSitesItemTermStoreGroupsItemSetsItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) Terms()(*ItemSitesItemTermStoreGroupsItemSetsItemTermsRequestBuilder) {
    return NewItemSitesItemTermStoreGroupsItemSetsItemTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoreGroupsItemSetsSetItemRequestBuilder) TermsById(id string)(*ItemSitesItemTermStoreGroupsItemSetsItemTermsTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return NewItemSitesItemTermStoreGroupsItemSetsItemTermsTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
