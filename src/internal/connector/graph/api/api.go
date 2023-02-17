package api

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

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

// IsNextLinkValid separate check to investigate whether error is
func IsNextLinkValid(next string) bool {
	return !strings.Contains(next, `users//`)
}
