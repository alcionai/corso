package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Lists() Lists {
	return Lists{c}
}

// Lists is an interface-compliant provider of the client.
type Lists struct {
	Client
}

// PostDrive creates a new list of type drive.  Specifically used to create
// documentLibraries for SharePoint Sites.
func (c Lists) PostDrive(
	ctx context.Context,
	siteID, driveName string,
) (models.Driveable, error) {
	list := models.NewList()
	list.SetDisplayName(&driveName)
	list.SetDescription(ptr.To("corso auto-generated restore destination"))

	li := models.NewListInfo()
	li.SetTemplate(ptr.To("documentLibrary"))
	list.SetList(li)

	// creating a list of type documentLibrary will result in the creation
	// of a new drive owned by the given site.
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists()

	newList, err := builder.Post(ctx, list, nil)
	if graph.IsErrItemAlreadyExistsConflict(err) {
		return nil, clues.StackWC(ctx, graph.ErrItemAlreadyExistsConflict, err)
	}

	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating documentLibrary list")
	}

	// drive information is not returned by the list creation.
	drive, err := builder.
		ByListId(ptr.Val(newList.GetId())).
		Drive().
		Get(ctx, nil)

	return drive, graph.Wrap(ctx, err, "fetching created documentLibrary").OrNil()
}

// SharePoint lists represent lists on a site. Inherits additional properties from
// baseItem: https://learn.microsoft.com/en-us/graph/api/resources/baseitem?view=graph-rest-1.0
// The full details concerning SharePoint Lists can
// be found at: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
// Note additional calls are required for the relationships that exist outside of the object properties.

// GetListById is a utility function to populate a SharePoint.List with objects associated with a given siteID.
// @param siteID the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func (c Lists) GetListByID(ctx context.Context, siteID, listID string) (models.Listable, error) {
	list, err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "fetching list")
	}

	cols, cTypes, lItems, err := c.getListContents(ctx, siteID, listID)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting list contents")
	}

	list.SetColumns(cols)
	list.SetContentTypes(cTypes)
	list.SetItems(lItems)

	return list, nil
}

// getListContents utility function to retrieve associated M365 relationships
// which are not included with the standard List query:
// - Columns, ContentTypes, ListItems
func (c Lists) getListContents(ctx context.Context, siteID, listID string) (
	[]models.ColumnDefinitionable,
	[]models.ContentTypeable,
	[]models.ListItemable,
	error,
) {
	cols, err := c.getColumns(ctx, siteID, listID, "")
	if err != nil {
		return nil, nil, nil, err
	}

	cTypes, err := c.getContentTypes(ctx, siteID, listID)
	if err != nil {
		return nil, nil, nil, err
	}

	lItems, err := c.getListItems(ctx, siteID, listID)
	if err != nil {
		return nil, nil, nil, err
	}

	return cols, cTypes, lItems, nil
}

// getListItems utility for retrieving ListItem data and the associated relationship
// data. Additional call append data to the tracked items, and do not create additional collections.
// Additional Call:
// * Fields
func (c Lists) getListItems(ctx context.Context, siteID, listID string) ([]models.ListItemable, error) {
	var (
		prefix = c.Stable.
			Client().
			Sites().
			BySiteId(siteID).
			Lists().
			ByListId(listID)
		builder = prefix.Items()
		itms    = make([]models.ListItemable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, err
		}

		for _, itm := range resp.GetValue() {
			newPrefix := prefix.Items().ByListItemId(ptr.Val(itm.GetId()))

			fields, err := newPrefix.Fields().Get(ctx, nil)
			if err != nil {
				return nil, graph.Wrap(ctx, err, "fetching list fields")
			}

			itm.SetFields(fields)

			itms = append(itms, itm)
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsItemItemsRequestBuilder(link, c.Stable.Adapter())
	}

	return itms, nil
}

// getColumns utility function to return columns from a site.
// An additional call required to check for details concerning the SourceColumn.
// For additional details:  https://learn.microsoft.com/en-us/graph/api/resources/columndefinition?view=graph-rest-1.0
// TODO: Refactor on if/else (dadams39)
func (c Lists) getColumns(ctx context.Context, siteID, listID, cTypeID string) ([]models.ColumnDefinitionable, error) {
	cs := make([]models.ColumnDefinitionable, 0)

	prefixBuilder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID)

	if len(cTypeID) == 0 {
		builder := prefixBuilder.Columns()

		for {
			resp, err := builder.Get(ctx, nil)
			if err != nil {
				return nil, graph.Wrap(ctx, err, "getting list columns")
			}

			cs = append(cs, resp.GetValue()...)

			link, ok := ptr.ValOK(resp.GetOdataNextLink())
			if !ok {
				break
			}

			builder = sites.NewItemListsItemColumnsRequestBuilder(link, c.Stable.Adapter())
		}
	} else {
		builder := prefixBuilder.ContentTypes().ByContentTypeId(cTypeID).Columns()

		for {
			resp, err := builder.Get(ctx, nil)
			if err != nil {
				return nil, graph.Wrap(ctx, err, "getting content columns")
			}

			cs = append(cs, resp.GetValue()...)

			link, ok := ptr.ValOK(resp.GetOdataNextLink())
			if !ok {
				break
			}

			builder = sites.NewItemListsItemContentTypesItemColumnsRequestBuilder(link, c.Stable.Adapter())
		}
	}

	return cs, nil
}

// getContentTypes retrieves all data for content type. Additional queries required
// for the following:
// - ColumnLinks
// - Columns
// Expand queries not used to retrieve the above. Possibly more than 20.
// Known Limitations: https://learn.microsoft.com/en-us/graph/known-issues#query-parameters
func (c Lists) getContentTypes(ctx context.Context, siteID, listID string) ([]models.ContentTypeable, error) {
	var (
		cTypes  = make([]models.ContentTypeable, 0)
		builder = c.Stable.
			Client().
			Sites().
			BySiteId(siteID).
			Lists().
			ByListId(listID).
			ContentTypes()
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, err
		}

		for _, cont := range resp.GetValue() {
			id := ptr.Val(cont.GetId())

			links, err := c.getColumnLinks(ctx, siteID, listID, id)
			if err != nil {
				return nil, err
			}

			cont.SetColumnLinks(links)

			cs, err := c.getColumns(ctx, siteID, listID, id)
			if err != nil {
				return nil, err
			}

			cont.SetColumns(cs)

			cTypes = append(cTypes, cont)
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsItemContentTypesRequestBuilder(link, c.Stable.Adapter())
	}

	return cTypes, nil
}

func (c Lists) getColumnLinks(ctx context.Context,
	siteID, listID, cTypeID string,
) ([]models.ColumnLinkable, error) {
	var (
		prefixBuilder = c.Stable.
				Client().
				Sites().
				BySiteId(siteID).
				Lists().
				ByListId(listID)
		builder = prefixBuilder.ContentTypes().ByContentTypeId(cTypeID).ColumnLinks()
		links   = make([]models.ColumnLinkable, 0)
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting column links")
		}

		links = append(links, resp.GetValue()...)

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsItemContentTypesItemColumnLinksRequestBuilder(link, c.Stable.Adapter())
	}

	return links, nil
}
