package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// generic handler for paging item ids in a container
// ---------------------------------------------------------------------------

type itemPager interface {
	// getNextPage will fetch what is created by builder inside and
	// sets builder to the next page.
	// Return values are items, is-last-page, delta-url, error
	getNextPage(context.Context) ([]getIDAndAddtler, bool, string, error)

	// reset is used to reset the delta url or to switch to using
	// non-delta fetch
	reset(bool)
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

func GetAddedAndRemovedItemIDsFromPager(
	ctx context.Context,
	oldDelta string,
	pgr itemPager,
) ([]string, []string, DeltaUpdate, error) {
	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err == nil {
		return added, removed, DeltaUpdate{deltaURL, false}, nil
	}

	// if we fail with quota exceeded error, retry with non delta
	if graph.IsErrQuotaExceeded(err) {
		pgr.reset(true)

		logger.CtxErr(ctx, err).Info("switching to non-delta pagination")

		added, removed, deltaURL, err = getItemsAddedAndRemovedFromContainer(ctx, pgr)
		if err == nil {
			// TODO(meain): Should we return resetDelta here? Will
			// returning reset delta ensure we create a non
			// incremental backup?
			return added, removed, DeltaUpdate{deltaURL, true}, nil
		}
	}

	// if we failed with invalidDelta error and prev delta was not
	// empty, reset and retry
	if graph.IsErrInvalidDelta(err) && len(oldDelta) != 0 {
		pgr.reset(false)

		logger.CtxErr(ctx, err).Info("resetting delta pagination")

		added, removed, deltaURL, err = getItemsAddedAndRemovedFromContainer(ctx, pgr)
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, true}, nil
		}
	}

	return nil, nil, DeltaUpdate{}, err
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
		items      []getIDAndAddtler
		lastPage   bool
		err        error
	)

	itemCount := 0
	page := 0

	for {
		// get the next page of data, check for standard errors
		items, lastPage, deltaURL, err = pager.getNextPage(ctx)
		if err != nil {
			return nil, nil, deltaURL, graph.Stack(ctx, err)
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

		if lastPage {
			break
		}
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", itemCount)

	return addedIDs, removedIDs, deltaURL, nil
}
