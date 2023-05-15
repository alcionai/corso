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
	// getPage get a page with the specified options from graph
	getPage(context.Context) (api.PageLinker, error)
	// setNext is used to pass in the next url got from graph
	setNext(string)
	// reset is used to clear delta url in delta pagers. When
	// reset is called, we reset the state(delta url) that we
	// currently have and start a new delta query without the token.
	reset(context.Context)
	// valuesIn gets us the values in a page
	valuesIn(api.PageLinker) ([]getIDAndAddtler, error)
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
		return nil, clues.New(fmt.Sprintf("type does not comply with the GetValue() interface: %T", a))
	}

	items := gv.GetValue()
	r := make([]getIDAndAddtler, 0, len(items))

	for _, item := range items {
		var a any = item

		ri, ok := a.(getIDAndAddtler)
		if !ok {
			return nil, clues.New(fmt.Sprintf("type does not comply with the getIDAndAddtler interface: %T", item))
		}

		r = append(r, ri)
	}

	return r, nil
}

func getAddedAndRemovedItemIDs(
	ctx context.Context,
	service graph.Servicer,
	pager itemPager,
	deltaPager itemPager,
	oldDelta string,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	var (
		pgr        itemPager
		resetDelta bool
	)

	if canMakeDeltaQueries {
		pgr = deltaPager
		resetDelta = len(oldDelta) == 0
	} else {
		pgr = pager
		resetDelta = true
	}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	// note: happy path, not the error condition
	if err == nil {
		return added, removed, DeltaUpdate{deltaURL, resetDelta}, err
	}

	// If we already tried with a non-delta url, we can return
	if !canMakeDeltaQueries {
		return nil, nil, DeltaUpdate{}, err
	}

	// return error if invalid not delta error or oldDelta was empty
	if !graph.IsErrInvalidDelta(err) || len(oldDelta) == 0 {
		return nil, nil, DeltaUpdate{}, err
	}

	// reset deltaPager
	pgr.reset(ctx)

	added, removed, deltaURL, err = getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, true}, nil
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
		nextLink   string
		deltaLink  string
	)

	itemCount := 0
	page := 0

	for {
		// get the next page of data, check for standard errors
		resp, err := pager.getPage(ctx)
		if err != nil {
			return nil, nil, deltaURL, graph.Stack(ctx, err)
		}

		// each category type responds with a different interface, but all
		// of them comply with GetValue, which is where we'll get our item data.
		items, err := pager.valuesIn(resp)
		if err != nil {
			return nil, nil, "", graph.Stack(ctx, err)
		}

		itemCount += len(items)
		page++

		// Log every ~1000 items (the page size we use is 200)
		if page%5 == 0 {
			logger.Ctx(ctx).Infow("queried items", "count", itemCount)
		}

		// iterate through the items in the page
		for _, item := range items {
			// if the additional data contains a `@removed` key, the value will either
			// be 'changed' or 'deleted'.  We don't really care about the cause: both
			// cases are handled the same way in storage.
			if item.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
				addedIDs = append(addedIDs, ptr.Val(item.GetId()))
			} else {
				removedIDs = append(removedIDs, ptr.Val(item.GetId()))
			}
		}

		dresp, ok := resp.(api.DeltaPageLinker)
		if ok {
			nextLink, deltaLink = api.NextAndDeltaLink(dresp)
		} else {
			nextLink = api.NextLink(resp)
			deltaLink = "" // to make sure we don't use an old value
		}

		// the deltaLink is kind of like a cursor for overall data state.
		// once we run through pages of nextLinks, the last query will
		// produce a deltaLink instead (if supported), which we'll use on
		// the next backup to only get the changes since this run.
		if len(deltaLink) > 0 {
			deltaURL = deltaLink
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
