package api

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
)

// ---------------------------------------------------------------------------
// generic handler for paging item ids in a container
// ---------------------------------------------------------------------------

type itemPager interface {
	getPage(context.Context) (pageLinker, error)
	setNext(string)
	valuesIn(pageLinker) ([]getIDAndAddtler, error)
}

type pageLinker interface {
	GetOdataDeltaLink() *string
	GetOdataNextLink() *string
}

type getIDAndAddtler interface {
	GetId() *string
	GetAdditionalData() map[string]any
}

// uses a models interface compliant with { GetValues() []T }
// to transform its results into a slice of getIDer interfaces.
// Generics used here to handle the variation of msoft interfaces
// that all _almost_ comply with GetValue, but all return a different
// interface.
func toValues[T any](a any) ([]getIDAndAddtler, error) {
	gv, ok := a.(interface{ GetValue() []T })
	if !ok {
		return nil, errors.Errorf("response of type [%T] does not comply with the GetValue() interface", a)
	}

	items := gv.GetValue()
	r := make([]getIDAndAddtler, 0, len(items))

	for _, item := range items {
		var a any
		a = item

		ri, ok := a.(getIDAndAddtler)
		if !ok {
			return nil, errors.Errorf("item of type [%T] does not comply with the getIDAndAddtler interface", item)
		}

		r = append(r, ri)
	}

	return r, nil
}

// generic controller for retrieving all item ids in a container.
func getItemsAddedAndRemovedFromContainer(
	ctx context.Context,
	pager itemPager,
	errUpdater func(error),
) ([]string, []string, string, error) {
	var (
		addedIDs   = []string{}
		removedIDs = []string{}
		deltaURL   string
	)

	for {
		// get the next page of data, check for standard errors
		resp, err := pager.getPage(ctx)
		if err != nil {
			if err := graph.IsErrDeletedInFlight(err); err != nil {
				return nil, nil, deltaURL, err
			}

			if err := graph.IsErrInvalidDelta(err); err != nil {
				return nil, nil, deltaURL, err
			}

			return nil, nil, deltaURL, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		// each category type responds with a different interface, but all
		// of them comply with GetValue, which is where we'll get our item data.
		items, err := pager.valuesIn(resp)
		if err != nil {
			return nil, nil, "", err
		}

		// iterate through the items in the page
		for _, item := range items {
			// if the additional data conains a `@removed` key, the value will either
			// be 'changed' or 'deleted'.  We don't really care about the cause: both
			// cases are handled the same way in storage.
			if item.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
				addedIDs = append(addedIDs, *item.GetId())
			} else {
				removedIDs = append(removedIDs, *item.GetId())
			}
		}

		// the deltaLink is kind of like a cursor for overall data state.
		// once we run through pages of nextLinks, the last query will
		// produce a deltaLink instead (if supported), which we'll use on
		// the next backup to only get the changes since this run.
		delta := resp.GetOdataDeltaLink()
		if delta != nil && len(*delta) > 0 {
			deltaURL = *delta
		}

		// the nextLink is our page cursor within this query.
		// if we have more data to retrieve, we'll have a
		// nextLink instead of a deltaLink.
		nextLink := resp.GetOdataNextLink()
		if nextLink == nil || len(*nextLink) == 0 {
			break
		}

		pager.setNext(*nextLink)
	}

	return addedIDs, removedIDs, deltaURL, nil
}
