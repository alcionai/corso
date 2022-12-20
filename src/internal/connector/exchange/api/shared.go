package api

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/pkg/errors"
)

// ---------------------------------------------------------------------------
// generic handler for paging item ids in a container
// ---------------------------------------------------------------------------

type itemPager interface {
	getPage(context.Context) (pageLinker, error)
	idsIn(pageLinker) ([]getIDer, error)
	setNext(string)
}

type pageLinker interface {
	GetOdataDeltaLink() *string
	GetOdataNextLink() *string
}

type getIDer interface {
	GetId() *string
}

// uses a models interface compliant with { GetValues() []T }
// to transform its results into a slice of getIDer interfaces.
// Generics used here to handle the variation of msoft interfaces
// that all _almost_ comply with GetValue, but all return a different
// interface.
func toIders(a any) ([]getIDer, error) {
	gv, ok := a.(interface{ GetValue() []getIDer })
	if !ok {
		return nil, errors.New("response does not comply with GetValue interface")
	}

	return gv.GetValue(), nil
}

// generic controller for retrieving all item ids in a container.
func getContainerIDs(
	ctx context.Context,
	pager itemPager,
	errUpdater func(error),
) ([]string, []string, string, error) {
	var (
		addedIDs   = []string{}
		removedIDs = []string{}
		deltaURL   string
	)

	for {
		resp, err := pager.getPage(ctx)
		if err != nil {
			if err := graph.IsErrDeletedInFlight(err); err != nil {
				return nil, nil, deltaURL, err
			}

			if err := graph.IsErrInvalidDelta(err); err != nil {
				return nil, nil, deltaURL, err
			}

			return nil, nil, deltaURL, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		items, err := pager.idsIn(resp)
		if err != nil {
			return nil, nil, "", err
		}

		for _, item := range items {
			if item.GetId() == nil {
				errUpdater(errors.Errorf("item with nil ID"))

				// TODO: Handle fail-fast.
				continue
			}

			if item.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
				addedIDs = append(addedIDs, *item.GetId())
			} else {
				removedIDs = append(removedIDs, *item.GetId())
			}
		}

		delta := resp.GetOdataDeltaLink()
		if delta != nil && len(*delta) > 0 {
			deltaURL = *delta
		}

		nextLink := resp.GetOdataNextLink()
		if nextLink == nil || len(*nextLink) == 0 {
			break
		}

		pager.setNext(*nextLink)
	}

	return addedIDs, removedIDs, deltaURL, nil
}
