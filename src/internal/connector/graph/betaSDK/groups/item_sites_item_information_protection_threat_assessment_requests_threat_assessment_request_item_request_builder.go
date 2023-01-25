package groups

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters get threatAssessmentRequests from groups
type ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal instantiates a new ThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    m := &ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder instantiates a new ThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property threatAssessmentRequests for groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get threatAssessmentRequests from groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property threatAssessmentRequests in groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property threatAssessmentRequests for groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get threatAssessmentRequests from groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable), nil
}
// Patch update the navigation property threatAssessmentRequests in groups
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) Patch(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderPatchRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ThreatAssessmentRequestable), nil
}
// Results provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) Results()(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResultsById provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) ResultsById(id string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentResult%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
