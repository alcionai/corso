package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/informationProtection/policy/labels/microsoft.graph.evaluateRemoval";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-1.0
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable), nil
}
