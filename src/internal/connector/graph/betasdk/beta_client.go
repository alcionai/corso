package betasdk

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
    i2eb6d683e5428ebf135f99959034915f5e407cd2f29750b77adb5df72d90e064 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/groups"
    ibf5c397449f66e6fa06ac9bc6395bf62fb3c888bc28f9cd7cc6bd015dffe3817 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/admin"
)

// BetaClient the main entry point of the SDK, exposes the configuration and the fluent API.
type BetaClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Admin the admin property
func (m *BetaClient) Admin()(*ibf5c397449f66e6fa06ac9bc6395bf62fb3c888bc28f9cd7cc6bd015dffe3817.AdminRequestBuilder) {
    return ibf5c397449f66e6fa06ac9bc6395bf62fb3c888bc28f9cd7cc6bd015dffe3817.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBetaClient instantiates a new BetaClient and sets the default values.
func NewBetaClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaClient) {
    m := &BetaClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
    }
    return m
}
// Groups the groups property
func (m *BetaClient) Groups()(*i2eb6d683e5428ebf135f99959034915f5e407cd2f29750b77adb5df72d90e064.GroupsRequestBuilder) {
    return i2eb6d683e5428ebf135f99959034915f5e407cd2f29750b77adb5df72d90e064.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/alcionai/corso/src/internal/connector/graph/betasdk.groups.item collection
func (m *BetaClient) GroupsById(id string)(*i2eb6d683e5428ebf135f99959034915f5e407cd2f29750b77adb5df72d90e064.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return i2eb6d683e5428ebf135f99959034915f5e407cd2f29750b77adb5df72d90e064.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *BetaClient) Sites()(*i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.SitesRequestBuilder) {
    return i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the collection of site entities.
func (m *BetaClient) SitesById(id string)(*i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i1a3c1a5501c5e41b7fd169f2d4c768dce9b096ac28fb5431bf02afcc57295411.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
