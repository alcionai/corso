package api

type PageLinker interface {
	GetOdataNextLink() *string
}

type DeltaPageLinker interface {
	PageLinker
	GetOdataDeltaLink() *string
}

func NextLink(pl PageLinker) string {
	next := pl.GetOdataNextLink()
	if next == nil || len(*next) == 0 {
		return ""
	}

	return *next
}

func NextAndDeltaLink(pl DeltaPageLinker) (string, string) {
	next := NextLink(pl)

	delta := pl.GetOdataDeltaLink()
	if delta == nil || len(*delta) == 0 {
		return next, ""
	}

	return next, *delta
}
