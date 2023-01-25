package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SiteItemRequestBuilder provides operations to manage the collection of site entities.
type SiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SiteItemRequestBuilderGetQueryParameters retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
type SiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SiteItemRequestBuilderGetQueryParameters
}
// SiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Analytics()(*ItemAnalyticsRequestBuilder) {
    return NewItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Columns()(*ItemColumnsRequestBuilder) {
    return NewItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*ItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ContentTypes()(*ItemContentTypesRequestBuilder) {
    return NewItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*ItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return NewItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
func (m *SiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in sites by key (id)
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) DrivesById(id string)(*ItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ExternalColumns()(*ItemExternalColumnsRequestBuilder) {
    return NewItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*ItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/site-get?view=graph-rest-1.0
func (m *SiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *SiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *SiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*ItemGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *SiteItemRequestBuilder) GetByPathWithPath(path *string)(*ItemGetByPathWithPathRequestBuilder) {
    return NewItemGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) InformationProtection()(*ItemInformationProtectionRequestBuilder) {
    return NewItemInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Items()(*ItemItemsRequestBuilder) {
    return NewItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ItemsById(id string)(*ItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return NewItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Lists()(*ItemListsRequestBuilder) {
    return NewItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) ListsById(id string)(*ItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return NewItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Operations()(*ItemOperationsRequestBuilder) {
    return NewItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) OperationsById(id string)(*ItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return NewItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Pages()(*ItemPagesRequestBuilder) {
    return NewItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) PagesById(id string)(*ItemPagesSitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage%2Did"] = id
    }
    return NewItemPagesSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in sites by key (id)
func (m *SiteItemRequestBuilder) Patch(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Permissions()(*ItemPermissionsRequestBuilder) {
    return NewItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*ItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Sites()(*ItemSitesRequestBuilder) {
    return NewItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) SitesById(id string)(*ItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return NewItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) TermStore()(*ItemTermStoreRequestBuilder) {
    return NewItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
