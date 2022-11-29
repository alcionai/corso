package sharepoint

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnlinks"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnpositions"
	tc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columns"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/columns"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items"
	"github.com/pkg/errors"
)

// list.go contains additional functions to help retrieve SharePoint List data from M365
// SharePoint lists represent lists on a site. Inherits additional properties from
// baseItem: https://learn.microsoft.com/en-us/graph/api/resources/baseitem?view=graph-rest-1.0
// The full details concerning SharePoint Lists can
// be found at: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
// Note additional calls are required for the relationships that exist outside of the object properties.

// loadLists is a utility function to populate the List object.
// @param identifier the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func loadLists(
	ctx context.Context,
	gs graph.Service,
	identifier string,
) ([]models.Listable, error) {
	var (
		prefix  = gs.Client().SitesById(identifier)
		builder = prefix.Lists()
		listing = make([]models.Listable, 0)
		errs    error
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			id := *entry.GetId()
			// Retrieve List column data
			columnBuilder := prefix.ListsById(id).Columns()
			cols, err := fetchColumns(ctx, gs, identifier, columnBuilder)
			if err != nil {
				errs = support.WrapAndAppend(identifier, err, errs)
				continue
			}
			entry.SetColumns(cols)

			cTypes, err := fetchContentTypes(ctx, gs, identifier, prefix.ListsById(id).ContentTypes())
			if err != nil {
				errs = support.WrapAndAppend(identifier, err, errs)
				continue
			}

			entry.SetContentTypes(cTypes)

			items, err := fetchListItems(ctx, gs, identifier, id)
			if err != nil {
				errs = support.WrapAndAppend(identifier, err, errs)
			}

			entry.SetItems(items)

			listing = append(listing, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = lists.NewListsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return listing, nil
}

//
// Additional calls for
// * Analytics
// * DriveItem --> Item content
// * Fields
func fetchListItems(
	ctx context.Context,
	gs graph.Service,
	siteID, listID string,
) ([]models.ListItemable, error) {
	var (
		prefix  = gs.Client().SitesById(siteID).ListsById(listID)
		builder = prefix.Items()
		itms    = make([]models.ListItemable, 0)
		errs    error
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, itm := range resp.GetValue() {
			id := *itm.GetId()
			ana, err := prefix.ItemsById(id).Analytics().Get(ctx, nil)
			if err != nil {
				errs = errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}
			itm.SetAnalytics(ana)

			dItem, err := prefix.ItemsById(id).DriveItem().Get(ctx, nil)
			if err != nil {
				errs = errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}
			itm.SetDriveItem(dItem)

			fields, err := prefix.ItemsById(id).Fields().Get(ctx, nil)
			if err != nil {
				errs = errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			itm.SetFields(fields)

			itms = append(itms, itm)

		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = items.NewItemsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}
	if errs != nil {
		return nil, errs
	}

	return itms, nil
}

// Need to send with builder... how
// fetchColumns utility function to populate columns makes additional calls to
// retrieve Source C
func fetchColumns(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	cb *columns.ColumnsRequestBuilder,
) ([]models.ColumnDefinitionable, error) {
	var (
		builder = cb
		errs    error
	)

	cs := make([]models.ColumnDefinitionable, 0)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			source, err := gs.Client().
				SitesById(identifier).
				ColumnsById(*entry.GetId()).
				SourceColumn().
				Get(ctx, nil)
			if err != nil {
				errs = support.WrapAndAppend(
					"loadColumn unable to retrieve source: "+support.ConnectorStackErrorTrace(err),
					err,
					errs,
				)
				continue
			}
			entry.SetSourceColumn(source)

			cs = append(cs, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = columns.NewColumnsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return cs, nil
}

func fetchContentColumns(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	cb *tc.ColumnsRequestBuilder,
) ([]models.ColumnDefinitionable, error) {
	var (
		builder = cb
		errs    error
	)

	cs := make([]models.ColumnDefinitionable, 0)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			source, err := gs.Client().
				SitesById(identifier).
				ColumnsById(*entry.GetId()).
				SourceColumn().
				Get(ctx, nil)
			if err != nil {
				errs = support.WrapAndAppend(
					"load Content Column unable to retrieve source: "+support.ConnectorStackErrorTrace(err),
					err,
					errs,
				)
				continue
			}
			entry.SetSourceColumn(source)

			cs = append(cs, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = tc.NewColumnsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return cs, nil
}

// fetchContentTypes retrieves all data for content type. Additional queries required
// for the following:
// - ColumnLinks
// - ColumnPositions
// - BaseTypes
// - Columns
func fetchContentTypes(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	ctrb *contenttypes.ContentTypesRequestBuilder,
) ([]models.ContentTypeable, error) {
	var (
		cTypes  = make([]models.ContentTypeable, 0)
		builder = ctrb
		errs    error
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, cont := range resp.GetValue() {
			id := *cont.GetId()

			links, err := fetchColumnLinks(ctx, gs, identifier, id)
			if err != nil {
				errs = support.WrapAndAppend("unable to add column links to list", err, errs)
				break
			}
			cont.SetColumnLinks(links)

			q2, _ := gs.Client().SitesById(identifier).
				ContentTypesById(id).ColumnPositions().Get(ctx, nil)
			if q2 != nil {
				cont.SetColumnPositions(q2.GetValue())
			}

			q3, _ := gs.Client().SitesById(identifier).
				ContentTypesById(id).BaseTypes().Get(ctx, nil)
			if q3 != nil {
				cont.SetBaseTypes(q3.GetValue())
			}

			// Can we print Columns or another call?
			cBuilder := gs.Client().
				SitesById(identifier).
				ContentTypesById(*cont.GetId()).
				Columns()

			cs, err := fetchContentColumns(ctx, gs, identifier, cBuilder)
			if err != nil {
				errs = support.WrapAndAppend("unable to populate columns for contentType", err, errs)
			}
			cont.SetColumns(cs)

			cTypes = append(cTypes, cont)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = contenttypes.NewContentTypesRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return cTypes, nil
}

func fetchColumnLinks(ctx context.Context, gs graph.Service, siteID, cTypeID string) ([]models.ColumnLinkable, error) {
	var (
		builder = gs.Client().SitesById(siteID).ContentTypesById(cTypeID).ColumnLinks()
		links   = make([]models.ColumnLinkable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, link := range resp.GetValue() {
			links = append(links, link)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = columnlinks.NewColumnLinksRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return links, nil
}

func fetchColumnPositions(ctx context.Context, gs graph.Service, siteID, cTypeID string) ([]models.ColumnDefinitionable, error) {
	var (
		builder   = gs.Client().SitesById(siteID).ContentTypesById(cTypeID).ColumnPositions()
		positions = make([]models.ColumnDefinitionable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, pos := range resp.GetValue() {
			positions = append(positions, pos)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = columnpositions.NewColumnPositionsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return positions, nil

}
