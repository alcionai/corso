package mock

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type DeltaNextLinks struct {
	Next  *string
	Delta *string
}

func (dnl *DeltaNextLinks) GetOdataNextLink() *string {
	return dnl.Next
}

func (dnl *DeltaNextLinks) GetOdataDeltaLink() *string {
	return dnl.Delta
}

type PagerResult[T any] struct {
	Values    []T
	NextLink  *string
	DeltaLink *string
	Err       error
}

// ---------------------------------------------------------------------------
// non-delta pager
// ---------------------------------------------------------------------------

type Pager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *Pager[T]) GetPage(context.Context) (api.PageLinker, error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	link := DeltaNextLinks{Next: p.ToReturn[idx].NextLink}

	return &link, p.ToReturn[idx].Err
}

func (p *Pager[T]) SetNext(string) {}

func (p *Pager[T]) ValuesIn(api.PageLinker) ([]T, error) {
	idx := p.getIdx
	if idx > 0 {
		// Return values lag by one since we increment in GetPage().
		idx--
	}

	if len(p.ToReturn) <= idx {
		return nil, clues.New("index out of bounds").
			With("index", idx, "values", p.ToReturn)
	}

	return p.ToReturn[idx].Values, nil
}

// ---------------------------------------------------------------------------
// delta pager
// ---------------------------------------------------------------------------

type DeltaPager[T any] struct {
	ToReturn []PagerResult[T]
	getIdx   int
}

func (p *DeltaPager[T]) GetPage(context.Context) (api.DeltaPageLinker, error) {
	if len(p.ToReturn) <= p.getIdx {
		return nil, clues.New("index out of bounds").
			With("index", p.getIdx, "values", p.ToReturn)
	}

	idx := p.getIdx
	p.getIdx++

	link := DeltaNextLinks{
		Next:  p.ToReturn[idx].NextLink,
		Delta: p.ToReturn[idx].DeltaLink,
	}

	return &link, p.ToReturn[idx].Err
}

func (p *DeltaPager[T]) SetNext(string)        {}
func (p *DeltaPager[T]) Reset(context.Context) {}

func (p *DeltaPager[T]) ValuesIn(api.PageLinker) ([]T, error) {
	idx := p.getIdx
	if idx > 0 {
		// Return values lag by one since we increment in GetPage().
		idx--
	}

	if len(p.ToReturn) <= idx {
		return nil, clues.New("index out of bounds").
			With("index", idx, "values", p.ToReturn)
	}

	return p.ToReturn[idx].Values, nil
}
