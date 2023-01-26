package betasdk

import (
	i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// BetaClient the main entry point of the SDK, exposes the configuration and the fluent API.
type BetaClient struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter *msgraphsdk.GraphRequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// NewBetaClient instantiates a new BetaClient and sets the default values.
// func NewBetaClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaClient) {
func NewBetaClient(requestAdapter *msgraphsdk.GraphRequestAdapter) *BetaClient {
	m := &BetaClient{}
	m.pathParameters = make(map[string]string)
	m.urlTemplate = "{+baseurl}"
	m.requestAdapter = requestAdapter

	if m.requestAdapter.GetBaseUrl() == "" {
		m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
	}
	return m
}

// Sites the sites property
func (m *BetaClient) Sites() *i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.SitesRequestBuilder {
	return i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SitesById provides operations to manage the collection of site entities.
func (m *BetaClient) SitesById(id string) *i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.SiteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["site%2Did"] = id
	}
	return i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
