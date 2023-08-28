package mock

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type PagerResult[T any] struct {
	DeltaLink *string
	NextLink  *string
	Values    []T
	Err       error
}

func (pr PagerResult[T]) GetValue() []T {
	return pr.Values
}

func (pr PagerResult[T]) GetOdataNextLink() *string {
	return pr.NextLink
}

func (pr PagerResult[T]) GetOdataDeltaLink() *string {
	return pr.DeltaLink
}

// ---------------------------------------------------------------------------
// non-delta pager
// ---------------------------------------------------------------------------

type Pager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *Pager[T]) GetPage(context.Context) (api.LinkValuer[T], error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	return &p.ToReturn[idx], p.ToReturn[idx].Err
}

func (p *Pager[T]) SetNext(string) {}

// ---------------------------------------------------------------------------
// delta pager
// ---------------------------------------------------------------------------

type DeltaPager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *DeltaPager[T]) GetPage(context.Context) (api.DeltaLinkValuer[T], error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	return &p.ToReturn[idx], p.ToReturn[idx].Err
}

func (p *DeltaPager[T]) SetNext(string) {}
func (p *DeltaPager[T]) Reset()         {}
