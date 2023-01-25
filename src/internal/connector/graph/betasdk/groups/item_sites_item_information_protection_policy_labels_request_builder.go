package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
    i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
)

// ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
type ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetQueryParameters get a collection of information protection labels available to the user or to the organization.
type ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionPolicyLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) {
    m := &ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemInformationProtectionPolicyLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionPolicyLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) Count()(*ItemSitesItemInformationProtectionPolicyLabelsCountRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a collection of information protection labels available to the user or to the organization.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to labels for groups
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// EvaluateApplication provides operations to call the evaluateApplication method.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) EvaluateApplication()(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateClassificationResults provides operations to call the evaluateClassificationResults method.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) EvaluateClassificationResults()(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateRemoval provides operations to call the evaluateRemoval method.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) EvaluateRemoval()(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtractLabel provides operations to call the extractLabel method.
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) ExtractLabel()(*ItemSitesItemInformationProtectionPolicyLabelsExtractLabelRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get a collection of information protection labels available to the user or to the organization.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionpolicy-list-labels?view=graph-rest-1.0
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderGetRequestConfiguration)(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelCollectionResponseable), nil
}
// Post create new navigation property to labels for groups
func (m *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilder) Post(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsRequestBuilderPostRequestConfiguration)(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateInformationProtectionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.InformationProtectionLabelable), nil
}
