package sharepoint

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists"
	"github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/columns"
	"github.com/pkg/errors"
)

func loadLists(
	ctx context.Context,
	gs graph.Service,
	identifier string,
) ([]models.Listable, error) {
	var (
		builder = gs.Client().SitesById(identifier).Lists()
		errs    error
	)
	//.Get(ctx, nil)

	listing := make([]models.Listable, 0)
	prefix := gs.Client().SitesById(identifier)

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return nil, support.WrapAndAppend(support.ConnectorStackErrorTrace(err), err, errs)
		}

		for _, entry := range resp.GetValue() {
			id := *entry.GetId()
			// Retrieve column data
			columnBuilder := prefix.ListsById(id).Columns()
			cols, _ := fetchColumns(ctx, gs, identifier, id)
			entry.SetColumns(cols)
			// get contentTypes

			if q2 != nil {
				cTypes, err := loadContentTypes(ctx, gs, identifier, q2)
				if err != nil {
					return nil, err
				}

				entry.SetContentTypes(cTypes)
			}

			q3, err := prefix.ListsById(id).Items().Get(ctx, nil)

			if err != nil {
				return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			if q3 != nil {
				items, _ := loadListItems(ctx, gs, identifier, id, q3)
				entry.SetItems(items)
			}

			listing = append(listing, entry)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = lists.NewListsRequestBuilder(*resp.GetOdataNextLink(), gs.Adapter())
	}

	return listing, nil
}

// Need to send with builder... how
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

func fetchContentTypes(
	ctx context.Context,
	gs graph.Service,
	identifier string,
	resp models.ContentTypeCollectionResponseable,
) ([]models.ContentTypeable, error) {
	cTypes := make([]models.ContentTypeable, 0)
	q2, _ := prefix.ListsById(id).ContentTypes().Get(ctx, nil)

	for _, cont := range resp.GetValue() {
		id := *cont.GetId()

		q1, _ := gs.Client().SitesById(identifier).
			ContentTypesById(id).ColumnLinks().Get(ctx, nil)
		if q1 != nil {
			cont.SetColumnLinks(q1.GetValue())
		}

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
		query, _ := gs.Client().
			SitesById(identifier).
			ContentTypesById(*cont.GetId()).
			Columns().
			Get(ctx, nil)

		if query != nil {
			columns, _ := loadColumns(ctx, gs, identifier, query)
			cont.SetColumns(columns)
		}

		cTypes = append(cTypes, cont)
	}

	return cTypes, nil
}
