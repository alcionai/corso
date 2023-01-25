package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder builds and executes requests for operations under \admin\edge\internetExplorerMode\siteLists\{browserSiteList-id}\sites
type EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) {
    m := &EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}/sites";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) Count()(*EdgeInternetExplorerModeSiteListsItemSitesCountRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemSitesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
