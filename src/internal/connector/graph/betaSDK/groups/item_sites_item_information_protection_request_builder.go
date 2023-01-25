package groups

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.site entity.
type ItemSitesItemInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters get informationProtection from groups
type ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) Bitlocker()(*ItemSitesItemInformationProtectionBitlockerRequestBuilder) {
    return NewItemSitesItemInformationProtectionBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemSitesItemInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    m := &ItemSitesItemInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get informationProtection from groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property informationProtection in groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DecryptBuffer provides operations to call the decryptBuffer method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) DecryptBuffer()(*ItemSitesItemInformationProtectionDecryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property informationProtection for groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
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
// EncryptBuffer provides operations to call the encryptBuffer method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) EncryptBuffer()(*ItemSitesItemInformationProtectionEncryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection from groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Patch(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.InformationProtectionable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) Policy()(*ItemSitesItemInformationProtectionPolicyRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityLabels()(*ItemSitesItemInformationProtectionSensitivityLabelsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*ItemSitesItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityPolicySettings()(*ItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignDigest provides operations to call the signDigest method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SignDigest()(*ItemSitesItemInformationProtectionSignDigestRequestBuilder) {
    return NewItemSitesItemInformationProtectionSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*ItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifySignature provides operations to call the verifySignature method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) VerifySignature()(*ItemSitesItemInformationProtectionVerifySignatureRequestBuilder) {
    return NewItemSitesItemInformationProtectionVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
