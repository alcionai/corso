package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
type ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters get dataLossPreventionPolicies from sites
type ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters
}
// ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal instantiates a new DataLossPreventionPoliciesRequestBuilder and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    m := &ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/informationProtection/dataLossPreventionPolicies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemInformationProtectionDataLossPreventionPoliciesRequestBuilder instantiates a new DataLossPreventionPoliciesRequestBuilder and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Count()(*ItemInformationProtectionDataLossPreventionPoliciesCountRequestBuilder) {
    return NewItemInformationProtectionDataLossPreventionPoliciesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get dataLossPreventionPolicies from sites
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to dataLossPreventionPolicies for sites
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyable, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Evaluate provides operations to call the evaluate method.
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Evaluate()(*ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) {
    return NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get dataLossPreventionPolicies from sites
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDataLossPreventionPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyCollectionResponseable), nil
}
// Post create new navigation property to dataLossPreventionPolicies for sites
func (m *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Post(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyable, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDataLossPreventionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DataLossPreventionPolicyable), nil
}
