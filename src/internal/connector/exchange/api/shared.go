package api

import (
	"context"
	"fmt"
	"os"

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
	getPage(context.Context) (api.PageLinker, error)
	setNext(string)
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
	user, directoryID, oldDelta string,
	pagerGetter func(context.Context, graph.Servicer, string, string, bool) (itemPager, error),
	deltaPagerGetter func(context.Context, graph.Servicer, string, string, string, bool) (itemPager, error),
	immutableIDs bool,
) ([]string, []string, DeltaUpdate, error) {
	pgr, err := deltaPagerGetter(ctx, service, user, directoryID, oldDelta, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating delta pager")
	}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	// note: happy path, not the error condition
	if err == nil {
		return added, removed, DeltaUpdate{deltaURL, len(oldDelta) != 0}, err
	}

	// return error if invalid delta error or if we did a non-delta fetch
	if !graph.IsErrInvalidDelta(err) && !graph.IsErrQuotaExceeded(err) {
		return nil, nil, DeltaUpdate{}, err
	}

	if graph.IsErrQuotaExceeded(err) {
		pgr, err = pagerGetter(ctx, service, user, directoryID, immutableIDs)
		if err != nil {
			return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating pager")
		}
	} else {
		if len(oldDelta) == 0 {
			// if we have already tried with empty delta, don't retry
			return nil, nil, DeltaUpdate{}, err
		}

		// Create mailDeltaPager without previous delta
		pgr, err = deltaPagerGetter(ctx, service, user, directoryID, "", immutableIDs)
		if err != nil {
			return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating delta pager without previous delta")
		}
	}

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

		nextLink, deltaLink := api.NextAndDeltaLink(resp)
		if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
			if !api.IsNextLinkValid(nextLink) || api.IsNextLinkValid(deltaLink) {
				logger.Ctx(ctx).Infof("Received invalid link from M365:\nNext Link: %s\nDelta Link: %s\n", nextLink, deltaLink)
			}
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
