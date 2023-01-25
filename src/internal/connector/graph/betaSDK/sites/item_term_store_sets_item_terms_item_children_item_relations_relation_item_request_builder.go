package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef "betasdk/models/termstore"
)

// ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}
// ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) {
    m := &ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/terms/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property relations for sites
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property relations in sites
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property relations for sites
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// FromTerm provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm()(*ItemTermStoreSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) {
    return NewItemTermStoreSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable), nil
}
// Patch update the navigation property relations in sites
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable, requestConfiguration *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie593ec024c01600085d102049951027d413f0d8e0f660388dd0e1804066bc3ef.Relationable), nil
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Set()(*ItemTermStoreSetsItemTermsItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemTermStoreSetsItemTermsItemChildrenItemRelationsItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToTerm provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm()(*ItemTermStoreSetsItemTermsItemChildrenItemRelationsItemToTermRequestBuilder) {
    return NewItemTermStoreSetsItemTermsItemChildrenItemRelationsItemToTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
