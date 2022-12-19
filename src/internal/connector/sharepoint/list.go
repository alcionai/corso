package sharepoint

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssite "github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
)

type listTuple struct {
	name string
	id   string
}

func preFetchListOptions() *mssite.SitesItemListsRequestBuilderGetRequestConfiguration {
	selecting := []string{"id", "displayName"}
	queryOptions := mssite.SitesItemListsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mssite.SitesItemListsRequestBuilderGetRequestConfiguration{
		QueryParameters: &queryOptions,
	}

	return options
}

func preFetchListIDs(
	ctx context.Context,
	gs graph.Servicer,
	siteID string,
) ([]listTuple, error) {
	var (
		builder    = gs.Client().SitesById(siteID).Lists()
		options    = preFetchListOptions()
		listTuples = make([]listTuple, 0)
		errs       error
	)

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			temp := listTuple{id: *entry.GetId()}

			name := entry.GetDisplayName()
			if name != nil {
				temp.name = *name
			} else {
				temp.name = *entry.GetId()
			}

			listTuples = append(listTuples, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = mssite.NewSitesItemListsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return listTuples, nil
}

// list.go contains additional functions to help retrieve SharePoint List data from M365
// SharePoint lists represent lists on a site. Inherits additional properties from
// baseItem: https://learn.microsoft.com/en-us/graph/api/resources/baseitem?view=graph-rest-1.0
// The full details concerning SharePoint Lists can
// be found at: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
// Note additional calls are required for the relationships that exist outside of the object properties.

// loadSiteLists is a utility function to populate a collection of SharePoint.List
// objects associated with a given siteID.
// @param siteID the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func loadSiteLists(
	ctx context.Context,
	gs graph.Servicer,
	siteID string,
	listIDs []string,
) ([]models.Listable, error) {
	var (
		results = make([]models.Listable, 0)
		errs    error
	)

	for _, listID := range listIDs {
		entry, err := gs.Client().SitesById(siteID).ListsById(listID).Get(ctx, nil)
		if err != nil {
			errs = support.WrapAndAppend(
				listID,
				errors.Wrap(err, support.ConnectorStackErrorTrace(err)),
				errs,
			)
		}

		cols, cTypes, lItems, err := fetchListContents(ctx, gs, siteID, listID)
		if err == nil {
			entry.SetColumns(cols)
			entry.SetContentTypes(cTypes)
			entry.SetItems(lItems)
		} else {
			errs = support.WrapAndAppend("unable to fetchRelationships during loadSiteLists", err, errs)
			continue
		}

		results = append(results, entry)
	}

	if errs != nil {
		return nil, errs
	}

	return results, nil
}

// fetchListContents utility function to retrieve associated M365 relationships
// which are not included with the standard List query:
// - Columns, ContentTypes, ListItems
func fetchListContents(
	ctx context.Context,
	service graph.Servicer,
	siteID, listID string,
) (
	[]models.ColumnDefinitionable,
	[]models.ContentTypeable,
	[]models.ListItemable,
	error,
) {
	var errs error

	cols, err := fetchColumns(ctx, service, siteID, listID, "")
	if err != nil {
		errs = support.WrapAndAppend(siteID, err, errs)
	}

	cTypes, err := fetchContentTypes(ctx, service, siteID, listID)
	if err != nil {
		errs = support.WrapAndAppend(siteID, err, errs)
	}

	lItems, err := fetchListItems(ctx, service, siteID, listID)
	if err != nil {
		errs = support.WrapAndAppend(siteID, err, errs)
	}

	if errs != nil {
		return nil, nil, nil, errs
	}

	return cols, cTypes, lItems, nil
}

// fetchListItems utility for retrieving ListItem data and the associated relationship
// data. Additional call append data to the tracked items, and do not create additional collections.
// Additional Call:
// * Fields
func fetchListItems(
	ctx context.Context,
	gs graph.Servicer,
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
	gs graph.Servicer,
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
	gs graph.Servicer,
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
	gs graph.Servicer,
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
