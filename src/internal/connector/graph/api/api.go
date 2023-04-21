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

func NextAndDeltaLink(pl PageLinker) (string, string) {
	dpl, ok := pl.(DeltaPageLinker)
	if ok {
		// return delta link if available
		return NextLink(pl), ptr.Val(dpl.GetOdataDeltaLink())
	}

	return NextLink(pl), ""
}
