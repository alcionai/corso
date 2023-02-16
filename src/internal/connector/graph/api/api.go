package api

import "github.com/alcionai/corso/src/internal/common/ptr"

type PageLinker interface {
	GetOdataNextLink() *string
}

type DeltaPageLinker interface {
	PageLinker
	GetOdataDeltaLink() *string
}

func NextAndDeltaLink(pl DeltaPageLinker) (string, string) {
	next := ptr.Val(pl.GetOdataNextLink())
	delta := ptr.Val(pl.GetOdataDeltaLink())

	return next, delta
}
