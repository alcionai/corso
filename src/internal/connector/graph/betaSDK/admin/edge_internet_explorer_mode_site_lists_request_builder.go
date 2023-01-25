package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EdgeInternetExplorerModeSiteListsRequestBuilder builds and executes requests for operations under \admin\edge\internetExplorerMode\siteLists
type EdgeInternetExplorerModeSiteListsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdgeInternetExplorerModeSiteListsRequestBuilderInternal instantiates a new SiteListsRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsRequestBuilder) {
    m := &EdgeInternetExplorerModeSiteListsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge/internetExplorerMode/siteLists";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdgeInternetExplorerModeSiteListsRequestBuilder instantiates a new SiteListsRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeSiteListsRequestBuilderInternal(urlParams, requestAdapter)
}
