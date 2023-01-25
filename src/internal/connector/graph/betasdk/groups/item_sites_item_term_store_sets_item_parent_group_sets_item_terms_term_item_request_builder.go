package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
    i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/termstore"
)

// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetQueryParameters all the terms under the set.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Children()(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilder) {
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) ChildrenById(id string)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) {
    m := &ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property terms for groups
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all the terms under the set.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property terms in groups
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property terms for groups
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all the terms under the set.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable), nil
}
// Patch update the navigation property terms in groups
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Patch(ctx context.Context, body i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Relations()(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) RelationsById(id string)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsTermItemRequestBuilder) Set()(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemSetRequestBuilder) {
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
