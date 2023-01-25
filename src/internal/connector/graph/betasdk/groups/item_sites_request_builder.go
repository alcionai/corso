package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesRequestBuilder builds and executes requests for operations under \groups\{group-id}\sites
type ItemSitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Add provides operations to call the add method.
func (m *ItemSitesRequestBuilder) Add()(*ItemSitesAddRequestBuilder) {
    return NewItemSitesAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemSitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewItemSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesRequestBuilder) {
    m := &ItemSitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewItemSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSitesRequestBuilder) Count()(*ItemSitesCountRequestBuilder) {
    return NewItemSitesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delta provides operations to call the delta method.
func (m *ItemSitesRequestBuilder) Delta()(*ItemSitesDeltaRequestBuilder) {
    return NewItemSitesDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Remove provides operations to call the remove method.
func (m *ItemSitesRequestBuilder) Remove()(*ItemSitesRemoveRequestBuilder) {
    return NewItemSitesRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
