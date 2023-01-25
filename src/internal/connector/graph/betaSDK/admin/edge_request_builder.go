package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EdgeRequestBuilder builds and executes requests for operations under \admin\edge
type EdgeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdgeRequestBuilderInternal instantiates a new EdgeRequestBuilder and sets the default values.
func NewEdgeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeRequestBuilder) {
    m := &EdgeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdgeRequestBuilder instantiates a new EdgeRequestBuilder and sets the default values.
func NewEdgeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeRequestBuilderInternal(urlParams, requestAdapter)
}
// InternetExplorerMode the internetExplorerMode property
func (m *EdgeRequestBuilder) InternetExplorerMode()(*EdgeInternetExplorerModeRequestBuilder) {
    return NewEdgeInternetExplorerModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
