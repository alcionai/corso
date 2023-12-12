package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

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
	cols, err := c.GetListColumns(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	cTypes, err := c.GetContentTypes(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	for i := 0; i < len(cTypes); i++ {
		columnLinks, err := c.GetColumnLinks(ctx, siteID, listID, ptr.Val(cTypes[i].GetId()), CallConfig{})
		if err != nil {
			return nil, nil, nil, err
		}

		cTypes[i].SetColumnLinks(columnLinks)

		cTypeColumns, err := c.GetCTypesColumns(ctx, siteID, listID, ptr.Val(cTypes[i].GetId()), CallConfig{})
		if err != nil {
			return nil, nil, nil, err
		}

		cTypes[i].SetColumns(cTypeColumns)
	}

	lItems, err := c.GetListItems(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	return cols, cTypes, lItems, nil
}
