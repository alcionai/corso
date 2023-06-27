package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// common interfaces and funcs
// ---------------------------------------------------------------------------

type PageLinker interface {
	GetOdataNextLink() *string
}

type DeltaPageLinker interface {
	PageLinker
	GetOdataDeltaLink() *string
}

// IsNextLinkValid separate check to investigate whether error is
func IsNextLinkValid(next string) bool {
	return !strings.Contains(next, `users//`)
}

func NextLink(pl PageLinker) string {
	return ptr.Val(pl.GetOdataNextLink())
}

func NextAndDeltaLink(pl DeltaPageLinker) (string, string) {
	return NextLink(pl), ptr.Val(pl.GetOdataDeltaLink())
}

type Valuer[T any] interface {
	GetValue() []T
}

type PageLinkValuer[T any] interface {
	PageLinker
	Valuer[T]
}

// EmptyDeltaLinker is used to convert PageLinker to DeltaPageLinker
type EmptyDeltaLinker[T any] struct {
	PageLinkValuer[T]
}

func (EmptyDeltaLinker[T]) GetOdataDeltaLink() *string {
	return ptr.To("")
}

func (e EmptyDeltaLinker[T]) GetValue() []T {
	return e.PageLinkValuer.GetValue()
}

// ---------------------------------------------------------------------------
// generic handler for non-delta item paging in a container
// ---------------------------------------------------------------------------

type itemPager[T any] interface {
	// getPage get a page with the specified options from graph
	getPage(context.Context) (PageLinkValuer[T], error)
	// setNext is used to pass in the next url got from graph
	setNext(string)
}

func enumerateItems[T any](
	ctx context.Context,
	pager itemPager[T],
) ([]T, error) {
	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		nextLink = "do-while"
	)

	for len(nextLink) > 0 {
		// get the next page of data, check for standard errors
		resp, err := pager.getPage(ctx)
		if err != nil {
			return nil, graph.Stack(ctx, err)
		}

		result = append(result, resp.GetValue()...)
		nextLink = NextLink(resp)

		pager.setNext(nextLink)
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", len(result))

	return result, nil
}

// ---------------------------------------------------------------------------
// generic handler for delta-based ittem paging in a container
// ---------------------------------------------------------------------------

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

type itemIDPager interface {
	// getPage get a page with the specified options from graph
	getPage(context.Context) (DeltaPageLinker, error)
	// setNext is used to pass in the next url got from graph
	setNext(string)
	// reset is used to clear delta url in delta pagers. When
	// reset is called, we reset the state(delta url) that we
	// currently have and start a new delta query without the token.
	reset(context.Context)
	// valuesIn gets us the values in a page
	valuesIn(PageLinker) ([]getIDAndAddtler, error)
}

type getIDAndAddtler interface {
	GetId() *string
	GetAdditionalData() map[string]any
}

func getAddedAndRemovedItemIDs(
	ctx context.Context,
	service graph.Servicer,
	pager itemIDPager,
	deltaPager itemIDPager,
	oldDelta string,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	var (
		pgr        itemIDPager
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
	pager itemIDPager,
) ([]string, []string, string, error) {
	var (
		addedIDs   = []string{}
		removedIDs = []string{}
		deltaURL   string
		itemCount  int
		page       int
	)

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

		nextLink, deltaLink := NextAndDeltaLink(resp)

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
