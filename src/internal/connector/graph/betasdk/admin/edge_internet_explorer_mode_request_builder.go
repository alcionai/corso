package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EdgeInternetExplorerModeRequestBuilder builds and executes requests for operations under \admin\edge\internetExplorerMode
type EdgeInternetExplorerModeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdgeInternetExplorerModeRequestBuilderInternal instantiates a new InternetExplorerModeRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeRequestBuilder) {
    m := &EdgeInternetExplorerModeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge/internetExplorerMode";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdgeInternetExplorerModeRequestBuilder instantiates a new InternetExplorerModeRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeRequestBuilderInternal(urlParams, requestAdapter)
}
// SiteLists the siteLists property
func (m *EdgeInternetExplorerModeRequestBuilder) SiteLists()(*EdgeInternetExplorerModeSiteListsRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteListsById gets an item from the github.com/alcionai/corso/src/internal/connector/graph/betasdk.admin.edge.internetExplorerMode.siteLists.item collection
func (m *EdgeInternetExplorerModeRequestBuilder) SiteListsById(id string)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSiteList%2Did"] = id
    }
    return NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
