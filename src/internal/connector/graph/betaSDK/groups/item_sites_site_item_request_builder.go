package groups

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesSiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type ItemSitesSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesSiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type ItemSitesSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesSiteItemRequestBuilderGetQueryParameters
}
// ItemSitesSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Analytics()(*ItemSitesItemAnalyticsRequestBuilder) {
    return NewItemSitesItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Columns()(*ItemSitesItemColumnsRequestBuilder) {
    return NewItemSitesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ColumnsById(id string)(*ItemSitesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemSitesSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesSiteItemRequestBuilder) {
    m := &ItemSitesSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypes()(*ItemSitesItemContentTypesRequestBuilder) {
    return NewItemSitesItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypesById(id string)(*ItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesSiteItemRequestBuilder) Drive()(*ItemSitesItemDriveRequestBuilder) {
    return NewItemSitesItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Drives()(*ItemSitesItemDrivesRequestBuilder) {
    return NewItemSitesItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) DrivesById(id string)(*ItemSitesItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewItemSitesItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumns()(*ItemSitesItemExternalColumnsRequestBuilder) {
    return NewItemSitesItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumnsById(id string)(*ItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, error) {
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
func (m *ItemSitesSiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemSitesItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *ItemSitesSiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*ItemSitesItemGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemSitesItemGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *ItemSitesSiteItemRequestBuilder) GetByPathWithPath(path *string)(*ItemSitesItemGetByPathWithPathRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) InformationProtection()(*ItemSitesItemInformationProtectionRequestBuilder) {
    return NewItemSitesItemInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Items()(*ItemSitesItemItemsRequestBuilder) {
    return NewItemSitesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ItemsById(id string)(*ItemSitesItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return NewItemSitesItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Lists()(*ItemSitesItemListsRequestBuilder) {
    return NewItemSitesItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ListsById(id string)(*ItemSitesItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return NewItemSitesItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Onenote()(*ItemSitesItemOnenoteRequestBuilder) {
    return NewItemSitesItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Operations()(*ItemSitesItemOperationsRequestBuilder) {
    return NewItemSitesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) OperationsById(id string)(*ItemSitesItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return NewItemSitesItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Pages()(*ItemSitesItemPagesRequestBuilder) {
    return NewItemSitesItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) PagesById(id string)(*ItemSitesItemPagesSitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage%2Did"] = id
    }
    return NewItemSitesItemPagesSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) Patch(ctx context.Context, body ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.Siteable, error) {
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
func (m *ItemSitesSiteItemRequestBuilder) Permissions()(*ItemSitesItemPermissionsRequestBuilder) {
    return NewItemSitesItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) PermissionsById(id string)(*ItemSitesItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemSitesItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Sites()(*ItemSitesItemSitesRequestBuilder) {
    return NewItemSitesItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) SitesById(id string)(*ItemSitesItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return NewItemSitesItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) TermStore()(*ItemSitesItemTermStoreRequestBuilder) {
    return NewItemSitesItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
