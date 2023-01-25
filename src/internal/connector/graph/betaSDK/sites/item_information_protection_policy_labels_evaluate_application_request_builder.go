package sites

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/informationProtection/policy/labels/microsoft.graph.evaluateApplication";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-1.0
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) Post(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemInformationProtectionPolicyLabelsEvaluateApplicationResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateItemInformationProtectionPolicyLabelsEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationProtectionPolicyLabelsEvaluateApplicationResponseable), nil
}
