package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef "betasdk/models/termstore"
)

// ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetQueryParameters
}
// ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Children()(*ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenRequestBuilder) {
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) ChildrenById(id string)(*ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) {
    m := &ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property children for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation children terms of set in term [store].
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property children in sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property children for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get children terms of set in term [store].
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property children in sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Patch(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Termable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Relations()(*ItemTermStoreSetsItemParentGroupSetsItemChildrenItemRelationsRequestBuilder) {
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) RelationsById(id string)(*ItemTermStoreSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder) Set()(*ItemTermStoreSetsItemParentGroupSetsItemChildrenItemSetRequestBuilder) {
    return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
