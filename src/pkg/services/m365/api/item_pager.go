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
// common interfaces
// ---------------------------------------------------------------------------

type DeltaPager[T any] interface {
	DeltaGetPager[T]
	Resetter
	SetNextLinker
}

type Pager[T any] interface {
	GetPager[T]
	SetNextLinker
}

type DeltaGetPager[T any] interface {
	GetPage(context.Context) (DeltaLinkValuer[T], error)
}

type GetPager[T any] interface {
	GetPage(context.Context) (LinkValuer[T], error)
}

type Valuer[T any] interface {
	GetValue() []T
}

type GetNextLinker interface {
	GetOdataNextLink() *string
}

type GetDeltaLinker interface {
	GetNextLinker
	GetOdataDeltaLink() *string
}

type LinkValuer[T any] interface {
	Valuer[T]
	GetNextLinker
}

type DeltaLinkValuer[T any] interface {
	LinkValuer[T]
	GetDeltaLinker
}

type SetNextLinker interface {
	SetNext(nextLink string)
}

type Resetter interface {
	Reset()
}

// ---------------------------------------------------------------------------
// common funcs
// ---------------------------------------------------------------------------

// IsNextLinkValid separate check to investigate whether error is
func IsNextLinkValid(next string) bool {
	return !strings.Contains(next, `users//`)
}

func NextLink(gnl GetNextLinker) string {
	return ptr.Val(gnl.GetOdataNextLink())
}

func NextAndDeltaLink(gdl GetDeltaLinker) (string, string) {
	return NextLink(gdl), ptr.Val(gdl.GetOdataDeltaLink())
}

// EmptyDeltaLinker is used to convert PageLinker to DeltaPageLinker
type EmptyDeltaLinker[T any] struct {
	LinkValuer[T]
}

func (EmptyDeltaLinker[T]) GetOdataDeltaLink() *string {
	return ptr.To("")
}

func (e EmptyDeltaLinker[T]) GetValue() []T {
	return e.GetValue()
}

// ---------------------------------------------------------------------------
// generic handler for non-delta item paging in a container
// ---------------------------------------------------------------------------

func enumerateItems[T any](
	ctx context.Context,
	pager Pager[T],
) ([]T, error) {
	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		nextLink = "do-while"
	)

	for len(nextLink) > 0 {
		// get the next page of data, check for standard errors
		resp, err := pager.GetPage(ctx)
		if err != nil {
			return nil, graph.Stack(ctx, err)
		}

		result = append(result, resp.GetValue()...)
		nextLink = NextLink(resp)

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", len(result))

	return result, nil
}

// ---------------------------------------------------------------------------
// generic handler for delta-based item paging in a container
// ---------------------------------------------------------------------------

func enumerateDeltaItems[T any](
	ctx context.Context,
	pager DeltaPager[T],
) ([]T, error) {
	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		nextLink = "do-while"
	)

	for len(nextLink) > 0 {
		// get the next page of data, check for standard errors
		resp, err := pager.GetPage(ctx)
		if err != nil {
			return nil, graph.Stack(ctx, err)
		}

		result = append(result, resp.GetValue()...)
		nextLink = NextLink(resp)

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", len(result))

	return result, nil
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

type getIDAndAddtler interface {
	GetId() *string
	GetAdditionalData() map[string]any
}

func getAddedAndRemovedItemIDs[T any](
	ctx context.Context,
	service graph.Servicer,
	pager DeltaPager[T],
	deltaPager DeltaPager[T],
	oldDelta string,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	var (
		pgr        DeltaPager[T]
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
	pgr.Reset()

	added, removed, deltaURL, err = getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, true}, nil
}

// generic controller for retrieving all item ids in a container.
func getItemsAddedAndRemovedFromContainer[T any](
	ctx context.Context,
	pager DeltaPager[T],
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
		resp, err := pager.GetPage(ctx)
		if err != nil {
			return nil, nil, deltaURL, graph.Stack(ctx, err)
		}

		// each category type responds with a different interface, but all
		// of them comply with GetValue, which is where we'll get our item data.
		items := resp.GetValue()

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

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Infow("completed enumeration", "count", itemCount)

	return addedIDs, removedIDs, deltaURL, nil
}
