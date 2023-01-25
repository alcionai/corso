package betasdk

import (
    i9d9936085e596cbee3d3ea43f0107031f3ce3c13f794ab652b3161796b79f287 "betasdk/sites"
    ie003035fcfa8fa2ed6512191a3ea5bc68bda5b53921bb2c93f27715bb8a253a5 "betasdk/admin"
    ie67197848a429ece31dfb4a51132c19957ded9669dc32ee5b9224d55d9ad935c "betasdk/groups"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
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
func (m *BetaClient) Admin()(*ie003035fcfa8fa2ed6512191a3ea5bc68bda5b53921bb2c93f27715bb8a253a5.AdminRequestBuilder) {
    return ie003035fcfa8fa2ed6512191a3ea5bc68bda5b53921bb2c93f27715bb8a253a5.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBetaClient instantiates a new BetaClient and sets the default values.
func NewBetaClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, backingStore *IBackingStoreFactory)(*BetaClient) {
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
    m.requestAdapter.EnableBackingStore(backingStore);
    return m
}
// Groups the groups property
func (m *BetaClient) Groups()(*ie67197848a429ece31dfb4a51132c19957ded9669dc32ee5b9224d55d9ad935c.GroupsRequestBuilder) {
    return ie67197848a429ece31dfb4a51132c19957ded9669dc32ee5b9224d55d9ad935c.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the BetaSdk.groups.item collection
func (m *BetaClient) GroupsById(id string)(*ie67197848a429ece31dfb4a51132c19957ded9669dc32ee5b9224d55d9ad935c.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return ie67197848a429ece31dfb4a51132c19957ded9669dc32ee5b9224d55d9ad935c.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *BetaClient) Sites()(*i9d9936085e596cbee3d3ea43f0107031f3ce3c13f794ab652b3161796b79f287.SitesRequestBuilder) {
    return i9d9936085e596cbee3d3ea43f0107031f3ce3c13f794ab652b3161796b79f287.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the collection of site entities.
func (m *BetaClient) SitesById(id string)(*i9d9936085e596cbee3d3ea43f0107031f3ce3c13f794ab652b3161796b79f287.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i9d9936085e596cbee3d3ea43f0107031f3ce3c13f794ab652b3161796b79f287.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
