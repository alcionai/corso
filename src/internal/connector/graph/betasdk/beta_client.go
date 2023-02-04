package betasdk

import (
	absser "github.com/microsoft/kiota-abstractions-go"
	kioser "github.com/microsoft/kiota-abstractions-go/serialization"
	kform "github.com/microsoft/kiota-serialization-form-go"
	kw "github.com/microsoft/kiota-serialization-json-go"
	ktext "github.com/microsoft/kiota-serialization-text-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
)

// BetaClient the main entry point of the SDK, exposes the configuration and the fluent API.
// Minimal Beta Connector:
// Details on how the Code was generated is present in `kioter-lock.json`.
// NOTE: kiota gen file is altered to indicate what files are included in the created
//
// Changes to Sites Directory:
// Access files send requests with an adapter's with ASync() support.
// This feature is not enabled in v1.0. Manually changed in remaining files.
// Additionally, only calls that begin as client.SitesBy(siteID).Pages() have an endpoint.
//
// The use case specific to Pages(). All other requests should be routed to the /internal/connector/graph.Servicer
// Specifics on `betaClient.SitesById(siteID).Pages` are located: sites/site_item_request_builder.go
//
// Required model files are identified as `modelFiles` in kiota-lock.json. Directory -> betasdk/models
// Required access files are identified as `sitesFiles` in kiota-lock.json. Directory -> betasdk/sites
//
// BetaClient minimal msgraph-beta-sdk-go for connecting to msgraph-beta-sdk-go
// for retrieving `SharePoint.Pages`. Code is generated from kiota.dev.
// requestAdapter is registered with the following the serializers:
// --  "Microsoft.Kiota.Serialization.Json.JsonParseNodeFactory",
// --  "Microsoft.Kiota.Serialization.Text.TextParseNodeFactory",
// --  "Microsoft.Kiota.Serialization.Form.FormParseNodeFactory"
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
	absser.RegisterDefaultSerializer(func() kioser.SerializationWriterFactory {
		return kw.NewJsonSerializationWriterFactory()
	})
	absser.RegisterDefaultSerializer(func() kioser.SerializationWriterFactory {
		return ktext.NewTextSerializationWriterFactory()
	})
	absser.RegisterDefaultSerializer(func() kioser.SerializationWriterFactory {
		return kform.NewFormSerializationWriterFactory()
	})
	absser.RegisterDefaultDeserializer(func() kioser.ParseNodeFactory {
		return kw.NewJsonParseNodeFactory()
	})
	absser.RegisterDefaultDeserializer(func() kioser.ParseNodeFactory {
		return ktext.NewTextParseNodeFactory()
	})
	absser.RegisterDefaultDeserializer(func() kioser.ParseNodeFactory {
		return kform.NewFormParseNodeFactory()
	})

	if m.requestAdapter.GetBaseUrl() == "" {
		m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
	}
	return m
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

// Adapter() helper method to export Adapter for iterating
func (m *BetaClient) Adapter() *msgraphsdk.GraphRequestAdapter {
	return m.requestAdapter
}
