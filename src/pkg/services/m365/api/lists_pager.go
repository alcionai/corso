package api

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// lists pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Listable] = &listsPageCtrl{}

type listsPageCtrl struct {
	siteID  string
	gs      graph.Servicer
	builder *sites.ItemListsRequestBuilder
	options *sites.ItemListsRequestBuilderGetRequestConfiguration
}

func (p *listsPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *listsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Listable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *listsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewListsPager(
	siteID string,
	cc CallConfig,
) *listsPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists()

	options := &sites.ItemListsRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &listsPageCtrl{
		siteID:  siteID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetLists fetches all lists in the site.
func (c Lists) GetLists(
	ctx context.Context,
	siteID string,
	cc CallConfig,
) ([]models.Listable, error) {
	pager := c.NewListsPager(siteID, cc)
	items, err := pagers.BatchEnumerateItems[models.Listable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// list items pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ListItemable] = &listItemsPageCtrl{}

type listItemsPageCtrl struct {
	siteID  string
	listID  string
	gs      graph.Servicer
	builder *sites.ItemListsItemItemsRequestBuilder
	options *sites.ItemListsItemItemsRequestBuilderGetRequestConfiguration
}

func (p *listItemsPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsItemItemsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *listItemsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ListItemable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *listItemsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewListItemsPager(
	siteID string,
	listID string,
	cc CallConfig,
) *listItemsPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Items()

	options := &sites.ItemListsItemItemsRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsItemItemsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &listItemsPageCtrl{
		siteID:  siteID,
		listID:  listID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetListItems fetches all list items in the list.
func (c Lists) GetListItems(
	ctx context.Context,
	siteID string,
	listID string,
	cc CallConfig,
) ([]models.ListItemable, error) {
	pager := c.NewListItemsPager(siteID, listID, cc)
	items, err := pagers.BatchEnumerateItems[models.ListItemable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// columns pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ColumnDefinitionable] = &columnsPageCtrl{}

type columnsPageCtrl struct {
	siteID  string
	listID  string
	gs      graph.Servicer
	builder *sites.ItemListsItemColumnsRequestBuilder
	options *sites.ItemListsItemColumnsRequestBuilderGetRequestConfiguration
}

func (p *columnsPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsItemColumnsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *columnsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ColumnDefinitionable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *columnsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewColumnsPager(
	siteID string,
	listID string,
	cc CallConfig,
) *columnsPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Columns()

	options := &sites.ItemListsItemColumnsRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsItemColumnsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &columnsPageCtrl{
		siteID:  siteID,
		listID:  listID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetListColumns fetches all list columns in the list.
func (c Lists) GetListColumns(
	ctx context.Context,
	siteID string,
	listID string,
	cc CallConfig,
) ([]models.ColumnDefinitionable, error) {
	pager := c.NewColumnsPager(siteID, listID, cc)
	items, err := pagers.BatchEnumerateItems[models.ColumnDefinitionable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// content types pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ContentTypeable] = &contentTypesPageCtrl{}

type contentTypesPageCtrl struct {
	siteID  string
	listID  string
	gs      graph.Servicer
	builder *sites.ItemListsItemContentTypesRequestBuilder
	options *sites.ItemListsItemContentTypesRequestBuilderGetRequestConfiguration
}

func (p *contentTypesPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsItemContentTypesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contentTypesPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ContentTypeable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *contentTypesPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewContentTypesPager(
	siteID string,
	listID string,
	cc CallConfig,
) *contentTypesPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		ContentTypes()

	options := &sites.ItemListsItemContentTypesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsItemContentTypesRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &contentTypesPageCtrl{
		siteID:  siteID,
		listID:  listID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetContentTypes fetches all content types in the list.
func (c Lists) GetContentTypes(
	ctx context.Context,
	siteID string,
	listID string,
	cc CallConfig,
) ([]models.ContentTypeable, error) {
	pager := c.NewContentTypesPager(siteID, listID, cc)
	items, err := pagers.BatchEnumerateItems[models.ContentTypeable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// content types columns pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ColumnDefinitionable] = &cTypesColumnsPageCtrl{}

type cTypesColumnsPageCtrl struct {
	siteID  string
	listID  string
	cTypeID string
	gs      graph.Servicer
	builder *sites.ItemListsItemContentTypesItemColumnsRequestBuilder
	options *sites.ItemListsItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration
}

func (p *cTypesColumnsPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsItemContentTypesItemColumnsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *cTypesColumnsPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ColumnDefinitionable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *cTypesColumnsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewCTypesColumnsPager(
	siteID string,
	listID string,
	cTypeID string,
	cc CallConfig,
) *cTypesColumnsPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		ContentTypes().
		ByContentTypeId(cTypeID).
		Columns()

	options := &sites.ItemListsItemContentTypesItemColumnsRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsItemContentTypesItemColumnsRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &cTypesColumnsPageCtrl{
		siteID:  siteID,
		listID:  listID,
		cTypeID: cTypeID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetCTypesColumns fetches all columns in the content type.
func (c Lists) GetCTypesColumns(
	ctx context.Context,
	siteID string,
	listID string,
	cTypeID string,
	cc CallConfig,
) ([]models.ColumnDefinitionable, error) {
	pager := c.NewCTypesColumnsPager(siteID, listID, cTypeID, cc)
	items, err := pagers.BatchEnumerateItems[models.ColumnDefinitionable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// column links pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ColumnLinkable] = &columnLinksPageCtrl{}

type columnLinksPageCtrl struct {
	siteID  string
	listID  string
	cTypeID string
	gs      graph.Servicer
	builder *sites.ItemListsItemContentTypesItemColumnLinksRequestBuilder
	options *sites.ItemListsItemContentTypesItemColumnLinksRequestBuilderGetRequestConfiguration
}

func (p *columnLinksPageCtrl) SetNextLink(nextLink string) {
	p.builder = sites.NewItemListsItemContentTypesItemColumnLinksRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *columnLinksPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ColumnLinkable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *columnLinksPageCtrl) ValidModTimes() bool {
	return true
}

func (c Lists) NewColumnLinksPager(
	siteID string,
	listID string,
	cTypeID string,
	cc CallConfig,
) *columnLinksPageCtrl {
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		ContentTypes().
		ByContentTypeId(cTypeID).
		ColumnLinks()

	options := &sites.ItemListsItemContentTypesItemColumnLinksRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemListsItemContentTypesItemColumnLinksRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	return &columnLinksPageCtrl{
		siteID:  siteID,
		listID:  listID,
		cTypeID: cTypeID,
		builder: builder,
		gs:      c.Stable,
		options: options,
	}
}

// GetColumnLinks fetches all column links in the content type.
func (c Lists) GetColumnLinks(
	ctx context.Context,
	siteID string,
	listID string,
	cTypeID string,
	cc CallConfig,
) ([]models.ColumnLinkable, error) {
	pager := c.NewColumnLinksPager(siteID, listID, cTypeID, cc)
	items, err := pagers.BatchEnumerateItems[models.ColumnLinkable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}
