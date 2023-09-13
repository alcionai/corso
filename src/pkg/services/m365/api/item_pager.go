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

type GetPager[T any] interface {
	GetPage(context.Context) (T, error)
}

type NextLinkValuer[T any] interface {
	NextLinker
	Valuer[T]
}

type NextLinker interface {
	GetOdataNextLink() *string
}

type SetNextLinker interface {
	SetNextLink(nextLink string)
}

type DeltaLinker interface {
	NextLinker
	GetOdataDeltaLink() *string
}

type DeltaLinkValuer[T any] interface {
	DeltaLinker
	Valuer[T]
}

type Valuer[T any] interface {
	GetValue() []T
}

type Resetter interface {
	Reset(context.Context)
}

// ---------------------------------------------------------------------------
// common funcs
// ---------------------------------------------------------------------------

// IsNextLinkValid separate check to investigate whether error is
func IsNextLinkValid(next string) bool {
	return !strings.Contains(next, `users//`)
}

func NextLink(pl NextLinker) string {
	return ptr.Val(pl.GetOdataNextLink())
}

func NextAndDeltaLink(pl DeltaLinker) (string, string) {
	return NextLink(pl), ptr.Val(pl.GetOdataDeltaLink())
}

// ---------------------------------------------------------------------------
// non-delta item paging
// ---------------------------------------------------------------------------

type Pager[T any] interface {
	GetPager[NextLinkValuer[T]]
	SetNextLinker
}

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
		page, err := pager.GetPage(ctx)
		if err != nil {
			return nil, graph.Stack(ctx, err)
		}

		result = append(result, page.GetValue()...)
		nextLink = NextLink(page)

		pager.SetNextLink(nextLink)
	}

	logger.Ctx(ctx).Infow("completed delta item enumeration", "result_count", len(result))

	return result, nil
}

// ---------------------------------------------------------------------------
// generic handler for delta-based item paging
// ---------------------------------------------------------------------------

type DeltaPager[T any] interface {
	GetPager[DeltaLinkValuer[T]]
	Resetter
	SetNextLinker
}

func deltaEnumerateItems[T any](
	ctx context.Context,
	pager DeltaPager[T],
	prevDeltaLink string,
) ([]T, DeltaUpdate, error) {
	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		newDeltaLink     = ""
		invalidPrevDelta = len(prevDeltaLink) == 0
		nextLink         = "do-while"
	)

	// Loop through all pages returned by Graph API.
	for len(nextLink) > 0 {
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))
		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("invalid previous delta", "delta_link", prevDeltaLink)

			invalidPrevDelta = true
			result = make([]T, 0)

			// Reset tells the pager to try again after ditching its delta history.
			pager.Reset(ctx)

			continue
		}

		if err != nil {
			return nil, DeltaUpdate{}, graph.Wrap(ctx, err, "retrieving page")
		}

		result = append(result, page.GetValue()...)

		nl, deltaLink := NextAndDeltaLink(page)
		if len(deltaLink) > 0 {
			newDeltaLink = deltaLink
		}

		nextLink = nl
		pager.SetNextLink(nextLink)
	}

	logger.Ctx(ctx).Debugw("completed delta item enumeration", "result_count", len(result))

	du := DeltaUpdate{
		URL:   newDeltaLink,
		Reset: invalidPrevDelta,
	}

	return result, du, nil
}

// ---------------------------------------------------------------------------
// shared enumeration runner funcs
// ---------------------------------------------------------------------------

func getAddedAndRemovedItemIDs[T any](
	ctx context.Context,
	pager Pager[T],
	deltaPager DeltaPager[T],
	prevDeltaLink string,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	if canMakeDeltaQueries {
		ts, du, err := deltaEnumerateItems[T](ctx, deltaPager, prevDeltaLink)
		if err != nil && (!graph.IsErrInvalidDelta(err) || len(prevDeltaLink) == 0) {
			return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
		}

		if err == nil {
			a, r, err := addedAndRemovedByAddtlData(ts)
			return a, r, du, graph.Stack(ctx, err).OrNil()
		}
	}

	du := DeltaUpdate{Reset: true}

	ts, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
	}

	a, r, err := addedAndRemovedByAddtlData[T](ts)

	return a, r, du, graph.Stack(ctx, err).OrNil()
}

type getIDAndAddtler interface {
	GetId() *string
	GetAdditionalData() map[string]any
}

func addedAndRemovedByAddtlData[T any](items []T) ([]string, []string, error) {
	added, removed := []string{}, []string{}

	for _, item := range items {
		giaa, ok := any(item).(getIDAndAddtler)
		if !ok {
			return nil, nil, clues.New("item does not provide id and additional data getters").
				With("item_type", fmt.Sprintf("%T", item))
		}

		// if the additional data contains a `@removed` key, the value will either
		// be 'changed' or 'deleted'.  We don't really care about the cause: both
		// cases are handled the same way in storage.
		if giaa.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
			added = append(added, ptr.Val(giaa.GetId()))
		} else {
			removed = append(removed, ptr.Val(giaa.GetId()))
		}
	}

	return added, removed, nil
}
