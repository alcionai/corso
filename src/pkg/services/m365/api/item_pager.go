package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// common structs
// ---------------------------------------------------------------------------

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queryable folder.
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}

type NextPage[T any] struct {
	Items []T
	// Reset is only true on the iteration where the delta pager's Reset()
	// is called.  Callers can use it to reset any data aggregation they
	// currently use.  After that loop, it will be false again, though the
	// DeltaUpdate will still contain the expected value.
	Reset bool
}

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

type ValidModTimer interface {
	ValidModTimes() bool
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
	ValidModTimer
}

func enumerateItems[T any](
	ctx context.Context,
	pager Pager[T],
	ch chan<- NextPage[T],
) error {
	defer close(ch)

	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		nextLink = "do-while"
	)

	for len(nextLink) > 0 {
		// get the next page of data, check for standard errors
		page, err := pager.GetPage(ctx)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		ch <- NextPage[T]{Items: page.GetValue()}

		nextLink = NextLink(page)

		pager.SetNextLink(nextLink)
	}

	logger.Ctx(ctx).Infow("completed delta item enumeration", "result_count", len(result))

	return nil
}

func batchEnumerateItems[T any](
	ctx context.Context,
	pager Pager[T],
) ([]T, error) {
	var (
		ch      = make(chan NextPage[T])
		results = []T{}
	)

	go func() {
		for np := range ch {
			results = append(results, np.Items...)
		}
	}()

	err := enumerateItems[T](ctx, pager, ch)

	return results, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// generic handler for delta-based item paging
// ---------------------------------------------------------------------------

type DeltaPager[T any] interface {
	GetPager[DeltaLinkValuer[T]]
	Resetter
	SetNextLinker
	ValidModTimer
}

// enumerates pages of items, streaming each page to the provided channel.
// the DeltaUpdate, reset notifications, and any errors are also fed to the
// same channel.
// Returns false if conditions disallow making delta calls for the provided
// pager.  Returns true otherwise, even in the event of an error.
func deltaEnumerateItems[T any](
	ctx context.Context,
	pager DeltaPager[T],
	ch chan<- NextPage[T],
	prevDeltaLink string,
) (DeltaUpdate, error) {
	defer close(ch)

	var (
		result = make([]T, 0)
		// stubbed initial value to ensure we enter the loop.
		newDeltaLink     = ""
		invalidPrevDelta = len(prevDeltaLink) == 0
		nextLink         = "do-while"
		consume          = graph.SingleGetOrDeltaLC
	)

	if invalidPrevDelta {
		// Delta queries with no previous token cost more.
		consume = graph.DeltaNoTokenLC
	}

	// Loop through all pages returned by Graph API.
	for len(nextLink) > 0 {
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, consume))
		if graph.IsErrDeltaNotSupported(err) {
			logger.Ctx(ctx).Infow("delta queries not supported")

			pager.Reset(ctx)
			ch <- NextPage[T]{Reset: true}

			return DeltaUpdate{}, clues.Stack(err)
		}

		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("invalid previous delta", "delta_link", prevDeltaLink)

			invalidPrevDelta = true

			// Set limiter consumption rate to non-delta.
			consume = graph.DeltaNoTokenLC

			// Reset tells the pager to try again after ditching its delta history.
			pager.Reset(ctx)
			ch <- NextPage[T]{Reset: true}

			continue
		}

		if err != nil {
			return DeltaUpdate{}, clues.Stack(err)
		}

		ch <- NextPage[T]{Items: page.GetValue()}

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

	return du, nil
}

func batchDeltaEnumerateItems[T any](
	ctx context.Context,
	pager DeltaPager[T],
	prevDeltaLink string,
) ([]T, DeltaUpdate, error) {
	var (
		ch      = make(chan NextPage[T])
		results = []T{}
	)

	go func() {
		for np := range ch {
			if np.Reset {
				results = []T{}
			}

			results = append(results, np.Items...)
		}
	}()

	du, err := deltaEnumerateItems[T](ctx, pager, ch, prevDeltaLink)

	return results, du, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// shared enumeration runner funcs
// ---------------------------------------------------------------------------

type addedAndRemovedHandler[T any] func(items []T) (map[string]time.Time, []string, error)

func getAddedAndRemovedItemIDs[T any](
	ctx context.Context,
	pager Pager[T],
	deltaPager DeltaPager[T],
	prevDeltaLink string,
	canMakeDeltaQueries bool,
	aarh addedAndRemovedHandler[T],
) (map[string]time.Time, bool, []string, DeltaUpdate, error) {
	if canMakeDeltaQueries {
		ts, du, err := batchDeltaEnumerateItems[T](ctx, deltaPager, prevDeltaLink)
		if err != nil && !graph.IsErrInvalidDelta(err) && !graph.IsErrDeltaNotSupported(err) {
			return nil, false, nil, DeltaUpdate{}, graph.Stack(ctx, err)
		}

		if err == nil {
			a, r, err := aarh(ts)
			return a, deltaPager.ValidModTimes(), r, du, graph.Stack(ctx, err).OrNil()
		}
	}

	du := DeltaUpdate{Reset: true}

	ts, err := batchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, false, nil, DeltaUpdate{}, graph.Stack(ctx, err)
	}

	a, r, err := aarh(ts)

	return a, pager.ValidModTimes(), r, du, graph.Stack(ctx, err).OrNil()
}

type getIDer interface {
	GetId() *string
}

// for added and removed by additionalData[@removed]

type getIDAndAddtler interface {
	getIDer
	GetAdditionalData() map[string]any
}

type getModTimer interface {
	GetLastModifiedDateTime() *time.Time
}

func addedAndRemovedByAddtlData[T any](
	items []T,
) (map[string]time.Time, []string, error) {
	added := map[string]time.Time{}
	removed := []string{}

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
			var modTime time.Time

			if mt, ok := giaa.(getModTimer); ok {
				modTime = ptr.Val(mt.GetLastModifiedDateTime())
			}

			added[ptr.Val(giaa.GetId())] = modTime
		} else {
			removed = append(removed, ptr.Val(giaa.GetId()))
		}
	}

	return added, removed, nil
}

// for added and removed by GetDeletedDateTime()

type getIDAndDeletedDateTimer interface {
	getIDer
	GetDeletedDateTime() *time.Time
}

func addedAndRemovedByDeletedDateTime[T any](
	items []T,
) (map[string]time.Time, []string, error) {
	added := map[string]time.Time{}
	removed := []string{}

	for _, item := range items {
		giaddt, ok := any(item).(getIDAndDeletedDateTimer)
		if !ok {
			return nil, nil, clues.New("item does not provide id and deleted date time getters").
				With("item_type", fmt.Sprintf("%T", item))
		}

		if giaddt.GetDeletedDateTime() == nil {
			var modTime time.Time

			if mt, ok := giaddt.(getModTimer); ok {
				// Make sure to get a non-zero mod time if the item doesn't have one for
				// some reason. Otherwise we can hit an issue where kopia has a
				// different mod time for the file than the details does. This occurs
				// due to a conversion kopia does on the time from
				// time.Time -> nanoseconds for serialization. During incremental
				// backups, kopia goes from nanoseconds -> time.Time but there's an
				// overflow which yields a different timestamp.
				// https://github.com/gohugoio/hugo/issues/6161#issuecomment-725915786
				modTime = ptr.OrNow(mt.GetLastModifiedDateTime())
			}

			added[ptr.Val(giaddt.GetId())] = modTime
		} else {
			removed = append(removed, ptr.Val(giaddt.GetId()))
		}
	}

	return added, removed, nil
}
