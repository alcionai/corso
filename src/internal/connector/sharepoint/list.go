package sharepoint

import (
	"context"
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/columns"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/basetypes"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions"
	tc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columns"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/items"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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

			cTypes, err := fetchContentTypes(ctx, gs, identifier, id)
			if err != nil {
				errs = support.WrapAndAppend(identifier, err, errs)
				continue
			}

			entry.SetContentTypes(cTypes)

			lItems, err := fetchListItems(ctx, gs, identifier, id)
			if err != nil {
				errs = support.WrapAndAppend(identifier, err, errs)
				continue
			}

			entry.SetItems(lItems)

			listing = append(listing, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = lists.NewListsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	if errs != nil {
		return nil, errs
	}

	return listing, nil
}

// fetchListItems utility for retrieving ListItem data and the associated relationship
// data. Additional call append data to the tracked items, and do not create additional collections.
// Additional Call:
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
			newPrefix := prefix.ItemsById(*itm.GetId())

			fields, err := newPrefix.Fields().Get(ctx, nil)
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
		return nil, errors.Wrap(errs, "fetchListItem unsuccessful")
	}

	return itms, nil
}

// fetchColumns utility function to return columns from a site.
// An additional call required to check for details concerning the SourceColumn.
// For additional details:  https://learn.microsoft.com/en-us/graph/api/resources/columndefinition?view=graph-rest-1.0
func fetchColumns(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	cb *columns.ColumnsRequestBuilder,
) ([]models.ColumnDefinitionable, error) {
	var (
		builder = cb
		errs    error
		cs      = make([]models.ColumnDefinitionable, 0)
	)

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

// fetchContentColumns
// Same as fetchColumns. However, the Get() function calls are from different libraries.
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
// - Columns
// The following two are not included:
// - ColumnPositions
// - BaseTypes
// These relationships are not included as they following error from the API:
// itemNotFound Item not found: error status code received from the API
// Current as of github.com/microsoftgraph/msgraph-sdk-go v0.40.0
// TODO: Verify functionality after version upgrade or remove (dadams39)
func fetchContentTypes(
	ctx context.Context,
	gs graph.Service,
	siteID, listID string,
) ([]models.ContentTypeable, error) {
	var (
		cTypes  = make([]models.ContentTypeable, 0)
		builder = gs.Client().SitesById(siteID).ListsById(listID).ContentTypes()
		errs    error
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, cont := range resp.GetValue() {
			id := *cont.GetId()

			links, err := fetchColumnLinks(ctx, gs, siteID, listID, id)
			if err != nil {
				errs = support.WrapAndAppend("unable to add column links to list", err, errs)
				break
			}

			cont.SetColumnLinks(links)

			// positions, err := fetchColumnPositions(ctx, gs, siteID, listID, id)
			// if err != nil {
			// 	errs = support.WrapAndAppend("unable to add column definitionable to list", err, errs)
			// 	break
			// }

			// cont.SetColumnPositions(positions)

			cs, err := fetchContentColumns(ctx,
				gs,
				siteID,
				gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(*cont.GetId()).Columns())
			if err != nil {
				errs = support.WrapAndAppend("unable to populate columns for contentType", err, errs)
			}

			cont.SetColumns(cs)

			// bTypes, err := fetchContentBaseTypes(ctx, gs, siteID, listID, id)
			// if err != nil {
			// 	errs = support.WrapAndAppend("unable to add baseTypes to List", err, errs)
			// 	break
			// }

			// cont.SetBaseTypes(bTypes)

			cTypes = append(cTypes, cont)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = contenttypes.NewContentTypesRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	if errs != nil {
		return nil, errs
	}

	return cTypes, nil
}

func fetchContentBaseTypes(
	ctx context.Context,
	gs graph.Service,
	siteID, listID, cTypeID string,
) ([]models.ContentTypeable, error) {
	var (
		errs    error
		cTypes  = make([]models.ContentTypeable, 0)
		builder = gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(cTypeID).BaseTypes()
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		fmt.Println("Pass initial query?")

		for _, obj := range resp.GetValue() {
			cBuilder := gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(*obj.GetId()).Columns()

			cs, err := fetchContentColumns(ctx, gs, siteID, cBuilder)
			if err != nil {
				errs = support.WrapAndAppend("columns not added on fetch baseType", err, errs)
			}

			obj.SetColumns(cs)

			lnk, err := fetchColumnLinks(ctx, gs, siteID, listID, *obj.GetId())
			if err != nil {
				errs = support.WrapAndAppend("columnLink failure on fetch baseType", err, errs)
			}

			obj.SetColumnLinks(lnk)

			pos, err := fetchColumnPositions(ctx, gs, siteID, listID, *obj.GetId())
			if err != nil {
				errs = support.WrapAndAppend("column position not added on fetch baseType", err, errs)
			}

			obj.SetColumnPositions(pos)

			cTypes = append(cTypes, obj)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = basetypes.NewBaseTypesRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	if errs != nil {
		return nil, errs
	}

	return cTypes, nil
}

func fetchColumnLinks(
	ctx context.Context,
	gs graph.Service,
	siteID, listID, cTypeID string,
) ([]models.ColumnLinkable, error) {
	var (
		builder = gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(cTypeID).ColumnLinks()
		links   = make([]models.ColumnLinkable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		links = append(links, resp.GetValue()...)

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = columnlinks.NewColumnLinksRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return links, nil
}

func fetchColumnPositions(
	ctx context.Context,
	gs graph.Service,
	siteID, listID, cTypeID string,
) ([]models.ColumnDefinitionable, error) {
	var (
		builder   = gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(cTypeID).ColumnPositions()
		positions = make([]models.ColumnDefinitionable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		positions = append(positions, resp.GetValue()...)

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = columnpositions.NewColumnPositionsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return positions, nil
}
