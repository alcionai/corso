package graph

import (
	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// query parameter aggregation
// ---------------------------------------------------------------------------

type QueryParams struct {
	Category      path.CategoryType
	ResourceOwner string
	Credentials   account.M365Config
	DeltaPaths    DeltaPaths
}

func (qp QueryParams) Validate() error {
	if qp.Category == path.UnknownCategory {
		return clues.New("unknown category")
	}

	if len(qp.ResourceOwner) == 0 {
		return clues.New("missing resource owner")
	}

	return nil
}

// ---------------------------------------------------------------------------
// delta path aggregation
//
// TODO: probably needs to be owned somewhere besides graph, but for now
// this allows the struct to be centralized and re-used through GC.
// ---------------------------------------------------------------------------

type CatDeltaPaths map[path.CategoryType]DeltaPaths

type DeltaPaths map[string]DeltaPath

func (dps DeltaPaths) AddDelta(k, d string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Delta = d
	dps[k] = dp
}

func (dps DeltaPaths) AddPath(k, p string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Path = p
	dps[k] = dp
}

type DeltaPath struct {
	Delta string
	Path  string
}
