package groups

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the **key** property. For information about how to read the **key** property, see Get bitlockerRecoveryKey.
type ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters
}
// NewItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal instantiates a new RecoveryKeysRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) {
    m := &ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/bitlocker/recoveryKeys{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder instantiates a new RecoveryKeysRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) Count()(*ItemSitesItemInformationProtectionBitlockerRecoveryKeysCountRequestBuilder) {
    return NewItemSitesItemInformationProtectionBitlockerRecoveryKeysCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the **key** property. For information about how to read the **key** property, see Get bitlockerRecoveryKey.
func (m *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the **key** property. For information about how to read the **key** property, see Get bitlockerRecoveryKey.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bitlocker-list-recoverykeys?view=graph-rest-1.0
func (m *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.BitlockerRecoveryKeyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateBitlockerRecoveryKeyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.BitlockerRecoveryKeyCollectionResponseable), nil
}
