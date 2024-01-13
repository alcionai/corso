package mock

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

var (
	_ pagers.NonDeltaHandler[any] = &Pager[any]{}
	_ pagers.DeltaHandler[any]    = &DeltaPager[any]{}
)

type DeltaNextLinkValues[T any] struct {
	Next   *string
	Delta  *string
	Values []T
}

func (dnl *DeltaNextLinkValues[T]) GetValue() []T {
	return dnl.Values
}

func (dnl *DeltaNextLinkValues[T]) GetOdataNextLink() *string {
	return dnl.Next
}

func (dnl *DeltaNextLinkValues[T]) GetOdataDeltaLink() *string {
	return dnl.Delta
}

type PagerResult[T any] struct {
	Values     []T
	NextLink   *string
	DeltaLink  *string
	ResetDelta bool
	Err        error
}

// ---------------------------------------------------------------------------
// non-delta pager
// ---------------------------------------------------------------------------

type Pager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *Pager[T]) GetPage(
	context.Context,
) (pagers.NextLinkValuer[T], error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	link := DeltaNextLinkValues[T]{
		Next:   p.ToReturn[idx].NextLink,
		Values: p.ToReturn[idx].Values,
	}

	return &link, p.ToReturn[idx].Err
}

func (p *Pager[T]) SetNextLink(string)  {}
func (p *Pager[T]) ValidModTimes() bool { return true }

// ---------------------------------------------------------------------------
// delta pager
// ---------------------------------------------------------------------------

type DeltaPager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *DeltaPager[T]) GetPage(
	context.Context,
) (pagers.DeltaLinkValuer[T], error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	link := DeltaNextLinkValues[T]{
		Next:   p.ToReturn[idx].NextLink,
		Delta:  p.ToReturn[idx].DeltaLink,
		Values: p.ToReturn[idx].Values,
	}

	return &link, p.ToReturn[idx].Err
}

func (p *DeltaPager[T]) SetNextLink(string)    {}
func (p *DeltaPager[T]) Reset(context.Context) {}
func (p *DeltaPager[T]) ValidModTimes() bool   { return true }
