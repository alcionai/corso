package sharepoint

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssite "github.com/microsoftgraph/msgraph-sdk-go/sites"
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
// @param siteID the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func loadLists(
	ctx context.Context,
	gs graph.Service,
	siteID string,
) ([]models.Listable, error) {
	var (
		prefix  = gs.Client().SitesById(siteID)
		builder = prefix.Lists()
		results = make([]models.Listable, 0)
		errs    error
	)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			id := *entry.GetId()

			cols, err := fetchColumns(ctx, gs, siteID, id, "")
			if err != nil {
				errs = support.WrapAndAppend(siteID, err, errs)
				continue
			}

			entry.SetColumns(cols)

			cTypes, err := fetchContentTypes(ctx, gs, siteID, id)
			if err != nil {
				errs = support.WrapAndAppend(siteID, err, errs)
				continue
			}

			entry.SetContentTypes(cTypes)

			lItems, err := fetchListItems(ctx, gs, siteID, id)
			if err != nil {
				errs = support.WrapAndAppend(siteID, err, errs)
				continue
			}

			entry.SetItems(lItems)

			results = append(results, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = mssite.NewSitesItemListsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	if errs != nil {
		return nil, errs
	}

	return results, nil
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

		builder = mssite.NewSitesItemListsItemItemsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	if errs != nil {
		return nil, errors.Wrap(errs, "fetchListItem unsuccessful")
	}

	return itms, nil
}

// fetchColumns utility function to return columns from a site.
// An additional call required to check for details concerning the SourceColumn.
// For additional details:  https://learn.microsoft.com/en-us/graph/api/resources/columndefinition?view=graph-rest-1.0
// TODO: Refactor on if/else (dadams39)
func fetchColumns(
	ctx context.Context,
	gs graph.Service,
	siteID, listID, cTypeID string,
) ([]models.ColumnDefinitionable, error) {
	cs := make([]models.ColumnDefinitionable, 0)

	if len(cTypeID) == 0 {
		builder := gs.Client().SitesById(siteID).ListsById(listID).Columns()

		for {
			resp, err := builder.Get(ctx, nil)
			if err != nil {
				return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, nil)
			}

			cs = append(cs, resp.GetValue()...)

			if resp.GetOdataNextLink() == nil {
				break
			}

			builder = mssite.NewSitesItemListsItemColumnsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
		}
	} else {
		builder := gs.Client().SitesById(siteID).ListsById(listID).ContentTypesById(cTypeID).Columns()

		for {
			resp, err := builder.Get(ctx, nil)
			if err != nil {
				return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			cs = append(cs, resp.GetValue()...)

			if resp.GetOdataNextLink() == nil {
				break
			}

			builder = mssite.NewSitesItemListsItemContentTypesItemColumnsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
		}
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
// TODO: Verify functionality after version upgrade or remove (dadams39) Check Stubs
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
			// TODO: stub for columPositions

			cs, err := fetchColumns(ctx, gs, siteID, listID, id)
			if err != nil {
				errs = support.WrapAndAppend("unable to populate columns for contentType", err, errs)
			}

			cont.SetColumns(cs)
			// TODO: stub for BaseTypes

			cTypes = append(cTypes, cont)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = mssite.NewSitesItemListsItemContentTypesRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
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

		builder = mssite.
			NewSitesItemListsItemContentTypesItemColumnLinksRequestBuilder(
				*resp.GetOdataNextLink(),
				gs.Adapter(),
			)
	}

	return links, nil
}
