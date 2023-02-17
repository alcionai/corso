package api

import "github.com/alcionai/corso/src/internal/common/ptr"

type PageLinker interface {
	GetOdataNextLink() *string
}

type DeltaPageLinker interface {
	PageLinker
	GetOdataDeltaLink() *string
}

func NextLink(pl PageLinker) string {
	return ptr.Val(pl.GetOdataNextLink())
}

func NextAndDeltaLink(pl DeltaPageLinker) (string, string) {
	return NextLink(pl), ptr.Val(pl.GetOdataDeltaLink())
}
