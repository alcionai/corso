package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder builds and executes requests for operations under \admin\edge\internetExplorerMode\siteLists\{browserSiteList-id}
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    m := &EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Sites the sites property
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Sites()(*EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) SitesById(id string)(*EdgeInternetExplorerModeSiteListsItemSitesBrowserSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSite%2Did"] = id
    }
    return NewEdgeInternetExplorerModeSiteListsItemSitesBrowserSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
