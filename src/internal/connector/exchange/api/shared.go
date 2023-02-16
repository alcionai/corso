package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// generic handler for paging item ids in a container
// ---------------------------------------------------------------------------

type itemPager interface {
	getPage(context.Context) (api.DeltaPageLinker, error)
	setNext(string)
	valuesIn(api.DeltaPageLinker) ([]getIDAndAddtler, error)
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
		return nil, clues.Wrap(fmt.Errorf("%T", a), "does not comply with the GetValue() interface")
	}

	items := gv.GetValue()
	r := make([]getIDAndAddtler, 0, len(items))

	for _, item := range items {
		var a any = item

		ri, ok := a.(getIDAndAddtler)
		if !ok {
			return nil, clues.Wrap(fmt.Errorf("%T", item), "does not comply with the getIDAndAddtler interface")
		}

		r = append(r, ri)
	}

	return r, nil
}

// generic controller for retrieving all item ids in a container.
func getItemsAddedAndRemovedFromContainer(
	ctx context.Context,
	pager itemPager,
) ([]string, []string, string, error) {
	var (
		addedIDs   = []string{}
		removedIDs = []string{}
		deltaURL   string
	)

	itemCount := 0
	page := 0

	for {
		// get the next page of data, check for standard errors
		resp, err := pager.getPage(ctx)
		if err != nil {
			return nil, nil, deltaURL, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		// each category type responds with a different interface, but all
		// of them comply with GetValue, which is where we'll get our item data.
		items, err := pager.valuesIn(resp)
		if err != nil {
			return nil, nil, "", clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		itemCount += len(items)
		page++

		// Log every ~1000 items (the page size we use is 200)
		if page%5 == 0 {
			logger.Ctx(ctx).Infow("queried items", "count", itemCount)
		}

		// iterate through the items in the page
		for _, item := range items {
			// if the additional data conains a `@removed` key, the value will either
			// be 'changed' or 'deleted'.  We don't really care about the cause: both
			// cases are handled the same way in storage.
			if item.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
				addedIDs = append(addedIDs, ptr.Val(item.GetId()))
			} else {
				removedIDs = append(removedIDs, ptr.Val(item.GetId()))
			}
		}

		nextLink, delta := api.NextAndDeltaLink(resp)

		// the deltaLink is kind of like a cursor for overall data state.
		// once we run through pages of nextLinks, the last query will
		// produce a deltaLink instead (if supported), which we'll use on
		// the next backup to only get the changes since this run.
		if len(delta) > 0 {
			deltaURL = delta
		}

		// the nextLink is our page cursor within this query.
		// if we have more data to retrieve, we'll have a
		// nextLink instead of a deltaLink.
		if len(nextLink) == 0 {
			break
		}

		pager.setNext(nextLink)
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", itemCount)

	return addedIDs, removedIDs, deltaURL, nil
}
