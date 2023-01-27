package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
//
//nolint:lll
type SiteItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// SiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type SiteItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *SiteItemRequestBuilderGetQueryParameters
}

// SiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type SiteItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
// REMOVED Analytics for minimial

// Columns provides operations to manage the columns property of the microsoft.graph.site entity.

// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.

// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
//
//nolint:lll,wsl
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SiteItemRequestBuilder {
	m := &SiteItemRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}{?%24select,%24expand}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
//
//nolint:lll,revive,wsl
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SiteItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}

// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.

// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
// CreateGetRequestInformation retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
//
//nolint:lll,wsl
func (m *SiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
//
//nolint:lll,errcheck,wsl
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body msmodel.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Removed Drive() for minimial connector

// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
// Removed Drives()

// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.

// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.

// // ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
//
//	func (m *SiteItemRequestBuilder) ExternalColumnsById(id string) *ItemExternalColumnsColumnDefinitionItemRequestBuilder {
//		urlTplParams := make(map[string]string)
//		for idx, item := range m.pathParameters {
//			urlTplParams[idx] = item
//		}
//		if id != "" {
//			urlTplParams["columnDefinition%2Did"] = id
//		}
//		return NewItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
//	}
//
// Get retrieve properties and relationships for a [site][] resource.A **site** resource represents a team site in SharePoint.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/site-get?view=graph-rest-1.0
//
//nolint:wsl,revive,lll
func (m *SiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration) (msmodel.Siteable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, msmodel.CreateSiteFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(msmodel.Siteable), nil
}

//nolint:lll
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
// GetByPathWithPath provides operations to call the getByPath method.
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.site entity.
// Items provides operations to manage the items property of the microsoft.graph.site entity.
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.

// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
func (m *SiteItemRequestBuilder) Pages() *ItemPagesRequestBuilder {
	return NewItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PagesById provides operations to manage the pages property of the microsoft.graph.site entity.
//
//nolint:revive,wsl
func (m *SiteItemRequestBuilder) PagesById(id string) *ItemPagesSitePageItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["sitePage%2Did"] = id
	}
	return NewItemPagesSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Patch update entity in sites by key (id)
//
//nolint:wsl,revive,lll
func (m *SiteItemRequestBuilder) Patch(
	ctx context.Context,
	body msmodel.Siteable,
	requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration,
) (msmodel.Siteable, error) {
	requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, msmodel.CreateSiteFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(msmodel.Siteable), nil
}

// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
// func (m *SiteItemRequestBuilder) Sites()
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
//
//nolint:revive,wsl
func (m *SiteItemRequestBuilder) SitesById(id string) *ItemSitesSiteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["site%2Did1"] = id
	}
	return NewItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
