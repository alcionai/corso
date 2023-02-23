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

// IsNextLinkValid separate check to investigate whether error is
func IsNextLinkValid(next string) bool {
	return !strings.Contains(next, `users//`)
}

func NextLink(pl PageLinker) string {
	return ptr.Val(pl.GetOdataNextLink())
}

func NextAndDeltaLink(pl DeltaPageLinker) (string, string) {
	return NextLink(pl), ptr.Val(pl.GetOdataDeltaLink())
}
