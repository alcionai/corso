package sites

import i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"

// SitesRequestBuilder builds and executes requests for operations under \sites
type SitesRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// Add provides operations to call the add method.

// NewSitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SitesRequestBuilder {
	m := &SitesRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewSitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SitesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSitesRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *SitesRequestBuilder) Count() *CountRequestBuilder {
	return NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Delta provides operations to call the delta method.
/*
func (m *SitesRequestBuilder) Delta() *DeltaRequestBuilder {
	return NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Remove provides operations to call the remove method.
func (m *SitesRequestBuilder) Remove() *RemoveRequestBuilder {
	return NewRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
*/