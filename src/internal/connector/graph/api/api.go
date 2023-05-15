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

type Valuer[T any] interface {
	GetValue() []T
}

type PageLinkValuer[T any] interface {
	PageLinker
	Valuer[T]
}

// emptyDeltaLinker is used to convert PageLinker to DeltaPageLinker
type emptyDeltaLinker[T any] struct {
	PageLinkValuer[T]
}

func (emptyDeltaLinker[T]) GetOdataDeltaLink() *string {
	empty := ""
	return &empty
}

func EmptyDeltaLinker[T any](e PageLinkValuer[T]) emptyDeltaLinker[T] {
	return emptyDeltaLinker[T]{e}
}

func (e emptyDeltaLinker[T]) GetValue() []T {
	return e.PageLinkValuer.GetValue()
}
